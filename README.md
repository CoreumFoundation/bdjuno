# BDJuno

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/forbole/bdjuno/Tests)](https://github.com/forbole/bdjuno/actions?query=workflow%3ATests)
[![Go Report Card](https://goreportcard.com/badge/github.com/forbole/bdjuno)](https://goreportcard.com/report/github.com/forbole/bdjuno)
![Codecov branch](https://img.shields.io/codecov/c/github/forbole/bdjuno/cosmos/v0.40.x)

BDJuno (shorthand for BigDipper Juno) is the [Juno](https://github.com/forbole/juno) implementation
for [BigDipper](https://github.com/forbole/big-dipper).

It extends the custom Juno behavior by adding different handlers and custom operations to make it easier for BigDipper
showing the data inside the UI.

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).

## Usage

To know how to setup and run BDJuno, please refer to
the [docs website](https://docs.bigdipper.live/cosmos-based/parser/overview/).

## Testing

If you want to test the code, you can do so by running

```shell
$ make test-unit
```

**Note**: Requires [Docker](https://docker.com).

This will:

1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.

## Local launch in docker compose

1. Create `.bdjuno` directory in the root of the repo and put there next two files:
    1. `config.yaml` which you can copy from `config-sample.yaml`. Do next change there:
        1. Replace `YOUR_CHAIN_ACC_PREFIX` with your account prefix(`coredev` is an example for devnet, you can use it
           too). This will let you have separate accounts with the same private keys for different chains.
        2. Replace `YOUR_NODE_IP` with `cored` address.
    2. `genesis.json` which you can find at cored `/genesis` endpoint(`http://127.0.0.1:26557/genesis?` for example).
       Take nested object:
       ```json
       {
           "genesis": {this object}
       }

```
   
2. Build and start
```bash
docker-compose build && docker-compose up
```

* Open [hasura UI](http://localhost:8080/console) and check that it works correctly.
  The password is defined in docker-compose.yaml and set to "myadminsecretkey" by default.

### Remarks

In case you run the bdjuno with the connection to old-running node you might face the error in the logs

```
error while getting staking pool: rpc error: code = Internal desc = UnmarshalJSON cannot decode empty bytes
```

This is expected since the node doesn't store all staking pool for all heights.

## Custom params integration

In case you want to integrate the indexing if the custom types, for example parameters,
that [PR](https://github.com/CoreumFoundation/bdjuno/pull/4)
can be taken as a reference implementation.

## Hasura API

### Endpoints

The primary hasura API is GraphQL API. All GraphQL Query endpoints are public and available for the usage.

| Chain               | Endpoint                                       |           
|---------------------|------------------------------------------------| 
| **testnet**         | https://hasura.testnet-1.coreum.dev/v1/graphql | 
| **devnet**          | https://hasura.devnet-1.coreum.dev/v1/graphql  |   
| **znet (localnet)** | http://localhost:8080/v1/graphql               |  

### GraphQL Schema

The GraphQL schema is located [here](./hasura/api/schema.graphql). It describes all supported queries.

In order to export it you can run the hasura locally and execute script

```
npm install -g graphqurl # install the gq
gq http://localhost:8080/v1/graphql -H "X-Hasura-Admin-Secret: myadminsecretkey" --introspect > schema.graphql  # export schema
```

Pay attention that `myadminsecretkey` is the secret set for this repo local environment, and can be different for other.

### GraphQL Playground

You can install run the [playground](https://github.com/graphql/graphql-playground) locally to test the queries.
Use [api](./hasura/api) as data source. Pay attention that [.graphqlconfig](./hasura/api/.graphqlconfig) file contains
settings for the `znet(localnet)`, so if
you want to use different you need to update it.

### Examples

#### Get transactions by address

Pay attention that example is for the `znet(localnet)`, so if you want to use different you need the address and
endpoint for the corresponding chain.

* Example for the playground

```graphql
{
    messages_by_address(args: {addresses: "{devcore1gafptgc4vmemklgzxg0d9acmj7p2lv55w7nfxj}", limit: "50", offset: "0", types: "{}"}) {
        value
        type
        transaction_hash
    }
}
```

* Raw request example

```bash
curl --location 'http://localhost:8080/v1/graphql' \
--header 'Content-Type: text/plain' \
--data '{"operationName":"GetMessagesByAddress","variables":{"limit":50,"offset":0,"types":"{}","address":"{devcore1gafptgc4vmemklgzxg0d9acmj7p2lv55w7nfxj}"},"query":"query GetMessagesByAddress($address: _text, $limit: bigint = 50, $offset: bigint = 0, $types: _text = \"{}\") {\n  messagesByAddress: messages_by_address(\n    args: {addresses: $address, types: $types, limit: $limit, offset: $offset}\n  ) {\n    transaction {\n      height\n      hash\n      success\n      messages\n      logs\n      block {\n        height\n        timestamp\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"}'
```

