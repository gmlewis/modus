// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
const MODEL_NAME = "minilm"

///|
const MY_PRODUCTS = "myProducts"

///|
let search_methods : Array[String] = ["searchMethod1", "searchMethod2"]

///|
pub fn embed(text : Array[String]) -> Array[Array[Float]]!Error {
  let model : @experimental.EmbeddingsModel = @models.get_model!(MODEL_NAME)
  let input = model.create_input!(text)
  let output = model.invoke!(input)
  output.predictions
}

///|
pub fn add_product(description : String) -> Array[String]!Error {
  let res = @collections.upsert!(MY_PRODUCTS, description)
  res.keys
}

///|
pub fn add_products(descriptions : Array[String]) -> Array[String]!Error {
  let res = @collections.upsert_batch!(MY_PRODUCTS, descriptions)
  res.keys
}

///|
pub fn delete_product(key : String) -> String!Error {
  let res = @collections.remove!(MY_PRODUCTS, key)
  res.status
}

///|
pub fn get_product(key : String) -> String!Error {
  @collections.get_text!(MY_PRODUCTS, key)
}

///|
pub fn get_products() -> Map[String, String]!Error {
  @collections.get_texts!(MY_PRODUCTS)
}

///|
pub fn get_product_vector(key : String) -> Array[Float]!Error {
  @collections.get_vector!(MY_PRODUCTS, search_methods[0], key)
}

///|
pub fn get_namespaces() -> Array[String]!Error {
  @collections.get_namespaces!(MY_PRODUCTS)
}

///|
pub fn get_labels(key : String) -> Array[String]!Error {
  @collections.get_labels!(MY_PRODUCTS, key)
}

///|
pub fn recompute_search_methods() -> String!Error {
  for method_ in search_methods {
    let res = @collections.recompute_search_method!(MY_PRODUCTS, method_)
    if res.status != "success" {
      return res.status
    }
  }
  "success"
}

///|
pub fn compute_distance_between_products(
  key1 : String,
  key2 : String
) -> Double!Error {
  let res = @collections.compute_distance!(
    MY_PRODUCTS,
    search_methods[0],
    key1,
    key2,
  )
  res.distance
}

///|
pub fn search_products(
  product_description : String,
  max_items : Int
) -> @collections.CollectionSearchResult!Error {
  @collections.search!(MY_PRODUCTS, search_methods[0], product_description, opts=[
    @collections.with_limit(max_items),
    @collections.with_return_text(true),
  ])
}

///|
pub fn search_products_with_key(
  product_key : String,
  max_items : Int
) -> @collections.CollectionSearchResult!Error {
  let vec = get_product_vector!(product_key)
  @collections.search_by_vector!(MY_PRODUCTS, search_methods[0], vec, opts=[
    @collections.with_limit(max_items),
    @collections.with_return_text(true),
  ])
}
