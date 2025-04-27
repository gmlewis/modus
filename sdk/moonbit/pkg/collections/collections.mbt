// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub const SUCCESS = "success"

///|
pub const ERROR = "error"

///|
pub(all) struct CollectionMutationResult {
  collection : String
  status : String
  error : String
  operation : String
  keys : Array[String]
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct SearchMethodMutationResult {
  collection : String
  status : String
  error : String
  operation : String
  search_method : String
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct CollectionSearchResult {
  collection : String
  status : String
  error : String
  search_method : String
  objects : Array[CollectionSearchResultObject]
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct CollectionSearchResultObject {
  namespace : String
  key : String
  text : String
  labels : Array[String]
  distance : Double
  score : Double
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct CollectionClassificationResult {
  collection : String
  status : String
  error : String
  search_method : String
  labels_result : Array[CollectionClassificationLabelObject]
  cluster : Array[CollectionClassificationResultObject]
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct CollectionClassificationLabelObject {
  label : String
  confidence : Double
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct CollectionClassificationResultObject {
  key : String
  labels : Array[String]
  distance : Double
  score : Double
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct NamespaceOptions {
  mut namespace : String
}

///|
pub type NamespaceOption (NamespaceOptions) -> Unit

///|
pub fn with_namespace(namespace : String) -> NamespaceOption {
  fn(o) { o.namespace = namespace }
}

///|
pub(all) type! CollectionError String derive(Show, Eq)

///|
pub fn upsert_batch(
  collection : String,
  texts : Array[String],
  keys~ : Array[String] = [],
  labels_arr~ : Array[Array[String]] = [[]],
  opts~ : Array[NamespaceOption] = []
) -> CollectionMutationResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard texts.length() > 0 else { raise CollectionError("Text is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_upsert!(collection, ns_opts.namespace, keys, texts, labels_arr)
}

///|
pub fn upsert(
  collection : String,
  text : String,
  key~ : String = "",
  labels~ : Array[String] = [],
  opts~ : Array[NamespaceOption] = []
) -> CollectionMutationResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard text != "" else { raise CollectionError("Text is required") }
  let keys = if key == "" { [] } else { [key] }
  let labels_arr = if labels.length() == 0 { [[]] } else { [labels] }
  upsert_batch!(collection, [text], keys~, labels_arr~, opts~)
}

///|
pub fn remove(
  collection : String,
  key : String,
  opts~ : Array[NamespaceOption] = []
) -> CollectionMutationResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard key != "" else { raise CollectionError("Key is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_delete!(collection, ns_opts.namespace, key)
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
  opts~ : Array[SearchOption] = []
) -> CollectionSearchResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard search_method != "" else {
    raise CollectionError("Search method is required")
  }
  guard text != "" else { raise CollectionError("Text is required") }

  //
  let s_opts = { namespaces: [], limit: 10, return_text: false }
  for opt in opts {
    (opt._)(s_opts)
  }

  //
  host_search!(
    collection,
    s_opts.namespaces,
    search_method,
    text,
    s_opts.limit,
    s_opts.return_text,
  )
}

///|
pub fn search_by_vector(
  collection : String,
  search_method : String,
  vector : Array[Float],
  opts~ : Array[SearchOption] = []
) -> CollectionSearchResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard search_method != "" else {
    raise CollectionError("Search method is required")
  }
  guard vector.length() > 0 else { raise CollectionError("Vector is required") }

  //
  let s_opts = { namespaces: [], limit: 10, return_text: false }
  for opt in opts {
    (opt._)(s_opts)
  }

  //
  host_search_by_vector!(
    collection,
    s_opts.namespaces,
    search_method,
    vector,
    s_opts.limit,
    s_opts.return_text,
  )
}

///|
pub fn nn_classify(
  collection : String,
  search_method : String,
  text : String,
  opts~ : Array[NamespaceOption] = []
) -> CollectionClassificationResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard search_method != "" else {
    raise CollectionError("Search method is required")
  }
  guard text != "" else { raise CollectionError("Text is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_classify_text!(collection, ns_opts.namespace, search_method, text)
}

///|
pub fn recompute_search_method(
  collection : String,
  search_method : String,
  opts~ : Array[NamespaceOption] = []
) -> SearchMethodMutationResult!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard search_method != "" else {
    raise CollectionError("Search method is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_recompute_index!(collection, ns_opts.namespace, search_method)
}

///|
pub fn compute_distance(
  collection : String,
  search_method : String,
  key1 : String,
  key2 : String,
  opts~ : Array[NamespaceOption] = []
) -> CollectionSearchResultObject!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard search_method != "" else {
    raise CollectionError("Search method is required")
  }
  guard key1 != "" else { raise CollectionError("Key1 is required") }
  guard key2 != "" else { raise CollectionError("Key2 is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_compute_distance!(
    collection,
    ns_opts.namespace,
    search_method,
    key1,
    key2,
  )
}

///|
pub fn get_text(
  collection : String,
  key : String,
  opts~ : Array[NamespaceOption] = []
) -> String!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard key != "" else { raise CollectionError("Key is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_get_text!(collection, ns_opts.namespace, key)
}

///|
pub fn get_texts(
  collection : String,
  opts~ : Array[NamespaceOption] = []
) -> Map[String, String]!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_dump_texts!(collection, ns_opts.namespace)
}

///|
pub fn get_namespaces(collection : String) -> Array[String]!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }

  //
  host_get_namespaces!(collection)
}

///|
pub fn get_vector(
  collection : String,
  search_method : String,
  key : String,
  opts~ : Array[NamespaceOption] = []
) -> Array[Float]!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard search_method != "" else {
    raise CollectionError("Search method is required")
  }
  guard key != "" else { raise CollectionError("Key is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_get_vector!(collection, ns_opts.namespace, search_method, key)
}

///|
pub fn get_labels(
  collection : String,
  key : String,
  opts~ : Array[NamespaceOption] = []
) -> Array[String]!Error {
  guard collection != "" else {
    raise CollectionError("Collection name is required")
  }
  guard key != "" else { raise CollectionError("Key is required") }

  //
  let ns_opts = { namespace: "" }
  for opt in opts {
    (opt._)(ns_opts)
  }

  //
  host_get_labels!(collection, ns_opts.namespace, key)
}
