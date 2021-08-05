package main

//go:generate bang $GOFILE:$GOLINE
// dest: example_generated.go
// vars:
//   package: main
//   items:
//    - a
//    - b
//    - c
// template: >
//   package {{ .package }}
//
//   {{ range $k, $v := .items }}
//     func f{{ $k }}() {
//       print("{{ $v }}")
//     }
//   {{ end }}

//go:generate bang -dest=example_generated_2.go $GOFILE:$GOLINE
// vars:
//   package: main
//   items:
//    - a
//    - b
//    - c
// template: >
//   package {{.package}}
//
//   {{ range $k, $v := .items }}
//     func f2{{ $k }}() {
//       print("{{ $v }}")
//     }
//   {{ end }}

func main() {
	// noop
}
