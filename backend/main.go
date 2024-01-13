package main

import (
	_ "backend/internal/packed"

	_ "backend/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
