package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "Lab_4_client/src/controlers"

func main() {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	r := gin.Default()
	//r.Use(gin.Logger())

	r.GET("/categories", controlers.GetAll)
	r.POST("/add/categories", controlers.AddCategory)
	r.GET("/categories/:id", controlers.GetByIb)
	r.DELETE("/delete/:id", controlers.DeleteCategory)
	r.GET("/product/:id", controlers.GetProductById)
	r.GET("/product/category/:categoryName", controlers.GetProductByCategory)
	r.POST("/add/products/:id", controlers.AddProduct)
	r.PUT("/edit/category/:id", controlers.PutCategory)
	//r.Use(gin.Recovery())

	r.Run(":8080")
}
