#!/bin/bash

# Loop to send 10 POST requests with incremented id
for ((id=1; id<=10; id++)); do
    curl --location -k 'https://echo.local.com/customers' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id": '$id',
        "firstName": "echo",
        "lastName": "local.com",
        "age": 1,
        "email": "echo@local.com"
    }'
done
