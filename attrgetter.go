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

package slogctx

import (
	"context"
	"log/slog"
)

type (
	// AttrGetter returns one, or more, [log/slog.Attr] instances using
	// information from a [context.Context].
	AttrGetter interface {
		GetAttrs(ctx context.Context) []slog.Attr
	}

	// AttrGetterFunc is a [AttrGetter] implemented as a single function.
	AttrGetterFunc func(context.Context) []slog.Attr
)

// GetAttrs returns the [log/slog.Attr] instances from the backing
// function.
func (f AttrGetterFunc) GetAttrs(ctx context.Context) []slog.Attr {
	return f(ctx)
}
