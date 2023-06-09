package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"JFT/transport/gql/graph/model"
	"context"
	"fmt"
)

// CreateNft is the resolver for the createNFT field.
func (r *mutationResolver) CreateNft(ctx context.Context, input model.CreateNft) (*model.Nft, error) {
	panic(fmt.Errorf("not implemented: CreateNft - createNFT"))
}

// DeleteNft is the resolver for the deleteNFT field.
func (r *mutationResolver) DeleteNft(ctx context.Context, input model.GetNft) (*model.Nft, error) {
	panic(fmt.Errorf("not implemented: DeleteNft - deleteNFT"))
}

// UpdateNft is the resolver for the updateNFT field.
func (r *mutationResolver) UpdateNft(ctx context.Context, input model.UpdateNft) (*model.Nft, error) {
	panic(fmt.Errorf("not implemented: UpdateNft - updateNFT"))
}

// Nft is the resolver for the nft field.
func (r *queryResolver) Nft(ctx context.Context, input model.GetNft) (*model.Nft, error) {
	panic(fmt.Errorf("not implemented: Nft - nft"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
