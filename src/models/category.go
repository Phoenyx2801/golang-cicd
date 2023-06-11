package models

type Category struct {
	id         int    `json:"id"`
	name       string `json:"name"`
	itemsCount int    `json:"itemsCount"`
}

type CategoryPost struct {
	name string `json:"title"`
}

type Product struct {
	id         int    `json:"id"`
	title      string `json:"title"`
	price      int    `json:"price"`
	categoryId int    `json:"categoryId"`
}
