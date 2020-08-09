## Qualify Offer Calculator

Application that scrapes a data set and calculates the qualifing offer for 2017, based on [2016 player salary data](https://questionnaire-148920.appspot.com/swe/data.html).

The qualifying offer is the average of the highest 125 salaries; players without a salary posted are ignored.

## Running the application
### Docker
Application is available as a public docker image at [`barrytam20/qualifying-offer-calculator:1.0.0`](https://hub.docker.com/repository/docker/barrytam20/qualifying-offer-calculator). To run, please ensure you have the [docker](https://www.docker.com/get-started) client installed and running.

To run, execute `docker run barrytam20/qualifying-offer-calculator:1.0.0`
```
docker run barrytam20/qualifying-offer-calculator:1.0.0
Unable to find image 'barrytam20/qualifying-offer-calculator:1.0.0' locally
1.0.0: Pulling from barrytam20/qualifying-offer-calculator
8e402f1a9c57: Already exists 
d8b2755dcc6c: Pull complete 
ddea03ec8827: Pull complete 
Digest: sha256:33eccbdbe2aeb9f474cf550bdff09d2bbb8f5db053fdfa69fc3ab0d68292286e
Status: Downloaded newer image for barrytam20/qualifying-offer-calculator:1.0.0
loading data from https://questionnaire-148920.appspot.com/swe/data.html
qualifying offer is: $16,667,965.57
number of player salaries processed: 1164
number of players missing salary data: 44
```

### Running locally
Application can also be run locally with [go](https://golang.org/dl/). Ensure you have at least version 1.13 of go installed, clone this repository, and run the following commands at the root of the project to view the results
1. `go mod tidy` 
1. `go run main.go`

```
$ go mod tidy
$ go run main.go
loading data from https://questionnaire-148920.appspot.com/swe/data.html
qualifying offer is: $16,557,794.14
number of player salaries processed: 1166
number of players missing salary data: 42
```