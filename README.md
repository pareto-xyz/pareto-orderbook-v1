# Pareto Order Book v1

**[Disclaimer: This repository is no longer maintained and is meant for primarily educational purposes.]**

Part of the series detailed in this [whitepaper](https://github.com/pareto-xyz/pareto-order-book-whitepaper/blob/main/how_to_orderbook.pdf). 

Implementation of an off-chain order book and a price-time matching algorithm in Go. The order book supports market and limit orders, and attempts to do many of the core operations of adding, editing, and canceling orders in O(1). Orders can hold options or expiring futures contracts, though not both.

## Usage

The API and server are in `main.go`, and include endpoints for creating new orders. Authentication is done via an API key, which can be created using an endpoint. API keys should be saved and never shared. A set of special API keys are reserved for admin functions such as pausing the orderbook. 

## Setup

This orderbook is dependent on Pareto's smart contracts. Clone [this repo](https://github.com/pareto-xyz/pareto-core-v1) and run the following command in its root directory while running a hardhat node (to do this, run `npx harhat node` in a separate terminal):
```
npx hardhat run ./scripts/deploy.localhost.ts --network localhost
```
This will return an output containing the following:
```
...
Deployed ETH margin contract:  0x610178dA211FEF7D417bC0e6FeD39F05609AD788
```

Set environment variable `MARGIN_CONTRACT` to the proxy contract address. Make sure this is not the underlying contract address, but rather the proxy on top. It is important that this global variable is up-to-date.

We will also need to compile the margin contract with `go-ethereum`. We wrote a helper script to do that:
```
python scripts/compile.py <PATH TO pareto-core-v1 REPO> ./contract/
```
which will generate all the ABI and binary files, as well as Go functions for all the smart contract external functions in `./contract/margin.go`. You will need to keep this file up-to-date and regenerate bindings each time the smart contract is updated.

## Commands

### Run Tests
```
go test ./auth
go test ./controller
go test ./orderbook
go test ./shared
```

### Run Linting
```
go lint ./auth
go lint ./controller
go lint ./orderbook
go lint ./shared
```

### Run Static Checks
```
staticcheck ./auth
staticcheck ./controller
staticcheck ./orderbook
staticcheck ./shared
```

### Notes
Boilerplate based on [i25959341/orderbook](https://github.com/i25959341/orderbook).
