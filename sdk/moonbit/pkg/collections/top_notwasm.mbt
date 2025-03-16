///| Upsert documents into a collection
pub fn host_upsert(
  _collection : String,
  _namespace : String,
  _keys : Array[String],
  _texts : Array[String],
  _labels : Array[Array[String]]
) -> CollectionMutationResult? {
  @ffi.silence_warning()
  abort("not wasm")
}

///| Delete a document from a collection
pub fn host_delete(
  _collection : String,
  _namespace : String,
  _key : String
) -> CollectionMutationResult? {
  abort("not wasm")
}

///| Search documents in a collection
pub fn host_search(
  _collection : String,
  _namespaces : Array[String],
  _search_method : String,
  _text : String,
  _limit : Int,
  _return_text : Bool
) -> CollectionSearchResult? {
  abort("not wasm")
}

///| Classify text using a collection
pub fn host_classify_text(
  _collection : String,
  _namespace : String,
  _search_method : String,
  _text : String
) -> CollectionClassificationResult? {
  abort("not wasm")
}

///| Recompute index for a collection
pub fn host_recompute_index(
  _collection : String,
  _namespace : String,
  _search_method : String
) -> SearchMethodMutationResult? {
  abort("not wasm")
}

///| Compute distance between two documents
pub fn host_compute_distance(
  _collection : String,
  _namespace : String,
  _search_method : String,
  _key1 : String,
  _key2 : String
) -> CollectionSearchResultObject? {
  abort("not wasm")
}

///| Get text for a document
pub fn host_get_text(
  _collection : String,
  _namespace : String,
  _key : String
) -> String? {
  abort("not wasm")
}

///| Dump all texts from a collection namespace
pub fn host_dump_texts(_collection : String, _namespace : String) -> String {
  abort("not wasm")
}

///| option<map<string, string>>;
/// Get all namespaces in a collection
pub fn host_get_namespaces(_collection : String) -> Array[String]? {
  abort("not wasm")
}

///| Get vector for a document
pub fn host_get_vector(
  _collection : String,
  _namespace : String,
  _search_method : String,
  _key : String
) -> FixedArray[Float]? {
  abort("not wasm")
}

///| Get labels for a document
pub fn host_get_labels(
  _collection : String,
  _namespace : String,
  _key : String
) -> Array[String]? {
  abort("not wasm")
}

///| Search by vector
pub fn host_search_by_vector(
  _collection : String,
  _namespaces : Array[String],
  _search_method : String,
  _vector : FixedArray[Float],
  _limit : Int,
  _return_text : Bool
) -> CollectionSearchResult? {
  abort("not wasm")
}
