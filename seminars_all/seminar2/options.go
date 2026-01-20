package main

import (
	"fmt"
	"time"
)

func main() {

}

type options struct {
	port int
	host string
	timeout time.Duration
	optinsmaxConn int
}

type Option func(opts *options) error

var (
	defaultPort = 8080
	defaultTimeout = 1 * time.Second
)

func New(opts ...Option) (*Server, error) {//"...Option" означает что мы можем передать как ничего так и несколько значений в функцию
	svr := &Server{}

	o := options{
		port: defaultPort,
		timeout: defaultTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return svr, err
		}
	}

	return svr, nil
}