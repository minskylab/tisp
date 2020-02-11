package statemanager

import (
	p "github.com/minskylab/tisp/repository"
)

const DefaultPrismaEndpoint = "http://localhost:4466"
const DeafaultPrismaSecret = "mysecret42"

type StateManager struct {
	client *p.Client
}

func NewStateManager(endpoint, secret string) (*StateManager, error) {
	return &StateManager{
		client: p.New(&p.Options{
			Endpoint: endpoint,
			Secret:   secret,
		}),
	}, nil
}

