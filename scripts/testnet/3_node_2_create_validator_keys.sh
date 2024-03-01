#!/bin/bash

# on node/validator 2 (awspp)
cp ./networks/testnet/genesis.json ~/.superstructd/config/ && superstructd keys add validator2
