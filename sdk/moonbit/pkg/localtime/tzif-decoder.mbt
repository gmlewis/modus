///|
struct TimeZone {
  name : String
  offset : Int
  transition_times : Array[Int]
  type_indices : Array[Int]
}

///|
struct TZifHeader {
  magic : String
  version : String
  tzh_ttisgmtcnt : Int
  tzh_ttisstdcnt : Int
  tzh_leapcnt : Int
  tzh_timecnt : Int
  tzh_typecnt : Int
  tzh_charcnt : Int
}

///|
struct TZifData {
  transition_times : Array[Int]
  type_indices : Array[Int]
  timezone_names : Array[String]
  offsets : Array[Int]
}

///|
fn decode_tzif_header(data : Array[Byte], offset : Int) -> (TZifHeader, Int) {
  let magic = @base64.array2str(data[offset:offset + 4])
  let version = @base64.array2str(data[offset + 4:offset + 5])
  let tzh_ttisgmtcnt = to_int_be(data[offset + 20:offset + 24])
  let tzh_ttisstdcnt = to_int_be(data[offset + 24:offset + 28])
  let tzh_leapcnt = to_int_be(data[offset + 28:offset + 32])
  let tzh_timecnt = to_int_be(data[offset + 32:offset + 36])
  let tzh_typecnt = to_int_be(data[offset + 36:offset + 40])
  let tzh_charcnt = to_int_be(data[offset + 40:offset + 44])
  let header = {
    magic,
    version,
    tzh_ttisgmtcnt,
    tzh_ttisstdcnt,
    tzh_leapcnt,
    tzh_timecnt,
    tzh_typecnt,
    tzh_charcnt,
  }
  (header, offset + 44)
}

///|
fn decode_tzif_data(
  data : Array[Byte],
  offset : Int,
  header : TZifHeader
) -> (TZifData, Int) {
  let transition_times = Array::make(header.tzh_timecnt, 0)
  let type_indices = Array::make(header.tzh_timecnt, 0)
  let timezone_names = Array::make(header.tzh_typecnt, "")
  let offsets = Array::make(header.tzh_typecnt, 0)
  for i in 0..<header.tzh_timecnt {
    transition_times[i] = to_int_be(data[offset + i * 4:offset + (i + 1) * 4])
  }
  let offset = offset + header.tzh_timecnt * 4
  for i in 0..<header.tzh_timecnt {
    type_indices[i] = data[offset + i].to_int()
  }
  let offset = offset + header.tzh_timecnt
  for i in 0..<header.tzh_typecnt {
    let offset_index = offset + i * 6
    let gmtoff = to_int_be(data[offset_index:offset_index + 4])
    let isdst = data[offset_index + 4].to_int()
    let abbrind = data[offset_index + 5].to_int()
    offsets[i] = gmtoff
    timezone_names[i] = @base64.array2str(
      data[offset + header.tzh_typecnt * 6 + abbrind:offset +
      header.tzh_typecnt * 6 +
      abbrind +
      1],
    )
  }
  let offset = offset + header.tzh_typecnt * 6 + header.tzh_charcnt
  let tzif_data = { transition_times, type_indices, timezone_names, offsets }
  (tzif_data, offset)
}

///|
pub fn load_location_from_tzdata(name : String, data : Array[Byte]) -> TimeZone {
  let (header, offset) = decode_tzif_header(data, 0)
  let (tzif_data, _) = decode_tzif_data(data, offset, header)
  {
    name,
    offset: tzif_data.offsets[0],
    transition_times: tzif_data.transition_times,
    type_indices: tzif_data.type_indices,
  }
}

// Helper function to convert byte array to integer (big-endian)
///|
fn to_int_be(bytes : ArrayView[Byte]) -> Int {
  let mut result = 0
  for i in 0..<bytes.length() {
    result = (result << 8) | bytes[i].to_int()
  }
  result
}

// // Example usage
// ///|
// fn main() {
//   let tzif_data : Array[Byte] = load_tzif_file("example.tzif") // Assume this function loads the file into a byte array
//   let timezone = load_location_from_tzdata("Example", tzif_data)
//   println("Timezone Name: ", timezone.name)
//   println("Timezone Offset: ", timezone.offset)
//   println("Transition Times: ", timezone.transition_times)
//   println("Type Indices: ", timezone.type_indices)
// }
