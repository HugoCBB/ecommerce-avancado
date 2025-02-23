package routers

import (
	"github.com/HugoCBB/Api/controllers/produtos"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	v1 := r.Group("/api/produtos/")
	{
		v1.GET("/", produtos.GetProdutos)
		v1.POST("/adicionar", produtos.PostProdutos)
		v1.DELETE("/deletar/:id", produtos.DeleteProdutos)
	}

	r.Run(":8000")

}
