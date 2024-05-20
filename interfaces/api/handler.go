package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-or/go-solid-sample/app/services"
	"github.com/guilherme-or/go-solid-sample/app/usecases"
	"github.com/guilherme-or/go-solid-sample/domain/repositories"
	"github.com/guilherme-or/go-solid-sample/infra/persistance"
	"github.com/guilherme-or/go-solid-sample/interfaces"
	"github.com/guilherme-or/go-solid-sample/interfaces/api/controllers"
)

const PORT = 8080

var (
	categoryRepo repositories.CategoryRepository = persistance.NewLocalCategoryRepository()
	productRepo  repositories.ProductRepository  = persistance.NewLocalProductRepository()

	categoryService services.CategoryService = services.CategoryService{CategoryRepo: categoryRepo}
	productService  services.ProductService  = services.ProductService{ProductRepo: productRepo}

	categoryUseCases = usecases.CategoryUseCases{CategoryService: categoryService}
	productUseCases  = usecases.ProductUseCases{ProductService: productService}

	categoryController *controllers.CategoryController = &controllers.CategoryController{CategoryUseCases: categoryUseCases}
	productController *controllers.ProductController = &controllers.ProductController{ProductUseCases: productUseCases}
)

// Interface

type Api interface {
	interfaces.UserInterface
}

// Implementation

type GinApi struct {
	Router *gin.Engine
}

func NewApi() Api {
	api := &GinApi{Router: gin.Default()}
	
	categoryController.RegisterHandlers(api.Router)
	productController.RegisterHandlers(api.Router)

	return api
}

func (a *GinApi) Run() {
	a.Router.Run(fmt.Sprintf(":%d", PORT))
}