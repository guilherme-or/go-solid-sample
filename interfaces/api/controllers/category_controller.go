package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-or/go-solid-sample/app/usecases"
	"github.com/guilherme-or/go-solid-sample/domain/entities"
)

const categoryPattern = "/categories"

type CategoryController struct {
	CategoryUseCases usecases.CategoryUseCases
}

func (c *CategoryController) RegisterHandlers(ctx *gin.Engine) {
	ctx.GET(categoryPattern, c.GetAllCategories)
	ctx.GET(fmt.Sprintf("%v/:id", categoryPattern), c.GetCategoryById)
	ctx.POST(categoryPattern, c.CreateCategory)
	ctx.PUT(categoryPattern, c.UpdateCategory)
	ctx.DELETE(fmt.Sprintf("%v/:id", categoryPattern), c.DeleteCategory)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	ctx.JSON(200, c.CategoryUseCases.GetAllCategories())
}

func (c *CategoryController) GetCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	id64 := int64(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	category, err := c.CategoryUseCases.GetCategoryById(id64)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": fmt.Sprintf("Category with id %v not found", id64),
		})
	}

	ctx.JSON(200, category)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category entities.Category
	ctx.BindJSON(&category)

	newCategory, err := c.CategoryUseCases.SaveCategory(category)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	ctx.JSON(201, newCategory)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var category entities.Category
	ctx.BindJSON(&category)

	err := c.CategoryUseCases.AlterCategory(category)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	id64 := int64(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	err = c.CategoryUseCases.RemoveCategoryById(id64)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": fmt.Sprintf("Category with id %v not found", id64),
		})
	}

	ctx.Status(204)
}
