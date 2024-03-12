package generator

type Generator interface {
	Make()
	FileExists() bool
}

func Make(m Generator) {
	m.Make()
}
