package service

import (
	"errors"
	"math/rand"
)

func Start() (string, error) {
	if rand.Intn(2) == 0 {
		return "", errors.New("unexpected")
	}
	return "Started", nil
}

func Stop() (string, error) {
	if rand.Intn(2) == 0 {
		return "", errors.New("unexpected")
	}
	return "Stopped", nil
}
