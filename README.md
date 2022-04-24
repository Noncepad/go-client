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

## Test Connection


```bash
curl -X POST -vv \
     -H "Content-Type: application/json" \
     -d '{"id":1,"method":"Arith.Multiply","params":[{"A":1,"B":2}]}' \
     --url http://localhost:8080/jsonrpc
```

## Send Transactions

```bash
curl -X POST -vv \
     -H "Content-Type: application/json" \
     -d '{"id":1,"method":"Basic.SendTx","params":[{"commitment":"processed","tx":["...serialized tx 1","...serialized tx 2",...]}]}' \
     --url http://localhost:8080/jsonrpc
```

The 3 options for `commitment` are `processed`,`confirmed`,and `finalized`.  The transactions must be serialized in base64.


The response will be a job id.  The send_tx action takes over 30 seconds.

```json
{
     "id":334
}
```

## Get SendTx


```bash
curl -X POST -vv \
     -H "Content-Type: application/json" \
     -d '{"id":1,"method":"Basic.GetSendTxResult","params":[{"id":334}]}' \
     --url http://localhost:8080/jsonrpc
```

The job id is `334` in this example.

The response will be:

```json
{
    "signature":["...sig 1","... sig 2"]
}
```

Signatures are formatted in base64.