///|modus:import modus_collections computeDistance(collection_name : String, namespace : String, search_method : String, key1 : String, key2 : String) -> CollectionSearchResultObject?
fn wasmImportComputeDistance(
  collection_name : Int,
  namespace : Int,
  search_method : Int,
  key1 : Int,
  key2 : Int
) -> Int = "modus_collections" "computeDistance"

///|modus:import modus_collections delete(collection_name : String, namespace : String, key : String) -> CollectionMutationResult?
fn wasmImportDelete(collection_name : Int, namespace : Int, key : Int) -> Int = "modus_collections" "delete"

///|modus:import modus_collections getNamespaces(collection_name : String) -> Array[String]?
fn wasmImportGetNamespaces(collection_name : Int) -> Int = "modus_collections" "getNamespaces"

///|modus:import modus_collections getText(collection_name : String, namespace : String, key : String) -> String?
fn wasmImportGetText(collection_name : Int, namespace : Int, key : Int) -> Int = "modus_collections" "getText"

///|modus:import modus_collections dumpTexts(collection_name : String, namespace : String) -> Map[String, String]?
fn wasmImportDumpTexts(collection_name : Int, namespace : Int) -> Int = "modus_collections" "dumpTexts"

///|modus:import modus_collections getVector(collection_name : String, namespace : String, search_method : String, key : String) -> Array[Float]?
fn wasmImportGetVector(
  collection_name : Int,
  namespace : Int,
  search_method : Int,
  key : Int
) -> Int = "modus_collections" "getVector"

///|modus:import modus_collections getLabels(collection_name : String, namespace : String, key : String) -> Array[String]?
fn wasmImportGetLabels(
  collection_name : Int,
  namespace : Int,
  key : Int
) -> Int = "modus_collections" "getLabels"

///|modus:import modus_collections classifyText(collection_name : String, namespace : String, search_method : String, text : String) -> CollectionClassificationResult?
fn wasmImportClassifyText(
  collection_name : Int,
  namespace : Int,
  search_method : Int,
  text : Int
) -> Int = "modus_collections" "classifyText"

///|modus:import modus_collections recomputeIndex(collection_name : String, namespace : String, search_method : String) -> SearchMethodMutationResult?
fn wasmImportRecomputeIndex(
  collection_name : Int,
  namespace : Int,
  search_method : Int
) -> Int = "modus_collections" "recomputeIndex"

///|modus:import modus_collections search(collection_name : String, namespaces : Array[String], search_method : String, text : String, limit : Int, return_text : Bool) -> CollectionSearchResult?
fn wasmImportSearch(
  collection_name : Int,
  namespaces : Int,
  search_method : Int,
  text : Int,
  limit : Int,
  return_text : Int
) -> Int = "modus_collections" "search"

///|modus:import modus_collections searchByVector(collection_name : String, namespaces : Array[String], search_method : String, vector : Array[Float], limit : Int, return_text : Bool) -> CollectionSearchResult?
fn wasmImportSearchByVector(
  collection_name : Int,
  namespaces : Int,
  search_method : Int,
  vector : Int,
  limit : Int,
  return_text : Int
) -> Int = "modus_collections" "searchByVector"

///|modus:import modus_collections upsert(collection_name : String, namespace : String, keys : Array[String], texts : Array[String], labels : Array[Array[String]]) -> CollectionMutationResult?
fn wasmImportUpsert(
  collection_name : Int,
  namespace : Int,
  keys : Int,
  texts : Int,
  labels : Int
) -> Int = "modus_collections" "upsert"

///|
fn[A, B] cast(a : A) -> B = "%identity"

///| Upsert documents into a collection
pub fn host_upsert(
  collection_name : String,
  namespace : String,
  keys : Array[String],
  texts : Array[String],
  labels : Array[Array[String]]
) -> CollectionMutationResult raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let keys : Int = cast(keys)
  let texts : Int = cast(texts)
  let labels : Int = cast(labels)
  let result = wasmImportUpsert(collection_name, namespace, keys, texts, labels)
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : CollectionMutationResult? = cast(result)
  result.unwrap()
}

///| Delete a document from a collection
pub fn host_delete(
  collection_name : String,
  namespace : String,
  key : String
) -> CollectionMutationResult raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let key : Int = cast(key)
  let result = wasmImportDelete(collection_name, namespace, key)
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : CollectionMutationResult? = cast(result)
  result.unwrap()
}

