# Very basic gRPC -> HTTP gateway

An example of usage:
```
curl --location --request POST 'http://127.0.0.1:8081/v1/example/echo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "value": "henlo"
}'
```
