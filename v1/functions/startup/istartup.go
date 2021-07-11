package startup

import "github.com/gndw/hello-golang/v1/functions/stderror"

type iStartup interface {
	Start() *stderror.Error
}