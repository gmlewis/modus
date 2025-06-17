///| Upsert documents into a collection
pub fn host_upsert(
  _collection_name : String,
  _namespace : String,
  _keys : Array[String],
  _texts : Array[String],
  _labels : Array[Array[String]]
) -> CollectionMutationResult raise Error {
  fail("not wasm")
}

///| Delete a document from a collection
pub fn host_delete(
  _collection_name : String,
  _namespace : String,
  _key : String
) -> CollectionMutationResult raise Error {
  fail("not wasm")
}

///| Search documents in a collection
pub fn host_search(
  _collection_name : String,
  _namespaces : Array[String],
  _search_method : String,
  _text : String,
  _limit : Int,
  _return_text : Bool
) -> CollectionSearchResult raise Error {
  fail("not wasm")
}

///| Classify text using a collection
pub fn host_classify_text(
  _collection_name : String,
  _namespace : String,
  _search_method : String,
  _text : String
) -> CollectionClassificationResult raise Error {
  fail("not wasm")
}

///| Recompute index for a collection
pub fn host_recompute_index(
  _collection_name : String,
  _namespace : String,
  _search_method : String
) -> SearchMethodMutationResult raise Error {
  fail("not wasm")
}

///| Compute distance between two documents
pub fn host_compute_distance(
  _collection_name : String,
  _namespace : String,
  _search_method : String,
  _key1 : String,
  _key2 : String
) -> CollectionSearchResultObject raise Error {
  fail("not wasm")
}

///| Get text for a document
pub fn host_get_text(
  _collection_name : String,
  _namespace : String,
  _key : String
) -> String raise Error {
  fail("not wasm")
}

///| Dump all texts from a collection namespace
pub fn host_dump_texts(
  _collection_name : String,
  _namespace : String
) -> Map[String, String] raise Error {
  fail("not wasm")
}

///| Get all namespaces in a collection
pub fn host_get_namespaces(
  _collection_name : String
) -> Array[String] raise Error {
  fail("not wasm")
}

///| Get vector for a document
pub fn host_get_vector(
  _collection_name : String,
  _namespace : String,
  _search_method : String,
  _key : String
) -> Array[Float] raise Error {
  fail("not wasm")
}

///| Get labels for a document
pub fn host_get_labels(
  _collection_name : String,
  _namespace : String,
  _key : String
) -> Array[String] raise Error {
  fail("not wasm")
}

///| Search by vector
pub fn host_search_by_vector(
  _collection_name : String,
  _namespaces : Array[String],
  _search_method : String,
  _vector : Array[Float],
  _limit : Int,
  _return_text : Bool
) -> CollectionSearchResult raise Error {
  fail("not wasm")
}
