///| Upsert documents into a collection
pub fn host_upsert(
  collection : String,
  namespace : String,
  keys : Array[String],
  texts : Array[String],
  labels : Array[Array[String]]
) -> CollectionMutationResult? {
  let cleanupList : Array[@ffi.Cleanup] = []
  let ignoreList : Array[&@ffi.Any] = []
  let address = @ffi.malloc(keys.length() * 8)
  for index = 0; index < keys.length(); index = index + 1 {
    let element : String = keys[index]
    let base = address + index * 8
    @ffi.store32(base + 4, element.iter().count())
    @ffi.store32(base + 0, @ffi.str2ptr(element))
    ignoreList.push(element)
  }
  let address2 = @ffi.malloc(texts.length() * 8)
  for index3 = 0; index3 < texts.length(); index3 = index3 + 1 {
    let element0 : String = texts[index3]
    let base1 = address2 + index3 * 8
    @ffi.store32(base1 + 4, element0.iter().count())
    @ffi.store32(base1 + 0, @ffi.str2ptr(element0))
    ignoreList.push(element0)
  }
  let address10 = @ffi.malloc(labels.length() * 8)
  for index11 = 0; index11 < labels.length(); index11 = index11 + 1 {
    let element4 : Array[String] = labels[index11]
    let base5 = address10 + index11 * 8
    let address8 = @ffi.malloc(element4.length() * 8)
    for index9 = 0; index9 < element4.length(); index9 = index9 + 1 {
      let element6 : String = element4[index9]
      let base7 = address8 + index9 * 8
      @ffi.store32(base7 + 4, element6.iter().count())
      @ffi.store32(base7 + 0, @ffi.str2ptr(element6))
      ignoreList.push(element6)
    }
    @ffi.store32(base5 + 4, element4.length())
    @ffi.store32(base5 + 0, address8)
    cleanupList.push({
      address: address8,
      size: element4.length() * 8,
      align: 4,
    })
  }
  let return_area = @ffi.malloc(44)
  wasmImportUpsert(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    address,
    keys.length(),
    address2,
    texts.length(),
    address10,
    labels.length(),
    return_area,
  )
  let lifted : CollectionMutationResult? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      ignore(@ffi.load32(return_area + 16))
      let result16 = @ffi.ptr2str(@ffi.load32(return_area + 12))
      ignore(@ffi.load32(return_area + 24))
      let result17 = @ffi.ptr2str(@ffi.load32(return_area + 20))
      ignore(@ffi.load32(return_area + 32))
      let result18 = @ffi.ptr2str(@ffi.load32(return_area + 28))
      let array : Array[String] = []
      for index22 = 0
          index22 < @ffi.load32(return_area + 40)
          index22 = index22 + 1 {
        let base20 = @ffi.load32(return_area + 36) + index22 * 8
        ignore(@ffi.load32(base20 + 4))
        let result21 = @ffi.ptr2str(@ffi.load32(base20 + 0))
        array.push(result21)
      }
      @ffi.free(@ffi.load32(return_area + 36))
      Option::Some(CollectionMutationResult::{
        collection: result,
        status: result16,
        error: result17,
        operation: result18,
        keys: array,
      })
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  @ffi.free(address)
  @ffi.free(address2)
  @ffi.free(address10)
  @ffi.free(return_area)
  cleanupList.each(fn(cleanup) { @ffi.free(cleanup.address) })
  ignore(ignoreList)
  return lifted
}

