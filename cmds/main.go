package main

import (
	"fmt"
	"godemo/configs"
)

func main() {
	configs.Init()

	fmt.Println("envs: ", configs.Env.Elastic.Addr, configs.Env.UseDifferentNameNeedTag)
}
