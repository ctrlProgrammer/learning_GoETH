package main

import (
	LETH "learningGO/ETHcontracts/src"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	client, err := ethclient.Dial("https://polygon-mainnet.g.alchemy.com/v2/IfBbeQVlWjzJGdhHsO_q3y7TvCgAg9aM")

	if err != nil {
		log.Fatal(err)
	}

	LETH.DefineRouters(router, client)

	router.Run("localhost:9090")
}
