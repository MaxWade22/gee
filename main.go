package main

import "Gee/router"

func main() {
	r := router.Router()
	r.Run(":8080")
}
