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
	"time"

	"github.com/pfflabs/slogctx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating an AttrGetter from a lookup function", func() {

	When("the key is an empty string", func() {
		key := ""
		lookup := slogctx.IgnoreZero(func(_ context.Context) any {
			return nil
		})

		It("panics", func() {
			Expect(func() { slogctx.Attr(key, lookup) }).To(
				PanicWith("key is empty"),
			)
		})
	})

	When("the lookup function is nil", func() {

		It("panics", func() {
			Expect(func() { slogctx.Attr[any](fooAttrName, nil) }).To(
				PanicWith("lookup is nil"),
			)
		})
	})

	When("the lookup function returns a bool value and a status bool", func() {
		lookup := func(ctx context.Context) (value bool, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(bool)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := true
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Bool(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns a duration value and a status bool", func() {
		lookup := func(ctx context.Context) (value time.Duration, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(time.Duration)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := 30 * time.Second
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Duration(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns a float value and a status bool", func() {
		lookup := func(ctx context.Context) (value float64, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(float64)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := 3.14
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Float64(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns an int value and a status bool", func() {
		lookup := func(ctx context.Context) (value int, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(int)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		ctx := context.WithValue(context.Background(), fooCtxKey, fooAttrValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(fooAttr),
				),
			)
		})
	})

	When("the lookup function returns an int64 value and a status bool", func() {
		lookup := func(ctx context.Context) (value int64, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(int64)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := int64(42)
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Int64(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns a string value and a status bool", func() {
		lookup := func(ctx context.Context) (value string, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(string)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := "hello!"
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.String(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns a time value and a status bool", func() {
		lookup := func(ctx context.Context) (value time.Time, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(time.Time)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := time.Now()
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Time(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns an uint64 value and a status bool", func() {
		lookup := func(ctx context.Context) (value uint64, ok bool) {
			value, ok = ctx.Value(fooCtxKey).(uint64)
			return
		}
		getter := slogctx.Attr(fooAttrName, lookup)
		fooValue := uint64(42)
		ctx := context.WithValue(context.Background(), fooCtxKey, fooValue)
		ret := getter.GetAttrs(ctx)

		It("returns the expected attribute", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Uint64(fooAttrName, fooValue)),
				),
			)
		})
	})

	When("the lookup function returns a value", func() {
		lookup := func(ctx context.Context) int {
			return ctx.Value(fooCtxKey).(int)
		}
		ctx := context.WithValue(context.Background(), fooCtxKey, fooAttrValue)

		Context("and the 'zero' value should be logged", func() {
			getter := slogctx.Attr(fooAttrName, slogctx.LogZero(lookup))
			ret := getter.GetAttrs(ctx)

			It("returns the expected attribute", func() {
				Expect(ret).To(
					And(
						HaveLen(1),
						ContainElement(fooAttr),
					),
				)
			})
		})

		Context("and the 'zero' value should not be logged", func() {
			getter := slogctx.Attr(fooAttrName, slogctx.IgnoreZero(lookup))
			ret := getter.GetAttrs(ctx)

			It("returns the expected attribute", func() {
				Expect(ret).To(
					And(
						HaveLen(1),
						ContainElement(fooAttr),
					),
				)
			})
		})
	})
})

var _ = When("passing nil to IgnoreZero", func() {

	It("panics", func() {
		Expect(func() { slogctx.IgnoreZero[any](nil) }).To(
			PanicWith("lookup is nil"),
		)
	})
})

var _ = When("passing nil to LogZero", func() {

	It("panics", func() {
		Expect(func() { slogctx.LogZero[any](nil) }).To(
			PanicWith("lookup is nil"),
		)
	})
})
