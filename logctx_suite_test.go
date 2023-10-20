// Copyright 2023 pfflabs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slogctx_test

import (
	"context"
	"log/slog"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pfflabs/slogctx"
)

func TestSrc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "logctx suite")
}

/*
 * Constants, mocks, dummies, spies, etc. used in the test suite.
 */

type ctxKey string

const fooCtxKey = ctxKey("foo")
const fooAttrName = "foo"
const fooAttrValue = 42

var fooAttr = slog.Int(fooAttrName, fooAttrValue)
var fooGetter = slogctx.Attr(fooAttrName, func(ctx context.Context) (int, bool) {
	v, ok := ctx.Value(fooCtxKey).(int)
	return v, ok
})

const barCtxKey = ctxKey("bar")
const barAttrName = "bar"
const barAttrValue = "oom"

var barAttr = slog.String(barAttrName, barAttrValue)
var barGetter = slogctx.Attr(barAttrName, func(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(barCtxKey).(string)
	return v, ok
})

const pifCtxKey = ctxKey("pif")
const pifAttrName = "pif"
const pifAttrValue = true

var pifAttr = slog.Bool(pifAttrName, pifAttrValue)
var pifGetter = slogctx.Attr(pifAttrName, func(ctx context.Context) (bool, bool) {
	v, ok := ctx.Value(pifCtxKey).(bool)
	return v, ok
})

const groupName = "ziz"

var noopGetter = slogctx.Attr("noop", func(_ context.Context) (any, bool) {
	return nil, false
})

type (
	EnableSpy struct {
		Ctx    context.Context
		Level  slog.Level
		Return bool
	}

	HandleSpy struct {
		Ctx context.Context
		Rec slog.Record
		Err error
	}

	WithAttrsSpy struct {
		Attrs []slog.Attr
	}

	WithGroupSpy struct {
		Name string
	}

	HandlerSpy struct {
		*EnableSpy
		*HandleSpy
		*WithAttrsSpy
		*WithGroupSpy
	}
)

func NewHandlerSpy() *HandlerSpy {
	return &HandlerSpy{
		EnableSpy:    &EnableSpy{},
		HandleSpy:    &HandleSpy{},
		WithAttrsSpy: &WithAttrsSpy{},
		WithGroupSpy: &WithGroupSpy{},
	}
}

func (s *EnableSpy) Enabled(ctx context.Context, level slog.Level) bool {
	s.Ctx = ctx
	s.Level = level
	return s.Return
}

func (s *HandleSpy) Handle(ctx context.Context, rec slog.Record) error {
	s.Ctx = ctx
	s.Rec = rec
	return s.Err
}

func (s *WithAttrsSpy) WithAttrs(attrs []slog.Attr) slog.Handler {
	s.Attrs = attrs
	return nil
}

func (s *WithGroupSpy) WithGroup(name string) slog.Handler {
	s.Name = name
	return nil
}

func (s *HandlerSpy) WithAttrs(attrs []slog.Attr) slog.Handler {
	s.WithAttrsSpy.WithAttrs(attrs)
	return s
}

func (s *HandlerSpy) WithGroup(name string) slog.Handler {
	s.WithGroupSpy.WithGroup(name)
	return s
}

func GetAttrs(rec slog.Record) []slog.Attr {
	attrs := make([]slog.Attr, 0, rec.NumAttrs())
	rec.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr)
		return true
	})
	return attrs
}
