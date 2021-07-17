package log

type Interface interface {
	Printf(string, ...interface{})
	Tracef(string, ...interface{})
	Traceln(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
	Warningf(string, ...interface{})
	Warningln(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
}