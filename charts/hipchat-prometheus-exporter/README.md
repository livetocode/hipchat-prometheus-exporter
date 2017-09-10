# hipchat-prometheus-exporter Helm Chart

A Prometheus Exporter that will create a metric for the number of messages produced by a room.

## Chart Details
This chart will do the following:

* Deploy the service
* Create a secret for the AuthToken
* Auto-register the service with Prometheus by leveraging the Prometheus annotations.

## Requirements

You must create a AuthToken in the HipChat admin pages:

Goto https://MyCompany.hipchat.com/account/api

Then create a new Token with at least the *"View Room"* scope.

You will then have to provide it either in the values.yaml or as a command-line parameter (--set hipchat.authToken=123).

## Installing the Chart

To install the chart with the release name `my-release` and define several rooms, as well as the AuthToken:

```bash
$ helm install --name my-release charts/hipchat-prometheus-exporter --set hipchat.rooms[0]=room1,hipchat.rooms[1]=room2,hipchat.authToken=myToken,hipchat.verbose=true
```

## Configuration

The following tables lists the configurable parameters of the Jenkins chart and their default values.


| Parameter                         | Description                          | Default                                                                      |
| --------------------------------- | ------------------------------------ | ---------------------------------------------------------------------------- |
| `hipchat.authToken`               | The AuthToken for calling the APIs   |                                                                              |
| `hipchat.rooms`                   | A list of room names                 | room1, room2                                                                 |
| `hipchat.interval`                | The time interval between 2 scrapes  | 30s                                                                          |
| `hipchat.verbose`                 | Should we log the API results?       | false                                                                        |
| `image.repository`                | Image name                           | `livetocode/hipchat-prometheus-exporter`                                     |
| `image.tag`                       | Image tag                            | `latest`                                                                     |
| `image.pullPolicy`                | Image pull policy                    | `Always`                                                                     |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
$ helm install --name my-release -f values.yaml charts/hipchat-prometheus-exporter
```

And for upgrading:
```bash
helm upgrade my-release charts/hipchat-prometheus-exporter/ --set hipchat.authToken=myToken,hipchat.rooms[0]=room1,hipchat.rooms[1]=room2,hipchat.verbose=true
```

> **Tip**: You can use the default [values.yaml](values.yaml)