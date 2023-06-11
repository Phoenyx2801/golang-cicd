package controlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context) {

	id := c.Param("id")

	var requestBody interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Make the POST request to the desired URL
	apiURL := fmt.Sprintf("https://localhost:5001/api/Category/categories/%s/products", id)
	fmt.Println(requestBody)
	if err := ForwardRequest(apiURL, requestBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "New product added")

}
