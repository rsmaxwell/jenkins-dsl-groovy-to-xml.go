# jenkins-job-to-xml

## Description

Reads Jenkins job definition in groovy from stdin, and outputs to stdout the equivilent XML Jenkins job definition



## Build

generate static resources

`go generate ./...`


Build and install the program

`go install ./...`


## Usage

`cat test.groovy | job-to-xml`

