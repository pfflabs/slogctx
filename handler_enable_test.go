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
	"os"

	"github.com/pfflabs/slogctx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Querying the handler's enabled status", func() {
	targetOpts := &slog.HandlerOptions{Level: slog.LevelInfo}
	levelAboveMin := slog.LevelWarn
	levelBelowMin := slog.LevelDebug
	target := slog.NewTextHandler(os.Stderr, targetOpts)
	handler := slogctx.NewHandler(target, noopGetter)
	ctx := context.Background()

	When("the query level meets the target handler's minimum level", func() {
		level := levelAboveMin

		It("returns that the handler is enabled", func() {
			Expect(handler.Enabled(ctx, level)).To(BeTrue())
		})
	})

	When("the query level does not meet the target handler's minimum level", func() {
		level := levelBelowMin

		It("returns that the handler is disabled", func() {
			Expect(handler.Enabled(ctx, level)).To(BeFalse())
		})
	})
})
