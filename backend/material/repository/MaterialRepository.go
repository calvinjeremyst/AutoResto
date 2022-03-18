package repository

import (
	"github.com/gin-gonic/gin"
	mt "github.com/AutoResto/material/entity"
)

type MaterialRepository interface {

	GetMaterialServiceDB()(mt.Material,error)
	SearchMaterialServiceDB()(mt.Material,error)
	InsertMaterialServiceDB(c* gin.Context)error
	InsertMaterialHelperServiceDB(c* gin.Context) error
	UpdateMaterialServiceDB(c* gin.Context) error
	DeleteMaterialServiceDB(c* gin.Context) error
}