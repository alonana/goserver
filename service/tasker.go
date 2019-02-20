package service

import (
	stderrors "errors"
	"github.com/alonana/goserver/logging"
	"github.com/go-errors/errors"
	"math/rand"
)

func Start() (string, error) {
	if rand.Intn(2) == 0 {
		return "", stderrors.New("unexpected")
	}
	return "Started", nil
}

func Stop() (string, error) {
	logging.AppLogger.Error(errors.New("YOYO"))
	if rand.Intn(2) == 0 {
		return "", errors.New("unexpected")
	}
	return "Stopped", nil
}