///| Delete a document from a collection
pub fn host_delete(
  collection : String,
  namespace : String,
  key : String
) -> CollectionMutationResult? {
  let return_area = @ffi.malloc(44)
  wasmImportDelete(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(key),
    key.iter().count(),
    return_area,
  )
  let lifted : CollectionMutationResult? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      ignore(@ffi.load32(return_area + 16))
      let result2 = @ffi.ptr2str(@ffi.load32(return_area + 12))
      ignore(@ffi.load32(return_area + 24))
      let result3 = @ffi.ptr2str(@ffi.load32(return_area + 20))
      ignore(@ffi.load32(return_area + 32))
      let result4 = @ffi.ptr2str(@ffi.load32(return_area + 28))
      let array : Array[String] = []
      for index = 0; index < @ffi.load32(return_area + 40); index = index + 1 {
        let base6 = @ffi.load32(return_area + 36) + index * 8
        ignore(@ffi.load32(base6 + 4))
        let result7 = @ffi.ptr2str(@ffi.load32(base6 + 0))
        array.push(result7)
      }
      @ffi.free(@ffi.load32(return_area + 36))
      Option::Some(CollectionMutationResult::{
        collection: result,
        status: result2,
        error: result3,
        operation: result4,
        keys: array,
      })
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(key)
  @ffi.free(return_area)
  return lifted
}

///| Search documents in a collection
pub fn host_search(
  collection : String,
  namespaces : Array[String],
  search_method : String,
  text : String,
  limit : Int,
  return_text : Bool
) -> CollectionSearchResult? {
  let cleanupList : Array[@ffi.Cleanup] = []
  let ignoreList : Array[&@ffi.Any] = []
  let address = @ffi.malloc(namespaces.length() * 8)
  for index = 0; index < namespaces.length(); index = index + 1 {
    let element : String = namespaces[index]
    let base = address + index * 8
    @ffi.store32(base + 4, element.iter().count())
    @ffi.store32(base + 0, @ffi.str2ptr(element))
    ignoreList.push(element)
  }
  let return_area = @ffi.malloc(44)
  wasmImportSearch(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    address,
    namespaces.length(),
    @ffi.str2ptr(search_method),
    search_method.iter().count(),
    @ffi.str2ptr(text),
    text.iter().count(),
    limit,
    if return_text {
      1
    } else {
      0
    },
    return_area,
  )
  let lifted : CollectionSearchResult? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      ignore(@ffi.load32(return_area + 16))
      let result4 = @ffi.ptr2str(@ffi.load32(return_area + 12))
      ignore(@ffi.load32(return_area + 24))
      let result5 = @ffi.ptr2str(@ffi.load32(return_area + 20))
      ignore(@ffi.load32(return_area + 32))
      let result6 = @ffi.ptr2str(@ffi.load32(return_area + 28))
      let array16 : Array[CollectionSearchResultObject] = []
      for index17 = 0
          index17 < @ffi.load32(return_area + 40)
          index17 = index17 + 1 {
        let base8 = @ffi.load32(return_area + 36) + index17 * 48
        ignore(@ffi.load32(base8 + 4))
        let result9 = @ffi.ptr2str(@ffi.load32(base8 + 0))
        ignore(@ffi.load32(base8 + 12))
        let result10 = @ffi.ptr2str(@ffi.load32(base8 + 8))
        ignore(@ffi.load32(base8 + 20))
        let result11 = @ffi.ptr2str(@ffi.load32(base8 + 16))
        let array : Array[String] = []
        for index15 = 0
            index15 < @ffi.load32(base8 + 28)
            index15 = index15 + 1 {
          let base13 = @ffi.load32(base8 + 24) + index15 * 8
          ignore(@ffi.load32(base13 + 4))
          let result14 = @ffi.ptr2str(@ffi.load32(base13 + 0))
          array.push(result14)
        }
        @ffi.free(@ffi.load32(base8 + 24))
        array16.push(CollectionSearchResultObject::{
          namespace: result9,
          key: result10,
          text: result11,
          labels: array,
          distance: @ffi.loadf64(base8 + 32),
          score: @ffi.loadf64(base8 + 40),
        })
      }
      @ffi.free(@ffi.load32(return_area + 36))
      Option::Some(CollectionSearchResult::{
        collection: result,
        status: result4,
        error: result5,
        search_method: result6,
        objects: array16,
      })
    }
    _ => panic()
  }
  ignore(collection)
  @ffi.free(address)
  ignore(search_method)
  ignore(text)
  @ffi.free(return_area)
  cleanupList.each(fn(cleanup) { @ffi.free(cleanup.address) })
  ignore(ignoreList)
  return lifted
}

