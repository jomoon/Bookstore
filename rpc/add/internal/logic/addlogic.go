package logic

import (
	"context"
	"fmt"

	"Bookstore/rpc/add/add"
	"Bookstore/rpc/add/internal/svc"
	"Bookstore/rpc/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	model, err := l.svcCtx.Model.Insert(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	lastInsertId, err := model.LastInsertId()
	fmt.Printf("last insert id %d", lastInsertId)
	return &add.AddResp{
		Ok: true,
	}, nil
}
