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

## Endpoint
There is one endpoint : GET /fizzbuzz
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

## Curl Exemples
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
