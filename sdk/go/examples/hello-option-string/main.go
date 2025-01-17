/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gmlewis/modus/sdk/go/pkg/console"
)

func HelloOptionString(nilString, emptyString bool) *string {
     switch {
     case nilString:
     return nil
     case emptyString:
     s := ""
     return &s
     default:
     s := "Hello, World!"
     return &s
     }
}

