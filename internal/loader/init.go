package loader

type Loader struct {
	// Scalability #2
	Resolution string
}

func NewLoader(res string) *Loader {
	return &Loader{
		Resolution: res,
	}
}

type Resolution string

var AllAvailableResolutions map[Resolution]struct{} = map[Resolution]struct{}{
	"sddefault":     {},
	"hqdefault":     {},
	"mqdefault":     {},
	"lqdefault":     {},
	"maxresdefault": {},
}
