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
	//currentPath, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}

	router := gin.Default()
	router.GET("/", getAlbums)

	//router.Run("localhost:8080")
	//
	server := http3.Server{
		Addr:    "0.0.0.0:443",
		Port:    443,
		Handler: router.Handler(),
	}

	//err := server.ListenAndServe()

	err := server.ListenAndServeTLS("cert/certificate.crt", "cert/private.key")
	//router.Run()
	if err != nil {
		log.Printf("Server error: %v", err)
	}

}
