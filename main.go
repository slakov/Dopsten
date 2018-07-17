package main

import (
	//"fmt"
	"log"
	"net/http"
	"math/big"

	"./donate"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/balance", balance)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func balance(c echo.Context) error {

	conn, err := ethclient.Dial("https://ropsten.infura.io/3fkcvSlU6y3nB3FH6fux")

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractInstance, err := donate.NewDonate(common.HexToAddress("0xfa7ca0caef9b29c7b38fc85358fd12690a2e52a4"), conn)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	res, _ := contractInstance.AmountRaised(nil)
	x := big.NewInt(1)
	times, _ := contractInstance.TimeArray(nil, x)
	//res2 := times[len(times)-1]
	
	return c.String(http.StatusOK, "I have: "+res.String()+" "+times.String())

}
