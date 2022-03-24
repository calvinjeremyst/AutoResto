package repository

import (
	rc "github.com/AutoResto/recipe/entity"
	"github.com/gin-gonic/gin"
)

type RecipeRepository interface {
	SelectAllRecipeServiceDB(c *gin.Context) (rc.RecipeDetail, error)
	InsertRecipeService(c *gin.Context) error
	InsertRecipeDetailService(c *gin.Context) error
	SelectMenuRecipeServiceDB(c *gin.Context) (rc.Recipe, error)
}
