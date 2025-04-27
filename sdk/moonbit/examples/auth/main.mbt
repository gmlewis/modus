// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| This is a sample struct that represents the claims in a JWT token.
pub(all) struct ExampleClaims {
  exp : Int64
  iat : Int64
  iss : String
  jti : String
  nbf : Int64
  sub : String
  userId : String
} derive(Show, Eq)

///| GetJWTClaims is a function that returns the claims in a JWT token,
/// and parses them into the Claims struct.
pub fn get_jwt_claims() -> ExampleClaims!Error {
  @auth.get_jwt_claims!()
}

///|
pub impl ToJson for ExampleClaims with to_json(self) {
  {
    "exp": self.exp.to_double().to_json(),
    "iat": self.iat.to_double().to_json(),
    "iss": self.iss.to_json(),
    "jti": self.jti.to_json(),
    "nbf": self.nbf.to_double().to_json(),
    "sub": self.sub.to_json(),
    "user-id": self.userId.to_json(),
  }.to_json()
}

///|
pub impl @json.FromJson for ExampleClaims with from_json(json, path) {
  guard json is Object(obj) else {
    raise @json.JsonDecodeError((path, "Expected object"))
  }
  guard obj["exp"] is Some(Number(exp)) else {
    raise @json.JsonDecodeError((path, "Expected exp to be Int64"))
  }
  let exp = exp.to_int64()
  guard obj["iat"] is Some(Number(iat)) else {
    raise @json.JsonDecodeError((path, "Expected iat to be Int64"))
  }
  let iat = iat.to_int64()
  guard obj["iss"] is Some(String(iss)) else {
    raise @json.JsonDecodeError((path, "Expected iss to be String"))
  }
  guard obj["jti"] is Some(String(jti)) else {
    raise @json.JsonDecodeError((path, "Expected jti to be String"))
  }
  guard obj["nbf"] is Some(Number(nbf)) else {
    raise @json.JsonDecodeError((path, "Expected nbf to be Int64"))
  }
  let nbf = nbf.to_int64()
  guard obj["sub"] is Some(String(sub)) else {
    raise @json.JsonDecodeError((path, "Expected sub to be String"))
  }
  guard obj["user-id"] is Some(String(userId)) else {
    raise @json.JsonDecodeError((path, "Expected user-id to be String"))
  }
  { exp, iat, iss, jti, nbf, sub, userId }
}
