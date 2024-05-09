package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	router := gin.Default()
	router.GET("/", getAlbums)

	err := http3.ListenAndServe("0.0.0.0:443", "cert/certificate.crt", "cert/private.key", router)

	if err != nil {
		log.Printf("Server error: %v", err)
	}

}
