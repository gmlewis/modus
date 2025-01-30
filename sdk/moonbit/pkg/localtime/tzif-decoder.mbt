///|
struct TZifHeader {
  magic : Int
  version : Int
  tzh_timecnt : Int
}

///|
fn decode_tzif_header(
  data : FixedArray[Byte],
  offset : Int
) -> (TZifHeader, Int) {
  let magic = to_int_be(data, offset, offset + 4)
  let version = to_int_be(data, offset + 4, offset + 5)
  let tzh_timecnt = to_int_be(data, offset + 32, offset + 36)
  let header = { magic, version, tzh_timecnt }
  (header, offset + 44)
}

///|
fn decode_tzif_data(
  data : FixedArray[Byte],
  offset : Int,
  header : TZifHeader
) -> Int {
  let offset = offset + header.tzh_timecnt * 4
  let offset = offset + header.tzh_timecnt
  to_int_be(data, offset, offset + 4)
}

///|
fn load_location_from_tzdata(
  name : String,
  data : FixedArray[Byte]
) -> @time.Zone!Error {
  let (header, byteoffset) = decode_tzif_header(data, 0)
  let want = 0x545A6966
  if header.magic != want {
    raise LocaltimeError("invalid TZif magic: \{header.magic}, want \{want}")
  }
  let want = 2
  if header.version != want {
    raise LocaltimeError(
      "invalid TZif version: \{header.version}, want \{want}",
    )
  }
  let offset = decode_tzif_data(data, byteoffset, header)
  @time.fixed_zone!(name, offset)
}

///| Helper function to convert byte array to integer (big-endian)
fn to_int_be(bytes : FixedArray[Byte], start : Int, end : Int) -> Int {
  let mut result = 0
  for i in start..<end {
    result = (result << 8) | bytes[i].to_int()
  }
  result
}
