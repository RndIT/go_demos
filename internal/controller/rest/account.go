package rest

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/RndIT/go_demos/internal/domain/models"
)


type accountAPI struct {
}

func NewAccountAPI() *accountAPI {
	return &accountAPI{}
}

func (as *accountService) GetAccountsList(ctx context.Context) []string {
	x := make([]string, 10)
	for i := 0; i < 10; i++ {
		x[i] = strconv.Itoa(i)
	}
	return x
}