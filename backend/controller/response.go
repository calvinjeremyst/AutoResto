package controller

import (
	"net/http"

	entityMaterial "github.com/AutoResto/domain/material/entity"
	entityMenu "github.com/AutoResto/domain/menu/entity"
	entityRecipe "github.com/AutoResto/domain/recipe/entity"
	"github.com/gin-gonic/gin"
)

// Menu Response
func SendMenuSuccessResponse(c *gin.Context, ur entityMenu.MenuResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendMenuErrorResponse(c *gin.Context, ur entityMenu.MenuResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

func SendRecipeSuccessResponse(c *gin.Context, ur entityRecipe.RecipeResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendRecipeErrorResponse(c *gin.Context, ur entityRecipe.RecipeResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

func SendRecipeDetailSuccessResponse(c *gin.Context, ur entityRecipe.RecipeDetailResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendRecipeDetailErrorResponse(c *gin.Context, ur entityRecipe.RecipeDetailResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

func SendMaterialSuccessResponse(c *gin.Context, ur entityMaterial.MaterialResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendMaterialErrorResponse(c *gin.Context, ur entityMaterial.MaterialResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
