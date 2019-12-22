# Introduction
This was a task given to me when I applied for a job at a company in March 2018. The work is done in Golang.

# The Task
An HTTP REST Service that is responsible for geo location and fare calculation. This service has two main functions. 

1) Determine the provided points region. If it is in Istanbul, the service returns string Asia or Europe. If it is not in Istanbul, the service returns (400) Bad Request.

2) Get two coordinates and calculate the estimated distance (in meters), duration (in seconds) and fare (in kurus). 

Both functions get JSON and return JSON objects. Below are example calculation parameters. (Google Maps SDK can be used to get the estimated duration and distance data).

# Examples

## Request 1
GET/ request for region:
input: {"lat":41.057808, "lng":29.008149}
output: {"region":"Europe"}

## Request 2
GET/ request for fare calculation:
input: {
     "from":{ "lat":41.057808, "lng":29.008149},
     "to":{ "lat":41.062756, "lng":29.011577}
}
output: {"duration":100, "distance":1000, "fare":1000}

## Fare Calculation Parameters
100m fee: 0.20 TRY
Opening fee: 5 TRY
Minimum fee: 10 TRY (is taken if the calculated fee is less than 10 TRY)
