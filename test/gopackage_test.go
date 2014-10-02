/*
 * gopackage
 * https://github.com/chrisenytc/gopackage
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package main

import (
	"github.com/chrisenytc/gopackage"
	"testing"
)

func TestReturnMsg(t *testing.T) {
	msg := "Testing Message func"
	testMsg := gopackage.ReturnMsg(msg)
	if msg != testMsg {
		t.Errorf("'%s' is different of '%s'", msg, testMsg)
	}
}
