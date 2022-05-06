# Observer Service

The observer service provides a read-only point of contact to interact with the Flow network. It implements parts of the [Access API](https://docs.onflow.org/access-api/). It only implements a subset. Users, who need to change the state of the network should opt for running an Access node.

It is a GRPC server which connects to a staked access node via libp2p, and forwards execute requests via GRPC. It is a scalable service.

At a high level it does the following:

1. Forwards all read-only Script related calls (`ExecuteScriptAtLatestBlock`, `ExecuteScriptAtBlockID` and `ExecuteScriptAtBlockHeight`) to one of the access nodes that forwards it to one of the execution services.
2. Follows updates to the blockchain and locally caches transactions, collections, and sealed blocks.
3. Replies to client API calls for information such as `GetBlockByID`, `GetCollectionByID`, `GetTransaction` etc.

***NOTE**: The Observer service does not participate in the Flow protocol*

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

## Table of Contents

- [Getting started](#getting-started)
- [Documentation](#documentation)
- [Installation](#installation)
    - [Clone Repository](#clone-repository)
    - [Install Dependencies](#install-dependencies)
- [Development Workflow](#development-workflow)
    - [Building](#building)

    - [Local Network](#local-network)
    - [Stopping](#stop)
    - [Testing](#testing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Getting started

- To install all dependencies and tools, see the [project setup](#installation) guide
- To dig into more documentation about Flow, see the [documentation](#documentation)
- To learn how to contribute, see the [contributing guide](/CONTRIBUTING.md)
- To see information on developing Flow, see the [development workflow](#development-workflow)

## Documentation

You can find an overview of the Flow architecture on the [documentation website](https://www.onflow.org/primer).

Development on Flow is divided into work streams. Each work stream has a home directory containing high-level
documentation for the stream, as well as links to documentation for relevant components used by that work stream.

The following table lists all work streams and links to their home directory and documentation:

| Work Stream      | Home directory                             |
|------------------|--------------------------------------------|
| Access Node      | [/cmd/access](/cmd/access)                 |
| Collection Node  | [/cmd/collection](/cmd/collection)         |
| Consensus Node   | [/cmd/consensus](/cmd/consensus)           |
| Execution Node   | [/cmd/execution](/cmd/execution)           |
| Verification Node | [/cmd/verification](/cmd/verification)     |
| Observer Service | [/cmd/observer](/cmd/observer)             |
| HotStuff         | [/consensus/hotstuff](/consensus/hotstuff) |
| Storage          | [/storage](/storage)                       |
| Ledger           | [/ledger](/ledger)                         |
| Networking       | [/network](/network/)                      |
| Cryptography     | [/crypto](/crypto)                         |

## Installation

### Clone Repository

- Clone this repository
- Clone this repository's submodules:

    ```bash
    git submodule update --init --recursive
    ```

### Install Dependencies

- Install [Go](https://golang.org/doc/install) (Flow supports Go 1.16 and later)
- Install [CMake](https://cmake.org/install/), which is used for building the crypto library
- Install [Docker](https://docs.docker.com/get-docker/), which is used for running a local network and integration tests
- Make sure the [`GOPATH`](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable) and `GOBIN` environment variables
  are set, and `GOBIN` is added to your path:

    ```bash
    export GOPATH=$(go env GOPATH)
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:$GOBIN
    ```

  Add these to your shell profile to persist them for future runs.
- Then, run the following command in dep:

    ```bash
    make install-tools
    ```

Note: if there is error about "relic" or "crypto", trying force removing the relic build and reinstall the tools again:

```bash
rm -rf crypto/relic
make install-tools
```

## Development Workflow

### Build Project Into Image

```bash
make install
```

### Build Service

The recommended way to build and run Observer Service for local development is using Docker.
Start the service in a docker container:

```bash
make start
```

### Stop the container

```bash
make stop
```

### Testing

Run the unit test suite:

```bash
make test
```