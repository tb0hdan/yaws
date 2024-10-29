#!/bin/bash

curl -H 'Content-type: application/json' -H 'X-Customer-id: 123' -X POST -d '{"order_id":"bfcf445e-bbcb-4c3b-aea7-57b1e7e2ca9f", "payment_status":"paid"}' http://localhost:8080/api/v1/payment-webhook