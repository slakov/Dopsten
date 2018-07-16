package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./donate"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

	contractInstance, err := donate.NewDonate(common.HexToAddress("0x08eb6ecdcd5ad8b19c9af94e78383200af2f3442"), conn)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	res, _ := contractInstance.AmountRaised(nil)
	fmt.Println(res)

	return c.String(http.StatusOK, "I have: "+res.String())

}
