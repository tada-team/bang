package {{ .package }}

{{ range $k, $v := .items }}
    func {{ $.prefix }}{{ $k }}() {
        print("{{ $v }}")
    }
{{ end }}
