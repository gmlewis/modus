///|
test "@auth.ExampleClaims/json_serialization" {
  let claims : @auth_example.ExampleClaims = {
    exp: 1735689600L, // 2025-01-01
    iat: 1672531200L, // 2023-01-01
    iss: "test-issuer",
    jti: "test-jti",
    nbf: 1672531200L, // 2023-01-01
    sub: "test-subject",
    userId: "test-user-id",
  }
  let decoded : ExampleClaims = @json.from_json!(claims.to_json())
  inspect!(
    decoded,
    content=
      #|{exp: 1735689600, iat: 1672531200, iss: "test-issuer", jti: "test-jti", nbf: 1672531200, sub: "test-subject", userId: "test-user-id"}
    ,
  )
}

///|
test "parse JSON claims" {
  let response =
    #|{
    #|  "exp": 1742237709,
    #|  "iat": 1742235909,
    #|  "jti": "test-jti",
    #|  "iss": "test-issuer",
    #|  "nbf": 1672531200,
    #|  "sub": "test-sub",
    #|  "user-id": "test-user-id"
    #|}
  // The following line fails with this error message:
  // test gmlewis/modus/examples/auth_example/main_test.mbt::parse JSON claims failed: JsonDecodeError(($.exp, "Int64::from_json: expected number in string representation"))
  // See: https://github.com/moonbitlang/core/pull/1417
  let json = @json.parse!(response)
  let decoded : ExampleClaims = @json.from_json!(json)
  inspect!(
    decoded,
    content=
      #|{exp: 1742237709, iat: 1742235909, iss: "test-issuer", jti: "test-jti", nbf: 1672531200, sub: "test-sub", userId: "test-user-id"}
    ,
  )
}
