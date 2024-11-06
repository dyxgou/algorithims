package test

import (
	"context"
	"testing"
)

type TestingSuite struct {
	T          *testing.T
	Ctx        context.Context
	Cancel     context.CancelFunc
	cleanFuncs []func()
}

func NewTestingSuite(t *testing.T) *TestingSuite {
	ctx, cancel := context.WithCancel(context.Background())

	ts := &TestingSuite{
		T:          t,
		Ctx:        ctx,
		Cancel:     cancel,
		cleanFuncs: make([]func(), 0, 10),
	}

	return ts
}

func (ts *TestingSuite) AppendCleanUp(f func()) {
	ts.cleanFuncs = append(ts.cleanFuncs, f)
}

func (ts *TestingSuite) CleanUp() {
	for _, cleanFunc := range ts.cleanFuncs {
		cleanFunc()
	}
}
