package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {

	r := gin.Default()
	m := melody.New()

	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/channel/:name", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "chan.html")
	})

	r.GET("/channel/:name/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	r.Run(":5000")

}
