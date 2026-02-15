package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const NFT_CONTRACT_ADDRESS = "0x816070929010A3D202D8A6B89f92BeE33B7e8769"

// NFTVerificationResponse represents the response structure
type NFTVerificationResponse struct {
	Address         string `json:"address"`
	ContractAddress string `json:"contract_address"`
	HasNFT          bool   `json:"has_nft"`
	Status          string `json:"status"` // "YES" or "NO"
}

// VerifyNFTOwnership checks if an address holds the specified NFT
func VerifyNFTOwnership(c *gin.Context) {
	address := c.Param("address")

	if address == "" {
		c.JSON(400, gin.H{"error": "address parameter is required"})
		return
	}

	// Call Blockscout API
	url := fmt.Sprintf(
		"https://blockscout-api.injective.network/api?module=account&action=tokennfttx&address=%s&contractaddress=%s&page=1&offset=1&sort=desc",
		address,
		NFT_CONTRACT_ADDRESS,
	)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to query NFT ownership", "details": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(500, gin.H{"error": "failed to parse response", "details": err.Error()})
		return
	}

	// Check if result array has any entries
	hasNFT := false
	status := "NO"

	if resultArray, ok := result["result"].([]interface{}); ok {
		if len(resultArray) > 0 {
			hasNFT = true
			status = "YES"
		}
	}

	response := NFTVerificationResponse{
		Address:         address,
		ContractAddress: NFT_CONTRACT_ADDRESS,
		HasNFT:          hasNFT,
		Status:          status,
	}

	c.JSON(200, response)
}

// BatchVerifyNFTOwnership checks multiple addresses at once
func BatchVerifyNFTOwnership(c *gin.Context) {
	var req struct {
		Addresses []string `json:"addresses" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if len(req.Addresses) > 50 {
		c.JSON(400, gin.H{"error": "maximum 50 addresses allowed per batch"})
		return
	}

	results := make([]NFTVerificationResponse, 0, len(req.Addresses))

	for _, address := range req.Addresses {
		url := fmt.Sprintf(
			"https://blockscout-api.injective.network/api?module=account&action=tokennfttx&address=%s&contractaddress=%s&page=1&offset=1&sort=desc",
			address,
			NFT_CONTRACT_ADDRESS,
		)

		resp, err := http.Get(url)
		if err != nil {
			results = append(results, NFTVerificationResponse{
				Address:         address,
				ContractAddress: NFT_CONTRACT_ADDRESS,
				HasNFT:          false,
				Status:          "ERROR",
			})
			continue
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		resp.Body.Close()

		hasNFT := false
		status := "NO"

		if resultArray, ok := result["result"].([]interface{}); ok {
			if len(resultArray) > 0 {
				hasNFT = true
				status = "YES"
			}
		}

		results = append(results, NFTVerificationResponse{
			Address:         address,
			ContractAddress: NFT_CONTRACT_ADDRESS,
			HasNFT:          hasNFT,
			Status:          status,
		})
	}

	c.JSON(200, gin.H{
		"results": results,
		"count":   len(results),
	})
}
