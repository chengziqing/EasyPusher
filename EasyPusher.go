/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.8
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: EasyPusher.i

package EasyPusher

/*

#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -llibEasyPusher

#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
extern void _wrap_Swig_free_EasyPusher_993691126ef1af58(uintptr_t arg1);
extern uintptr_t _wrap_GO_EasyPusher_Activate_EasyPusher_993691126ef1af58(swig_type_1 arg1);
extern uintptr_t _wrap_GO_EasyPusher_Create_EasyPusher_993691126ef1af58(void);
extern uintptr_t _wrap_GO_EasyPusher_StartStream_EasyPusher_993691126ef1af58(uintptr_t arg1, swig_type_2 arg2, uintptr_t arg3, swig_type_3 arg4, swig_type_4 arg5, swig_type_5 arg6, uintptr_t arg7, uintptr_t arg8, uintptr_t arg9);
extern uintptr_t _wrap_GO_EasyPusher_StopStream_EasyPusher_993691126ef1af58(uintptr_t arg1);
extern uintptr_t _wrap_GO_EasyPusher_PushFrame_EasyPusher_993691126ef1af58(uintptr_t arg1, uintptr_t arg2);
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
	C._wrap_Swig_free_EasyPusher_993691126ef1af58(C.uintptr_t(_swig_i_0))
}

func GO_EasyPusher_Activate(arg1 string) (_swig_ret Easy_I32) {
	var swig_r Easy_I32
	_swig_i_0 := arg1
	swig_r = (Easy_I32)(SwigcptrEasy_I32(C._wrap_GO_EasyPusher_Activate_EasyPusher_993691126ef1af58(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func GO_EasyPusher_Create() (_swig_ret Easy_Pusher_Handle) {
	var swig_r Easy_Pusher_Handle
	swig_r = (Easy_Pusher_Handle)(SwigcptrEasy_Pusher_Handle(C._wrap_GO_EasyPusher_Create_EasyPusher_993691126ef1af58()))
	return swig_r
}

func GO_EasyPusher_StartStream(arg1 Easy_Pusher_Handle, arg2 string, arg3 Easy_U16, arg4 string, arg5 string, arg6 string, arg7 EASY_MEDIA_INFO_T, arg8 Easy_U32, arg9 Easy_Bool) (_swig_ret Easy_U32) {
	var swig_r Easy_U32
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	_swig_i_6 := arg7.Swigcptr()
	_swig_i_7 := arg8.Swigcptr()
	_swig_i_8 := arg9.Swigcptr()
	swig_r = (Easy_U32)(SwigcptrEasy_U32(C._wrap_GO_EasyPusher_StartStream_EasyPusher_993691126ef1af58(C.uintptr_t(_swig_i_0), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)), C.uintptr_t(_swig_i_2), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_3)), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_4)), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_5)), C.uintptr_t(_swig_i_6), C.uintptr_t(_swig_i_7), C.uintptr_t(_swig_i_8))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg4
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg5
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg6
	}
	return swig_r
}

func GO_EasyPusher_StopStream(arg1 Easy_Pusher_Handle) (_swig_ret Easy_U32) {
	var swig_r Easy_U32
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (Easy_U32)(SwigcptrEasy_U32(C._wrap_GO_EasyPusher_StopStream_EasyPusher_993691126ef1af58(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func GO_EasyPusher_PushFrame(arg1 Easy_Pusher_Handle, arg2 EASY_AV_Frame) (_swig_ret Easy_U32) {
	var swig_r Easy_U32
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Easy_U32)(SwigcptrEasy_U32(C._wrap_GO_EasyPusher_PushFrame_EasyPusher_993691126ef1af58(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

type SwigcptrEasy_U16 uintptr
type Easy_U16 interface {
	Swigcptr() uintptr
}

func (p SwigcptrEasy_U16) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrEASY_MEDIA_INFO_T uintptr
type EASY_MEDIA_INFO_T interface {
	Swigcptr() uintptr
}

func (p SwigcptrEASY_MEDIA_INFO_T) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrEasy_I32 uintptr
type Easy_I32 interface {
	Swigcptr() uintptr
}

func (p SwigcptrEasy_I32) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrEasy_Pusher_Handle uintptr
type Easy_Pusher_Handle interface {
	Swigcptr() uintptr
}

func (p SwigcptrEasy_Pusher_Handle) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrEasy_Bool uintptr
type Easy_Bool interface {
	Swigcptr() uintptr
}

func (p SwigcptrEasy_Bool) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrEASY_AV_Frame uintptr
type EASY_AV_Frame interface {
	Swigcptr() uintptr
}

func (p SwigcptrEASY_AV_Frame) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrEasy_U32 uintptr
type Easy_U32 interface {
	Swigcptr() uintptr
}

func (p SwigcptrEasy_U32) Swigcptr() uintptr {
	return uintptr(p)
}