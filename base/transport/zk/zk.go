package zk

import (
	"context"
	"time"

	"github.com/go-kit/kit/transport"
	"github.com/pkg/errors"
)

type (

	// ErrorHandler is wrapper on top of kit.transport.ErrorHandler
	ErrorHandler interface{ transport.ErrorHandler }

	Decoder func(context.Context, interface{}) (interface{}, error)

	ReconnectOnErr func(error) bool
	DelayOnErr     func(error) time.Duration
)

// consumer Errors
var (
	ErrCreatingConsumer = errors.New("error creating consumer")
)
