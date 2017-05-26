// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.go
// DO NOT EDIT (except nc_double.go).

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteUint16s writes data as the entire data for variable v.
func (v Var) WriteUint16s(data []uint16) error {
	if err := okData(v, USHORT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_ushort(C.int(v.ds), C.int(v.id), (*C.ushort)(unsafe.Pointer(&data[0]))))
}

// ReadUint16s reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadUint16s(data []uint16) error {
	if err := okData(v, USHORT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_ushort(C.int(v.ds), C.int(v.id), (*C.ushort)(unsafe.Pointer(&data[0]))))
}

// WriteUint16s sets the value of attribute a to val.
func (a Attr) WriteUint16s(val []uint16) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_ushort(C.int(a.v.ds), C.int(a.v.id), cname,
		C.nc_type(USHORT), C.size_t(len(val)), (*C.ushort)(unsafe.Pointer(&val[0]))))
}

// ReadUint16s reads the entire attribute value into val.
func (a Attr) ReadUint16s(val []uint16) (err error) {
	if err := okData(a, USHORT, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_ushort(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.ushort)(unsafe.Pointer(&val[0]))))
	return
}

// Uint16sReader is a interface that allows reading a sequence of values of fixed length.
type Uint16sReader interface {
	Len() (n uint64, err error)
	ReadUint16s(val []uint16) (err error)
}

// GetUint16s reads the entire data in r and returns it.
func GetUint16s(r Uint16sReader) (data []uint16, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]uint16, n)
	err = r.ReadUint16s(data)
	return
}

// testReadUint16s writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteUint16s(v Var, n uint64) error {
	data := make([]uint16, n)
	for i := 0; i < int(n); i++ {
		data[i] = uint16(i + 10)
	}
	return v.WriteUint16s(data)
}

// testReadUint16s reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadUint16s(v Var, n uint64) error {
	data := make([]uint16, n)
	if err := v.ReadUint16s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := uint16(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v\n", i, data[i], val)
		}
	}
	return nil
}