package respository

import (
	"github.com/gin-gonic/gin"
	mn "github.com/AutoResto/menu/entity"
)

type MenuRepository interface {

	SelectMenuServiceDB(c* gin.Context)(mn.Menu,error)
	SearchMenuServiceDB(c *gin.Context)(mn.Menu,error)
	UpdateMenuServiceDB(c *gin.Context) error
	InsertMenuService(c *gin.Context) error
	DeleteMenuServiceDB(c *gin.Context) error


}