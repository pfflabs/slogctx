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

type groupAttrGetter struct {
	key string
	AttrGetter
}

func (g *groupAttrGetter) GetAttrs(ctx context.Context) []slog.Attr {
	attrs := g.AttrGetter.GetAttrs(ctx)
	if len(attrs) == 0 {
		return nil
	}

	return []slog.Attr{
		{Key: g.key, Value: slog.GroupValue(attrs...)},
	}
}

// Group returns a [AttrGetter] that groups one or more [AttrGetter]
// instances.
//
// Panics if key is empty, receives zero [AttrGetter] instances, or any
// [AttrGetter] references are nil.
func Group(key string, attrGetters ...AttrGetter) AttrGetter {
	validateKey(key)

	return &groupAttrGetter{
		key:        key,
		AttrGetter: concat(attrGetters),
	}
}

func validateKey(key string) {
	if len(key) > 0 {
		return
	}
	panic("key is empty")
}
