```graphql
type Reservation {

}

type Query {
  reservations(cid: Int): [Reservation]
}

type Mutation {
  reserve(cid: Int, itemId: Int): Boolean
}
```
