package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type MaterialRepository interface {
	GetMaterialServiceDB() (*sql.Rows, error)
	GetNameMaterialService()(*sql.Rows,error)
	SearchMaterialServiceDB() (*sql.Rows, error)
	InsertMaterialServiceDB(c *gin.Context) error
	InsertMaterialHelperServiceDB(c *gin.Context) error
	UpdateMaterialServiceDB(c *gin.Context) error
	DeleteMaterialServiceDB(c *gin.Context) error
}
