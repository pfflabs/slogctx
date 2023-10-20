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

// # Quickstart
//
//	import (
//		// ...
//		"log/slog"
//		"github.com/pfflabs/slogctx"
//	)
//
//	// ...
//
//	h := slog.NewTextHandler(os.Stdout, nil)
//	h = slogctx.NewHandler(h,
//		// apkg.FromCtx is a func(context.Context) (string, bool)
//		slogctx.Attr("a", apkg.FromCtx),
//		// bpkg.FromCtx is a func(context.Context) (uint64, bool)
//		slogctx.Attr("b", bpkg.FromCtx),
//	)
//	log := slog.New(h)
//
// # Usage
//
// The [Handler] wraps a [log/slog.Handler] and simply extract attributes from
// a [context.Context] and passes those attributes to the wrapped handler for
// formatting and output.
//
// Use [Attr] to create [AttrGetter] instances from a function that returns a
// value and a bool which indicates whether the value was found in a
// [context.Context].
//
//	// cpkg.FromCtx is a func(context.Context) (time.Time, bool)
//	g := slogctx.Attr("c", cpkg.FromCtx),
//
// [Attr] can also be used with functions that do not return a bool to indicate
// whether a value was found in a [context.Context] by also using the
// [LogZero] and [IgnoreZero] functions. These functions configure whether a
// value's 'zero value' should be included in the logged output or omitted,
// respectively.
//
//	// dpkg.FromCtx is a func(context.Context) float64
//	g := slogctx.Attr("d", slogctx.IgnoreZero(dpkg.FromCtx))
//	// or
//	g := slogctx.Attr("d", slogctx.LogZero(dpkg.FromCtx))
//
// Use [Group] to assign all retrieved attributes to a named group.
//
//	g := slogctx.Group("config",
//		slogctx.Attr("hostname", clientpkg.HostnameFromCtx),
//		slogctx.Attr("port", clientpkg.PortFromCtx),
//		slogctx.Attr("timeout", clientpkg.TimeoutFromCtx),
//	)
//	// Will log attributes as "config.hostname", etc. or however the target
//	// handler formats grouped attributes.
//
// Use [AttrGetterFunc] to customize the display of types that are not
// automatically recognized by [Attr] (e.g. user defined type).
//
//	g := slogctx.Group("browser",
//		slogctx.AttrGetterFunc(func(ctx context.Context) []slog.Attr {
//			browser, ok := browserpkg.FromCtx(ctx)
//			if !ok {
//				return nil
//			}
//
//			return []slog.Attr{
//				slog.String("user-agent", browser.UserAgent()),
//				slog.Bool("js", browser.JSEnabled()),
//				// ...
//			}
//		})
package slogctx