///| Search documents in a collection
pub fn host_search(
  collection_name : String,
  namespaces : Array[String],
  search_method : String,
  text : String,
  limit : Int,
  return_text : Bool
) -> CollectionSearchResult raise Error {
  let collection_name : Int = cast(collection_name)
  let namespaces : Int = cast(namespaces)
  let search_method : Int = cast(search_method)
  let text : Int = cast(text)
  let limit : Int = cast(limit)
  let return_text : Int = cast(return_text)
  let result = wasmImportSearch(
    collection_name, namespaces, search_method, text, limit, return_text,
  )
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : CollectionSearchResult? = cast(result)
  result.unwrap()
}

///| Classify text using a collection
pub fn host_classify_text(
  collection_name : String,
  namespace : String,
  search_method : String,
  text : String
) -> CollectionClassificationResult raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let search_method : Int = cast(search_method)
  let text : Int = cast(text)
  let result = wasmImportClassifyText(
    collection_name, namespace, search_method, text,
  )
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : CollectionClassificationResult? = cast(result)
  result.unwrap()
}

///| Recompute index for a collection
pub fn host_recompute_index(
  collection_name : String,
  namespace : String,
  search_method : String
) -> SearchMethodMutationResult raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let search_method : Int = cast(search_method)
  let result = wasmImportRecomputeIndex(
    collection_name, namespace, search_method,
  )
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : SearchMethodMutationResult? = cast(result)
  result.unwrap()
}

///| Compute distance between two documents
pub fn host_compute_distance(
  collection_name : String,
  namespace : String,
  search_method : String,
  key1 : String,
  key2 : String
) -> CollectionSearchResultObject raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let search_method : Int = cast(search_method)
  let key1 : Int = cast(key1)
  let key2 : Int = cast(key2)
  let result = wasmImportComputeDistance(
    collection_name, namespace, search_method, key1, key2,
  )
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : CollectionSearchResultObject? = cast(result)
  result.unwrap()
}

///| Get text for a document
pub fn host_get_text(
  collection_name : String,
  namespace : String,
  key : String
) -> String raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let key : Int = cast(key)
  let result = wasmImportGetText(collection_name, namespace, key)
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : String? = cast(result)
  result.unwrap()
}

///| Dump all texts from a collection namespace
pub fn host_dump_texts(
  collection_name : String,
  namespace : String
) -> Map[String, String] raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let result = wasmImportDumpTexts(collection_name, namespace)
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : Map[String, String]? = cast(result)
  result.unwrap()
}

///| Get all namespaces in a collection
pub fn host_get_namespaces(
  collection_name : String
) -> Array[String] raise Error {
  let collection_name : Int = cast(collection_name)
  let result = wasmImportGetNamespaces(collection_name)
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : Array[String]? = cast(result)
  result.unwrap()
}

///| Get vector for a document
pub fn host_get_vector(
  collection_name : String,
  namespace : String,
  search_method : String,
  key : String
) -> Array[Float] raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let search_method : Int = cast(search_method)
  let key : Int = cast(key)
  let result = wasmImportGetVector(
    collection_name, namespace, search_method, key,
  )
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : Array[Float]? = cast(result)
  result.unwrap()
}

///| Get labels for a document
pub fn host_get_labels(
  collection_name : String,
  namespace : String,
  key : String
) -> Array[String] raise Error {
  let collection_name : Int = cast(collection_name)
  let namespace : Int = cast(namespace)
  let key : Int = cast(key)
  let result = wasmImportGetLabels(collection_name, namespace, key)
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : Array[String]? = cast(result)
  result.unwrap()
}

///| Search by vector
pub fn host_search_by_vector(
  collection_name : String,
  namespaces : Array[String],
  search_method : String,
  vector : Array[Float],
  limit : Int,
  return_text : Bool
) -> CollectionSearchResult raise Error {
  let collection_name : Int = cast(collection_name)
  let namespaces : Int = cast(namespaces)
  let search_method : Int = cast(search_method)
  let vector : Int = cast(vector)
  let limit : Int = cast(limit)
  let return_text : Int = cast(return_text)
  let result = wasmImportSearchByVector(
    collection_name, namespaces, search_method, vector, limit, return_text,
  )
  guard result != 0 else {
    raise CollectionError(
      "failed to get result. Check the logs for more information.",
    )
  }
  let result : CollectionSearchResult? = cast(result)
  result.unwrap()
}
