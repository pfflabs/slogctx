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

// Handler wraps a target [log/slog.Handler] and extracts attributes from
// a [context.Context] and passes those attributes to the target handler
// for formatting and output.
type Handler struct {
	attrGetter AttrGetter
	target     slog.Handler
}

var _ slog.Handler = (*Handler)(nil)

// NewHandler returns a new Handler that will add attributes taken from the
// provided context then delegates handling to the target.
//
// Panics if target handler is nil, receives zero [AttrGetter] instances, or
// any [AttrGetter] references are nil.
func NewHandler(target slog.Handler, attrGetters ...AttrGetter) *Handler {
	if target == nil {
		panic("target is nil")
	}

	return &Handler{
		attrGetter: concat(attrGetters),
		target:     target,
	}
}

// Enabled returns whether the target handler is enabled for the context
// and level.
func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.target.Enabled(ctx, level)
}

// Handle delegates handling the record and any attributes gathered from the
// context to the target handler.
func (h *Handler) Handle(ctx context.Context, rec slog.Record) error {
	if attrs := h.attrGetter.GetAttrs(ctx); len(attrs) > 0 {
		// From https://pkg.go.dev/log/slog#hdr-Working_with_Records:
		// "Before modifying a Record, use Record.Clone to create a copy"
		rec = rec.Clone()
		rec.AddAttrs(attrs...)
	}

	return h.target.Handle(ctx, rec)
}

// WithAttrs returns a handler that will include the given attributes when
// handling records.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		attrGetter: h.attrGetter,
		target:     h.target.WithAttrs(attrs),
	}
}

// WithGroup returns a handler that will group attributes when handling
// records.
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		attrGetter: h.attrGetter,
		target:     h.target.WithGroup(name),
	}
}
