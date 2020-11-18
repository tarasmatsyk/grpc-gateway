# Very basic gRPC -> HTTP gateway

An example of usage:
```
curl --location --request POST 'http://127.0.0.1:8081/v1/example/echo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "value": "henlo"
}'
```

# Points to improve

- Run proto generation on all internal proto files and ignore third_party; At the moment every proto has to be specified in the build cmd
- Figure out how middleware behaves between Swagger/gRPC
