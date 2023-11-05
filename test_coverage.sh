#!/bin/bash

COVERAGE_THRESHOLD="80"
COVERAGE_OUTPUT=$(go test ./... -covermode=atomic 2>&1)

# Initialize a flag to check if any coverage data was found
COVERAGE_FOUND=false

# Iterate over each line of the output
while read -r line; do
    CURRENT_COVERAGE=$(echo "$line" | grep -oE 'coverage: [0-9.]+%' | awk -F'[ %]' '{print $2}')
    PACKAGE=$(echo "$line" | grep -oE 'admins/[a-zA-Z0-9_/]+' | awk -F'/' '{print $2}')
    if [ -n "$CURRENT_COVERAGE" ]; then
        COVERAGE_FOUND=true
        if (( $(awk -v current="$CURRENT_COVERAGE" -v threshold="$COVERAGE_THRESHOLD" 'BEGIN { print (current < threshold) }') )); then

            echo "Code coverage is less than $COVERAGE_THRESHOLD% for the package: $PACKAGE"
            # Set red color for the output
            tput setaf 1
            echo "Coverage: $CURRENT_COVERAGE"
            exit 1
        else
            # Set green color for the output
            tput setaf 2
            echo "Code coverage is greater than or equal to $COVERAGE_THRESHOLD% for the package: $PACKAGE"
            echo "Coverage: $CURRENT_COVERAGE"
        fi
    fi
done <<< "$COVERAGE_OUTPUT"

if [ "$COVERAGE_FOUND" = false ]; then
    echo "No coverage data found in the test output."
    exit 1
fi
