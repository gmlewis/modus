///|modus:import modus_http_client fetch(request : Request?) -> Response?
fn wasmImportfetch(request : Int) -> Int = "modus_http_client" "fetch"

///|
pub fn host_fetch(request : Request) -> Response raise Error {
  let request : Int = cast(Some(request))
  let response = wasmImportfetch(request)
  if response == 0 {
    raise HttpError("HTTP fetch failed. Check the logs for more information.")
  }
  let response : Response? = cast(response)
  response.unwrap()
}

///|
fn[A, B] cast(a : A) -> B = "%identity"
