package routes

import (
	"github.com/gofiber/fiber/v2"
	"product-service/internal/handler"
	"product-service/internal/middleware"
)

func ProductRoutes(app *fiber.App, productHandler *handler.ProductHandler, categoryHandler *handler.CategoryHandler) {

	// ---------- PRODUCT ROUTES ----------
	product := app.Group("/api/products")
	product.Use(middleware.AuthMiddleware())

	// Admin-only product routes
	adminProduct := product.Group("/admin")
	adminProduct.Use(middleware.AdminOnly())
	adminProduct.Post("/", productHandler.Create)
	adminProduct.Put("/:id", productHandler.UpdateProduct)
	adminProduct.Delete("/:id", productHandler.DeleteProduct)

	// Public (user-accessible) product routes
	product.Get("/", productHandler.GetAllProducts)
	product.Get("/:id", productHandler.GetProductById)

	// ---------- CATEGORY ROUTES ----------
	category := app.Group("/api/categories")
	category.Use(middleware.AuthMiddleware())

	// Admin-only category routes
	adminCategory := category.Group("/admin")
	adminCategory.Use(middleware.AdminOnly())
	adminCategory.Post("/", categoryHandler.CreateCategory)
	// adminCategory.Put("/:id", categoryHandler.UpdateCategory) // Added for completeness
	// adminCategory.Delete("/:id", categoryHandler.DeleteCategory) // Added for completeness

	// Public (user-accessible) category routes
	category.Get("/", categoryHandler.GetAllCategories)
	category.Get("/:id", categoryHandler.GetCategoryById)
}
