package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Estructura del producto
type Producto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

// store es una base de datos en memoria
type Store struct {
	Productos []Producto
}

func main() {

	// Carga la base de datos en memoria
	store := Store{}
	store.LoadStore()

	engine := gin.Default()

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		})

		grupoProducto := group.Group("/producto")
		{
			grupoProducto.GET("", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"data": store.Productos,
				})
			})

			grupoProducto.GET("/search/:parametroPrecio", func(ctx *gin.Context) {

				precioParametro := ctx.Param("parametroPrecio")

				precioCasteado, err := strconv.ParseFloat(precioParametro, 64)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"mensaje": "parametro invalido",
					})
					return
				}

				var result []Producto
				for _, producto := range store.Productos {
					if producto.Price > precioCasteado {
						result = append(result, producto)
					}
				}

				ctx.JSON(http.StatusOK, gin.H{
					"data": result,
				})

			})
		
	}

	//Metodo post para product params
	grupoProducto.POST("/productparams", addProductParams) 
			

	//Metodo get para product params
	grupoProducto.GET("/products/:id", getProductByID) 
	

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

//punto 1: product params
ffunc addProductParams(ctx *gin.Context) {
    var product Producto

    // Recupera los datos del producto de los parámetros de la solicitud
    product.Id = ctx.Query("id")
    product.Name = ctx.Query("name")
    product.Quantity, _ = strconv.Atoi(ctx.Query("quantity"))
    product.CodeValue = ctx.Query("code_value")
    product.IsPublished, _ = strconv.ParseBool(ctx.Query("is_published"))
    expiration, _ := time.Parse(time.RFC3339, ctx.Query("expiration"))
    product.Expiration = expiration
    product.Price, _ = strconv.ParseFloat(ctx.Query("price"), 64)

    // Agrega el producto a la lista
    products = append(products, product)

    // Devuelve el producto en formato JSON
    ctx.JSON(http.StatusCreated, product)
}

// Creando la ruta product/id
func getProductByID(ctx *gin.Context) {
    id := ctx.Param("id")

    for _, product := range products {
        if product.ID == id {
            ctx.JSON(http.StatusOK, product)
            return
        }
    }

    ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}



// LoadStore carga la base de datos en memoria
func (s *Store) LoadStore() {
	s.Productos = []Producto{
		{
			Id:          "1",
			Name:        "Coco Cola",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          "2",
			Name:        "Pepsito",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          "3",
			Name:        "Fantastica",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},
	}
}