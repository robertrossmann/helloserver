package backend

type Backend struct {
	PackageRoot string
}

func NewBackend() *Backend {
	return &Backend{}
}
