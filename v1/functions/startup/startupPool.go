package startup

import "github.com/gndw/hello-golang/v1/functions/stderror"

type Startup struct {
	pools []iStartup
}

func (s *Startup) Start() (err *stderror.Error) {
	for _, s := range s.pools {
		err = s.Start()
		if (err != nil) {
			return err
		}
	}
	return nil
}