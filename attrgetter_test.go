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

var _ = Describe("Adapting a function into an AttrGetter", func() {
	getter := slogctx.AttrGetterFunc(func(c context.Context) []slog.Attr {
		v := c.Value(fooCtxKey)
		if v == nil {
			return nil
		}

		return []slog.Attr{slog.Int(fooAttrName, v.(int))}
	})
	ctx := context.WithValue(context.Background(), fooCtxKey, fooAttrValue)
	ret := getter.GetAttrs(ctx)

	It("returns a slice containing the attribute", func() {
		Expect(ret).To(
			And(
				HaveLen(1),
				ContainElement(fooAttr),
			),
		)
	})
})
