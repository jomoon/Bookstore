package logic

import (
	"context"
	"time"

	"Bookstore/api/internal/svc"
	"Bookstore/api/internal/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginerLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginerLogic {
	return LoginerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginerLogic) Loginer(req types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	seconds := l.svcCtx.Config.Auth.AccessExpire
	exp := iat + seconds
	claims["exp"] = exp
	claims["iat"] = iat
	claims["username"] = req.Username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	token_str, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Username: req.Username,
		Token:    token_str,
		Exp:      exp,
	}, nil
}