///| Classify text using a collection
pub fn host_classify_text(
  collection : String,
  namespace : String,
  search_method : String,
  text : String
) -> CollectionClassificationResult? {
  let return_area = @ffi.malloc(52)
  wasmImportClassifyText(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(search_method),
    search_method.iter().count(),
    @ffi.str2ptr(text),
    text.iter().count(),
    return_area,
  )
  let lifted : CollectionClassificationResult? = match
    @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      ignore(@ffi.load32(return_area + 16))
      let result2 = @ffi.ptr2str(@ffi.load32(return_area + 12))
      ignore(@ffi.load32(return_area + 24))
      let result3 = @ffi.ptr2str(@ffi.load32(return_area + 20))
      ignore(@ffi.load32(return_area + 32))
      let result4 = @ffi.ptr2str(@ffi.load32(return_area + 28))
      let array : Array[CollectionClassificationLabelObject] = []
      for index = 0; index < @ffi.load32(return_area + 40); index = index + 1 {
        let base6 = @ffi.load32(return_area + 36) + index * 16
        ignore(@ffi.load32(base6 + 4))
        let result7 = @ffi.ptr2str(@ffi.load32(base6 + 0))
        array.push(CollectionClassificationLabelObject::{
          label: result7,
          confidence: @ffi.loadf64(base6 + 8),
        })
      }
      @ffi.free(@ffi.load32(return_area + 36))
      let array16 : Array[CollectionClassificationResultObject] = []
      for index17 = 0
          index17 < @ffi.load32(return_area + 48)
          index17 = index17 + 1 {
        let base9 = @ffi.load32(return_area + 44) + index17 * 32
        ignore(@ffi.load32(base9 + 4))
        let result10 = @ffi.ptr2str(@ffi.load32(base9 + 0))
        let array14 : Array[String] = []
        for index15 = 0
            index15 < @ffi.load32(base9 + 12)
            index15 = index15 + 1 {
          let base12 = @ffi.load32(base9 + 8) + index15 * 8
          ignore(@ffi.load32(base12 + 4))
          let result13 = @ffi.ptr2str(@ffi.load32(base12 + 0))
          array14.push(result13)
        }
        @ffi.free(@ffi.load32(base9 + 8))
        array16.push(CollectionClassificationResultObject::{
          key: result10,
          labels: array14,
          distance: @ffi.loadf64(base9 + 16),
          score: @ffi.loadf64(base9 + 24),
        })
      }
      @ffi.free(@ffi.load32(return_area + 44))
      Option::Some(CollectionClassificationResult::{
        collection: result,
        status: result2,
        error: result3,
        search_method: result4,
        labels_result: array,
        cluster: array16,
      })
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(search_method)
  ignore(text)
  @ffi.free(return_area)
  return lifted
}

///| Recompute index for a collection
pub fn host_recompute_index(
  collection : String,
  namespace : String,
  search_method : String
) -> SearchMethodMutationResult? {
  let return_area = @ffi.malloc(44)
  wasmImportRecomputeIndex(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(search_method),
    search_method.iter().count(),
    return_area,
  )
  let lifted : SearchMethodMutationResult? = match
    @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      ignore(@ffi.load32(return_area + 16))
      let result2 = @ffi.ptr2str(@ffi.load32(return_area + 12))
      ignore(@ffi.load32(return_area + 24))
      let result3 = @ffi.ptr2str(@ffi.load32(return_area + 20))
      ignore(@ffi.load32(return_area + 32))
      let result4 = @ffi.ptr2str(@ffi.load32(return_area + 28))
      ignore(@ffi.load32(return_area + 40))
      let result5 = @ffi.ptr2str(@ffi.load32(return_area + 36))
      Option::Some(SearchMethodMutationResult::{
        collection: result,
        status: result2,
        error: result3,
        operation: result4,
        search_method: result5,
      })
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(search_method)
  @ffi.free(return_area)
  return lifted
}

