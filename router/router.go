package router

import (
	"Gee/gee"
	"net/http"
)

func Router() *gee.Engine {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello \n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"test"}
		c.String(http.StatusOK, names[100])
	})

	return r
}
