#!/bin/bash

cp ./networks/testnet-1/genesis.json ~/.superstructd/config/ && superstructd genesis gentx validator2 800000000stake --chain-id hyper-1
