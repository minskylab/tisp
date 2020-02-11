package statemanager

import (
	p "github.com/minskylab/tisp/repository"
)

type StateManager struct {
	client *p.Client
}

func NewStateManager() (*StateManager, error) {
	return &StateManager{
		client: p.New(&p.Options{
			Endpoint: "http://localhost:4466",
			Secret:   "mysecret42",
		}),
	}, nil
}

