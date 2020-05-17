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

	ProcessPayment(ctx)

	shutDown(ctx)
}

func startServer(ctx context.Context) {
	logger.Ctx(ctx).Info("server started")
}

func ProcessPayment(ctx context.Context) {
	logger.Ctx(ctx).Info("payment process started", zap.Object("details", &protolog.PaymentSvcLog{
		Amount: 5600,
		Method: "card",
	}))

	paymentId := "payment_12"

	dummyErr := errors.New("something bad happened")
	logger.Ctx(ctx).Error(dummyErr.Error(), zap.Object("details", &protolog.PaymentSvcLog{
		PaymentId: paymentId,
	}))

	// log success
	logger.Ctx(ctx).Info("payment processed", zap.Object("details", &protolog.PaymentSvcLog{
		PaymentId: paymentId,
		Amount: 5600,
		Method: "card",
		OrderId: "order_11",
	}))
}

func shutDown(ctx context.Context) {
	logger.Ctx(ctx).Info("shutting down")
}
