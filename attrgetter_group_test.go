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

	"github.com/pfflabs/slogctx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grouping AttrGetter instances", func() {

	When("the group attributes are found in the context", func() {
		getter := slogctx.Group(groupName, fooGetter, barGetter)
		ctx := context.WithValue(context.Background(), fooCtxKey, fooAttrValue)
		ctx = context.WithValue(ctx, barCtxKey, barAttrValue)
		ret := getter.GetAttrs(ctx)

		It("returns a slice containing the group of attributes", func() {
			Expect(ret).To(
				And(
					HaveLen(1),
					ContainElement(slog.Group(groupName, fooAttr, barAttr)),
				),
			)
		})
	})

	When("zero group attributes are found in the context", func() {
		getter := slogctx.Group(groupName, fooGetter, barGetter)
		ret := getter.GetAttrs(context.Background())

		It("returns an empty slice of attributes", func() {
			Expect(ret).To(BeEmpty())
		})
	})

	When("the group name is an empty string", func() {
		name := ""

		It("panics", func() {
			Expect(func() { slogctx.Group(name, fooGetter) }).To(
				PanicWith("key is empty"),
			)
		})
	})

	When("passing zero AttrGetter instances", func() {

		It("panics", func() {
			Expect(func() {
				slogctx.Group(groupName)
			}).To(
				PanicWith("received 0 AttrGetters"),
			)
		})
	})

	When("passing nil for the AttrGetter", func() {

		It("panics", func() {
			Expect(func() {
				slogctx.Group(groupName, nil)
			}).To(
				PanicWith("AttrGetter is nil"),
			)
		})
	})

	When("passing multiple AttrGetter arguments", func() {

		Context("and one of the AttrGetter references is nil", func() {

			It("panics", func() {
				Expect(func() {
					slogctx.Group(groupName, fooGetter, nil, barGetter)
				}).To(
					PanicWith("AttrGetter 2 of 3 is nil"),
				)
			})
		})
	})
})
