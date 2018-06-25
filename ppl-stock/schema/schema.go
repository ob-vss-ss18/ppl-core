package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/ob-vss-ss18/ppl-stock/models"
)

/*
 * This is the only exported type from this file. It defines the
 * schema of our api.
 */
var PPLStockSchema graphql.Schema

var exampleSki = models.Ski{
	Id:           10,
	Usage:        models.Langlauf,
	Category:     models.Beginner,
	Usertype:     models.Erwachsener,
	Gender:       models.Male,
	Manufacturer: "Fischer",
	Model:        "Super Ski 3000",
	Length:       2,
	Bodyheight:   3,
	Bodyweight:   4,
	Color:        "Rot",
	PriceNew:     34.99,
	Condition:    models.New,
	Status:       models.Available,
}

var exampleStick = models.Stick{
	Id:           5,
	Usage:        models.All_Mountain,
	Usertype:     models.Kinder,
	Gender:       models.Uni,
	Manufacturer: "Müller",
	Model:        "Super Stick 7000",
	Length:       25,
	Bodyheight:   13,
	GripKind:     "sticky",
	Color:        "Blau",
	PriceNew:     14.99,
	Condition:    models.Used,
	Status:       models.Available,
}

/*
 * Automatically called by go to initialize variables defined here.
 */
func init() {
	/*
	 * The query type which will define the fields which can be requested
	 * from this api. This is basically the definition of the api.
	 */
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				/*
				 * If someone requests field "ski" an id must be provided with the
				 * request. The requested subselections will be resolved using the
				 * definitions is "skiType" and the input param which is given
				 * by the resolve function.
				 *
				 * This could be a good spot to load the data from the database into
				 * a local struct. That is why there is a method called "loadSkiFrom
				 * Database".
				 */
				"Ski": &graphql.Field{
					Type: skiType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Description: "Id of the requested ski.",
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
						return loadSkiFromDatabase(parameter.Args["id"].(int)), nil
					},
				},

				"Stick": &graphql.Field{
					Type: stickType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Description: "Id of the requested stick.",
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
						return loadStickFromDatabase(parameter.Args["id"].(int)), nil
					},
				},
			},
		})

	/*
	 * Now create a new schema which has the hello world query created above as
	 * top level query.
	 */
	PPLStockSchema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

}

/*
 * Dummy method this could load a ski given by id from a database.
 */
func loadSkiFromDatabase(id int) models.Ski {
	result := models.Ski{}
	if exampleSki.Id == uint32(id) {
		result = exampleSki
	}
	return result
}

func loadStickFromDatabase(id int) models.Stick {
	result := models.Stick{}
	if exampleStick.Id == uint32(id) {
		result = exampleStick
	}
	return result
}

