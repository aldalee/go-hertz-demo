package main

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func hertzRun() {
	h := server.Default()
	h.Use(recovery.Recovery())

	register(h)
	h.Spin()
}
