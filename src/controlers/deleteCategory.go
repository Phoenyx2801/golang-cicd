package controlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	client := &http.Client{}

	apiUrl := fmt.Sprintf("https://localhost:5001/api/Category/categories/%s", id)
	req, err := http.NewRequest("DELETE", apiUrl, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	c.String(resp.StatusCode, resp.Status)

}
