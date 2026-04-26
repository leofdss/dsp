package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"dsp/internal/adapters/file"
	"dsp/internal/adapters/pipewire"
	"dsp/internal/app"
	"dsp/internal/domain"
)

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: dsp <config.{json,yaml,yml}>")
	}

	loader := file.ConfigLoader{}
	cfg, err := loader.Load(args[0])
	if err != nil {
		return err
	}

	processor := domain.NewGain(cfg.Gain)
	stream := pipewire.Stream{}
	engine := app.NewEngine(stream, processor)

	return engine.Run(ctx)
}
