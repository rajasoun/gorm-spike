#!/usr/bin/env bash

# Define constants for formatting
NC=$'\e[0m' # No Color
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'
YELLOW='\033[1;33m'
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'

#Function : Run unit tests and generate coverage report
function run_unit_tests() {
    echo "${BLUE}Running Unit Tests${NC}"
    go test -v ./... --coverprofile reports/outfile
    go tool cover -html=reports/outfile -o reports/cover.html
}


#Function: Open coverage report in default browser
function open_coverage_report() {
    echo "${BLUE}Opening coverage report${NC}"
    open reports/cover.html
}

#Function: main 
function main() {
    run_unit_tests
    open_coverage_report
}

main $@

