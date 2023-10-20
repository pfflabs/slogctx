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
	"reflect"
	"time"
)

type typeAttrGetter[T any] struct {
	key      string
	lookup   func(context.Context) (T, bool)
	makeAttr func(string, T) slog.Attr
}

func (l *typeAttrGetter[T]) GetAttrs(ctx context.Context) []slog.Attr {
	t, ok := l.lookup(ctx)
	if !ok {
		return nil
	}

	return []slog.Attr{l.makeAttr(l.key, t)}
}

// Attr returns an [AttrGetter] that takes its value from a [context.Context].
//
// Automatically recognizes types
//
//	any
//	bool
//	time.Duration
//	float64
//	int
//	int64
//	string
//	time.Time
//	uint64
//
// Unrecognized types are treated as type any.
//
// Panics if key is empty or lookup is nil.
func Attr[T any](
	key string,
	lookup func(context.Context) (value T, ok bool),
) AttrGetter {
	validateKey(key)
	if lookup == nil {
		panic("lookup is nil")
	}

	switch l := any(lookup).(type) {

	case func(context.Context) (bool, bool):
		return &typeAttrGetter[bool]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Bool,
		}

	case func(context.Context) (time.Duration, bool):
		return &typeAttrGetter[time.Duration]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Duration,
		}

	case func(context.Context) (float64, bool):
		return &typeAttrGetter[float64]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Float64,
		}

	case func(context.Context) (int, bool):
		return &typeAttrGetter[int]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Int,
		}

	case func(context.Context) (int64, bool):
		return &typeAttrGetter[int64]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Int64,
		}

	case func(context.Context) (string, bool):
		return &typeAttrGetter[string]{
			key:      key,
			lookup:   l,
			makeAttr: slog.String,
		}

	case func(context.Context) (time.Time, bool):
		return &typeAttrGetter[time.Time]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Time,
		}

	case func(context.Context) (uint64, bool):
		return &typeAttrGetter[uint64]{
			key:      key,
			lookup:   l,
			makeAttr: slog.Uint64,
		}

	default:
		return &typeAttrGetter[any]{
			key: key,
			lookup: func(ctx context.Context) (any, bool) {
				t, ok := lookup(ctx)
				return any(t), ok
			},
			makeAttr: slog.Any,
		}
	}
}

// IgnoreZero returns a function suitable for [Attr] that will return a
// false status if the value returned by lookup is the type's 'zero' value.
func IgnoreZero[T any](
	lookup func(context.Context) T,
) func(context.Context) (value T, ok bool) {
	if lookup == nil {
		panic("lookup is nil")
	}

	return func(ctx context.Context) (T, bool) {
		t := lookup(ctx)
		return t, !reflect.ValueOf(t).IsZero()
	}
}

// LogZero returns a function suitable for [Attr] that will always return
// a true status. Useful if the type's 'zero' value should be included in
// logged output.
func LogZero[T any](
	lookup func(context.Context) T,
) func(context.Context) (value T, ok bool) {
	if lookup == nil {
		panic("lookup is nil")
	}

	return func(ctx context.Context) (T, bool) {
		return lookup(ctx), true
	}
}
