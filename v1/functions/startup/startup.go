package startup

import "hello-golang/v1/functions/stderror"

type Startup struct {
	pools []IStartup
}

func (s *Startup) Start() (err *stderror.Error) {
	for _, s := range s.pools {
		err = s.Startup()
		if (err != nil) {
			return err
		}
	}
	return
}

func (s *Startup) Register(obj IStartup) (err *stderror.Error) {
	s.pools = append(s.pools, obj)
	return
}