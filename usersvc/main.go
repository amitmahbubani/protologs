package main

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/amitmahbubani/protologs/pkg/logger"
	"github.com/amitmahbubani/protologs/pkg/protolog"
)

func main() {
	var ctx context.Context
	ctx = context.Background()

	// init logger
	ctx, log := logger.New(context.TODO())
	_ = log

	startServer(ctx)

	CreateUser(ctx)

	shutDown(ctx)
}

func startServer(ctx context.Context) {
	logger.Ctx(ctx).Info("server started")
}

func CreateUser(ctx context.Context) {
	logger.Ctx(ctx).Info("create user request", zap.Object("details", &protolog.UserSvcLog{
		Name: "amit mahbubani",
		City: "bangalore",
	}))

	userId := "user_11"

	dummyErr := errors.New("something bad happened")
	logger.Ctx(ctx).Error("create user failed", zap.Object("details", &protolog.UserSvcLog{
		Errormsg: dummyErr.Error(),
	}))

	// log success
	logger.Ctx(ctx).Info("payment processed", zap.Object("details", &protolog.UserSvcLog{
		UserId: userId,
	}))
}

func shutDown(ctx context.Context) {
	logger.Ctx(ctx).Info("shutting down")
}
