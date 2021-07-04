package logic

import (
	"context"
	"fmt"

	"Bookstore/rpc/add/add"
	"Bookstore/rpc/add/internal/svc"
	"Bookstore/rpc/model"

	"github.com/tal-tech/go-zero/core/limit"
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

const (
	seconds = 1
	total   = 10
	quota   = 2
)

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	// 增加限流
	periodLimit := limit.NewPeriodLimit(seconds, quota, l.svcCtx.Config.Redis.NewRedis(), "bookstore")
	key := "add"
	code, err := periodLimit.Take(key)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	switch code {
	case limit.OverQuota:
		logx.Errorf("OverQuota key: %v", key)
		return nil, fmt.Errorf("OverQuota")
	case limit.Allowed:
		logx.Infof("AllowedQuota key: %v", key)
	case limit.HitQuota:
		logx.Errorf("HitQuota key: %v", key)
		// todo: maybe we need to let users know they hit the quota
		return nil, fmt.Errorf("HitQuota")
	default:
		logx.Errorf("DefaultQuota key: %v", key)
		// unknown response, we just let the sms go
		return nil, fmt.Errorf("DefaultQuota")
	}

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
