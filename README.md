# MicroServices using Nats.io and Golang (Nats request/reply)

## setup repository
```
docker-compose -f build.yml up
```
## Making Curl Call: Add two numbers
```curl
curl -X POST http://127.0.0.1:8080/sum -H "Content-Type: application/json" -H "Accept: application/json" -d "{\"a\":10,\"b\":5}"
```


## Making Curl Call: Subtract (a - b)
```curl
curl -X POST http://127.0.0.1:8080/substract -H "Content-Type: application/json" -H "Accept: application/json" -d "{\"a\":10,\"b\":5}"
```

## Making Curl Call: Multiply (a * b)
```curl
curl -X POST http://127.0.0.1:8080/multiply -H "Content-Type: application/json" -H "Accept: application/json" -d "{\"a\":10,\"b\":5}"
```