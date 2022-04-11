package state

import (
	"log"
	"net/http"

	iv "github.com/AutoResto/inventory/controller/inventory"
	modelMaterial "github.com/AutoResto/material/entity"
	mtservice "github.com/AutoResto/material/service"
	ow "github.com/AutoResto/owner/controller/owner"
	"github.com/gin-gonic/gin"
)

type hasItemState struct {
	checker *checker
}

func (h *hasItemState) requestItem(c *gin.Context) {
	ow.SearchMaterial(c)
}

func (h *hasItemState) addItem(c *gin.Context) {
	materialName := c.PostForm("material")
	//cek terlebih dahulu apakah sudah ada nama item yang tersedia di database
	rows, err := mtservice.NewMaterialRepository().SearchMaterialServiceDB(c)
	cek := true

	if err != nil {
		log.Print(err)
	}

	var item modelMaterial.Material
	var items []modelMaterial.Material
	for rows.Next() {
		if err := rows.Scan(&item.Id, &item.Name, &item.Quantity, &item.Unit); err != nil {
			log.Fatal(err.Error())
		} else {
			items = append(items, item)
		}
	}

	for i := 0; i < len(items); i++ {
		if items[i].Name == materialName { //misal sudah ada,tidak di insert
			cek = false
		}
	}
	//bila material baru sekali
	if cek {
		err = mtservice.NewMaterialRepository().InsertMaterialServiceDB(c)
	}
	//bila material terdaftar,jumlahkan quantity nya
	if !cek {
		err = mtservice.NewMaterialRepository().UpdateMaterialServiceDB(c)
	}
	var response modelMaterial.MaterialResponse
	if err == nil {
		response.Message = "Success"
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Failed"
		c.JSON(http.StatusBadRequest, response)

	}

}

func (h *hasItemState) dispenseItem(c *gin.Context) {
	iv.RemoveMaterial(c)
}
