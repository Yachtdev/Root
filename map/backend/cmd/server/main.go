package main

import (
	"context"
	"fmt"

	"server/internal/command"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	if stop != nil {
		defer stop()
	}

	if err := command.NewCommand().ExecuteContext(ctx); err != nil {
		fmt.Println(fmt.Sprintf(`Failed to execute command: %s`, err.Error()))
	}
}
