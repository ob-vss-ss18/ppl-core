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

query User {
  user(user: User): User
}
```
