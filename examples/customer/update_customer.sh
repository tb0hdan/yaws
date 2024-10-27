#!/bin/bash


curl -H 'Content-type: application/json' -X PUT -d '{"name": "John Doe", "email": "john.doe@gmail.com", "phone": "12345", "address": "123 Elm street, Los Angeles, CA, USA"}' http://localhost:8080/api/v1/customers/7