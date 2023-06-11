package controlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {

	resp, err := http.Get("https://localhost:5001/api/Category/categories")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, data)
}