///| Compute distance between two documents
pub fn host_compute_distance(
  collection : String,
  namespace : String,
  search_method : String,
  key1 : String,
  key2 : String
) -> CollectionSearchResultObject? {
  let return_area = @ffi.malloc(56)
  wasmImportComputeDistance(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(search_method),
    search_method.iter().count(),
    @ffi.str2ptr(key1),
    key1.iter().count(),
    @ffi.str2ptr(key2),
    key2.iter().count(),
    return_area,
  )
  let lifted : CollectionSearchResultObject? = match
    @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 12))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 8))
      ignore(@ffi.load32(return_area + 20))
      let result2 = @ffi.ptr2str(@ffi.load32(return_area + 16))
      ignore(@ffi.load32(return_area + 28))
      let result3 = @ffi.ptr2str(@ffi.load32(return_area + 24))
      let array : Array[String] = []
      for index = 0; index < @ffi.load32(return_area + 36); index = index + 1 {
        let base5 = @ffi.load32(return_area + 32) + index * 8
        ignore(@ffi.load32(base5 + 4))
        let result6 = @ffi.ptr2str(@ffi.load32(base5 + 0))
        array.push(result6)
      }
      @ffi.free(@ffi.load32(return_area + 32))
      Option::Some(CollectionSearchResultObject::{
        namespace: result,
        key: result2,
        text: result3,
        labels: array,
        distance: @ffi.loadf64(return_area + 40),
        score: @ffi.loadf64(return_area + 48),
      })
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(search_method)
  ignore(key1)
  ignore(key2)
  @ffi.free(return_area)
  return lifted
}

///| Get text for a document
pub fn host_get_text(
  collection : String,
  namespace : String,
  key : String
) -> String? {
  let return_area = @ffi.malloc(12)
  wasmImportGetText(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(key),
    key.iter().count(),
    return_area,
  )
  let lifted : String? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      Option::Some(result)
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(key)
  @ffi.free(return_area)
  return lifted
}

///| Dump all texts from a collection namespace
pub fn host_dump_texts(collection : String, namespace : String) -> String {
  let return_area = @ffi.malloc(8)
  wasmImportDumpTexts(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    return_area,
  )
  ignore(@ffi.load32(return_area + 4))
  let result = @ffi.ptr2str(@ffi.load32(return_area + 0))
  ignore(collection)
  ignore(namespace)
  @ffi.free(return_area)
  return result
}

///| option<map<string, string>>;
/// Get all namespaces in a collection
pub fn host_get_namespaces(collection : String) -> Array[String]? {
  let return_area = @ffi.malloc(12)
  wasmImportGetNamespaces(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    return_area,
  )
  let lifted : Array[String]? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      let array : Array[String] = []
      for index = 0; index < @ffi.load32(return_area + 8); index = index + 1 {
        let base3 = @ffi.load32(return_area + 4) + index * 8
        ignore(@ffi.load32(base3 + 4))
        let result = @ffi.ptr2str(@ffi.load32(base3 + 0))
        array.push(result)
      }
      @ffi.free(@ffi.load32(return_area + 4))
      Option::Some(array)
    }
    _ => panic()
  }
  ignore(collection)
  @ffi.free(return_area)
  return lifted
}

///| Get vector for a document
pub fn host_get_vector(
  collection : String,
  namespace : String,
  search_method : String,
  key : String
) -> FixedArray[Float]? {
  let return_area = @ffi.malloc(12)
  wasmImportGetVector(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(search_method),
    search_method.iter().count(),
    @ffi.str2ptr(key),
    key.iter().count(),
    return_area,
  )
  let lifted : FixedArray[Float]? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      let result = @ffi.ptr2float_array(
        @ffi.load32(return_area + 4),
        @ffi.load32(return_area + 8),
      )
      Option::Some(result)
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(search_method)
  ignore(key)
  @ffi.free(return_area)
  return lifted
}

