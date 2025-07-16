///|
pub fn host_fetch(_request : Request) -> Response!Error {
  @ffi.silence_warning()
  fail!("not wasm")
}
