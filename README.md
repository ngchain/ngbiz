<h1> <img src="./resources/ng_16x16.png" >biz</h1>

## What is ngchain?

The ngchain is a radically updating **brand-new blockchain network**, which is not a fork of ethereum or any other
chain.

The ngchain's **goal** is to build **a blockchain engine** which acts more **auditable, scalable, security-oriented**
and supports more network tasks with vm.

The ngchain uses modern models - Vault(Block), Account, Multi-type Tx, and the traditional Block model to build the
blockchain ecosystem. And ngchain strictly follow the idea of blockchain, locking items with hash to keep engine work
safely. Not only blockchain, but Vault(Block) will also link up to be a chain for account security and powerful
functions like fast ignition, free account state and so on. So It's safe and robust like bitcoin and ethereum but more
scalable and powerful on the node's operational capacity and p2p network's speed and performance.

## Status

[![Go Report Card](https://goreportcard.com/badge/github.com/ngchain/ngbiz)](
https://goreportcard.com/report/github.com/ngchain/ngbiz)
![CI](https://github.com/ngchain/ngbiz/workflows/CI/badge.svg)
![GitHub](https://img.shields.io/github/license/ngchain/ngbiz)
![GitHub last commit](https://img.shields.io/github/last-commit/ngchain/ngbiz)

## Features

- **Fast ignition**
- Less, or **no storage cost**(mem only)
- With **humanizing** account model, users can send tx with **memorable short number**
- **High security** with Sheet and Vault(Block) model
- Powerful and scalable types of tx
- Support **Multi-Tx**, sending coins to different places in the same time
- Powerful **WASM** VM support based on account state(contract).
- **Libp2p(ipfs)** powered p2p networking
- Available **anonymous** address for saving balance
- Using the **schnorr signature**, allowing Multi-Sig when sending and receiving
- ...

## Requirements

go version >= 1.17

**NOTICE**: go build on Windows you should use `-buildmode=exe` flag (go version >= 1.15)

## Build

### Go

```bash
# go will automatically sync the dependencies
# GCC is required because of high performance db & vm
go build ./cmd/ngbiz
```

## Usage

```bash
# dircetly run the binary
export GOLOG_FILE=ngbiz.log # disable stderr output and write to the ngbiz.log file
export GOLOG_LOG_LEVEL=debug # print more logs
./ngbiz

# ngwallet is a rpc client in dart for ngin's daemon, see https://github.com/ngchain/ngwallet-dart
./ngwallet register 10086
./ngwallet transact 10010 1.5 # send 1.5 NG to account 10010
./ngwallet transact QfUnsE4CNgnpVS4oC4WEYH8u7WWAs8AwMrFBknWWqGSYwBXU 1.5 # send 1.5 NG to address QfUn...
```

If you wanna start mining(proof of work), try `--mining <Thread Num>` flag

```bash
./ngbiz --mining 0 # zero means using all available cores
```

You can view more flags and options with `--help` flag

```bash
./ngbiz --help
```

Or you can choose to run in a docker

```bash
git clone https://github.com/ngchain/ngbiz && cd ngbiz
sudo docker build . -t ngbiz

# Run as a bootstrap node
sudo docker run -p 52520:52520 -p 52521:52521 -v .:/workspace -v ~/.ngkeys:~/.ngkeys ngbiz --bootstrap true

# Run as a mining node, 0 means using all cpu cores, --in-mem will disable writing into disk and make the miner lighter
sudo docker run -p 52520:52520 -p 52521:52521 -v .:/workspace -v ~/.ngkeys:~/.ngkeys ngbiz --mining 0 --in-mem
```

## Run a ngchain forknet

It's so easy to run an independent PoW chain on ngbiz codebase.

1. Modify the `GenesisAddressBase58` in `./ngtypes/defaults.go` and `protocolVersion` in `./ngp2p/defaults/defaults.go`

2. Generate a new signature for genesis generate tx, and genesis block's nonce (with `ngbiz gentools` toolset)

3. Run more than 2 bootstrap node with `--bootstrap` flag (without mining)

4. Write the bootstrap node to bootstrapNodes in `./ngp2p/bootstrap_nodes.go`

5. Run a mining node with `--mining 0` flag

6. Enjoy your fascinating PoW chain
