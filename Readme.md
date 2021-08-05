## Bang

Sugar for `go generate`. Reads yml data from .go source file and render template with variables.

Install:
```bash
go install github.com/tada-team/bang
```

Sample usage:

```go
package main

//go:generate bang $GOFILE:$GOLINE
// dest: main_generated.go
// vars:
//   package: main
//   items:
//    - foo
//    - bar
// template: >
//   package {{.package}}
//
//   {{ range $item := .items }}
//       func f{{ $item }}() {
//           print("{{ $item }}")
//       }
//   {{ end }}
//
```
