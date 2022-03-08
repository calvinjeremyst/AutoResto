package controller

import (
	"net/http"

	entityMaterial "github.com/AutoResto/module/material/entity"
	entityMenu "github.com/AutoResto/module/menu/entity"
	entityRecipe "github.com/AutoResto/module/recipe/entity"
	"github.com/gin-gonic/gin"
)

// Menu Response
func sendMenuSuccessResponse(c *gin.Context, ur entityMenu.MenuResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendMenuErrorResponse(c *gin.Context, ur entityMenu.MenuResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

func sendRecipeDetailSuccessResponse(c *gin.Context, ur entityRecipe.RecipeDetailResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendRecipeDetailErrorResponse(c *gin.Context, ur entityRecipe.RecipeDetailResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

func sendMaterialSuccessResponse(c *gin.Context, ur entityMaterial.MaterialResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendMaterialErrorResponse(c *gin.Context, ur entityMaterial.MaterialResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
