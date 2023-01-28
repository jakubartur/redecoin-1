
#!/bin/bash
clear
echo "node rede start"
sleep 1
wget https://raw.githubusercontent.com/redecoinproject/redecoin/master/genesis/genesis.json
wget https://github.com/redecoinproject/redecoin/releases/download/v1.0.1/geth
chmod 755 geth
./geth init genesis.json
./geth --syncmode full console
