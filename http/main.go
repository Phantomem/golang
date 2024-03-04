package main

import (
	"example/http/album"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

func main() {
	router := gin.Default()
	router.GET("/listAlbums", album.ListAlbums)
	router.Run("localhost:8080")
}
