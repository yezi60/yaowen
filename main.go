package main

import (
	"log"
	initialize "soleaf.xyz/yaowen/init"
)

func main() {

	initialize.InitDB()
	r := initialize.Routers()

	log.Panic(r.Run(":8082"))
}



