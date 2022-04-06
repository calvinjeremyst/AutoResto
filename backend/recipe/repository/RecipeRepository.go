package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type RecipeRepository interface {
	SelectAllRecipeServiceDB(c *gin.Context) (*sql.DB, error)
	InsertRecipeService(c *gin.Context) error
	InsertRecipeDetailService(c *gin.Context) error
	GetAvailabilityMenuDescription(c *gin.Context)(*sql.DB,error)
	SelectMenuRecipeServiceDB(c *gin.Context) (*sql.DB, error)
}
