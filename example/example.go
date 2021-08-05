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
//   package {{.package}}
//
//   func main() {
//        println("123: {{.vars.items}}")
//   }
//
//   {{range $k, $v := .items}}
//       func f{{$k}}() {
//        print("{{$v}}")
//       }
//   {{end}}

func main() {
	// noop
}
