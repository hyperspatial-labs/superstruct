#!/bin/bash

seeds="e3fedf4a977aaa5c750306ec0af78e84a38f67ea@172.27.3.8:26656,ee3d2da82d829a34ba7baf3582d561644b063790@172.27.3.138:26656"
peers="e3fedf4a977aaa5c750306ec0af78e84a38f67ea@172.27.3.8:26656,ee3d2da82d829a34ba7baf3582d561644b063790@172.27.3.138:26656"
sed -i.bak -e "s/^seeds *=.*/seeds = \"$seeds\"/; s/^persistent_peers *=.*/persistent_peers = \"$peers\"/" ~/.superstructd/config/config.toml