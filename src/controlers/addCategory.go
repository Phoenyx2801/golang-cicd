package controlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCategory(c *gin.Context) {

	var requestBody interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Make the POST request to the desired URL
	apiURL := "https://localhost:5001/api/Category/categories"
	fmt.Println(requestBody)
	if err := ForwardRequest(apiURL, requestBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func ForwardRequest(apiURL string, user interface{}) error {
	// Convert the User object to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Send the POST request to the other server
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return fmt.Errorf("API request failed")
	}

	return nil
}
