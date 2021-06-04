// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2021-06-04 12:45:41.451348 -0700 PDT m=&#43;0.002943946
// using data from coins.yml
package coin

import (
	"fmt"
)

// Coin is the native currency of a blockchain
type Coin struct {
	ID               uint
	Handle           string
	Symbol           string
	Name             string
	Decimals         uint
	BlockTime        int
	MinConfirmations int64
	SampleAddr       string
}

func (c *Coin) String() string {
	return fmt.Sprintf("[%s] %s (#%d)", c.Symbol, c.Name, c.ID)
}

const (
	ETHEREUM = 60
	CLASSIC = 61
	ICON = 74
	COSMOS = 118
	RIPPLE = 144
	STELLAR = 148
	POA = 178
	TRON = 195
	FIO = 235
	NIMIQ = 242
	IOTEX = 304
	ZILLIQA = 313
	AION = 425
	AETERNITY = 457
	KAVA = 459
	THETA = 500
	BINANCE = 714
	VECHAIN = 818
	CALLISTO = 820
	TOMOCHAIN = 889
	THUNDERTOKEN = 1001
	ONTOLOGY = 1024
	TEZOS = 1729
	KIN = 2017
	NEBULAS = 2718
	GOCHAIN = 6060
	WANCHAIN = 5718350
	WAVES = 5741564
	BITCOIN = 0
	LITECOIN = 2
	DOGE = 3
	DASH = 5
	VIACOIN = 14
	GROESTLCOIN = 17
	ZCASH = 133
	ZCOIN = 136
	BITCOINCASH = 145
	RAVENCOIN = 175
	QTUM = 2301
	ZELCASH = 19167
	DECRED = 42
	ALGORAND = 283
	NANO = 165
	DIGIBYTE = 20
	HARMONY = 1023
	KUSAMA = 434
	POLKADOT = 354
	SOLANA = 501
	NEAR = 397
	ELROND = 508
	SMARTCHAIN = 20000714
	FILECOIN = 461
	OASIS = 474
	MONACOIN = 22
	BITCOINGOLD = 156
	EOS = 194
	TERRA = 330
	BAND = 494
	NEO = 888
	CARDANO = 1815
	NULS = 8964
	POLYGON = 966
	THORCHAIN = 931
	MINT = 1281
)

