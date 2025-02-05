# Callisto

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/forbole/callisto/Tests)](https://github.com/forbole/callisto/actions?query=workflow%3ATests)
[![Go Report Card](https://goreportcard.com/badge/github.com/forbole/callisto)](https://goreportcard.com/report/github.com/forbole/callisto)
![Codecov branch](https://img.shields.io/codecov/c/github/forbole/callisto/cosmos/v0.40.x)

Callisto (shorthand for BigDipper Juno) is the [Juno](https://github.com/forbole/juno) implementation
for [BigDipper](https://github.com/forbole/big-dipper).

It extends the custom Juno behavior by adding different handlers and custom operations to make it easier for BigDipper
showing the data inside the UI.

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).

## Usage

To know how to setup and run Callisto, please refer to
the [docs website](https://docs.bigdipper.live/cosmos-based/parser/overview/).

## Testing

If you want to test the code, you can do so by running

```shell
$ make test
```

**Note**: Requires [Docker](https://docker.com).

This will:

1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.

## Local launch in docker compose

1. Create `.callisto` directory in the root of the repo and put there next two files:
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
make images
docker-compose up
```

* Open [hasura UI](http://localhost:8080/console) and check that it works correctly.
  The password is defined in docker-compose.yaml and set to "myadminsecretkey" by default.

### Remarks

In case you run the callisto with the connection to old-running node you might face the error in the logs

```
error while getting staking pool: rpc error: code = Internal desc = UnmarshalJSON cannot decode empty bytes
```

This is expected since the node doesn't store all staking pool for all heights.

## Custom params integration

In case you want to integrate the indexing if the custom types, for example parameters,
that [PR](https://github.com/CoreumFoundation/callisto/pull/4)
can be taken as a reference implementation.

## Hasura API

### GraphQL Schema

The GraphQL schema is located [here](./hasura/api/schema.graphql). It describes all supported queries.

In order to export it you can run the hasura locally and execute script

```
npm install -g graphqurl # install the gq
gq http://localhost:8080/v1/graphql -H "X-Hasura-Admin-Secret: myadminsecretkey" --introspect > schema.graphql  # export schema
```

Pay attention that `myadminsecretkey` is the secret set for this repo local environment, and can be different for
others.
