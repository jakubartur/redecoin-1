#!/usr/bin/env bash
clear
echo "******************************"
echo "* install node rede ethash    *"
echo "******************************"
sleep 3
rm -rf .rede
rm -f geth
rm -f genesis.json
wget https://raw.githubusercontent.com/redecoinproject/redecoin/master/genesis/genesis.json
wget https://github.com/redecoinproject/redecoin/releases/download/v1.0.1/geth
chmod 755 geth
./geth init genesis.json
./geth --syncmode full console

