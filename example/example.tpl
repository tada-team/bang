package {{ .package }}

{{ range $k, $v := .items }}
    func f3{{ $k }}() {
        print("{{ $v }}")
    }
{{ end }}
