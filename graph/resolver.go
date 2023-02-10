package graph

import "github.com/phanikumarps/graph/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
}

var Accounts = make(map[string]*model.Account)
