package main

import (
	"fmt"

	"github.com/felix021/idaas/entity"
)

func sync() {
	x := entity.Obj{}
	fmt.Println("%v", x)
}

func main() {

	sync()
}
