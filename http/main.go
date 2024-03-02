package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

func parseInt(str string) int {
	defaultValue := 0
	val, err := strconv.Atoi(str)
	if err != nil {
		val = defaultValue
	}
	return val
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &target)
}

func listAlbums(c *gin.Context) {
	limit := parseInt(c.DefaultQuery("limit", "10"))
	page := parseInt(c.DefaultQuery("page", "0"))
	log.Print(limit, page)
	albums := []Album{}
	err := getJson("https://jsonplaceholder.typicode.com/albums", &albums)
	if err != nil {
		log.Panic("Exception on listAlbums:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	var output []Album = albums[page*limit : page*limit+limit] // TODO fix pagination
	c.IndentedJSON(http.StatusOK, output)
}

func main() {
	router := gin.Default()
	router.GET("/listAlbums", listAlbums)
	router.Run("localhost:8080")
}
