package controlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetProductByCategory(c *gin.Context) {

	categoryName := c.Param("categoryName")

	apiUrl := fmt.Sprintf("https://localhost:5001/api/Category/categories/search?categoryName=%s", categoryName)
	resp, err := http.Get(apiUrl)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(data)

	c.JSON(http.StatusOK, data)
}
