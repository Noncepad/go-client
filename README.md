# Noncepad Client

This repository compiles a Golang client that lets users access the gRPC endpoints using a local RPC proxy.



# Installation


```bash
git clone https://github.com/Noncepad/go-client
cd go-client
make build
```

The RPC server binary will be at `./bin/rpc-server` .

# Usage

Set the following environmental variables:

```bash
set -a
export API_KEY=...put your API KEY here...
./bin/rpc-server
```

## Test RPC Connection


```bash
curl -X POST -vv \
     -H "Content-Type: application/json" \
     -d '{"id":1,"method":"Arith.Multiply","params":[{"A":1,"B":2}]}' \
     --url http://localhost:8080/jsonrpc
```