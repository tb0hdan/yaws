#!/bin/bash

curl -vsk -H 'Content-type: application/json' -X POST -d '{"id":"bfcf445e-bbcb-4c3b-aea7-57b1e7e2ca9f", "customer_id": 9, "products":[{"id":"bc805d99-105a-4218-b921-c5aa0a8def77", "quantity": 3}]}' http://localhost:8080/api/v1/orders