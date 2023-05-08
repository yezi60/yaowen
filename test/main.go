package main

import (
	"soleaf.xyz/yaowen/handler"
	initialize "soleaf.xyz/yaowen/init"
	"soleaf.xyz/yaowen/model"
)

func main() {
	initialize.InitDB()
	handler.Save(model.Data{
		Address: "219.228.135.127",
		Health:  232,
	})
}
