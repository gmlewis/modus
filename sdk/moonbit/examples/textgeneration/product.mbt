// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| The Product struct and the sample product will be used in the some of the examples.
pub(all) struct Product {
  id : String?
  name : String
  price : Double
  description : String
} derive(Show, Eq, FromJson, ToJson)

///|
let sample_product : Product = {
  id: Some("123"),
  name: "Shoes",
  price: 50.0,
  description: "Great shoes for walking.",
}

// since we'll be using this sample data often, we'll serialize it once and store it as a string.
///|
let sample_product_json : String = @json.stringify(sample_product.to_json())
