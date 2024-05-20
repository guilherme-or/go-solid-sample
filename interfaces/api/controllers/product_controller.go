package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-or/go-solid-sample/app/usecases"
	"github.com/guilherme-or/go-solid-sample/domain/entities"
)

type ProductController struct {
	ProductUseCases usecases.ProductUseCases
}

func (c *ProductController) RegisterHandlers(ctx *gin.Engine) {
	ctx.GET("/products", c.GetAllProducts)
	ctx.GET("/products/:id", c.GetProductById)
	ctx.POST("/products", c.CreateProduct)
	ctx.PUT("/products", c.UpdateProduct)
	ctx.DELETE("/products/:id", c.DeleteProduct)
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	ctx.JSON(200, c.ProductUseCases.GetAllProducts())
}

func (c *ProductController) GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	id64 := int64(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	product, err := c.ProductUseCases.GetProductById(id64)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": fmt.Sprintf("Product with id %v not found", id64),
		})
		return
	}

	ctx.JSON(200, product)
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product entities.Product
	ctx.BindJSON(&product)

	newProduct, err := c.ProductUseCases.SaveProduct(product)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	ctx.JSON(201, newProduct)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var product entities.Product
	ctx.BindJSON(&product)

	err := c.ProductUseCases.AlterProduct(product)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	id64 := int64(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	err = c.ProductUseCases.RemoveProductById(id64)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": fmt.Sprintf("Product with id %v not found", id64),
		})
		return
	}

	ctx.Status(204)
}
