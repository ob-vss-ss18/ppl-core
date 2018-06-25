package schema

import (
"github.com/graphql-go/graphql"
"github.com/ob-vss-ss18/ppl-stock/models"
)

var stickType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Stick",
	Description: "A stick.",

	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The id of the stick",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Id, nil
				}
				return nil, nil
			},
		},
		"usage": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The use case of the stick",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Usage, nil
				}
				return nil, nil
			},
		},
		"usertype": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The usertype of the stick",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Usertype, nil
				}
				return nil, nil
			},
		},
		"gender": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The gender by which the stick is intended to be used.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Gender, nil
				}
				return nil, nil
			},
		},
		"manufacturer": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The manufacturer of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Manufacturer, nil
				}
				return nil, nil
			},
		},
		"modell": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The model of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Model, nil
				}
				return nil, nil
			},
		},
		"length": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The length of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Length, nil
				}
				return nil, nil
			},
		},
		"bodyheight": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The best bodyheight for using this stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Bodyheight, nil
				}
				return nil, nil
			},
		},
		"grip_kind": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The grip_kind of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.GripKind, nil
				}
				return nil, nil
			},
		},
		"color": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The color of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Color, nil
				}
				return nil, nil
			},
		},
		"price_new": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "The new price of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.PriceNew, nil
				}
				return nil, nil
			},
		},
		"condition": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The condition of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Condition, nil
				}
				return nil, nil
			},
		},
		"availability": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The status/availability of the stick.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if stick, ok := parameter.Source.(models.Stick); ok {
					return stick.Status, nil
				}
				return nil, nil
			},
		},
	},
})

