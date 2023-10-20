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

var _ = Describe("Handling a record", func() {
	var (
		rec        slog.Record
		spyHandler *HandlerSpy
		getters    []slogctx.AttrGetter
		ctx        context.Context
		err        error
	)
	const testMessage = "this message is a test"

	BeforeEach(func() {
		spyHandler = NewHandlerSpy()
		rec = slog.NewRecord(time.Now(), slog.LevelInfo, testMessage, 0)
	})

	JustBeforeEach(func() {
		err = slogctx.NewHandler(spyHandler, getters...).Handle(ctx, rec)
	})

	When("no attributes are retrieved from the context", func() {

		BeforeEach(func() {
			ctx = context.Background()
			getters = []slogctx.AttrGetter{noopGetter}
		})

		It("does not return an error", func() {
			Expect(err).To(BeNil())
		})

		It("passes the same context to the target handler", func() {
			Expect(spyHandler.HandleSpy.Ctx).To(BeIdenticalTo(ctx))
		})

		It("does not add attributes to the record passed to the target handler", func() {
			Expect(spyHandler.HandleSpy.Rec.NumAttrs()).To(Equal(rec.NumAttrs()))
		})
	})

	When("one attribute is retrieved from the context", func() {

		BeforeEach(func() {
			ctx = context.WithValue(context.Background(), fooCtxKey, fooAttrValue)
			getters = []slogctx.AttrGetter{fooGetter}
		})

		It("does not return an error", func() {
			Expect(err).To(BeNil())
		})

		It("passes the same context to the target handler", func() {
			Expect(spyHandler.HandleSpy.Ctx).To(BeIdenticalTo(ctx))
		})

		It("appends the attribute to the record passed to the target handler", func() {
			Expect(spyHandler.HandleSpy.Rec.NumAttrs()).To(Equal(rec.NumAttrs() + 1))

			Expect(GetAttrs(spyHandler.HandleSpy.Rec)).To(
				ContainElement(
					slog.Attr{Key: fooAttrName, Value: slog.IntValue(fooAttrValue)},
				),
			)
		})
	})

	When("multiple attributes are retrieved from the context", func() {

		BeforeEach(func() {
			ctx = context.WithValue(context.Background(), fooCtxKey, fooAttrValue)
			ctx = context.WithValue(ctx, barCtxKey, barAttrValue)
			ctx = context.WithValue(ctx, pifCtxKey, pifAttrValue)
			getters = []slogctx.AttrGetter{fooGetter, barGetter, pifGetter}
		})

		It("does not return an error", func() {
			Expect(err).To(BeNil())
		})

		It("passes the same context to the target handler", func() {
			Expect(spyHandler.HandleSpy.Ctx).To(BeIdenticalTo(ctx))
		})

		It("appends the attributes to the record passed to the target handler", func() {
			Expect(spyHandler.HandleSpy.Rec.NumAttrs()).To(Equal(rec.NumAttrs() + 3))
			Expect(GetAttrs(spyHandler.HandleSpy.Rec)).To(
				ContainElements(
					fooAttr,
					barAttr,
					pifAttr,
				),
			)
		})
	})

	When("a group of attributes is retrieved from the context", func() {

		BeforeEach(func() {
			ctx = context.WithValue(context.Background(), fooCtxKey, fooAttrValue)
			ctx = context.WithValue(ctx, barCtxKey, barAttrValue)
			ctx = context.WithValue(ctx, pifCtxKey, pifAttrValue)
			getters = []slogctx.AttrGetter{
				slogctx.Group(groupName, fooGetter, barGetter, pifGetter),
			}
		})

		It("does not return an error", func() {
			Expect(err).To(BeNil())
		})

		It("passes the same context to the target handler", func() {
			Expect(spyHandler.HandleSpy.Ctx).To(BeIdenticalTo(ctx))
		})

		It("appends the attributes to the record passed to the target handler", func() {
			Expect(spyHandler.HandleSpy.Rec.NumAttrs()).To(Equal(rec.NumAttrs() + 1))
			Expect(GetAttrs(spyHandler.HandleSpy.Rec)).To(
				ContainElement(
					slog.Group(groupName, fooAttr, barAttr, pifAttr),
				),
			)
		})
	})
})
