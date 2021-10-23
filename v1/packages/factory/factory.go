package factory

import (
	"hello-golang/v1/packages/gochi"
	"hello-golang/v1/packages/logrus"
	"hello-golang/v1/services/http/router"
	"hello-golang/v1/services/log"
)


func GetRouter() func() (router.Interface, error) {
	return gochi.GetRouter
}

func GetLog() func() (log.Interface, error) {
	return logrus.GetLog
}