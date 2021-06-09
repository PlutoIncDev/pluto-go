# Pluto Go
Microservice Framework written for Golang. 

## Useful Commands
```bash
# Clean Go code
gofmt -s -w .
```

```bash
# apache bench test the API!
ab -n 10000 -c 100 -m POST localhost:8080/auth/login

# wrk: https://github.com/wg/wrk
```

```bash
# proto
protoc -I=./protos --go_out=./pb ./protos/test.proto
```

TODO: ADD BETTER LOGGING FOR STRUCTS SO THAT THEYRE PREFIXED WITH THE NAME OF THE PROVIDER!
