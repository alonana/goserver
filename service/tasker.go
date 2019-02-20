package service

import (
	stderrors "errors"
	"fmt"
	"github.com/go-errors/errors"
	"math/rand"
	"time"
)

const poolSize = 5

type jobRequest struct {
	time     string
	response chan string
}

var sendWorkChannel = make(chan jobRequest, poolSize)

func init() {
	for i := 0; i < poolSize; i++ {
		go poolThread(i)
	}
}

func poolThread(i int) {
	for {
		select {
		case s := <-sendWorkChannel:
			s.response <- fmt.Sprintf("work %v got %v\n", i, s.time)
		}
	}
}

func Start() (string, error) {
	if rand.Intn(2) == 0 {
		return "", stderrors.New("unexpected")
	}

	req := jobRequest{
		response: make(chan string),
		time:     time.Now().String(),
	}
	sendWorkChannel <- req

	return <-req.response, nil
}

func Stop() (string, error) {
	if rand.Intn(2) == 0 {
		return "", errors.New("unexpected")
	}
	return "Stopped", nil
}
