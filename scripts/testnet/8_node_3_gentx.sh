#!/bin/bash

cp genesis.json ~/.superstructd/config/ && superstructd gentx validator2 600000000stake --chain-id testnet-1

