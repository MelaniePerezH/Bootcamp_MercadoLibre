package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre" binding:"required"`
	Color     string `json:"color" binding:"required"`
	Precio    int    `json:"precio" binding:"required"`
	Stock     bool   `json:"stock" binding:"required"`
	Codigo    string `json:"codigo" binding:"required"`
	Publicado bool   `json:"publicado" binding:"required"`
	Fecha     string `json:"fecha" binding:"required"`
}

var productitos = []Product{}
var token string = "melanieesmuypro"

func getAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, productitos)
}
func postProduct(c *gin.Context) {
	tokenHardCodeado := c.GetHeader("token")
	if token != tokenHardCodeado {
		c.JSON(401, gin.H{"error": "No tiene los permisos para realizar esta operación"})
		return
	}
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	newProduct.Id = len(productitos) + 1
	productitos = append(productitos, newProduct)
	c.IndentedJSON(http.StatusCreated, productitos)
}
func getProductById(c *gin.Context) {
	tokenHardCodeado := c.GetHeader("token")
	if token != tokenHardCodeado {
		c.JSON(401, gin.H{"error": "No tiene los permisos para realizar esta operación"})
		return
	}
	id := c.Params.ByName("id")
	for _, p := range productitos {
		intUserId := 0
		_, err := fmt.Sscan(id, &intUserId)
		if err != nil {
			log.Fatal(err)
		}
		if p.Id == intUserId {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado"})
}
func getProductByNombre(c *gin.Context) {
	tokenHardCodeado := c.GetHeader("token")
	if token != tokenHardCodeado {
		c.JSON(401, gin.H{"error": "No tiene los permisos para realizar esta operación"})
		return
	}
	nombre := c.Param("nombre")
	for _, p := range productitos {
		if p.Nombre == nombre {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"code": "404", "message": "Theme no encontrado"})
}

func main() {
	myjson, err := ioutil.ReadFile("productos.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(myjson, &productitos)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/productos", getAll)
	router.POST("/productos/new", postProduct)
	router.GET("/productos/:id", getProductById)
	router.GET("/productos/get/:nombre", getProductByNombre)
	router.Run("localhost:8080")
}
