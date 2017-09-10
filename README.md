# hipchat-prometheus-exporter

## Description

This exporter will call the HipChat Room API in order to collect the number of messages sent for your selected rooms.

It will then create the following Prometheus metrics that you can use in a Dashboard for showing a room activity:

- hipchat_room_messages_total: the number of sent messages
- hipchat_room_errors_total: the number of errors while trying to fetch the stats

Note that each metric will have a **"name"** property containing the Room's name.

## Requirements

You must create a AuthToken in the HipChat admin pages:

Goto https://MyCompany.hipchat.com/account/api

Then create a new Token with at least the *"View Room"* scope.

## Build

To create a local docker image, execute:

```
./scripts/build-image.sh
```

## Run

Once you have the image built and your AuthTojen, you can run it in Docker locally for testing it:

```
./scripts/run-image.sh -authToken myToken -rooms room1,room2
```

And then you can access the metrics:

```
open http://localhost:8080
```


## Kubernetes

Use the Helm Chart for installing it.

See the [README](charts/hipchat-prometheus-exporter/README.md)


