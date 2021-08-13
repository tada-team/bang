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
//     func vars1example{{ $k }}() {
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
//     func vars2example{{ $k }}() {
//       print("{{ $v }}")
//     }
//   {{ end }}

//go:generate bang -verbose -template=example.tpl -dest=example_generated_3.go $GOFILE:$GOLINE
// vars:
//   package: main
//   prefix: vars3example
//   items:
//    - a
//    - b
//    - c

//go:generate bang -template=example.tpl -vars=vars.yaml -dest=example_generated_4.go $GOFILE:$GOLINE
