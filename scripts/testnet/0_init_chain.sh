#!/bin/bash

superstructd init superstruct-1 --chain-id hyper-1 -o

cp ~/.superstructd/config/genesis.json ./networks/testnet-1/genesis.json
