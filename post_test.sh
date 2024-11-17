#!/bin/bash

# Define the URL and file path
URL="http://localhost:8080/json"
FILE_PATH="internal/util/vendor.json"

# Send the POST request
curl -X POST -H "Content-Type: application/json" -d @$FILE_PATH $URL
