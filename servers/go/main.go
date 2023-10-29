package main

import (
	"fmt"
	"net/http"
	"os"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	maestro "github.com/maestro-org/go-sdk/client"
)

type APIError struct {
	Error string `json:"error" example:"Bad Request"` // Error returned from the API request
}

type LockReqBody struct {
	Sender                   string `json:"sender"`
	Beneficiary              string `json:"beneficiary"`
	AssetPolicyId            string `json:"asset_policy_id"`
	AssetTokenName           string `json:"asset_token_name"`
	TotalVestingQuantity     int64  `json:"total_vesting_quantity"`
	VestingPeriodStart       int64  `json:"vesting_period_start"`
	VestingPeriodEnd         int64  `json:"vesting_period_end"`
	FirstUnlockPossibleAfter int64  `json:"first_unlock_possible_after"`
	TotalInstallments        int64  `json:"total_installments"`
}

type LockRespBody struct {
	CborHex string `json:"cbor_hex"`
	TxHash  string `json:"tx_hash"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maestroClient := maestro.NewClient(os.Getenv("MAESTRO_API_KEY"), os.Getenv("CARDANO_NETWORK"))

	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Maestro!",
		})
	})

	router.POST("/vesting/lock", func(c *gin.Context) {

		var lockReqBody LockReqBody
		if err := c.BindJSON(&lockReqBody); err != nil {
			c.JSON(http.StatusBadRequest, APIError{Error: fmt.Sprintf("Invalid request body: %s", err.Error())})
			return
		}

		lockTx, err := maestroClient.LockAssets(maestro.LockBody(lockReqBody))
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIError{Error: fmt.Sprintf("Cannot lock assets: %s", err.Error())})
			return
		}

		lockRespBody := LockRespBody(*lockTx)
		c.JSON(http.StatusAccepted, lockRespBody)
	})

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
