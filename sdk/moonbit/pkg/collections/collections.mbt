// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub typealias CollectionStatus = @host.CollectionStatus

// ///|
// pub(all) enum CollectionStatus {
//   Success
//   Error
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias CollectionMutationResult = @host.CollectionMutationResult

// ///|
// pub(all) struct CollectionMutationResult {
//   collection : String
//   status : CollectionStatus
//   error : String
//   operation : String
//   keys : Array[String]
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias SearchMethodMutationResult = @host.SearchMethodMutationResult

// ///|
// pub(all) struct SearchMethodMutationResult {
//   collection : String
//   status : CollectionStatus
//   error : String
//   operation : String
//   search_method : String
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias CollectionSearchResult = @host.CollectionSearchResult

// ///|
// pub(all) struct CollectionSearchResult {
//   collection : String
//   status : CollectionStatus
//   error : String
//   search_method : String
//   objects : Array[CollectionSearchResultObject]
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias CollectionSearchResultObject = @host.CollectionSearchResultObject

// ///|
// pub(all) struct CollectionSearchResultObject {
//   namespace : String
//   key : String
//   text : String
//   labels : Array[String]
//   distance : Double
//   score : Double
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias CollectionClassificationResult = @host.CollectionClassificationResult

// ///|
// pub(all) struct CollectionClassificationResult {
//   collection : String
//   status : String
//   error : String
//   search_method : String
//   labels_result : Array[CollectionClassificationLabelObject]
//   cluster : Array[CollectionClassificationResultObject]
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias CollectionClassificationLabelObject = @host.CollectionClassificationLabelObject

// ///|
// pub(all) struct CollectionClassificationLabelObject {
//   label : String
//   confidence : Double
// } derive(Show, Eq, FromJson, ToJson)

///|
pub typealias CollectionClassificationResultObject = @host.CollectionClassificationResultObject

// ///|
// pub(all) struct CollectionClassificationResultObject {
//   key : String
//   labels : Array[String]
//   distance : Double
//   score : Double
// } derive(Show, Eq, FromJson, ToJson)

///|
pub(all) type NamespaceOption (NamespaceOptions) -> Unit

///|
pub(all) struct NamespaceOptions {
  mut namespace : String
}

///|
pub fn with_namespace(namespace : String) -> NamespaceOption {
  fn(o) { o.namespace = namespace }
}

///|
pub(all) type! CollectionError String derive(Show, Eq)

///|
pub fn upsert_batch(
  collection : String,
  keys : Array[String],
  texts : Array[String],
  labels_arr : Array[Array[String]],
  opts : Array[NamespaceOption]
) -> CollectionMutationResult!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if texts.length() == 0 {
    raise CollectionError("texts is empty")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.upsert(
    collection,
    ns_opts.namespace,
    keys,
    texts,
    labels_arr,
  )
  if result.is_empty() {
    raise CollectionError("Failed to upsert")
  }

  //
  result.unwrap()
}

///|
pub fn upsert(
  collection : String,
  key : String,
  text : String,
  labels : Array[String],
  opts : Array[NamespaceOption]
) -> CollectionMutationResult!Error {
  upsert_batch!(collection, [key], [text], [labels], opts)
}

///|
pub fn remove(
  collection : String,
  key : String,
  opts : Array[NamespaceOption]
) -> CollectionMutationResult!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if key == "" {
    raise CollectionError("Key is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.delete(collection, ns_opts.namespace, key)
  if result.is_empty() {
    raise CollectionError("Failed to delete")
  }

  //
  result.unwrap()
}

///|
pub(all) type SearchOption (SearchOptions) -> Unit

///|
pub(all) struct SearchOptions {
  mut namespaces : Array[String]
  mut limit : Int
  mut return_text : Bool
}

///|
pub fn with_namespaces(namespaces : Array[String]) -> SearchOption {
  fn(o) { o.namespaces = namespaces }
}

///|
pub fn with_limit(limit : Int) -> SearchOption {
  fn(o) { o.limit = limit }
}

///|
pub fn with_return_text(return_text : Bool) -> SearchOption {
  fn(o) { o.return_text = return_text }
}

///|
pub fn search(
  collection : String,
  search_method : String,
  text : String,
  opts : Array[SearchOption]
) -> CollectionSearchResult!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if search_method == "" {
    raise CollectionError("Search method is required")
  }
  if text == "" {
    raise CollectionError("Text is required")
  }

  //
  let s_opts = { namespaces: [], limit: 10, return_text: false }
  for opt in opts {
    (opt._)(s_opts)
  }

  //
  let result = @host.search(
    collection,
    s_opts.namespaces,
    search_method,
    text,
    s_opts.limit,
    s_opts.return_text,
  )
  if result.is_empty() {
    raise CollectionError("Failed to search")
  }

  //
  result.unwrap()
}

