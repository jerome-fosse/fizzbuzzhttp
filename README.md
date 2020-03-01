FizzBuzz HTTP
=
FizzBuzz HTTP is a small project written in Go that expose a REST API that generate a fizzbuzz sequence with a list of numbers that are replaced by fizz if they can be divided by 3, by buzz if they can be divided by 5, and by fizzbuzz if they can be divided by 3 and 5.

# Installation
## Requirements

- Go 1.13
- Docker (any recent version) if you wish to build the docker image

## Makefile
GNU/Make is used to compile and install the project. Several commands are available.

- make build : Compile, test and build the project
- make install : Compile, test, build and install the project
- make docker-build : Build the docker image
- make integration-test : Build the docker image and run the integrations tests
- make docker-push : Build, run the integrations test and push the docker image to docker hub

# Execution
Simply run 
```
./fizzbuzzhttp
```
It will start FizzBuzz HTTP on the default port 8080.

fizzbuzzhttp accept two parameters

- port=\<number> : set the port fizzbuzzhttp will use
- verbose : display debug logs

Exemple :
```
./fizzbuzzhttp --port=8081 --verbose
```

# Endpoints
## 1 - GET /whoami
Show version number.

## 2 - GET /fizzbuzz
This endpoint generate a fizzbuzz sequence with 15 elements. Its behaviour can be customized with 5 query parameters :

- int1 : every numbers that can be divided by int1 will be replaced by fizz or an other word that can be set with a parameter.<br>
Default value is 3.
- int2 : every numbers that can be divided by int2 will be replaced by buzz or an other word that can be set with a parameter.<br>
Default value is 5.
- word1 : the word that will replace every number that can be divided by int1<br>
Default value is fizz.
- word2 : the word that will replace every number that can be divided by int2<br>
Default value is buzz.
- limit : a number that set the size of the fizzbuzz sequence<br>
Default value is 15

### Curl Examples
```
curl -X GET http://localhost:8080/fizzbuzz | jq

{
  "result": [
    "1",
    "2",
    "fizz",
    "4",
    "buzz",
    "fizz",
    "7",
    "8",
    "fizz",
    "buzz",
    "11",
    "fizz",
    "13",
    "14",
    "fizzbuzz"
  ]
}
```
```
curl -X GET "http://localhost:8080/fizzbuzz?int1=4&word1=Hello&int2=6&word2=World&limit=30" | jq

{
  "result": [
    "1",
    "2",
    "3",
    "Hello",
    "5",
    "World",
    "7",
    "Hello",
    "9",
    "10",
    "11",
    "HelloWorld",
    "13",
    "14",
    "15",
    "Hello",
    "17",
    "World",
    "19",
    "Hello",
    "21",
    "22",
    "23",
    "HelloWorld",
    "25",
    "26",
    "27",
    "Hello",
    "29",
    "World"
  ]
}
```
## 3 - GET /stats
This end point shows statistics per FizzBuzz query. For each Query the statistics available are the number of hits and the date of the last hit. This endpoint has one parameter :
- limit : the number of FizzBuzz queries to show. If not set all queries that have been requested since the server is up are shown with their statistics.

### Curl example
```
curl -X GET http://localhost:8080/stats?limit=4 | jq

[
  {
    "Query": "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]",
    "Hits": 24,
    "LastHit": "2020-03-01T18:52:01.491385755+01:00"
  },
  {
    "Query": "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]",
    "Hits": 20,
    "LastHit": "2020-03-01T18:52:01.48812416+01:00"
  },
  {
    "Query": "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]",
    "Hits": 14,
    "LastHit": "2020-03-01T18:52:01.494784265+01:00"
  },
  {
    "Query": "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]",
    "Hits": 10,
    "LastHit": "2020-03-01T18:52:01.492829939+01:00"
  }
]

```