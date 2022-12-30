package service

import (
	"context"
	"strconv"
)

type accountService struct {
}

func NewAccountService() *accountService {
	return &accountService{}
}

func (as *accountService) GetAccountsList(ctx context.Context) []string {
	x := make([]string, 10)
	for i := 0; i < 10; i++ {
		x[i] = strconv.Itoa(i)
	}
	return x
}
