#!/bin/bash

curl -H 'Content-type: application/json' -X POST -d '[{"id":"bc805d99-105a-4218-b921-c5aa0a8def77", "name": "Kettle, 2kW", "price": "101.95", "quantity": 10}]' http://localhost:8080/api/v1/products