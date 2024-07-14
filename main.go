package main

import (
	"oj/router"
)

func main() {
	r := router.Router()

	r.Run(":8090")
}
