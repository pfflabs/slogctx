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
	"fmt"
	"log/slog"
	"strings"
)

type concatAttrGetter []AttrGetter

func (c *concatAttrGetter) GetAttrs(ctx context.Context) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(*c))
	for _, g := range *c {
		attrs = append(attrs, g.GetAttrs(ctx)...)
	}

	return attrs
}

func concat(gs []AttrGetter) AttrGetter {
	n := len(gs)
	switch {
	case n == 0:
		panic("received 0 AttrGetters")

	case n == 1:
		g := gs[0]
		if g == nil {
			panic("AttrGetter is nil")
		}

		return g
	}

	validateAttrGetters(gs)

	c := concatAttrGetter(gs)
	return &c
}

func validateAttrGetters(gs []AttrGetter) {
	var (
		n       = len(gs)
		errMsgs []string
	)

	for i, g := range gs {
		if g != nil {
			continue
		}

		msg := fmt.Sprintf("AttrGetter %d of %d is nil", i+1, n)
		errMsgs = append(errMsgs, msg)
	}

	if len(errMsgs) == 0 {
		return
	}
	panic(strings.Join(errMsgs, ", "))
}
