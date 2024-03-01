
# val 1
superstructd genesis add-genesis-account xpz1c6ch3zqxkkcsqs0craj0mclvqxkfc6ua30xvte 1000000000stake,1000000000uxpz

# val 2
superstructd genesis add-genesis-account xpz1t3nqafmt8zhuvx5zeetnujpknsp42rzkr5jajp 1000000000stake,2000000000uxpz

# val 3
superstructd genesis add-genesis-account xpz13lmqfydemnu0x2y8xqwr6jehw9ct4rumumze8g 1000000000stake,3000000000uxpz


# superstructd tx distribution withdraw-rewards xpz1c6ch3zqxkkcsqs0craj0mclvqxkfc6ua30xvte --from validator1 --chain-id hyper-1

# superstructd tx bank send xpz13lmqfydemnu0x2y8xqwr6jehw9ct4rumumze8g xpz1c6ch3zqxkkcsqs0craj0mclvqxkfc6ua30xvte 100xpz --chain-id hyper-1 --fees 2000000uxpz --gas 200000