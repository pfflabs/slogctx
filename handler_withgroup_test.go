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
	"fmt"
	"log/slog"
	"strings"

	"github.com/pfflabs/slogctx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating a new handler with an attribute group", func() {
	const testMessage = "this message is a test"
	var buf *strings.Builder

	BeforeEach(func() {
		buf = &strings.Builder{}
		handler := slogctx.NewHandler(slog.NewTextHandler(buf, nil), noopGetter).
			WithGroup(groupName)
		slog.New(handler).Info(testMessage, fooAttr, barAttr, pifAttr)
	})

	When("a message is logged", func() {

		Specify("the output includes the message", func() {
			Expect(buf.String()).To(ContainSubstring(testMessage))
		})

		Specify("the output includes the attribute group", func() {
			fs := fmt.Sprintf("%s.%s=%d", groupName, fooAttrName, fooAttrValue)
			bs := fmt.Sprintf("%s.%s=%s", groupName, barAttrName, barAttrValue)
			ps := fmt.Sprintf("%s.%s=%t", groupName, pifAttrName, pifAttrValue)
			Expect(buf.String()).To(
				And(
					ContainSubstring(fs),
					ContainSubstring(bs),
					ContainSubstring(ps),
				),
			)
		})
	})
})
