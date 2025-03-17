// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| This is a sample struct that represents the claims in a JWT token.
pub struct ExampleClaims {
  exp : Int64 // `json:"exp"`
  iat : Int64 // `json:"iat"`
  iss : String // `json:"iss"`
  jti : String // `json:"jti"`
  nbf : Int64 // `json:"nbf"`
  sub : String // `json:"sub"`
  user_id : String // `json:"user-id"`
} derive(Show, Eq)

///|
pub impl ToJson for ExampleClaims with to_json(self) {
  {
    "exp": self.exp.to_json(),
    "iat": self.iat.to_json(),
    "iss": self.iss.to_json(),
    "jti": self.jti.to_json(),
    "nbf": self.nbf.to_json(),
    "sub": self.sub.to_json(),
    "user-id": self.user_id.to_json(),
  }
}

///|
pub impl @json.FromJson for ExampleClaims with from_json(self, _path) {
  let obj = self.as_object().unwrap()
  let exp : Int64 = @json.from_json!(obj.get("exp").unwrap())
  let iat : Int64 = @json.from_json!(obj.get("iat").unwrap())
  let iss : String = @json.from_json!(obj.get("iss").unwrap())
  let jti : String = @json.from_json!(obj.get("jti").unwrap())
  let nbf : Int64 = @json.from_json!(obj.get("nbf").unwrap())
  let sub : String = @json.from_json!(obj.get("sub").unwrap())
  let user_id : String = @json.from_json!(obj.get("user-id").unwrap())
  { exp, iat, iss, jti, nbf, sub, user_id }
}

///| GetJWTClaims is a function that returns the claims in a JWT token, and parses them into the Claims struct.
pub fn get_jwt_claims() -> ExampleClaims!Error {
  @auth.get_jwt_claims!()
}
