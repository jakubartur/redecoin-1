
#!/bin/bash
clear
echo "node rede start"
echo "cmd start"
echo "curl https://raw.githubusercontent.com/redecoinproject/redecoin/master/node.sh | bash"
sleep 2
wget https://raw.githubusercontent.com/redecoinproject/redecoin/master/genesis/genesis.json
wget https://github.com/redecoinproject/redecoin/releases/download/v1.0.1/geth
chmod 755 geth
./geth init genesis.json
wget https://raw.githubusercontent.com/redecoinproject/redecoin/master/start.sh
chmod 755 start.sh
./start.sh
