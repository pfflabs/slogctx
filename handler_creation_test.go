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
	"log/slog"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/pfflabs/slogctx"
)

var _ = Describe("Creating a new handler", func() {
	target := slog.NewTextHandler(os.Stderr, nil)
	getter := noopGetter

	When("the 'target' argument is nil", func() {
		target := slog.Handler(nil)

		It("causes a panic", func() {
			Expect(func() { slogctx.NewHandler(target, getter) }).
				To(PanicWith("target is nil"))
		})
	})

	When("passing zero AttrGetter instances", func() {

		It("causes a panic", func() {
			Expect(func() { slogctx.NewHandler(target) }).
				To(PanicWith("received 0 AttrGetters"))
		})
	})

	When("passing nil for the AttrGetter", func() {

		It("causes a panic", func() {
			Expect(func() { slogctx.NewHandler(target, nil) }).
				To(PanicWith("AttrGetter is nil"))
		})
	})

	When("passing multiple AttrGetter arguments", func() {

		Context("and one of the AttrGetter references is nil", func() {
			nilGetter := slogctx.AttrGetter(nil)

			It("causes a panic", func() {
				Expect(func() {
					slogctx.NewHandler(target, getter, nilGetter, fooGetter)
				}).To(PanicWith("AttrGetter 2 of 3 is nil"))
			})
		})

		Context("and none the arguments are nil", func() {
			target := slog.NewTextHandler(os.Stderr, nil)

			It("returns a handler", func() {
				Expect(slogctx.NewHandler(target, getter)).ToNot(BeNil())
			})
		})
	})
})
