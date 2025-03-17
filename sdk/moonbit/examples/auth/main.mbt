// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| This is a sample struct that represents the claims in a JWT token.
pub struct ExampleClaims {
  exp : Int64
  iat : Int64
  iss : String
  jti : String
  nbf : Int64
  sub : String
  userId : String
} derive(Show, Eq, FromJson, ToJson)

///| GetJWTClaims is a function that returns the claims in a JWT token,
/// and parses them into the Claims struct.
pub fn get_jwt_claims() -> ExampleClaims!Error {
  @auth.get_jwt_claims!()
}
