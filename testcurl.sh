#!/bin/bash

# POST Receipt
curl -d '{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}' http://localhost:8080/receipts/process

curl -d '{
             "retailer": "Target",
             "purchaseDate": "2022-01-02",
             "purchaseTime": "13:13",
             "total": "1.25",
             "items": [
                 {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
             ]
}' http://localhost:8080/receipts/process