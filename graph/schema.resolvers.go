package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/phanikumarps/graph/graph/model"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	mockAcc := model.Account{
		ID: input.ID,
		Name: &model.Name{
			First: input.Name.First,
			Last:  input.Name.Last,
		},
		ShippingAddress: &model.Address{
			Country:  input.ShippingAddress.Country,
			Street:   input.ShippingAddress.Street,
			State:    input.ShippingAddress.State,
			Zip:      input.ShippingAddress.Zip,
			Building: input.ShippingAddress.Building,
		},
		CreditCard: &model.CreditCard{
			Number:         input.CreditCard.Number,
			Pin:            input.CreditCard.Pin,
			ExpirationDate: input.CreditCard.ExpirationDate,
		},
	}
	Accounts[input.ID] = &mockAcc
	return &mockAcc, nil
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	if acc, ok := Accounts[id]; ok {
		return acc, nil
	} else {
		return nil, nil
	}
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, limit *int) ([]*model.Account, error) {
	accArray := make([]*model.Account, 0, len(Accounts))
	for _, v := range Accounts {
		accArray = append(accArray, v)
	}
	l := *limit
	if l > len(accArray) {
		l = len(accArray)
	}

	return accArray[:l], nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
