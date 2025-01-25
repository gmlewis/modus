/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

import (
	"fmt"
)

func HelloMaps0Items() map[string]string {
	return map[string]string{}
}

func HelloMaps1Item() map[string]string {
	return map[string]string{"a": "apple"}
}

func HelloMaps2Items() map[string]string {
	return map[string]string{"b": "banana", "c": "cherry"}
}

func HelloMaps3Items() map[string]string {
	return map[string]string{"d": "date", "e": "elderberry", "f": "fig"}
}

func HelloMaps4Items() map[string]string {
	return map[string]string{"g": "grape", "h": "honeydew", "i": "imbe", "j": "jackfruit"}
}

func HelloMaps5Items() map[string]string {
	return map[string]string{"k": "kiwi", "l": "lemon", "m": "mango", "n": "nectarine", "o": "orange"}
}

func HelloMaps6Items() map[string]string {
	return map[string]string{
		"p": "papaya",
		"q": "quince",
		"r": "raspberry",
		"s": "strawberry",
		"t": "tangerine",
		"u": "ugli",
	}
}

func HelloMapsNItems(n int) map[string]string {
	m := map[string]string{}
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key%v", i+1)
		value := fmt.Sprintf("value%v", i+1)
		m[key] = value
	}
	return m
}
