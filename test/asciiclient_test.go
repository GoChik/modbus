// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package test

import (
	"io/ioutil"
	"testing"

	"github.com/gochik/modbus"
)

func TestASCIIClient(t *testing.T) {
	f, err := ioutil.TempFile("", "asciidevice*")
	if err != nil {
		t.Fatal(err)
	}
	// Diagslave does not support broadcast id.
	handler := modbus.NewASCIIClientHandler(f)
	handler.SetSlave(17)
	ClientTestAll(t, modbus.NewClient(handler))
}

func TestASCIIClientAdvancedUsage(t *testing.T) {
	f, err := ioutil.TempFile("", "asciidevice_advanced*")
	if err != nil {
		t.Fatal(err)
	}
	handler := modbus.NewASCIIClientHandler(f)
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadDiscreteInputs(15, 2)
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
	results, err = client.ReadWriteMultipleRegisters(0, 2, 2, 2, []byte{1, 2, 3, 4})
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
}
