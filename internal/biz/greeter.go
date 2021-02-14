package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreeter(*Greeter) error
	UpdateGreeter(*Greeter) error
}

type GreeterUsecase struct {
	repo GreeterRepo
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo}
}

func (uc *GreeterUsecase) Create(g *Greeter) error {
	return uc.repo.CreateGreeter(g)
}

func (uc *GreeterUsecase) Update(g *Greeter) error {
	return uc.repo.UpdateGreeter(g)
}
