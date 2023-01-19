package api

type srvApi struct {
}

func NewAPI() *srvApi {
	return &srvApi{}
}

// Make sure we conform to ServerInterface
var _ ServerInterface = (*srvApi)(nil)
