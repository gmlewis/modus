// This files is provided so that `moon test --target js` can be run without errors.

///|
pub extern "js" fn extend16(value : Int) -> Int =
  #|(value) => value
//   #|(func (param i32) (result i32) local.get 0 i32.extend16_s)

///|
pub extern "js" fn extend8(value : Int) -> Int =
  #|(value) => value
//   #|(func (param i32) (result i32) local.get 0 i32.extend8_s)

///|
pub extern "js" fn store8(offset : Int, value : Int) =
  #|(offset, value) => {}
//   #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store8)

///|
pub extern "js" fn load8_u(offset : Int) -> Int =
  #|(offset) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.load8_u)

///|
pub extern "js" fn load8(offset : Int) -> Int =
  #|(offset) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.load8_s)

///|
pub extern "js" fn store16(offset : Int, value : Int) =
  #|(offset, value) => {}
//   #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store16)

///|
pub extern "js" fn load16(offset : Int) -> Int =
  #|(offset) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.load16_s)

///|
pub extern "js" fn load16_u(offset : Int) -> Int =
  #|(offset) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.load16_u)

///|
pub extern "js" fn store32(offset : Int, value : Int) =
  #|(offset, value) => {}
//   #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store)

///|
pub extern "js" fn load32(offset : Int) -> Int =
  #|(offset) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.load)

///|
pub extern "js" fn store64(offset : Int, value : Int64) =
  #|(offset, value) => {}
//   #|(func (param i32) (param i64) local.get 0 local.get 1 i64.store)

///|
pub extern "js" fn load64(offset : Int) -> Int64 =
  #|(offset) => 0
//   #|(func (param i32) (result i64) local.get 0 i64.load)

///|
pub extern "js" fn storef32(offset : Int, value : Float) =
  #|(offset, value) => {}
//   #|(func (param i32) (param f32) local.get 0 local.get 1 f32.store)

///|
pub extern "js" fn loadf32(offset : Int) -> Float =
  #|(offset) => 0
//   #|(func (param i32) (result f32) local.get 0 f32.load)

///|
pub extern "js" fn storef64(offset : Int, value : Double) =
  #|(offset, value) => {}
//   #|(func (param i32) (param f64) local.get 0 local.get 1 f64.store)

///|
pub extern "js" fn loadf64(offset : Int) -> Double =
  #|(offset) => 0
//   #|(func (param i32) (result f64) local.get 0 f64.load)

///|
pub extern "js" fn f32_to_i32(value : Float) -> Int =
  #|(value) => Math.floor(value)
//   #|(func (param f32) (result i32) local.get 0 f32.convert_i32_s)

///|
pub extern "js" fn f32_to_i64(value : Float) -> Int64 =
  #|(value) => Math.floor(value)
//   #|(func (param f32) (result i64) local.get 0 f32.convert_i64_s)

///|
extern "js" fn malloc_inline(size : Int) -> Int =
  #|(size) => 0
// #|(func (param i32) (result i32) local.get 0 call $moonbit.malloc)

///|
pub fn malloc(size : Int) -> Int {
  // let words = size / 4 + 1
  // let address = malloc_inline(8 + words * 4)
  // store32(address, 1)
  // store32(address + 4, (words << 8) | 246)
  // store8(address + words * 4 + 7, 3 - size % 4)
  // address + 8
  0
}

///|
pub extern "js" fn free(position : Int) =
  #|(position) => {}
//   #|(func (param i32) local.get 0 i32.const 8 i32.sub call $moonbit.decref)

///|
pub fn copy(dest : Int, src : Int) -> Unit {
  // let src = src - 8
  // let dest = dest - 8
  // let src_len = load32(src + 4) & 0xFFFFFF
  // let dest_len = load32(dest + 4) & 0xFFFFFF
  // let min = if src_len < dest_len { src_len } else { dest_len }
  // copy_inline(dest, src, min)
}

///|
extern "js" fn copy_inline(dest : Int, src : Int, len : Int) =
  #|(dest, src, len) => {}
// #|(func (param i32) (param i32) (param i32) local.get 0 local.get 1 local.get 2 memory.copy)

///|
pub extern "js" fn str2ptr(str : String) -> Int =
  #|(str) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn ptr2str(ptr : Int) -> String =
  #|(ptr) => ""
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 243 i32.store8 local.get 0 i32.const 8 i32.sub)

///|
pub extern "js" fn bytes2ptr(bytes : Bytes) -> Int =
  #|(bytes) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn ptr2bytes(ptr : Int) -> Bytes =
  #|(ptr) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

///|
pub extern "js" fn uint_array2ptr(array : FixedArray[UInt]) -> Int =
  #|(array) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn uint64_array2ptr(array : FixedArray[UInt64]) -> Int =
  #|(array) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn int_array2ptr(array : FixedArray[Int]) -> Int =
  #|(array) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn int64_array2ptr(array : FixedArray[Int64]) -> Int =
  #|(array) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn float_array2ptr(array : FixedArray[Float]) -> Int =
  #|(array) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn double_array2ptr(array : FixedArray[Double]) -> Int =
  #|(array) => 0
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "js" fn ptr2uint_array(ptr : Int) -> FixedArray[UInt] =
  #|(ptr) => []
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

///|
pub extern "js" fn ptr2int_array(ptr : Int) -> FixedArray[Int] =
  #|(ptr) => []
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

///|
pub extern "js" fn ptr2float_array(ptr : Int) -> FixedArray[Float] =
  #|(ptr) => []
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

///|
extern "js" fn ptr2uint64_array_ffi(ptr : Int) -> FixedArray[UInt64] =
  #|(ptr) => []
//  #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

///|
pub fn ptr2uint64_array(ptr : Int) -> FixedArray[UInt64] {
  // set_64_header_ffi(ptr - 4)
  // ptr2uint64_array_ffi(ptr)
  []
}

///|
extern "js" fn ptr2int64_array_ffi(ptr : Int) -> FixedArray[Int64] =
  #|(ptr) => []
//  #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

///|
pub fn ptr2int64_array(ptr : Int) -> FixedArray[Int64] {
  // set_64_header_ffi(ptr - 4)
  // ptr2int64_array_ffi(ptr)
  []
}

///|
extern "js" fn ptr2double_array_ffi(ptr : Int) -> FixedArray[Double] =
  #|(ptr) => []
//  #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

///|
pub fn ptr2double_array(ptr : Int) -> FixedArray[Double] {
  // set_64_header_ffi(ptr - 4)
  // ptr2double_array_ffi(ptr)
  []
}

///|
fn set_64_header_ffi(offset : Int) -> Unit {
  // let len = load32(offset)
  // store32(offset, len >> 1)
  // store8(offset, 241)
}

///|
pub(open) trait Any {}

///|
pub(all) struct Cleanup {
  address : Int
  size : Int
  align : Int
}
