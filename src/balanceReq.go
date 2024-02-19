package LETH

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type BalanceRequest struct {
	Address     string `json:"address"`
	Balance     string `json:"balance"`
	LastRequest int    `json:"lastRequest"`
}

var balanceRequests map[string]BalanceRequest
var MAX_TIME_CONFIG int = 1000 * 60 * 1

func ETHGetAddressBalance(address string, client *ethclient.Client) (*BalanceRequest, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		return nil, errors.New("not found")
	}

	return &BalanceRequest{
		Address:     address,
		Balance:     ParseETH(*balance),
		LastRequest: int(time.Now().UnixMilli()),
	}, nil
}

func GetAddressBalanceAPI(context *gin.Context, client *ethclient.Client) {
	accountAddress := context.Param("address")
	cachedRequest, ok := balanceRequests[accountAddress]

	if ok {
		if cachedRequest.LastRequest+MAX_TIME_CONFIG >= int(time.Now().UnixMilli()) {
			context.IndentedJSON(http.StatusOK, cachedRequest)
			return
		}
	}

	request, err := ETHGetAddressBalance(accountAddress, client)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": true})
		return
	}

	balanceRequests[accountAddress] = *request

	context.IndentedJSON(http.StatusOK, *request)
}

func DefineRouters(router *gin.Engine, client *ethclient.Client) {
	balanceRequests = make(map[string]BalanceRequest)

	router.GET("/balance/:address", func(context *gin.Context) {
		GetAddressBalanceAPI(context, client)
	})
}
