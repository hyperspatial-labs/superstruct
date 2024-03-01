#!/bin/bash
superstructd genesis collect-gentxs

cp ~/.superstructd/config/genesis.json ./networks/testnet-1/genesis.json
