# slogctx

[![Test Status][tests-badge]][tests-link]
[![Coverage Status][cov-badge]][cov-link]
[![License][license-badge]][license-link]
[![Go Reference][docs-badge]][docs-link]

Package slogctx works with the [slog] structed logging package to add information from a [context.Context] as attributes.

## Installation

```sh
go get -u github.com/pfflabs/slogctx
```

---

Released under the [Apache 2.0 License].

[tests-badge]: https://github.com/pfflabs/slogctx/actions/workflows/run-tests.yml/badge.svg?branch=main
[tests-link]: https://github.com/pfflabs/slogctx/actions/workflows/run-tests.yml
[cov-badge]: https://coveralls.io/repos/github/pfflabs/slogctx/badge.svg?branch=main
[cov-link]: https://coveralls.io/github/pfflabs/slogctx?branch=main
[license-badge]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[license-link]: https://github.com/pfflabs/slogctx/blob/main/LICENSE
[docs-badge]: https://pkg.go.dev/badge/github.com/pfflabs/slogctx.svg
[docs-link]: https://pkg.go.dev/github.com/pfflabs/slogctx
[slog]: https://pkg.go.dev/log/slog
[context.Context]: https://pkg.go.dev/context#Context
[Apache 2.0 License]: https://www.apache.org/licenses/LICENSE-2.0
