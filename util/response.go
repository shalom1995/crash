package util

import (
	"FlashLoan/msgs"
	"github.com/gin-gonic/gin"
	"math/big"
)

type Context struct {
	C *gin.Context
}

func (c *Context) Response(httpCode int, code int, data interface{}) {
	c.C.Header("Access-Control-Allow-Origin", "*")
	c.C.Header("Access-Control-Allow-Methods", "*")
	c.C.Header("Access-Control-Allow-Headers", "*")

	c.C.JSON(httpCode, gin.H{
		"Code": code,
		"Msg":  msgs.MsgReturn[code],
		"Data": data,
	})
	return
}

type ERC721 struct {
	ID           uint   `json:"id"`
	ContractAddr string `json:"contract_addr"`
	TokenID      string `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	TokenURI     string `json:"token_uri"`
	ContractName string `json:"contract_name"`
	ProviderLink string `json:"provider_link"`
}

type ERC1155 struct {
	ID           uint   `json:"id"`
	ContractAddr string `json:"contract_addr"`
	TokenID      string `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	Balance      string `json:"balance"`
	TokenURI     string `json:"token_uri"`
	ContractName string `json:"contract_name"`
	ProviderLink string `json:"provider_link"`
}

type NFT struct {
	Id           uint   `json:"id"`
	ContractAddr string `json:"contract_addr"`
	ContractName string `json:"contract_name"`
	TokenType    string `json:"token_type"`
	TokenID      string `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	Balance      string `json:"balance"`
	TokenURI     string `json:"token_uri"`
	ProviderLink string `json:"provider_link"`

	Name        interface{} `json:"name"`
	Description interface{} `json:"description"`
	Image       interface{} `json:"image"`
	Metadata    interface{} `json:"metadata"`
}

type Contract struct {
	ContractAddr string `json:"contract_addr"`
	TokenID      string `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	Balance      string `json:"balance"`
}
type Contracts []*Contract

func (s Contracts) Len() int { return len(s) }

func (s Contracts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Contracts) Less(i, j int) bool {
	bi := big.NewInt(0)
	bi.SetString(s[i].TokenID, 0)
	bj := big.NewInt(0)
	bj.SetString(s[j].TokenID, 0)

	f := bi.Cmp(bj)
	return f == -1
}
