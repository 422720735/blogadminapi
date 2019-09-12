package main

import (
	"blogadminapi/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":4000")
}
