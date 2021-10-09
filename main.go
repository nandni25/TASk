package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
        ID     string  `json:"id"`
        NAME  string  `json:"Name"`
        EMAIL string  `json:"Email"`
        PASSWORD  float64 `json:"Password"`
}
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, albums)
}
// albums slice to seed record album data.
var albums = []album{
        {ID: "1", NAME: "ABC", EMAIL: "abc@gmail.com", PASSWORD: 2334},
        {ID: "2", NAME: "DEF", EMAIL: "def@gamil.com", PASSWORD: 4456},
        {ID: "3", NAME: "HIJ", EMAIL: "hij@gmail.com", PASSWORD: 3456},
}
func main() {
        router := gin.Default()
        router.GET("/albums", getAlbums)

        router.Run("localhost:8080")
}

