package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"flag"
	"errors"
	"strings"
	"net/http"
	
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

var (
	messageCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hipchat_room_messages_total",
			Help: "Number of messages produced by a HipChat room.",
		},
		[]string{"name"},
	)
	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hipchat_room_errors_total",
			Help: "Number of errors while accessing a HipChat room.",
		},
		[]string{"name"},
	)
)

func processRoom(hipchatClient *hipchat.Client, roomName string, verbose bool) error {
	room, _, err := hipchatClient.Room.GetStatistics(roomName)
	if err != nil {
		log.Printf("Could not fetch stats of Room %s: %s\n", roomName, err)
		errorCounter.With(prometheus.Labels{"name": roomName}).Inc()
		return err
	} 
	if verbose {
		log.Printf("Room %s has %d messages\n", roomName, room.MessagesSent)
	}
	messageCounter.With(prometheus.Labels{"name": roomName}).Set(float64(room.MessagesSent))
	return nil
}

func processRooms(hipchatClient *hipchat.Client, roomNames string, verbose bool) error {
	var hasErrors = false
	roomNamesArray := strings.Split(roomNames, ",")
	for i := range roomNamesArray {
		if processRoom(hipchatClient, roomNamesArray[i], verbose) != nil {
			hasErrors = true
		}
	}
	if hasErrors {
		return errors.New("Could not process all rooms")
	}
	return nil
}	

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(messageCounter)
	prometheus.MustRegister(errorCounter)
}

func main() {
	roomNames := flag.String("rooms", "", "The list of the HipChat room names, separated by a comma.")
	authToken := flag.String("authtoken", "", "The access token used to call the HipChat API.")
	port      := flag.Int("port", 8080, "The http port the server will listen on")
	interval  := flag.Duration("interval", time.Second * 30, "The interval between 2 scrapes")
	verbose   := flag.Bool("verbose", true, "Should we log the requests?")
	
	// Validate flags
	flag.Parse()
	if *roomNames == "" {
		fmt.Println("You must specify a list of rooms!")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}
	if *authToken == "" {
		fmt.Println("You must specify the token argument!")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}
	hipchatClient := hipchat.NewClient(*authToken)
	// Get initial stats
	log.Println("Fetching initial metrics")
	if processRooms(hipchatClient, *roomNames, *verbose) != nil {
		log.Println("Could not fetch all the stats. Exiting now!")
		os.Exit(2)
	}
	// Start ticker to collect the HipChat room stats
	log.Println("Starting the ticker")
	ticker := time.NewTicker(*interval)
    go func() {
        for range ticker.C {
			processRooms(hipchatClient, *roomNames, *verbose)
        }
	}()
	defer ticker.Stop()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
