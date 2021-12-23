package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int64   `json:"id" binding:"required"`
	Nombre        string  `json:"nombre" binding:"required"`
	Color         string  `json:"color" binding:"required"`
	Precio        float64 `json:"precio" binding:"required"`
	Stock         int64   `json:"stock" binding:"required"`
	Codigo        string  `json:"codigo" binding:"required"`
	Publicado     bool    `json:"publicado" binding:"required"`
	FechaCreacion string  `json:"fechaCreacion" binding:"required"`
}

/*
func prueba() {
	//crear router
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello David!",
		})
	})
	//GET ALL PRODUCTOS
	jsonData, err := ioutil.ReadFile("./products.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonData)
	productos := []Producto{}
	if err := json.Unmarshal(jsonData, &productos); err != nil {
		panic(err)
	}

	router.GET("/getProducts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"PRODUCTOS": productos,
		})
	})
	router.Run()
}
**/
func HandlerHola(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello David!",
	})
}
func HandlerAllProductos(ctx *gin.Context) {
	/*
		jsonData, err := ioutil.ReadFile("./products.json")
		if err != nil {
			log.Fatal(err)
		}
		//productos := []Producto{}
		if err := json.Unmarshal(jsonData, &productos); err != nil {
			panic(err)
		}
	*/
	ctx.JSON(200, gin.H{
		"PRODUCTOS ": productos,
	})
}
func HandlerFilterProducts(ctx *gin.Context) {
	/*
		jsonData, err := ioutil.ReadFile("./products.json")
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err,
			})
		}
		//productos := []Producto{}
		if err := json.Unmarshal(jsonData, &productos); err != nil {
			panic(err)
		}
	*/
	var match = false
	var filteredProd []Producto
	for _, producto := range productos {
		if ctx.Query("id") == fmt.Sprint(producto.Id) {
			match = true
		} else if ctx.Query("id") != "" {
			match = false
		}
		if ctx.Query("nombre") == producto.Nombre {
			match = true
		} else if ctx.Query("nombre") != "" {
			match = false
		}
		if ctx.Query("color") == producto.Color {
			match = true
		} else if ctx.Query("color") != "" {
			match = false
		}
		if ctx.Query("precio") == fmt.Sprint(producto.Precio) {
			match = true
		} else if ctx.Query("precio") != "" {
			match = false
		}
		if ctx.Query("stock") == fmt.Sprint(producto.Stock) {
			match = true
		} else if ctx.Query("stock") != "" {
			match = false
		}
		if ctx.Query("codigo") == producto.Codigo {
			match = true
		} else if ctx.Query("codigo") != "" {
			match = false
		}
		if ctx.Query("publicado") == fmt.Sprint(producto.Publicado) {
			match = true
		} else if ctx.Query("publicado") != "" {
			match = false
		}
		if ctx.Query("fechaCreacion") == producto.FechaCreacion {
			match = true
		} else if ctx.Query("fechaCreacion") != "" {
			match = false
		}
		if match {
			filteredProd = append(filteredProd, producto)
		}
	}
	ctx.JSON(200, gin.H{
		"PRODUCTOS": filteredProd,
	})

}

func HandlerFindProduct(ctx *gin.Context) {
	/*
		jsonData, err := ioutil.ReadFile("./products.json")
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err,
			})
		}
		//productos := []Producto{}
		if err := json.Unmarshal(jsonData, &productos); err != nil {
			panic(err)
		}
	*/
	p := Producto{}
	encontrado := false
	for _, producto := range productos {
		if ctx.Param("id") == fmt.Sprint(producto.Id) {
			encontrado = true
			p = producto
		}
	}
	if encontrado {
		ctx.JSON(200, gin.H{
			"PRODUCTO": p,
		})
	} else {
		ctx.JSON(404, gin.H{
			"RESPONSE": "No existe el producto con ese ID",
		})
	}
}
func HandlerCreateProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	myToken := "123456"
	if token != myToken {
		ctx.JSON(401, gin.H{
			"error": "Invalid token",
		})
		return
	}
	var lastID int64
	for _, producto := range productos {
		lastID = producto.Id
	}
	var req Producto
	err2 := ctx.ShouldBindJSON(&req)
	if err2 != nil {
		ctx.JSON(400, gin.H{
			"error": err2,
		})
		return
	}
	req.Id = lastID + 1
	productos = append(productos, req)
	ctx.JSON(200, req)

}
func cargarProductos() {
	jsonData, err := ioutil.ReadFile("./products.json")
	if err != nil {
		log.Fatal(err)
	}
	//productos := []Producto{}
	if err := json.Unmarshal(jsonData, &productos); err != nil {
		panic(err)
	}
}

var productos []Producto

func main() {
	router := gin.Default()
	cargarProductos()
	router.GET("/hello", HandlerHola)
	router.GET("/allProducts", HandlerAllProductos)
	router.GET("/filterProducts", HandlerFilterProducts)
	router.GET("/product/:id", HandlerFindProduct)
	router.POST("/createProducto", HandlerCreateProduct)
	router.Run()
}
