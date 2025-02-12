/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package codegen

import (
	"bytes"
	"fmt"
	"maps"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

func PostProcess(config *config.Config, meta *metadata.Metadata) error {
	imports := meta.GetImports()
	types := getTypes(meta)

	body := &bytes.Buffer{}
	writeFuncUnpin(body)
	writeFuncNew(body, types, imports)
	writeFuncMake(body, types, imports)
	writeFuncReadMap(body, types, imports)
	writeFuncWriteMap(body, types, imports)

	header := &bytes.Buffer{}
	writePostProcessHeader(header, meta, imports)

	return writeBuffersToFile(filepath.Join(config.SourceDir, post_file), header, body)
}

func getTypes(meta *metadata.Metadata) []*metadata.TypeDefinition {
	types := utils.MapValues(meta.Types)
	sort.Slice(types, func(i, j int) bool {
		return types[i].Id < types[j].Id
	})
	return types
}

func writePostProcessHeader(b *bytes.Buffer, meta *metadata.Metadata, imports map[string]string) {
	b.WriteString("// Code generated by modus-moonbit-build. DO NOT EDIT.\n\n")
	b.WriteString(`///|
pub fn cabi_realloc(
  src_offset : Int,
  src_size : Int,
  _dst_alignment : Int,
  dst_size : Int
) -> Int {
  // malloc
  if src_offset == 0 && src_size == 0 {
    return malloc(dst_size)
  }
  // free
  if dst_size <= 0 {
    free(src_offset)
    return 0
  }
  // realloc
  let dst = malloc(dst_size)
  copy(dst, src_offset)
  free(src_offset)
  dst
}

// Generated by wit-bindgen 0.36.0. DO NOT EDIT!

// ///|
// pub extern "wasm" fn extend16(value : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.extend16_s)

// ///|
// pub extern "wasm" fn extend8(value : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.extend8_s)

///|
pub extern "wasm" fn store8(offset : Int, value : Int) =
  #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store8)

// ///|
// pub extern "wasm" fn load8_u(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load8_u)

// ///|
// pub extern "wasm" fn load8(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load8_s)

// ///|
// pub extern "wasm" fn store16(offset : Int, value : Int) =
//   #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store16)

// ///|
// pub extern "wasm" fn load16(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load16_s)

// ///|
// pub extern "wasm" fn load16_u(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load16_u)

///|
pub extern "wasm" fn store32(offset : Int, value : Int) =
  #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store)

///|
pub extern "wasm" fn load32(offset : Int) -> Int =
  #|(func (param i32) (result i32) local.get 0 i32.load)

// ///|
// pub extern "wasm" fn store64(offset : Int, value : Int64) =
//   #|(func (param i32) (param i64) local.get 0 local.get 1 i64.store)

// ///|
// pub extern "wasm" fn load64(offset : Int) -> Int64 =
//   #|(func (param i32) (result i64) local.get 0 i64.load)

// ///|
// pub extern "wasm" fn storef32(offset : Int, value : Float) =
//   #|(func (param i32) (param f32) local.get 0 local.get 1 f32.store)

// ///|
// pub extern "wasm" fn loadf32(offset : Int) -> Float =
//   #|(func (param i32) (result f32) local.get 0 f32.load)

// ///|
// pub extern "wasm" fn storef64(offset : Int, value : Double) =
//   #|(func (param i32) (param f64) local.get 0 local.get 1 f64.store)

// ///|
// pub extern "wasm" fn loadf64(offset : Int) -> Double =
//   #|(func (param i32) (result f64) local.get 0 f64.load)

// ///|
// // pub extern "wasm" fn f32_to_i32(value : Float) -> Int =
// //   #|(func (param f32) (result i32) local.get 0 f32.convert_i32_s)

// ///|
// // pub extern "wasm" fn f32_to_i64(value : Float) -> Int64 =
// //   #|(func (param f32) (result i64) local.get 0 f32.convert_i64_s)

///|
extern "wasm" fn malloc_inline(size : Int) -> Int =
  #|(func (param i32) (result i32) local.get 0 call $moonbit.malloc)

///|
pub fn malloc(size : Int) -> Int {
  let words = size / 4 + 1
  let address = malloc_inline(8 + words * 4)
  store32(address, 1)
  store32(address + 4, (words << 8) | 246)
  store8(address + words * 4 + 7, 3 - size % 4)
  address + 8
}

///|
pub extern "wasm" fn free(position : Int) =
  #|(func (param i32) local.get 0 i32.const 8 i32.sub call $moonbit.decref)

///|
pub fn copy(dest : Int, src : Int) -> Unit {
  let src = src - 8
  let dest = dest - 8
  let src_len = load32(src + 4) & 0xFFFFFF
  let dest_len = load32(dest + 4) & 0xFFFFFF
  let min = if src_len < dest_len { src_len } else { dest_len }
  copy_inline(dest, src, min)
}

///|
extern "wasm" fn copy_inline(dest : Int, src : Int, len : Int) =
  #|(func (param i32) (param i32) (param i32) local.get 0 local.get 1 local.get 2 memory.copy)

// ///|
// pub extern "wasm" fn str2ptr(str : String) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "wasm" fn ptr2str(ptr : Int) -> String =
  #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 243 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn bytes2ptr(bytes : Bytes) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn ptr2bytes(ptr : Int) -> Bytes =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn uint_array2ptr(array : FixedArray[UInt]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn uint64_array2ptr(array : FixedArray[UInt64]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn int_array2ptr(array : FixedArray[Int]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn int64_array2ptr(array : FixedArray[Int64]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn float_array2ptr(array : FixedArray[Float]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn double_array2ptr(array : FixedArray[Double]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn ptr2uint_array(ptr : Int) -> FixedArray[UInt] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn ptr2int_array(ptr : Int) -> FixedArray[Int] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn ptr2float_array(ptr : Int) -> FixedArray[Float] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// extern "wasm" fn ptr2uint64_array_ffi(ptr : Int) -> FixedArray[UInt64] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub fn ptr2uint64_array(ptr : Int) -> FixedArray[UInt64] {
//   set_64_header_ffi(ptr - 4)
//   ptr2uint64_array_ffi(ptr)
// }

// ///|
// extern "wasm" fn ptr2int64_array_ffi(ptr : Int) -> FixedArray[Int64] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub fn ptr2int64_array(ptr : Int) -> FixedArray[Int64] {
//   set_64_header_ffi(ptr - 4)
//   ptr2int64_array_ffi(ptr)
// }

// ///|
// extern "wasm" fn ptr2double_array_ffi(ptr : Int) -> FixedArray[Double] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub fn ptr2double_array(ptr : Int) -> FixedArray[Double] {
//   set_64_header_ffi(ptr - 4)
//   ptr2double_array_ffi(ptr)
// }

// ///|
// fn set_64_header_ffi(offset : Int) -> Unit {
//   let len = load32(offset)
//   store32(offset, len >> 1)
//   store8(offset, 241)
// }

// ///|
// pub(open) trait Any {}

// ///|
// pub(all) struct Cleanup {
//   address : Int
//   size : Int
//   align : Int
// }
`)

	/*
	     b.WriteString(`
	   pub fn write_map(key_type_name_ptr : Int, value_type_name_ptr : Int, keys_ptr : Int, values_ptr : Int) -> Int {
	     let key_type_name = ptr2str(key_type_name_ptr)
	     let value_type_name = ptr2str(value_type_name_ptr)
	     match (key_type_name, value_type_name) {
	   `)
	   	patterns, helpers := genWriteMapPatternsAndHelpers(meta)
	   	b.WriteString(patterns)
	   	b.WriteString(`
	     }
	   }
	   `)
	   	b.WriteString(helpers)
	*/

	// for pkg, name := range imports {
	// 	gmlPrintf("GML: codegen/postprocess.go: writePostProcessHeader: imports['%v']='%v'", pkg, name)
	// }

	if _, ok := imports["@time"]; ok {
		b.WriteString(`
///|
pub fn zoned_date_time_from_unix_seconds_and_nanos(second : Int64, nanos : Int64) -> @time.ZonedDateTime!Error {
  let nanosecond = (nanos % 1_000_000_000).to_int()
  @time.unix!(second, nanosecond~)
}

///|
pub fn duration_from_nanos(nanoseconds : Int64) -> @time.Duration!Error {
  @time.Duration::of!(nanoseconds~)
}
`)
	}

	// b.WriteString("package main\n\n")
	// b.WriteString("import (\n")
	// b.WriteString("\t\"unsafe\"\n")
	// for pkg, name := range imports {
	// 	if pkg != "" && pkg != "unsafe" && pkg != meta.Module {
	// 		if pkg == name || strings.HasSuffix(pkg, "/"+name) {
	// 			b.WriteString(fmt.Sprintf("\t\"%s\"\n", pkg))
	// 		} else {
	// 			b.WriteString(fmt.Sprintf("\t%s \"%s\"\n", name, pkg))
	// 		}
	// 	}
	// }
	// b.WriteString(")\n")
}

func genWriteMapPatternsAndHelpers(meta *metadata.Metadata) (string, string) {
	var patterns []string
	var helpers []string
	type keyValuePair struct {
		key   string
		value string
	}
	processed := map[string]*keyValuePair{}

	for typ := range meta.Types {
		if utils.IsMapType(typ) {
			keyTypeName, valueTypeName := utils.GetMapSubtypes(typ)
			mapTypeName := fmt.Sprintf("Map[%v, %v]", keyTypeName, valueTypeName)
			if _, ok := processed[mapTypeName]; ok {
				continue
			}
			processed[mapTypeName] = &keyValuePair{key: keyTypeName, value: valueTypeName}
		}
	}

	// Output the patterns and helpers in sorted order.
	keys := slices.Sorted(maps.Keys(processed))
	for _, kvPair := range keys {
		keyTypeName := processed[kvPair].key
		valueTypeName := processed[kvPair].value
		patterns = append(patterns, fmt.Sprintf("    (%q, %q) => write_map_helper_%v(keys_ptr, values_ptr)", keyTypeName, valueTypeName, len(helpers)))
		helpers = append(helpers, genWriteMapHelperFuncs(len(helpers), keyTypeName, valueTypeName))
	}

	patterns = append(patterns, "    _ => 0")
	return strings.Join(patterns, "\n"), strings.Join(helpers, "\n")
}

func genWriteMapHelperFuncs(fnNum int, keyTypeName, valueTypeName string) string {
	var helpers []string
	keyFn := fmt.Sprintf(`extern "wasm" fn ptr2write_map_helper_keys_%v(ptr : Int) -> Array[%v] =
   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)
`, fnNum, keyTypeName)

	valueFn := fmt.Sprintf(`extern "wasm" fn ptr2write_map_helper_values_%v(ptr : Int) -> Array[%v] =
   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)
`, fnNum, valueTypeName)

	mapFn := fmt.Sprintf(`extern "wasm" fn write_map_helper_map_%v_to_ptr(m : Map[%v, %v]) -> Int =
   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)
`, fnNum, keyTypeName, valueTypeName)

	helperFn := fmt.Sprintf(`fn write_map_helper_%v(keys_ptr: Int, values_ptr: Int) -> Int {
	  let keys = ptr2write_map_helper_keys_%[1]v(keys_ptr)
	  let values = ptr2write_map_helper_values_%[1]v(values_ptr)
    let m : Map[%[2]v, %[3]v] = Map::new(capacity=keys.length())
    for i in 0..<keys.length() {
      m[keys[i]] = values[i]
    }
    write_map_helper_map_%[1]v_to_ptr(m)
}
`, fnNum, keyTypeName, valueTypeName)
	helpers = append(helpers, keyFn, valueFn, mapFn, helperFn)
	return strings.Join(helpers, "\n\n")
}

func writeFuncUnpin(b *bytes.Buffer) {
	// b.WriteString(`
	// var __pins = make(map[unsafe.Pointer]int)
	//
	// //go:export __unpin
	// func __unpin(p unsafe.Pointer) {
	// 	n := __pins[p]
	// 	if n == 1 {
	// 		delete(__pins, p)
	// 	} else {
	// 		__pins[p] = n - 1
	// 	}
	// }`)
	//
	// 	b.WriteString("\n")
}

func writeFuncNew(b *bytes.Buffer, types []*metadata.TypeDefinition, imports map[string]string) {
	// 	buf := &bytes.Buffer{}
	// 	found := false
	//
	// 	buf.WriteString(`
	// //go:export __new
	// func __new(id int) unsafe.Pointer {
	// `)
	// 	buf.WriteString("\tswitch id {\n")
	// 	for _, t := range types {
	// 		if utils.IsSliceType(t.Name) && utils.IsMapType(t.Name) {
	// 			continue
	// 		}
	//
	// 		ptrName := utils.GetNameForType(t.Name, imports)
	// 		elementName := utils.GetUnderlyingType(ptrName)
	// 		found = true
	// 		buf.WriteString(fmt.Sprintf(`	case %d:
	// 		o := new(%s)
	// 		p := unsafe.Pointer(o)
	// 		__pins[p]++
	// 		return p
	// `, t.Id, elementName))
	// 	}
	// 	buf.WriteString("\t}\n\n")
	// 	buf.WriteString("\treturn nil\n}\n")
	//
	// 	if found {
	// 		_, _ = buf.WriteTo(b)
	// 	}
}

func writeFuncMake(b *bytes.Buffer, types []*metadata.TypeDefinition, imports map[string]string) {
	// b.WriteString(`
	// //go:export __make
	// func __make(id, size int) unsafe.Pointer {
	// 	switch id {
	// 	case 1:
	// 		o := make([]byte, size)
	// 		p := unsafe.Pointer(&o)
	// 		__pins[p]++
	// 		return p
	// 	case 2:
	// 		o := string(make([]byte, size))
	// 		p := unsafe.Pointer(&o)
	// 		__pins[p]++
	// 		return p
	// `)
	//
	// 	for _, t := range types {
	// 		name := utils.GetNameForType(t.Name, imports)
	// 		if utils.IsSliceType(name) || utils.IsMapType(name) {
	// 			b.WriteString(fmt.Sprintf(`	case %d:
	// 		o := make(%s, size)
	// 		p := unsafe.Pointer(&o)
	// 		__pins[p]++
	// 		return p
	// `, t.Id, name))
	// 		}
	// 	}
	//
	// 	b.WriteString("\t}\n\n")
	// 	b.WriteString("\treturn nil\n}\n")
}

func writeFuncReadMap(b *bytes.Buffer, types []*metadata.TypeDefinition, imports map[string]string) {
	// 	buf := &bytes.Buffer{}
	// 	found := false
	//
	// 	buf.WriteString(`
	// //go:export __read_map
	// func __read_map(id int, m unsafe.Pointer) uint64 {
	// `)
	// 	buf.WriteString("\tswitch id {\n")
	// 	for _, t := range types {
	// 		if utils.IsMapType(t.Name) {
	// 			found = true
	// 			typeName := utils.GetNameForType(t.Name, imports)
	// 			buf.WriteString(fmt.Sprintf(`	case %d:
	// 		return __doReadMap(*(*%s)(m))
	// `, t.Id, typeName))
	// 		}
	// 	}
	// 	buf.WriteString("\t}\n\n")
	// 	buf.WriteString("\treturn 0\n}\n")
	//
	// 	buf.WriteString(`
	// func __doReadMap[M ~map[K]V, K comparable, V any](m M) uint64 {
	// 	size := len(m)
	// 	keys := make([]K, size)
	// 	values := make([]V, size)
	//
	// 	i := 0
	// 	for k, v := range m {
	// 		keys[i] = k
	// 		values[i] = v
	// 		i++
	// 	}
	//
	// 	pKeys := uint32(uintptr(unsafe.Pointer(&keys)))
	// 	pValues := uint32(uintptr(unsafe.Pointer(&values)))
	// 	return uint64(pKeys)<<32 | uint64(pValues)
	// }
	// `)
	//
	// 	if found {
	// 		_, _ = buf.WriteTo(b)
	// 	}
}

func writeFuncWriteMap(b *bytes.Buffer, types []*metadata.TypeDefinition, imports map[string]string) {
	// 	buf := &bytes.Buffer{}
	// 	found := false
	//
	// 	buf.WriteString(`
	// //go:export __write_map
	// func __write_map(id int, m, keys, values unsafe.Pointer) {
	// `)
	// 	buf.WriteString("\tswitch id {\n")
	// 	for _, t := range types {
	// 		if strings.HasPrefix(t.Name, "map[") {
	// 			found = true
	// 			typeName := utils.GetNameForType(t.Name, imports)
	// 			kt, vt := utils.GetMapSubtypes(t.Name)
	// 			keyTypeName := utils.GetNameForType(kt, imports)
	// 			valTypeName := utils.GetNameForType(vt, imports)
	// 			buf.WriteString(fmt.Sprintf(`	case %d:
	// 		__doWriteMap(*(*%s)(m), *(*[]%s)(keys), *(*[]%s)(values))
	// `, t.Id, typeName, keyTypeName, valTypeName))
	// 		}
	// 	}
	// 	buf.WriteString("\t}\n")
	// 	buf.WriteString("}\n")
	//
	// 	buf.WriteString(`
	// func __doWriteMap[M ~map[K]V, K comparable, V any](m M, keys[]K, values[]V) {
	// 	for i := 0; i < len(keys); i++ {
	// 		m[keys[i]] = values[i]
	// 	}
	// }
	// `)
	//
	// 	if found {
	// 		_, _ = buf.WriteTo(b)
	// 	}
}
