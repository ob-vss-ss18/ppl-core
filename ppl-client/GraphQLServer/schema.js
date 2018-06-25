// schema.js

const schema = `
# declare custom scalars for date as GQDate
scalar GQDate

# customer type
type Customer {
    id: ID!
    firstName: String
    lastName: String
    dob: GQDate
    email: String
    street: String
    housenumber: String
	postalcode: String
	city: String
}

# ski type
type Ski {
    id: ID!
    manufacturer: String!
    modell: String!
    length: Int
    bodyheight: Int
    bodyweight: Int
    price_new: Float!
    availability: Boolean
  }

type Query {
    # Return a customer by id
    Customer(id: ID!): Customer
    # Return all customers
    Customers(limit: Int): [Customer]
    # Return a customer by id
    Ski(id: ID!): Ski
    # Return all customers
    Skis(limit: Int): [Ski]
}

type Mutation {
    # Create a customer
    createCustomer (firstName: String,lastName: String, dob: GQDate, email: String, street: String, housenumber: String, postalcode: String, city: String): Customer
    # Update a customer
    updateCustomer (id: ID!, firstName: String,lastName: String, dob: GQDate, email: String, street: String, housenumber: String, postalcode: String, city: String): Customer
    # Delete a customer
    deleteCustomer(id: ID!): Customer

    # Create a ski
    createSki (manufacturer: String, modell: String, length: Int, bodyheight: Int, bodyweight: Int, price_new: Float, availability: Boolean): Ski
    # Update a ski
    updateSki (id: ID!, manufacturer: String, modell: String, length: Int, bodyheight: Int, bodyweight: Int, price_new: Float, availability: Boolean): Ski
    # Delete a ski
    deleteSki(id: ID!): Ski
}
`;

module.exports.Schema = schema;