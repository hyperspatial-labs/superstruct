
# val 1
superstructd genesis add-genesis-account xpz1mk5ldkmh2089fhhau6xe026pz703n3vrl3t9m8 1000000000stake,1000000000uxpz

# val 2
superstructd genesis add-genesis-account xpz1uyx2vwn5jfyg299nz2rcgdgax92mcza5wz64tq 1000000000stake,2000000000uxpz

# val 3
superstructd genesis add-genesis-account xpz1n8zadvwxyv2u3l59xdnq48tdjf2htk6afrfff2 1000000000stake,3000000000uxpz
                                                                                    

# superstructd tx distribution withdraw-rewards xpz1c6ch3zqxkkcsqs0craj0mclvqxkfc6ua30xvte --from validator1 --chain-id hyper-1

# superstructd tx bank send xpz13lmqfydemnu0x2y8xqwr6jehw9ct4rumumze8g xpz1c6ch3zqxkkcsqs0craj0mclvqxkfc6ua30xvte 100xpz --chain-id hyper-1 --fees 2000000uxpz --gas 200000