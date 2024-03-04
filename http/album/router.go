package album

import (
	"example/http/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAlbums(c *gin.Context) {
	limit := helpers.ParseInt(c.DefaultQuery("limit", "10"))
	page := helpers.ParseInt(c.DefaultQuery("page", "0"))
	userId := helpers.ParseInt(c.DefaultQuery("userId", "0"))
	albums := []Album{}
	err := helpers.GetJson("https://jsonplaceholder.typicode.com/albums", &albums)
	if err != nil {
		log.Panic("Exception on listAlbums:", err)
		c.IndentedJSON(http.StatusInternalServerError, helpers.ResolveErrorHttpJsonResponse(err))
	}
	if userId != 0 {
		albums = helpers.Filter(albums, func(a Album) bool { return a.UserId == userId })
	}
	from, to := helpers.ResolvePaginationFromArrayLength(len(albums), limit, page)
	c.IndentedJSON(http.StatusOK, helpers.ResolveSuccessHttpJsonResponse(albums[from:to]))
}
