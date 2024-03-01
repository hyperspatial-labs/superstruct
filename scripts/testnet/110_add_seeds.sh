#!/bin/bash

seeds="da3f178494c92a5a0265fe249e3722fd298f9be0@172.27.3.8:26656,bf413f020ab2b39fc7b3403ffc0814f0a4de2494@172.27.3.138:26656"
peers="da3f178494c92a5a0265fe249e3722fd298f9be0@172.27.3.8:26656,bf413f020ab2b39fc7b3403ffc0814f0a4de2494@172.27.3.138:26656"
sed -i.bak -e "s/^seeds *=.*/seeds = \"$seeds\"/; s/^persistent_peers *=.*/persistent_peers = \"$peers\"/" ~/.superstructd/config/config.toml