package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Masterminds/sprig"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/starefossen/guestbook.go/src/posts"
)

var redisHost = flag.String("redishost", "192.168.99.102:6379", "Redis address")
var redisPass = flag.String("redispass", "", "Redis password")
var redisDB = flag.Int("redisdb", 0, "Redis database")

var rdb *redis.Client

func main() {
	router := gin.Default()
	router.SetFuncMap(sprig.HtmlFuncMap())
	router.HTMLRender = loadTemplates("./templates")
	// router.LoadHTMLGlob("templates/views/*")

	// nc, _ := nats.Connect(fmt.Sprintf("nats://%s:%d", "192.168.99.102", 4222))
	rdb = redis.NewClient(&redis.Options{
		Addr: *redisHost, Password: *redisPass, // no password set
		DB: *redisDB, // use default DB
	})
	ps := posts.PostStorage{Rdb: rdb}

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	router.GET("/", posts.GetPosts(&ps))
	router.GET("/post", posts.GetNewPost)
	router.POST("/post", posts.PostNewPost(&ps))

	host := ""
	port := "8080"

	// This is to prevent the macOS firewall from complaining!
	if gin.IsDebugging() {
		host = "localhost"
	}

	router.Run(fmt.Sprintf("%s:%s", host, port))
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(filepath.Base(include), sprig.HtmlFuncMap(), files...)
	}
	return r
}
