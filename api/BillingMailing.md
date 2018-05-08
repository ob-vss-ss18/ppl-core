```graphql
type Invoice {

}

type Query {
  invoices(id: Int): [Invoice]
}

type Mutation {
  createInvoice(invoice: Invoice!): Boolean
}

```
