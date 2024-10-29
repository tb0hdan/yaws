#!/bin/bash

curl -H 'Content-type: application/json' -X POST -d '{"status":"completed"}' http://localhost:8080/api/v1/orders/bfcf445e-bbcb-4c3b-aea7-57b1e7e2ca9f/status
