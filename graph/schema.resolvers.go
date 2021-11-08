package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lakshya.singhal/gqlgen-pokemon/graph/generated"
	"github.com/lakshya.singhal/gqlgen-pokemon/graph/model"
)


func (r *mutationResolver) CreateBattle(ctx context.Context, input model.NewBattle) (*model.Battle, error) {
	var newBattle = model.Battle{
		PokemonID: input.PokemonID,
		Opponent:  input.Opponent,
	}
	BattleData = append(BattleData, &newBattle)
	return &newBattle, nil
}

func (r *mutationResolver) DeleteBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}


func (r *queryResolver) GetBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListBattle(ctx context.Context) ([]*model.Battle, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

var PokemonData []*model.Pokemon
var BattleData []*model.Battle

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	newPokemon := model.Pokemon{
		Name: input.Name,
		Power: input.Power,
		//Description: input.Date,
	}
	PokemonData = append(PokemonData, &newPokemon)
	return &newPokemon, nil
}
func (r *queryResolver) ListPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	return PokemonData,nil
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented"))
}