/*
 * gopackage
 * https://github.com/chrisenytc/gopackage
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package main

/*
 * Dependencies
 */

import (
	"fmt"
	"github.com/chrisenytc/gopackage/lib"
)

/*
 * Main application
 */

func main() {
	msg := lib.ReturnMsg("Welcome to Go Package")
	fmt.Println(msg)
}