///| Get labels for a document
pub fn host_get_labels(
  collection : String,
  namespace : String,
  key : String
) -> Array[String]? {
  let return_area = @ffi.malloc(12)
  wasmImportGetLabels(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    @ffi.str2ptr(namespace),
    namespace.iter().count(),
    @ffi.str2ptr(key),
    key.iter().count(),
    return_area,
  )
  let lifted : Array[String]? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      let array : Array[String] = []
      for index = 0; index < @ffi.load32(return_area + 8); index = index + 1 {
        let base3 = @ffi.load32(return_area + 4) + index * 8
        ignore(@ffi.load32(base3 + 4))
        let result = @ffi.ptr2str(@ffi.load32(base3 + 0))
        array.push(result)
      }
      @ffi.free(@ffi.load32(return_area + 4))
      Option::Some(array)
    }
    _ => panic()
  }
  ignore(collection)
  ignore(namespace)
  ignore(key)
  @ffi.free(return_area)
  return lifted
}

///| Search by vector
pub fn host_search_by_vector(
  collection : String,
  namespaces : Array[String],
  search_method : String,
  vector : FixedArray[Float],
  limit : Int,
  return_text : Bool
) -> CollectionSearchResult? {
  let cleanupList : Array[@ffi.Cleanup] = []
  let ignoreList : Array[&@ffi.Any] = []
  let address = @ffi.malloc(namespaces.length() * 8)
  for index = 0; index < namespaces.length(); index = index + 1 {
    let element : String = namespaces[index]
    let base = address + index * 8
    @ffi.store32(base + 4, element.iter().count())
    @ffi.store32(base + 0, @ffi.str2ptr(element))
    ignoreList.push(element)
  }
  let return_area = @ffi.malloc(44)
  wasmImportSearchByVector(
    @ffi.str2ptr(collection),
    collection.iter().count(),
    address,
    namespaces.length(),
    @ffi.str2ptr(search_method),
    search_method.iter().count(),
    @ffi.float_array2ptr(vector),
    vector.length(),
    limit,
    if return_text {
      1
    } else {
      0
    },
    return_area,
  )
  let lifted : CollectionSearchResult? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      ignore(@ffi.load32(return_area + 16))
      let result4 = @ffi.ptr2str(@ffi.load32(return_area + 12))
      ignore(@ffi.load32(return_area + 24))
      let result5 = @ffi.ptr2str(@ffi.load32(return_area + 20))
      ignore(@ffi.load32(return_area + 32))
      let result6 = @ffi.ptr2str(@ffi.load32(return_area + 28))
      let array16 : Array[CollectionSearchResultObject] = []
      for index17 = 0
          index17 < @ffi.load32(return_area + 40)
          index17 = index17 + 1 {
        let base8 = @ffi.load32(return_area + 36) + index17 * 48
        ignore(@ffi.load32(base8 + 4))
        let result9 = @ffi.ptr2str(@ffi.load32(base8 + 0))
        ignore(@ffi.load32(base8 + 12))
        let result10 = @ffi.ptr2str(@ffi.load32(base8 + 8))
        ignore(@ffi.load32(base8 + 20))
        let result11 = @ffi.ptr2str(@ffi.load32(base8 + 16))
        let array : Array[String] = []
        for index15 = 0
            index15 < @ffi.load32(base8 + 28)
            index15 = index15 + 1 {
          let base13 = @ffi.load32(base8 + 24) + index15 * 8
          ignore(@ffi.load32(base13 + 4))
          let result14 = @ffi.ptr2str(@ffi.load32(base13 + 0))
          array.push(result14)
        }
        @ffi.free(@ffi.load32(base8 + 24))
        array16.push(CollectionSearchResultObject::{
          namespace: result9,
          key: result10,
          text: result11,
          labels: array,
          distance: @ffi.loadf64(base8 + 32),
          score: @ffi.loadf64(base8 + 40),
        })
      }
      @ffi.free(@ffi.load32(return_area + 36))
      Option::Some(CollectionSearchResult::{
        collection: result,
        status: result4,
        error: result5,
        search_method: result6,
        objects: array16,
      })
    }
    _ => panic()
  }
  ignore(collection)
  @ffi.free(address)
  ignore(search_method)
  ignore(vector)
  @ffi.free(return_area)
  cleanupList.each(fn(cleanup) { @ffi.free(cleanup.address) })
  ignore(ignoreList)
  return lifted
}
