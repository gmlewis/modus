// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
priv suberror AuthError String derive(Show, Eq)

///| get_jwt_claims returns the JWT claims from the environment.
pub fn[T : @json.FromJson] get_jwt_claims() -> T raise Error {
  let claims_str = @wasi.environ_get("CLAIMS")
  guard claims_str != "" else { raise AuthError("JWT claims not found") }
  get_jwt_claims_from_string(claims_str)
}

///| get_jwt_claims_from_string returns the JWT claims from the given string.
pub fn[T : @json.FromJson] get_jwt_claims_from_string(
  claims_str : String
) -> T raise Error {
  let claims_json = @json.parse(claims_str)
  let claims : T = @json.from_json(claims_json)
  claims
}
