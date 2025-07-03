package service

import (
	"context"
	"database/sql"
	"github.com/frochyzhang/ag-core/ag/ag_db"
	"github.com/frochyzhang/ag-core/ag/ag_ext"
	mw "github.com/frochyzhang/ag-core/ag/ag_ext"
	"log"
	"time"
)

type TmMiddlewareContext struct {
	tm ag_db.TransactionManager
}

func NewTmMiddlewareContext(tm ag_db.TransactionManager) TmMiddlewareContext {
	return TmMiddlewareContext{tm: tm}
}

func (tmCtx TmMiddlewareContext) GetOrder() int {
	return mw.MiddlewarePriorityHigh
}

func (tmCtx TmMiddlewareContext) GetMiddleware() ag_ext.Middleware {
	return tmCtx.TransactionMiddleware
}

func (tmCtx TmMiddlewareContext) TransactionMiddleware(
	method string,
	ctx context.Context,
	req interface{},
	next func(context.Context, interface{}) (interface{}, error),
) (res interface{}, err error) {
	start := time.Now()
	log.Printf("[%s] 准备添加事务", method)
	ctx, txCall := tmCtx.tm.WithTransaction(ctx, append(make([]*sql.TxOptions, 0), &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: true})...)
	defer func() {
		txCall(err)
	}()

	err = tmCtx.tm.Transaction(ctx, func(ctx context.Context) error {
		res, err = next(ctx, req)
		return err
	})

	if err != nil {
		log.Printf("[%s] 事务执行失败 %v: %v", method, time.Since(start), err)
	} else {
		log.Printf("[%s] 事务执行成功 %v", method, time.Since(start))
	}
	return
}

type GlobalTraceMiddleWare struct{}

func (g GlobalTraceMiddleWare) GetOrder() int {
	return mw.MiddlewarePriorityNormal
}

func NewGlobalTraceMiddleware() GlobalTraceMiddleWare {
	return GlobalTraceMiddleWare{}
}
func (g GlobalTraceMiddleWare) GetMiddleware() mw.Middleware {
	return func(method string, ctx context.Context, req interface{}, next func(context.Context, interface{}) (interface{}, error)) (interface{}, error) {
		log.Println("全局事务开启啦")
		res, err := next(ctx, req)
		log.Println("<UNK>")
		return res, err
	}
}
