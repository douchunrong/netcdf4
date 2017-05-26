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

// WriteFloat32s writes data as the entire data for variable v.
func (v Var) WriteFloat32s(data []float32) error {
	if err := okData(v, FLOAT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_float(C.int(v.ds), C.int(v.id), (*C.float)(unsafe.Pointer(&data[0]))))
}

// ReadFloat32s reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadFloat32s(data []float32) error {
	if err := okData(v, FLOAT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_float(C.int(v.ds), C.int(v.id), (*C.float)(unsafe.Pointer(&data[0]))))
}

// WriteFloat32s sets the value of attribute a to val.
func (a Attr) WriteFloat32s(val []float32) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_float(C.int(a.v.ds), C.int(a.v.id), cname,
		C.nc_type(FLOAT), C.size_t(len(val)), (*C.float)(unsafe.Pointer(&val[0]))))
}

// ReadFloat32s reads the entire attribute value into val.
func (a Attr) ReadFloat32s(val []float32) (err error) {
	if err := okData(a, FLOAT, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_float(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.float)(unsafe.Pointer(&val[0]))))
	return
}

// Float32sReader is a interface that allows reading a sequence of values of fixed length.
type Float32sReader interface {
	Len() (n uint64, err error)
	ReadFloat32s(val []float32) (err error)
}

// GetFloat32s reads the entire data in r and returns it.
func GetFloat32s(r Float32sReader) (data []float32, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]float32, n)
	err = r.ReadFloat32s(data)
	return
}

// testReadFloat32s writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteFloat32s(v Var, n uint64) error {
	data := make([]float32, n)
	for i := 0; i < int(n); i++ {
		data[i] = float32(i + 10)
	}
	return v.WriteFloat32s(data)
}

// testReadFloat32s reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadFloat32s(v Var, n uint64) error {
	data := make([]float32, n)
	if err := v.ReadFloat32s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := float32(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v\n", i, data[i], val)
		}
	}
	return nil
}