///|
pub fn host_fetch(_request : Request) -> Response raise Error {
  @ffi.silence_warning()
  fail("not wasm")
}