var Coins = map[uint]Coin{
	ETHEREUM: {
		ID:               60,
		Handle:           "ethereum",
		Symbol:           "ETH",
		Name:             "Ethereum",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1",
	},
	CLASSIC: {
		ID:               61,
		Handle:           "classic",
		Symbol:           "ETC",
		Name:             "Ethereum Classic",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "0xf3524415b6D873205B4c3Cda783527b2aC4daAA9",
	},
	ICON: {
		ID:               74,
		Handle:           "icon",
		Symbol:           "ICX",
		Name:             "ICON",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "hxee691e7bccc4eb11fee922896e9f51490e62b12e",
	},
	COSMOS: {
		ID:               118,
		Handle:           "cosmos",
		Symbol:           "ATOM",
		Name:             "Cosmos",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
	},
	RIPPLE: {
		ID:               144,
		Handle:           "ripple",
		Symbol:           "XRP",
		Name:             "Ripple",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
	},
	STELLAR: {
		ID:               148,
		Handle:           "stellar",
		Symbol:           "XLM",
		Name:             "Stellar",
		Decimals:         7,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
	},
	POA: {
		ID:               178,
		Handle:           "poa",
		Symbol:           "POA",
		Name:             "Poa",
		Decimals:         18,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "0x1fddEc96688e0538A316C64dcFd211c491ECf0d8",
	},
	TRON: {
		ID:               195,
		Handle:           "tron",
		Symbol:           "TRX",
		Name:             "Tron",
		Decimals:         6,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	},
	FIO: {
		ID:               235,
		Handle:           "fio",
		Symbol:           "FIO",
		Name:             "FIO",
		Decimals:         9,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "FIO5J2xdfWygeNdHZNZRzRws8YGbVxjUXtp4eP8KoGkGKoLFQ7CaU",
	},
	NIMIQ: {
		ID:               242,
		Handle:           "nimiq",
		Symbol:           "NIM",
		Name:             "Nimiq",
		Decimals:         5,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	},
	IOTEX: {
		ID:               304,
		Handle:           "iotex",
		Symbol:           "IOTX",
		Name:             "IoTeX",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
	},
	ZILLIQA: {
		ID:               313,
		Handle:           "zilliqa",
		Symbol:           "ZIL",
		Name:             "Zilliqa",
		Decimals:         12,
		BlockTime:        30000,
		MinConfirmations: 1,
		SampleAddr:       "zil1anrjcsj2ntklaa3arq4w3s6gw4l4hqrycs9egy",
	},
	AION: {
		ID:               425,
		Handle:           "aion",
		Symbol:           "AION",
		Name:             "Aion",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0xa07981da70ce919e1db5f051c3c386eb526e6ce8b9e2bfd56e3f3d754b0a17f3",
	},
	AETERNITY: {
		ID:               457,
		Handle:           "aeternity",
		Symbol:           "AE",
		Name:             "Aeternity",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
		SampleAddr:       "ak_2p5878zbFhxnrm7meL7TmqwtvBaqcBddyp5eGzZbovZ5FeVfcw",
	},
	KAVA: {
		ID:               459,
		Handle:           "kava",
		Symbol:           "KAVA",
		Name:             "Kava",
		Decimals:         6,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "kava13fxkk4730cqglgdv7w0mdelyx07myyq7h2nd3x",
	},
	THETA: {
		ID:               500,
		Handle:           "theta",
		Symbol:           "THETA",
		Name:             "Theta",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
	},
	BINANCE: {
		ID:               714,
		Handle:           "binance",
		Symbol:           "BNB",
		Name:             "BNB",
		Decimals:         8,
		BlockTime:        1000,
		MinConfirmations: 2,
		SampleAddr:       "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	},
	VECHAIN: {
		ID:               818,
		Handle:           "vechain",
		Symbol:           "VET",
		Name:             "VeChain Token",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "0xB5e883349e68aB59307d1604555AC890fAC47128",
	},
	CALLISTO: {
		ID:               820,
		Handle:           "callisto",
		Symbol:           "CLO",
		Name:             "Callisto",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0x39ec1c88a7a7c1a575e8c8f42eff7630d9278179",
	},
	TOMOCHAIN: {
		ID:               889,
		Handle:           "tomochain",
		Symbol:           "TOMO",
		Name:             "TOMO",
		Decimals:         18,
		BlockTime:        4000,
		MinConfirmations: 0,
		SampleAddr:       "0x7daa83030e3086477b79b6e757ca8608899fe783",
	},
	THUNDERTOKEN: {
		ID:               1001,
		Handle:           "thundertoken",
		Symbol:           "TT",
		Name:             "ThunderCore",
		Decimals:         18,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "0x0ad80a408eac4f17ba0a9de8a12d8736f60700c3",
	},
	ONTOLOGY: {
		ID:               1024,
		Handle:           "ontology",
		Symbol:           "ONT",
		Name:             "Ontology",
		Decimals:         0,
		BlockTime:        10000,
		MinConfirmations: 0,
		SampleAddr:       "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	},
	TEZOS: {
		ID:               1729,
		Handle:           "tezos",
		Symbol:           "XTZ",
		Name:             "Tezos",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
	},
	KIN: {
		ID:               2017,
		Handle:           "kin",
		Symbol:           "KIN",
		Name:             "Kin",
		Decimals:         5,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "GBHKUZ7C2SZ5N3X2S7O6TT6LNUWSEA2BXMSR5GTTSR6VZARSVAXIQNGH",
	},
	NEBULAS: {
		ID:               2718,
		Handle:           "nebulas",
		Symbol:           "NAS",
		Name:             "Nebulas",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a",
	},
	GOCHAIN: {
		ID:               6060,
		Handle:           "gochain",
		Symbol:           "GO",
		Name:             "GoChain GO",
		Decimals:         18,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "0x76c2F81716A8D198a00502Ae9a59126418899FDe",
	},
	WANCHAIN: {
		ID:               5718350,
		Handle:           "wanchain",
		Symbol:           "WAN",
		Name:             "Wanchain",
		Decimals:         18,
		BlockTime:        30000,
		MinConfirmations: 0,
		SampleAddr:       "0x36cEdc3A9d969306AF4F7CA2b83ABBf74095914d",
	},
	WAVES: {
		ID:               5741564,
		Handle:           "waves",
		Symbol:           "WAVES",
		Name:             "WAVES",
		Decimals:         8,
		BlockTime:        30000,
		MinConfirmations: 1,
		SampleAddr:       "3P7wz6TXienpw3BHe8eHUEuZWb6WE58kgnQ",
	},
	BITCOIN: {
		ID:               0,
		Handle:           "bitcoin",
		Symbol:           "BTC",
		Name:             "Bitcoin",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
		SampleAddr:       "bc1quvuarfksewfeuevuc6tn0kfyptgjvwsvrprk9d",
	},
	LITECOIN: {
		ID:               2,
		Handle:           "litecoin",
		Symbol:           "LTC",
		Name:             "Litecoin",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
		SampleAddr:       "ltc1qhd8fxxp2dx3vsmpac43z6ev0kllm4n53t5sk0u",
	},
	DOGE: {
		ID:               3,
		Handle:           "doge",
		Symbol:           "DOGE",
		Name:             "Dogecoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "DJRFZNg8jkUtjcpo2zJd92FUAzwRjitw6f",
	},
	DASH: {
		ID:               5,
		Handle:           "dash",
		Symbol:           "DASH",
		Name:             "Dash",
		Decimals:         8,
		BlockTime:        180000,
		MinConfirmations: 0,
		SampleAddr:       "XqHiz8EXYbTAtBEYs4pWTHh7ipEDQcNQeT",
	},
	VIACOIN: {
		ID:               14,
		Handle:           "viacoin",
		Symbol:           "VIA",
		Name:             "Viacoin",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
		SampleAddr:       "via1qnmsgjd6cvfprnszdgmyg9kewtjfgqflz67wwhc",
	},
	GROESTLCOIN: {
		ID:               17,
		Handle:           "groestlcoin",
		Symbol:           "GRS",
		Name:             "Groestlcoin",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "grs1qexwmshts5pdpeqglkl39zyl6693tmfwp0cue4j",
	},
	ZCASH: {
		ID:               133,
		Handle:           "zcash",
		Symbol:           "ZEC",
		Name:             "Zcash",
		Decimals:         8,
		BlockTime:        150000,
		MinConfirmations: 0,
		SampleAddr:       "t1YYnByMzdGhQv3W3rnjHMrJs6HH4Y231gy",
	},
	ZCOIN: {
		ID:               136,
		Handle:           "zcoin",
		Symbol:           "FIRO",
		Name:             "Firo",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
		SampleAddr:       "aEd5XFChyXobvEics2ppAqgK3Bgusjxtik",
	},
	BITCOINCASH: {
		ID:               145,
		Handle:           "bitcoincash",
		Symbol:           "BCH",
		Name:             "Bitcoin Cash",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
		SampleAddr:       "bitcoincash:qq07l6rr5lsdm3m80qxw80ku2ex0tj76vvsxpvmgme",
	},
	RAVENCOIN: {
		ID:               175,
		Handle:           "ravencoin",
		Symbol:           "RVN",
		Name:             "Raven",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "RGkwvrUors8DtmhKy5bddFwRCTZaunjpvo",
	},
	QTUM: {
		ID:               2301,
		Handle:           "qtum",
		Symbol:           "QTUM",
		Name:             "Qtum",
		Decimals:         8,
		BlockTime:        60000,
		MinConfirmations: 0,
		SampleAddr:       "QhceuaTdeCZtcxmVc6yyEDEJ7Riu5gWFoF",
	},
	ZELCASH: {
		ID:               19167,
		Handle:           "zelcash",
		Symbol:           "ZEL",
		Name:             "Zelcash",
		Decimals:         8,
		BlockTime:        120000,
		MinConfirmations: 0,
		SampleAddr:       "t1UKbRPzL4WN8Rs8aZ8RNiWoD2ftCMHKGUf",
	},
	DECRED: {
		ID:               42,
		Handle:           "decred",
		Symbol:           "DCR",
		Name:             "Decred",
		Decimals:         8,
		BlockTime:        300000,
		MinConfirmations: 0,
		SampleAddr:       "DsTxPUVFxXeNgu5fzozr4mTR4tqqMaKcvpY",
	},
	ALGORAND: {
		ID:               283,
		Handle:           "algorand",
		Symbol:           "ALGO",
		Name:             "Algorand",
		Decimals:         6,
		BlockTime:        20000,
		MinConfirmations: 0,
		SampleAddr:       "4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U",
	},
	NANO: {
		ID:               165,
		Handle:           "nano",
		Symbol:           "NANO",
		Name:             "Nano",
		Decimals:         30,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "nano_1trqphog5noig7z888asnjejcie8z1iopxyepcjdo1atps8whxiuwd51ehbw",
	},
	DIGIBYTE: {
		ID:               20,
		Handle:           "digibyte",
		Symbol:           "DGB",
		Name:             "DigiByte",
		Decimals:         8,
		BlockTime:        15000,
		MinConfirmations: 0,
		SampleAddr:       "D8NBg12kfW8uLjzCv7LYnPYCNhqvVtHaMQ",
	},
	HARMONY: {
		ID:               1023,
		Handle:           "harmony",
		Symbol:           "ONE",
		Name:             "Harmony",
		Decimals:         18,
		BlockTime:        5000,
		MinConfirmations: 0,
		SampleAddr:       "one1syjs6cnfwd9fgrhng03dyzs07suwtywwreczmk",
	},
	KUSAMA: {
		ID:               434,
		Handle:           "kusama",
		Symbol:           "KSM",
		Name:             "Kusama",
		Decimals:         12,
		BlockTime:        6000,
		MinConfirmations: 0,
		SampleAddr:       "HewiDTQv92L2bVtkziZC8ASxrFUxr6ajQ62RXAnwQ8FDVmg",
	},
	POLKADOT: {
		ID:               354,
		Handle:           "polkadot",
		Symbol:           "DOT",
		Name:             "Polkadot",
		Decimals:         10,
		BlockTime:        6000,
		MinConfirmations: 0,
		SampleAddr:       "13SkL2uACPqBzpKBh3d2n5msYNFB2QapA5vEDeKeLjG2LS3Y",
	},
	SOLANA: {
		ID:               501,
		Handle:           "solana",
		Symbol:           "SOL",
		Name:             "Solana",
		Decimals:         9,
		BlockTime:        500,
		MinConfirmations: 0,
		SampleAddr:       "boot1Z6jb15CLqpaMTn2CxktktwZpRAVAgHZEW6SxQ7",
	},
	NEAR: {
		ID:               397,
		Handle:           "near",
		Symbol:           "NEAR",
		Name:             "NEAR",
		Decimals:         18,
		BlockTime:        2000,
		MinConfirmations: 0,
		SampleAddr:       "NEAR6Y66fCzeKqWiwxoPox5oGeDN9VhNCu7CEQ9M86iniqoN9vg2X",
	},
	ELROND: {
		ID:               508,
		Handle:           "elrond",
		Symbol:           "eGLD",
		Name:             "Elrond",
		Decimals:         18,
		BlockTime:        6000,
		MinConfirmations: 0,
		SampleAddr:       "erd12tqtt5zcg6vpw65y4hkanvt49kzq695sr3ctuszjy92xw0ppzcssy2xd5r",
	},
	SMARTCHAIN: {
		ID:               20000714,
		Handle:           "smartchain",
		Symbol:           "BNB",
		Name:             "Smart Chain",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
		SampleAddr:       "0x35552c16704d214347f29Fa77f77DA6d75d7C752",
	},
	FILECOIN: {
		ID:               461,
		Handle:           "filecoin",
		Symbol:           "FIL",
		Name:             "Filecoin",
		Decimals:         18,
		BlockTime:        3000,
		MinConfirmations: 0,
		SampleAddr:       "t1nbb73vhk5dtmnsgeaetbo76daepqjtrfoccn74i",
	},
	OASIS: {
		ID:               474,
		Handle:           "oasis",
		Symbol:           "ROSE",
		Name:             "Oasis",
		Decimals:         9,
		BlockTime:        6000,
		MinConfirmations: 0,
		SampleAddr:       "oasis1qpxx577luv69znyr7dleue43mz6aq7fh4sm43cxu",
	},
	MONACOIN: {
		ID:               22,
		Handle:           "monacoin",
		Symbol:           "MONA",
		Name:             "Monacoin",
		Decimals:         8,
		BlockTime:        90000,
		MinConfirmations: 0,
		SampleAddr:       "PSEbmFMLvjUg6vKoZBj3zPk1KVbVca5Uun",
	},
	BITCOINGOLD: {
		ID:               156,
		Handle:           "bitcoingold",
		Symbol:           "BTG",
		Name:             "Bitcoin Gold",
		Decimals:         8,
		BlockTime:        600000,
		MinConfirmations: 0,
		SampleAddr:       "GR9ixMSWCPSo7yEPvNuobEWAdMKzY35LzQ",
	},
	EOS: {
		ID:               194,
		Handle:           "eos",
		Symbol:           "EOS",
		Name:             "EOS",
		Decimals:         4,
		BlockTime:        500,
		MinConfirmations: 0,
		SampleAddr:       "",
	},
	TERRA: {
		ID:               330,
		Handle:           "terra",
		Symbol:           "LUNA",
		Name:             "Terra",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "terra1fex9f78reuwhfsnc8sun6mz8rl9zwqh03fhwf3",
	},
	BAND: {
		ID:               494,
		Handle:           "band",
		Symbol:           "BAND",
		Name:             "BandChain",
		Decimals:         6,
		BlockTime:        2000,
		MinConfirmations: 0,
		SampleAddr:       "band12nmsm9khdsv0tywge43q3zwj8kkj3hvup9rltp",
	},
	NEO: {
		ID:               888,
		Handle:           "neo",
		Symbol:           "NEO",
		Name:             "NEO",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "AcxuqWhTureEQGeJgbmtSWNAtssjMLU7pb",
	},
	CARDANO: {
		ID:               1815,
		Handle:           "cardano",
		Symbol:           "ADA",
		Name:             "Cardano",
		Decimals:         6,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "addr1s3xuxwfetyfe7q9u3rfn6je9stlvcgmj8rezd87qjjegdtxm3y3f2mgtn87mrny9r77gm09h6ecslh3gmarrvrp9n4yzmdnecfxyu59jz29g8j",
	},
	NULS: {
		ID:               8964,
		Handle:           "nuls",
		Symbol:           "NULS",
		Name:             "NULS",
		Decimals:         8,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "",
	},
	POLYGON: {
		ID:               966,
		Handle:           "polygon",
		Symbol:           "MATIC",
		Name:             "Polygon",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "0x720E1fa107A1Df39Db4E78A3633121ac36Bec132",
	},
	THORCHAIN: {
		ID:               931,
		Handle:           "thorchain",
		Symbol:           "RUNE",
		Name:             "THORChain",
		Decimals:         18,
		BlockTime:        0,
		MinConfirmations: 0,
		SampleAddr:       "bnb1czhptncu9uhhlrlrk9q0rmy459gae2ua79mapn",
	},
	MINT: {
		ID:               1281,
		Handle:           "mint",
		Symbol:           "MNT",
		Name:             "Mint",
		Decimals:         18,
		BlockTime:        1000,
		MinConfirmations: 0,
		SampleAddr:       "0x720E1fa107A1Df39Db4E78A3633121ac36Bec132",
	},
}
func Ethereum() Coin {
	return Coins[ETHEREUM]
}
func Classic() Coin {
	return Coins[CLASSIC]
}
func Icon() Coin {
	return Coins[ICON]
}
func Cosmos() Coin {
	return Coins[COSMOS]
}
func Ripple() Coin {
	return Coins[RIPPLE]
}
func Stellar() Coin {
	return Coins[STELLAR]
}
func Poa() Coin {
	return Coins[POA]
}
func Tron() Coin {
	return Coins[TRON]
}
func Fio() Coin {
	return Coins[FIO]
}
func Nimiq() Coin {
	return Coins[NIMIQ]
}
func Iotex() Coin {
	return Coins[IOTEX]
}
func Zilliqa() Coin {
	return Coins[ZILLIQA]
}
func Aion() Coin {
	return Coins[AION]
}
func Aeternity() Coin {
	return Coins[AETERNITY]
}
func Kava() Coin {
	return Coins[KAVA]
}
func Theta() Coin {
	return Coins[THETA]
}
func Binance() Coin {
	return Coins[BINANCE]
}
func Vechain() Coin {
	return Coins[VECHAIN]
}
func Callisto() Coin {
	return Coins[CALLISTO]
}
func Tomochain() Coin {
	return Coins[TOMOCHAIN]
}
func Thundertoken() Coin {
	return Coins[THUNDERTOKEN]
}
func Ontology() Coin {
	return Coins[ONTOLOGY]
}
func Tezos() Coin {
	return Coins[TEZOS]
}
func Kin() Coin {
	return Coins[KIN]
}
func Nebulas() Coin {
	return Coins[NEBULAS]
}
func Gochain() Coin {
	return Coins[GOCHAIN]
}
func Wanchain() Coin {
	return Coins[WANCHAIN]
}
func Waves() Coin {
	return Coins[WAVES]
}
func Bitcoin() Coin {
	return Coins[BITCOIN]
}
func Litecoin() Coin {
	return Coins[LITECOIN]
}
func Doge() Coin {
	return Coins[DOGE]
}
func Dash() Coin {
	return Coins[DASH]
}
func Viacoin() Coin {
	return Coins[VIACOIN]
}
func Groestlcoin() Coin {
	return Coins[GROESTLCOIN]
}
func Zcash() Coin {
	return Coins[ZCASH]
}
func Zcoin() Coin {
	return Coins[ZCOIN]
}
func Bitcoincash() Coin {
	return Coins[BITCOINCASH]
}
func Ravencoin() Coin {
	return Coins[RAVENCOIN]
}
func Qtum() Coin {
	return Coins[QTUM]
}
func Zelcash() Coin {
	return Coins[ZELCASH]
}
func Decred() Coin {
	return Coins[DECRED]
}
func Algorand() Coin {
	return Coins[ALGORAND]
}
func Nano() Coin {
	return Coins[NANO]
}
func Digibyte() Coin {
	return Coins[DIGIBYTE]
}
func Harmony() Coin {
	return Coins[HARMONY]
}
func Kusama() Coin {
	return Coins[KUSAMA]
}
func Polkadot() Coin {
	return Coins[POLKADOT]
}
func Solana() Coin {
	return Coins[SOLANA]
}
func Near() Coin {
	return Coins[NEAR]
}
func Elrond() Coin {
	return Coins[ELROND]
}
func Smartchain() Coin {
	return Coins[SMARTCHAIN]
}
func Filecoin() Coin {
	return Coins[FILECOIN]
}
func Oasis() Coin {
	return Coins[OASIS]
}
func Monacoin() Coin {
	return Coins[MONACOIN]
}
func Bitcoingold() Coin {
	return Coins[BITCOINGOLD]
}
func Eos() Coin {
	return Coins[EOS]
}
func Terra() Coin {
	return Coins[TERRA]
}
func Band() Coin {
	return Coins[BAND]
}
func Neo() Coin {
	return Coins[NEO]
}
func Cardano() Coin {
	return Coins[CARDANO]
}
func Nuls() Coin {
	return Coins[NULS]
}
func Polygon() Coin {
	return Coins[POLYGON]
}
func Thorchain() Coin {
	return Coins[THORCHAIN]
}
func Mint() Coin {
	return Coins[MINT]
}