///|
pub fn search_by_vector(
  collection : String,
  search_method : String,
  vector : Array[Double],
  opts : Array[SearchOption]
) -> CollectionSearchResult!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if search_method == "" {
    raise CollectionError("Search method is required")
  }
  if vector.length() == 0 {
    raise CollectionError("Vector is required")
  }

  //
  let s_opts = { namespaces: [], limit: 10, return_text: false }
  for opt in opts {
    (opt._)(s_opts)
  }

  //
  let result = @host.search_by_vector(
    collection,
    s_opts.namespaces,
    search_method,
    vector,
    s_opts.limit,
    s_opts.return_text,
  )
  if result.is_empty() {
    raise CollectionError("Failed to search")
  }

  //
  result.unwrap()
}

///|
pub fn nn_classify(
  collection : String,
  search_method : String,
  text : String,
  opts : Array[NamespaceOption]
) -> CollectionClassificationResult!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if search_method == "" {
    raise CollectionError("Search method is required")
  }
  if text == "" {
    raise CollectionError("Text is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.classify_text(
    collection,
    ns_opts.namespace,
    search_method,
    text,
  )
  if result.is_empty() {
    raise CollectionError("Failed to classify")
  }

  //
  result.unwrap()
}

///|
pub fn recompute_search_method(
  collection : String,
  search_method : String,
  opts : Array[NamespaceOption]
) -> SearchMethodMutationResult!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if search_method == "" {
    raise CollectionError("Search method is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.recompute_index(
    collection,
    ns_opts.namespace,
    search_method,
  )
  if result.is_empty() {
    raise CollectionError("Failed to recompute")
  }

  //
  result.unwrap()
}

///|
pub fn compute_distance(
  collection : String,
  search_method : String,
  key1 : String,
  key2 : String,
  opts : Array[NamespaceOption]
) -> CollectionSearchResultObject!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if search_method == "" {
    raise CollectionError("Search method is required")
  }
  if key1 == "" {
    raise CollectionError("Key1 is required")
  }
  if key2 == "" {
    raise CollectionError("Key2 is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.compute_distance(
    collection,
    ns_opts.namespace,
    search_method,
    key1,
    key2,
  )
  if result.is_empty() {
    raise CollectionError("Failed to compute distance")
  }

  //
  result.unwrap()
}

///|
pub fn get_text(
  collection : String,
  key : String,
  opts : Array[NamespaceOption]
) -> String!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if key == "" {
    raise CollectionError("Key is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.get_text(collection, ns_opts.namespace, key)
  if result.is_empty() {
    raise CollectionError("Failed to get text")
  }

  //
  result.unwrap()
}

///|
pub fn get_texts(
  collection : String,
  opts : Array[NamespaceOption]
) -> Map[String, String]!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.dump_texts(collection, ns_opts.namespace)
  if result.is_empty() {
    raise CollectionError("Failed to get texts")
  }

  // TODO:
  // result.unwrap()
  {}
}

///|
pub fn get_namespaces(collection : String) -> Array[String]!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }

  //
  let result = @host.get_namespaces(collection)
  if result.is_empty() {
    raise CollectionError("Failed to get namespaces")
  }

  //
  result.unwrap()
}

///|
pub fn get_vector(
  collection : String,
  search_method : String,
  key : String,
  opts : Array[NamespaceOption]
) -> Array[Double]!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if search_method == "" {
    raise CollectionError("Search method is required")
  }
  if key == "" {
    raise CollectionError("Key is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.get_vector(
    collection,
    ns_opts.namespace,
    search_method,
    key,
  )
  if result.is_empty() {
    raise CollectionError("Failed to get vector")
  }

  //
  result.unwrap()
}

///|
pub fn get_labels(
  collection : String,
  key : String,
  opts : Array[NamespaceOption]
) -> Array[String]!Error {
  if collection == "" {
    raise CollectionError("Collection name is required")
  }
  if key == "" {
    raise CollectionError("Key is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  let result = @host.get_labels(collection, ns_opts.namespace, key)
  if result.is_empty() {
    raise CollectionError("Failed to get labels")
  }

  //
  result.unwrap()
}
