```graphql
type User {
  uid: ID!
  email: String!
  role: UserRole!
  token: String
}

enum UserRole {
  OWNER
  STAFF
  CUSTOMER
}

type Query {
  user(user: User): User
}

type Mutation {
  requestToken(email: String!): Boolean
  loginPwd(email: String!, password: String!): User
  loginToken(email: String!, token: String): User
}
```
