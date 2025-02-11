///|
struct TestClaims {
  sub : String
  name : String
  admin : Bool
} derive(Show, FromJson)

///|
test "get_jwt_claims_from_string" {
  let claims_str =
    #|{"sub":"1234567890","name":"John Doe","admin":true}
  let claims : TestClaims = get_jwt_claims_from_string!(claims_str)
  inspect!(
    claims,
    content=
      #|{sub: "1234567890", name: "John Doe", admin: true}
    ,
  )
}
