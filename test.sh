#!/bin/bash

curl -i -X POST -H 'Content-Type: application/json' -d '{ "id": 0, "toCaps" : "nasa" }' http://localhost:8080
