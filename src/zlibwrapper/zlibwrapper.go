/* ----------------------------------------------------------------------------
* This file was automatically generated by SWIG (http://www.swig.org).
* Version 3.0.7
*
* This file is not intended to be easily readable and contains a number of
* coding conventions designed to improve portability and efficiency. Do not make
* changes to this file unless you know what you are doing--modify the SWIG
* interface file instead.
* ----------------------------------------------------------------------------- */
// source: zlibwrapper.i
package zlibwrapper

/*
#cgo windows LDFLAGS: -L"E:\MyWork\golang\ConfRobber\pkg\windows_amd64" -lzlibwrapper -lzlibstatic
#cgo windows CFLAGS: -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#define intgo swig_intgo
typedef void *swig_voidp;
#include <stdint.h>
typedef long long intgo;
typedef unsigned long long uintgo;
typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;
typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
extern void _wrap_Swig_free_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern uintptr_t _wrap_new_BinaryData__SWIG_0_zlibwrapper_dd73c011c24953e7(void);
extern uintptr_t _wrap_new_BinaryData__SWIG_1_zlibwrapper_dd73c011c24953e7(swig_type_1 arg1);
extern swig_type_2 _wrap_BinaryData_size_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern swig_type_3 _wrap_BinaryData_capacity_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern void _wrap_BinaryData_reserve_zlibwrapper_dd73c011c24953e7(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_BinaryData_isEmpty_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern void _wrap_BinaryData_clear_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern void _wrap_BinaryData_add_zlibwrapper_dd73c011c24953e7(uintptr_t arg1, char arg2);
extern char _wrap_BinaryData_get_zlibwrapper_dd73c011c24953e7(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_BinaryData_set_zlibwrapper_dd73c011c24953e7(uintptr_t arg1, swig_intgo arg2, char arg3);
extern void _wrap_delete_BinaryData_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern uintptr_t _wrap_uncompress_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
extern uintptr_t _wrap_compress_zlibwrapper_dd73c011c24953e7(uintptr_t arg1);
#undef intgo
*/
import "C"
import "unsafe"
import _ "runtime/cgo"
import "sync"

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte
type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0))
}

type SwigcptrBinaryData uintptr

func (p SwigcptrBinaryData) Swigcptr() uintptr {
	return (uintptr)(p)
}
func (p SwigcptrBinaryData) SwigIsBinaryData() {
}
func NewBinaryData__SWIG_0() (_swig_ret BinaryData) {
	var swig_r BinaryData
	swig_r = (BinaryData)(SwigcptrBinaryData(C._wrap_new_BinaryData__SWIG_0_zlibwrapper_dd73c011c24953e7()))
	return swig_r
}
func NewBinaryData__SWIG_1(arg1 int64) (_swig_ret BinaryData) {
	var swig_r BinaryData
	_swig_i_0 := arg1
	swig_r = (BinaryData)(SwigcptrBinaryData(C._wrap_new_BinaryData__SWIG_1_zlibwrapper_dd73c011c24953e7(C.swig_type_1(_swig_i_0))))
	return swig_r
}
func NewBinaryData(a ...interface{}) BinaryData {
	argc := len(a)
	if argc == 0 {
		return NewBinaryData__SWIG_0()
	}
	if argc == 1 {
		return NewBinaryData__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}
func (arg1 SwigcptrBinaryData) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_BinaryData_size_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0)))
	return swig_r
}
func (arg1 SwigcptrBinaryData) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_BinaryData_capacity_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0)))
	return swig_r
}
func (arg1 SwigcptrBinaryData) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_BinaryData_reserve_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}
func (arg1 SwigcptrBinaryData) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_BinaryData_isEmpty_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0)))
	return swig_r
}
func (arg1 SwigcptrBinaryData) Clear() {
	_swig_i_0 := arg1
	C._wrap_BinaryData_clear_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0))
}
func (arg1 SwigcptrBinaryData) Add(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_BinaryData_add_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}
func (arg1 SwigcptrBinaryData) Get(arg2 int) (_swig_ret byte) {
	var swig_r byte
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (byte)(C._wrap_BinaryData_get_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}
func (arg1 SwigcptrBinaryData) Set(arg2 int, arg3 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_BinaryData_set_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.char(_swig_i_2))
}
func DeleteBinaryData(arg1 BinaryData) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_BinaryData_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0))
}

type BinaryData interface {
	Swigcptr() uintptr
	SwigIsBinaryData()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 byte)
	Get(arg2 int) (_swig_ret byte)
	Set(arg2 int, arg3 byte)
}

func Uncompress(arg1 BinaryData) (_swig_ret BinaryData) {
	var swig_r BinaryData
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (BinaryData)(SwigcptrBinaryData(C._wrap_uncompress_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0))))
	return swig_r
}
func Compress(arg1 BinaryData) (_swig_ret BinaryData) {
	var swig_r BinaryData
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (BinaryData)(SwigcptrBinaryData(C._wrap_compress_zlibwrapper_dd73c011c24953e7(C.uintptr_t(_swig_i_0))))
	return swig_r
}
