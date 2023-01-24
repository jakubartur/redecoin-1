// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

var RedeMainnetBootnodes = []string{
        "enode://aa6c289248b0e62a2407b256b3b9e0c94de4039bbbfe0b67cefe93dd9121a399c1948a897caa12b0d26cf6340d15c81c5ebdee0118337d1cefe177c39d0fd36d@37.187.131.105:30303",
        "enode://a01e9da1848e7325f2d4a6aae87865397c48e0570f51aef7ed2ffd17c77b4582ebdfffcf14733e882018f6597b283f126a9448f55dd0e636056e7dad11304674@51.38.133.225:30303",
        "enode://2d0d3c09323a8e89f91349a44f8461c3665ae884040e0fa283357bad96a210d52f67a8f1db50f1b12f7fbef6936112001799f8d09c22f76701a3b380229c6a74@54.38.53.146:30303",
        "enode://df381bea9841164b32c0d3edc1091c96b47a55fc2734b6bad1760d7c1fdbdbcf7aed6ad16d867253ef6346e08f96b6599de38dbe8959d3cc153ab8e9a776f8ff@51.38.65.97:30303",
        "enode://195b66dff27bcb32f38380bbabc1d4a2be4ddfc6ae732ce5d08a2d26dd0a14dac705f178a4ecc237a0dab9698d284a2a7a78bf7d41b30f834772759cbde1985e@139.99.238.22:30303",
        "enode://f0d54c61017afe51daf899391b46934b308477387c75d2a718e4edb6314e32437a6402d64b9426b5362e4c8230dcc8ac19bd309ea5fcfface00a9a9ea2f49f14@15.235.192.151:30303",
        "enode://4931b8d5bbeaf39e0a42236ded7ef04005b1c8eb8170cbfcec2405887b8825ab22d8c64551b34e1afa66913a52144be9186564127ab9302ce4f5d74c2bac1879@144.217.13.169:30303",
        "enode://f76b156425116f9bf8f37947015e3be849fd182e1939489745c05c9b4032c959fb4168f16ae2cefc6df37927d38061448a3680e0e461bb9d162b1b9e3e4183fd@146.59.95.189:30303",
        "enode://12c3276083eebc96a44003ce98236cd0bacdf80ec3854d888d3c616d9f79fbfd681775f3ee1bde901a72dc18acd0c4d1997586db2ea1d1608a6cbd590489a0ce@188.165.213.69:30533",
        "enode://17fc1cd87a6a8471653f46834e652f3ca6ecf07d6ad4531849cf641db7cea4b6f30d055a4301eb5f6c6276855615d8e18f67b9da179e6440bae1d695fe6a54fb@51.83.187.212:30303",
}
var MainnetBootnodes = []string{
	 "enode://aa6c289248b0e62a2407b256b3b9e0c94de4039bbbfe0b67cefe93dd9121a399c1948a897caa12b0d26cf6340d15c81c5ebdee0118337d1cefe177c39d0fd36d@37.187.131.105:30303",
        "enode://a01e9da1848e7325f2d4a6aae87865397c48e0570f51aef7ed2ffd17c77b4582ebdfffcf14733e882018f6597b283f126a9448f55dd0e636056e7dad11304674@51.38.133.225:30303",
        "enode://2d0d3c09323a8e89f91349a44f8461c3665ae884040e0fa283357bad96a210d52f67a8f1db50f1b12f7fbef6936112001799f8d09c22f76701a3b380229c6a74@54.38.53.146:30303",
        "enode://df381bea9841164b32c0d3edc1091c96b47a55fc2734b6bad1760d7c1fdbdbcf7aed6ad16d867253ef6346e08f96b6599de38dbe8959d3cc153ab8e9a776f8ff@51.38.65.97:30303",
        "enode://195b66dff27bcb32f38380bbabc1d4a2be4ddfc6ae732ce5d08a2d26dd0a14dac705f178a4ecc237a0dab9698d284a2a7a78bf7d41b30f834772759cbde1985e@139.99.238.22:30303",
        "enode://f0d54c61017afe51daf899391b46934b308477387c75d2a718e4edb6314e32437a6402d64b9426b5362e4c8230dcc8ac19bd309ea5fcfface00a9a9ea2f49f14@15.235.192.151:30303",
        "enode://4931b8d5bbeaf39e0a42236ded7ef04005b1c8eb8170cbfcec2405887b8825ab22d8c64551b34e1afa66913a52144be9186564127ab9302ce4f5d74c2bac1879@144.217.13.169:30303",
        "enode://f76b156425116f9bf8f37947015e3be849fd182e1939489745c05c9b4032c959fb4168f16ae2cefc6df37927d38061448a3680e0e461bb9d162b1b9e3e4183fd@146.59.95.189:30303",
        "enode://12c3276083eebc96a44003ce98236cd0bacdf80ec3854d888d3c616d9f79fbfd681775f3ee1bde901a72dc18acd0c4d1997586db2ea1d1608a6cbd590489a0ce@188.165.213.69:30533",
        "enode://17fc1cd87a6a8471653f46834e652f3ca6ecf07d6ad4531849cf641db7cea4b6f30d055a4301eb5f6c6276855615d8e18f67b9da179e6440bae1d695fe6a54fb@51.83.187.212:30303",
}
var RopstenBootnodes = []string{
}
var SepoliaBootnodes = []string{
}
var RinkebyBootnodes = []string{
}
var GoerliBootnodes = []string{
}
var KilnBootnodes = []string{
}
var V5Bootnodes = []string{
        "enr:-KO4QDUqgovT8zm6TujmteUNRV0UavPFdH0zmXtZWtfJCR7zW15ZGIQ1hOHUiqv9MKN8YsgFV5oF5NvJcbD0M3yKaFeGAYUV_V9Mg2V0aMfGhMs-gbKAgmlkgnY0gmlwhCW7g2mJc2VjcDI1NmsxoQKpz9duoDA3eBZO2HZmtLIaxEITLiCcAF63I1yUI2CbloRzbmFwwIN0Y3CCdl-DdWRwgnZf",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
//        case RedeGenesisHash:
//                net = "rede"
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
