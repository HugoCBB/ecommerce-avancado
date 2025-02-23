package produtos

import (
	"net/http"
	"time"

	"github.com/HugoCBB/Api/database"
	"github.com/HugoCBB/Api/models"
	"github.com/gin-gonic/gin"
)

func GetProdutos(c *gin.Context) {
	var p []models.Produtos
	database.DB.Find(&p)
	c.JSON(http.StatusOK, p)
}

func PostProdutos(c *gin.Context) {
	var p models.Produtos
	p.DataCriacao = time.Now()
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&p)
	c.JSON(http.StatusOK, p)

}

func DeleteProdutos(c *gin.Context) {
	var p models.Produtos
	id := c.Param("id")

	database.DB.Delete(&p, id)
	c.JSON(http.StatusOK, gin.H{"data": "Produto deletado com sucesso!"})

}
