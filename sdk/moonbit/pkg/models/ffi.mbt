// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
fn host_get_model_info(model_name : String) -> ModelInfo = "modus_models" "getModelInfo"

// //go:noescape
// //go:wasmimport modus_models getModelInfo
// func _hostGetModelInfo(modelName *string) unsafe.Pointer
//
// //modus:import modus_models getModelInfo
// func hostGetModelInfo(modelName *string) *ModelInfo {
// 	info := _hostGetModelInfo(modelName)
// 	if info == nil {
// 		return nil
// 	}
// 	return (*ModelInfo)(info)
// }
//
// //go:noescape
// //go:wasmimport modus_models invokeModel
// func hostInvokeModel(modelName *string, input *string) *string
//
