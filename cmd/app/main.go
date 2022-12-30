package main

import (
	"context"
	"fmt"

	"github.com/RndIT/go_demos/internal/domain/service"
)

func main() {
	// инициализация контекста
	ctx := context.Background()

	servAcc := service.NewAccountService()
	r := servAcc.GetAccountsList(ctx)
	fmt.Println(r)
	fmt.Println("main() stopped")
}
