package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func index(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Hello %s, New Release!!!", runtime.Version()))
}

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8080, "HTTP服务器端口")
	flag.Parse()
	logrus.Println("port: " + strconv.Itoa(port))

	router := gin.Default()
	router.GET("/", index)
	router.GET("/albums", getAlbums)

	router.Run(fmt.Sprintf(":%s", strconv.Itoa(port)))
}
