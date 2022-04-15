package respository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type MenuRepository interface {
	SelectMenuOwnerServiceDB(c *gin.Context) (*sql.DB, error)
	SelectMenuChefServiceDB(c *gin.Context)(*sql.DB,error)
	SearchMenuServiceDB(c *gin.Context) (*sql.DB, error)
	GetAvailabilityMenuServiceDB(c *gin.Context) (*sql.DB, error)
	GetMenuDetailByIdService(c *gin.Context)(*sql.DB,error)
	UpdateMenuServiceDB(c *gin.Context) error
	UpdateDetailMenuService(c *gin.Context) error
	InsertMenuService(c *gin.Context) error
	DeleteMenuServiceDB(c *gin.Context) error
}
