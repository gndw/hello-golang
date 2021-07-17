package logrus

import (
	"hello-golang/v1/services/log"

	"github.com/sirupsen/logrus"
)

type Instance struct {
}

func GetLog() (log.Interface, error) {
	ins := &Instance{}

	// TODO this is dev only ya
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
	})

	return ins, nil
}

func (x *Instance) Printf(str string, args ...interface{}) {
	logrus.Debugf(str, args...)
}

func (x *Instance) Tracef(str string, args ...interface{}) {
	logrus.Tracef(str, args...)
}

func (x *Instance) Traceln(args ...interface{}) {
	logrus.Traceln(args...)
}

func (x *Instance) Debugf(str string, args ...interface{}) {
	logrus.Debugf(str, args...)
}

func (x *Instance) Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

func (x *Instance) Infof(str string, args ...interface{}) {
	logrus.Infof(str, args...)
}

func (x *Instance) Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

func (x *Instance) Warningf(str string, args ...interface{}) {
	logrus.Warningf(str, args...)
}

func (x *Instance) Warningln(args ...interface{}) {
	logrus.Warningln(args...)
}

func (x *Instance) Errorf(str string, args ...interface{}) {
	logrus.Errorf(str, args...)
}

func (x *Instance) Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}