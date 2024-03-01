#!/bin/bash

cp genesis.json ~/.superstructd/config/ && superstructd gentx validator2 800000000stake --chain-id testnet-1
