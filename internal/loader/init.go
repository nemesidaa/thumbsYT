package loader

type Loader struct {
	// Scalability #2
	Resolution string
	SubDir     string
}

func NewLoader() *Loader {
	return &Loader{
		Resolution: "hqdefault",
		SubDir:     "tests",
	}
}
