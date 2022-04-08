package respository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type MenuRepository interface {
	SelectMenuServiceDB(c *gin.Context) (*sql.DB, error)
	SearchMenuServiceDB(c *gin.Context) (*sql.DB, error)
	GetAvailabilityMenuServiceDB(c *gin.Context) (*sql.DB, error)
	UpdateMenuServiceDB(c *gin.Context) error
	InsertMenuService(c *gin.Context) error
	DeleteMenuServiceDB(c *gin.Context) error
}
