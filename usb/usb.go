/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: usb.i

package usb

/*
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
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
extern void _wrap_Swig_free_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_usb_4f19f1d7d83a7073(swig_intgo arg1);

#cgo pkg-config: libusb-1.0
#cgo pkg-config: hidapi
#include <libusb.h>
#include <hidapi.h> 
extern void _wrap_TransferStatus_transferring_set_usb_4f19f1d7d83a7073(uintptr_t arg1, _Bool arg2);
extern _Bool _wrap_TransferStatus_transferring_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_TransferStatus_status_code_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_TransferStatus_status_code_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_TransferStatus_status_error_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_type_1 arg2);
extern swig_type_2 _wrap_TransferStatus_status_error_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_TransferStatus_buf_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_voidp arg2);
extern swig_voidp _wrap_TransferStatus_buf_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap_new_TransferStatus_usb_4f19f1d7d83a7073(void);
extern void _wrap_delete_TransferStatus_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap__swig_NewDirectorHIDPacketHandlerHIDPacketHandler_usb_4f19f1d7d83a7073(int);
extern void _wrap_DeleteDirectorHIDPacketHandler_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_delete_HIDPacketHandler_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_HIDPacketHandler_handleIncomingPacket_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_voidp arg2);
extern swig_intgo _wrap_PROTOCOL_UNKNOWN_Device_usb_4f19f1d7d83a7073(void);
extern swig_intgo _wrap_HALFKAY_Device_usb_4f19f1d7d83a7073(void);
extern swig_intgo _wrap_DFU_Device_usb_4f19f1d7d83a7073(void);
extern swig_intgo _wrap_FORMAT_UNKNOWN_Device_usb_4f19f1d7d83a7073(void);
extern swig_intgo _wrap_HEX_Device_usb_4f19f1d7d83a7073(void);
extern swig_intgo _wrap_BIN_Device_usb_4f19f1d7d83a7073(void);
extern void _wrap_Device_file_format_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_file_format_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_protocol_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_protocol_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_packet_handler_set_usb_4f19f1d7d83a7073(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Device_packet_handler_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_bootloader_set_usb_4f19f1d7d83a7073(uintptr_t arg1, _Bool arg2);
extern _Bool _wrap_Device_bootloader_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_pid_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_pid_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_port_number_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_port_number_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_vid_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_vid_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_fingerprint_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_fingerprint_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_friendly_name_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_type_3 arg2);
extern swig_type_4 _wrap_Device_friendly_name_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_model_set_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_type_5 arg2);
extern swig_type_6 _wrap_Device_model_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap_Device_usb_transfer_usb_4f19f1d7d83a7073(uintptr_t arg1, char arg2, char arg3, short arg4, short arg5, swig_voidp arg6, short arg7, swig_intgo arg8);
extern _Bool _wrap_Device_hid_listen_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern _Bool _wrap_Device_hid_open_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern _Bool _wrap_Device_usb_claim_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern swig_intgo _wrap_Device_check_connected_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern swig_intgo _wrap_Device_send_hid_packet_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_voidp arg2, swig_intgo arg3);
extern swig_intgo _wrap_Device_usb_auto_detach_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern swig_intgo _wrap_Device_usb_claim_interface_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_usb_set_configuration_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Device_get_firmware_format_usb_4f19f1d7d83a7073(swig_intgo arg1);
extern swig_intgo _wrap_Device_get_flashing_protocol_usb_4f19f1d7d83a7073(swig_intgo arg1);
extern _Bool _wrap_Device_is_bootloader_usb_4f19f1d7d83a7073(swig_intgo arg1);
extern _Bool _wrap_Device_is_interesting_usb_4f19f1d7d83a7073(swig_intgo arg1, swig_intgo arg2);
extern swig_type_7 _wrap_Device_get_friendly_name_usb_4f19f1d7d83a7073(swig_intgo arg1);
extern swig_type_8 _wrap_Device_get_model_usb_4f19f1d7d83a7073(swig_intgo arg1);
extern swig_type_9 _wrap_Device_get_dfu_string_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_Device_close_hid_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Device_usb_close_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap_new_Device_usb_4f19f1d7d83a7073(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern void _wrap_delete_Device_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap__swig_NewDirectorEventHandlerEventHandler_usb_4f19f1d7d83a7073(int);
extern void _wrap_DeleteDirectorEventHandler_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_delete_EventHandler_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_EventHandler_handleUSBConnectionEvent_usb_4f19f1d7d83a7073(uintptr_t arg1, _Bool arg2, uintptr_t arg3);
extern void _wrap_Enumerator_EventObject_set_usb_4f19f1d7d83a7073(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Enumerator_EventObject_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern uintptr_t _wrap_new_Enumerator_usb_4f19f1d7d83a7073(void);
extern void _wrap_delete_Enumerator_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Enumerator_ListenDevices_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Enumerator_StopListenDevices_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Enumerator_HandleEvents_usb_4f19f1d7d83a7073(uintptr_t arg1);
extern void _wrap_Enumerator_Devices_set_usb_4f19f1d7d83a7073(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Enumerator_Devices_get_usb_4f19f1d7d83a7073(uintptr_t arg1);
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


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0)))
	return swig_r
}

const HID_PACKET_SIZE int = 33
const USB_BUFFER_SIZE int = 2048
type SwigcptrTransferStatus uintptr

func (p SwigcptrTransferStatus) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrTransferStatus) SwigIsTransferStatus() {
}

func (arg1 SwigcptrTransferStatus) SetTransferring(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_TransferStatus_transferring_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrTransferStatus) GetTransferring() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_TransferStatus_transferring_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTransferStatus) SetStatus_code(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_TransferStatus_status_code_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrTransferStatus) GetStatus_code() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_TransferStatus_status_code_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrTransferStatus) SetStatus_error(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_TransferStatus_status_error_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), *(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrTransferStatus) GetStatus_error() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_TransferStatus_status_error_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrTransferStatus) SetBuf(arg2 *byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_TransferStatus_buf_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (arg1 SwigcptrTransferStatus) GetBuf() (_swig_ret *byte) {
	var swig_r *byte
	_swig_i_0 := arg1
	swig_r = (*byte)(C._wrap_TransferStatus_buf_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewTransferStatus() (_swig_ret TransferStatus) {
	var swig_r TransferStatus
	swig_r = (TransferStatus)(SwigcptrTransferStatus(C._wrap_new_TransferStatus_usb_4f19f1d7d83a7073()))
	return swig_r
}

func DeleteTransferStatus(arg1 TransferStatus) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_TransferStatus_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

type TransferStatus interface {
	Swigcptr() uintptr
	SwigIsTransferStatus()
	SetTransferring(arg2 bool)
	GetTransferring() (_swig_ret bool)
	SetStatus_code(arg2 int)
	GetStatus_code() (_swig_ret int)
	SetStatus_error(arg2 string)
	GetStatus_error() (_swig_ret string)
	SetBuf(arg2 *byte)
	GetBuf() (_swig_ret *byte)
}

type _swig_DirectorHIDPacketHandler struct {
	SwigcptrHIDPacketHandler
	v interface{}
}

func (p *_swig_DirectorHIDPacketHandler) Swigcptr() uintptr {
	return p.SwigcptrHIDPacketHandler.Swigcptr()
}

func (p *_swig_DirectorHIDPacketHandler) SwigIsHIDPacketHandler() {
}

func (p *_swig_DirectorHIDPacketHandler) DirectorInterface() interface{} {
	return p.v
}

func NewDirectorHIDPacketHandler(v interface{}) HIDPacketHandler {
	p := &_swig_DirectorHIDPacketHandler{0, v}
	p.SwigcptrHIDPacketHandler = SwigcptrHIDPacketHandler(C._wrap__swig_NewDirectorHIDPacketHandlerHIDPacketHandler_usb_4f19f1d7d83a7073(C.int(swigDirectorAdd(p))))
	return p
}

func DeleteDirectorHIDPacketHandler(arg1 HIDPacketHandler) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_DeleteDirectorHIDPacketHandler_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

//export Swiggo_DeleteDirector_HIDPacketHandler_usb_4f19f1d7d83a7073
func Swiggo_DeleteDirector_HIDPacketHandler_usb_4f19f1d7d83a7073(c int) {
	swigDirectorLookup(c).(*_swig_DirectorHIDPacketHandler).SwigcptrHIDPacketHandler = 0
	swigDirectorDelete(c)
}

type _swig_DirectorInterfaceHIDPacketHandlerHandleIncomingPacket interface {
	HandleIncomingPacket(*int8)
}

func (swig_p *_swig_DirectorHIDPacketHandler) HandleIncomingPacket(packet *int8) {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceHIDPacketHandlerHandleIncomingPacket); swig_ok {
		swig_g.HandleIncomingPacket(packet)
		return
	}
	panic("call to pure virtual method")
}

//export Swig_DirectorHIDPacketHandler_callback_handleIncomingPacket_usb_4f19f1d7d83a7073
func Swig_DirectorHIDPacketHandler_callback_handleIncomingPacket_usb_4f19f1d7d83a7073(swig_c int, packet *int8) {
	swig_p := swigDirectorLookup(swig_c).(*_swig_DirectorHIDPacketHandler)
	swig_p.HandleIncomingPacket(packet)
}

type SwigcptrHIDPacketHandler uintptr

func (p SwigcptrHIDPacketHandler) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrHIDPacketHandler) SwigIsHIDPacketHandler() {
}

func (p SwigcptrHIDPacketHandler) DirectorInterface() interface{} {
	return nil
}

func DeleteHIDPacketHandler(arg1 HIDPacketHandler) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_HIDPacketHandler_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrHIDPacketHandler) HandleIncomingPacket(arg2 *int8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_HIDPacketHandler_handleIncomingPacket_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
}

type HIDPacketHandler interface {
	Swigcptr() uintptr
	SwigIsHIDPacketHandler()
	DirectorInterface() interface{}
	HandleIncomingPacket(arg2 *int8)
}

type SwigcptrDevice uintptr

func (p SwigcptrDevice) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrDevice) SwigIsDevice() {
}

type DeviceFlash_protocol int
func _swig_getDevice_PROTOCOL_UNKNOWN_Device() (_swig_ret DeviceFlash_protocol) {
	var swig_r DeviceFlash_protocol
	swig_r = (DeviceFlash_protocol)(C._wrap_PROTOCOL_UNKNOWN_Device_usb_4f19f1d7d83a7073())
	return swig_r
}

var DevicePROTOCOL_UNKNOWN DeviceFlash_protocol = _swig_getDevice_PROTOCOL_UNKNOWN_Device()
func _swig_getDevice_HALFKAY_Device() (_swig_ret DeviceFlash_protocol) {
	var swig_r DeviceFlash_protocol
	swig_r = (DeviceFlash_protocol)(C._wrap_HALFKAY_Device_usb_4f19f1d7d83a7073())
	return swig_r
}

var DeviceHALFKAY DeviceFlash_protocol = _swig_getDevice_HALFKAY_Device()
func _swig_getDevice_DFU_Device() (_swig_ret DeviceFlash_protocol) {
	var swig_r DeviceFlash_protocol
	swig_r = (DeviceFlash_protocol)(C._wrap_DFU_Device_usb_4f19f1d7d83a7073())
	return swig_r
}

var DeviceDFU DeviceFlash_protocol = _swig_getDevice_DFU_Device()
type DeviceFirmware_format int
func _swig_getDevice_FORMAT_UNKNOWN_Device() (_swig_ret DeviceFirmware_format) {
	var swig_r DeviceFirmware_format
	swig_r = (DeviceFirmware_format)(C._wrap_FORMAT_UNKNOWN_Device_usb_4f19f1d7d83a7073())
	return swig_r
}

var DeviceFORMAT_UNKNOWN DeviceFirmware_format = _swig_getDevice_FORMAT_UNKNOWN_Device()
func _swig_getDevice_HEX_Device() (_swig_ret DeviceFirmware_format) {
	var swig_r DeviceFirmware_format
	swig_r = (DeviceFirmware_format)(C._wrap_HEX_Device_usb_4f19f1d7d83a7073())
	return swig_r
}

var DeviceHEX DeviceFirmware_format = _swig_getDevice_HEX_Device()
func _swig_getDevice_BIN_Device() (_swig_ret DeviceFirmware_format) {
	var swig_r DeviceFirmware_format
	swig_r = (DeviceFirmware_format)(C._wrap_BIN_Device_usb_4f19f1d7d83a7073())
	return swig_r
}

var DeviceBIN DeviceFirmware_format = _swig_getDevice_BIN_Device()
func (arg1 SwigcptrDevice) SetFile_format(arg2 DeviceFirmware_format) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_file_format_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetFile_format() (_swig_ret DeviceFirmware_format) {
	var swig_r DeviceFirmware_format
	_swig_i_0 := arg1
	swig_r = (DeviceFirmware_format)(C._wrap_Device_file_format_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetProtocol(arg2 DeviceFlash_protocol) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_protocol_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetProtocol() (_swig_ret DeviceFlash_protocol) {
	var swig_r DeviceFlash_protocol
	_swig_i_0 := arg1
	swig_r = (DeviceFlash_protocol)(C._wrap_Device_protocol_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetPacket_handler(arg2 HIDPacketHandler) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_Device_packet_handler_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetPacket_handler() (_swig_ret HIDPacketHandler) {
	var swig_r HIDPacketHandler
	_swig_i_0 := arg1
	swig_r = (HIDPacketHandler)(SwigcptrHIDPacketHandler(C._wrap_Device_packet_handler_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrDevice) SetBootloader(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_bootloader_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetBootloader() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Device_bootloader_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetPid(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_pid_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetPid() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Device_pid_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetPort_number(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_port_number_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetPort_number() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Device_port_number_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetVid(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_vid_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetVid() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Device_vid_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetFingerprint(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_fingerprint_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrDevice) GetFingerprint() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Device_fingerprint_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) SetFriendly_name(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_friendly_name_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrDevice) GetFriendly_name() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Device_friendly_name_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrDevice) SetModel(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Device_model_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrDevice) GetModel() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Device_model_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrDevice) Usb_transfer(arg2 byte, arg3 byte, arg4 uint16, arg5 uint16, arg6 *byte, arg7 uint16, arg8 int) (_swig_ret TransferStatus) {
	var swig_r TransferStatus
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	_swig_i_6 := arg7
	_swig_i_7 := arg8
	swig_r = (TransferStatus)(SwigcptrTransferStatus(C._wrap_Device_usb_transfer_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.char(_swig_i_1), C.char(_swig_i_2), C.short(_swig_i_3), C.short(_swig_i_4), C.swig_voidp(_swig_i_5), C.short(_swig_i_6), C.swig_intgo(_swig_i_7))))
	return swig_r
}

func (arg1 SwigcptrDevice) Hid_listen() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Device_hid_listen_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) Hid_open(arg2 int) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_Device_hid_open_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrDevice) Usb_claim() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Device_usb_claim_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) Check_connected() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Device_check_connected_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) Send_hid_packet(arg2 *byte, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_Device_send_hid_packet_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrDevice) Usb_auto_detach() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Device_usb_auto_detach_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrDevice) Usb_claim_interface(arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_Device_usb_claim_interface_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrDevice) Usb_set_configuration(arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_Device_usb_set_configuration_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func DeviceGet_firmware_format(arg1 DeviceFlash_protocol) (_swig_ret DeviceFirmware_format) {
	var swig_r DeviceFirmware_format
	_swig_i_0 := arg1
	swig_r = (DeviceFirmware_format)(C._wrap_Device_get_firmware_format_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func DeviceGet_flashing_protocol(arg1 int) (_swig_ret DeviceFlash_protocol) {
	var swig_r DeviceFlash_protocol
	_swig_i_0 := arg1
	swig_r = (DeviceFlash_protocol)(C._wrap_Device_get_flashing_protocol_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func DeviceIs_bootloader(arg1 int) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Device_is_bootloader_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func DeviceIs_interesting(arg1 int, arg2 int) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_Device_is_interesting_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func DeviceGet_friendly_name(arg1 int) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Device_get_friendly_name_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func DeviceGet_model(arg1 int) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Device_get_model_usb_4f19f1d7d83a7073(C.swig_intgo(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrDevice) Get_dfu_string(arg2 int) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_Device_get_dfu_string_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrDevice) Close_hid() {
	_swig_i_0 := arg1
	C._wrap_Device_close_hid_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrDevice) Usb_close() {
	_swig_i_0 := arg1
	C._wrap_Device_usb_close_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func NewDevice(arg1 Libusb_device, arg2 int, arg3 int) (_swig_ret Device) {
	var swig_r Device
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Device)(SwigcptrDevice(C._wrap_new_Device_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2))))
	return swig_r
}

func DeleteDevice(arg1 Device) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Device_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

type Device interface {
	Swigcptr() uintptr
	SwigIsDevice()
	SetFile_format(arg2 DeviceFirmware_format)
	GetFile_format() (_swig_ret DeviceFirmware_format)
	SetProtocol(arg2 DeviceFlash_protocol)
	GetProtocol() (_swig_ret DeviceFlash_protocol)
	SetPacket_handler(arg2 HIDPacketHandler)
	GetPacket_handler() (_swig_ret HIDPacketHandler)
	SetBootloader(arg2 bool)
	GetBootloader() (_swig_ret bool)
	SetPid(arg2 int)
	GetPid() (_swig_ret int)
	SetPort_number(arg2 int)
	GetPort_number() (_swig_ret int)
	SetVid(arg2 int)
	GetVid() (_swig_ret int)
	SetFingerprint(arg2 int)
	GetFingerprint() (_swig_ret int)
	SetFriendly_name(arg2 string)
	GetFriendly_name() (_swig_ret string)
	SetModel(arg2 string)
	GetModel() (_swig_ret string)
	Usb_transfer(arg2 byte, arg3 byte, arg4 uint16, arg5 uint16, arg6 *byte, arg7 uint16, arg8 int) (_swig_ret TransferStatus)
	Hid_listen() (_swig_ret bool)
	Hid_open(arg2 int) (_swig_ret bool)
	Usb_claim() (_swig_ret bool)
	Check_connected() (_swig_ret int)
	Send_hid_packet(arg2 *byte, arg3 int) (_swig_ret int)
	Usb_auto_detach() (_swig_ret int)
	Usb_claim_interface(arg2 int) (_swig_ret int)
	Usb_set_configuration(arg2 int) (_swig_ret int)
	Get_dfu_string(arg2 int) (_swig_ret string)
	Close_hid()
	Usb_close()
}

type _swig_DirectorEventHandler struct {
	SwigcptrEventHandler
	v interface{}
}

func (p *_swig_DirectorEventHandler) Swigcptr() uintptr {
	return p.SwigcptrEventHandler.Swigcptr()
}

func (p *_swig_DirectorEventHandler) SwigIsEventHandler() {
}

func (p *_swig_DirectorEventHandler) DirectorInterface() interface{} {
	return p.v
}

func NewDirectorEventHandler(v interface{}) EventHandler {
	p := &_swig_DirectorEventHandler{0, v}
	p.SwigcptrEventHandler = SwigcptrEventHandler(C._wrap__swig_NewDirectorEventHandlerEventHandler_usb_4f19f1d7d83a7073(C.int(swigDirectorAdd(p))))
	return p
}

func DeleteDirectorEventHandler(arg1 EventHandler) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_DeleteDirectorEventHandler_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

//export Swiggo_DeleteDirector_EventHandler_usb_4f19f1d7d83a7073
func Swiggo_DeleteDirector_EventHandler_usb_4f19f1d7d83a7073(c int) {
	swigDirectorLookup(c).(*_swig_DirectorEventHandler).SwigcptrEventHandler = 0
	swigDirectorDelete(c)
}

type _swig_DirectorInterfaceEventHandlerHandleUSBConnectionEvent interface {
	HandleUSBConnectionEvent(bool, Device)
}

func (swig_p *_swig_DirectorEventHandler) HandleUSBConnectionEvent(connected bool, dev Device) {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceEventHandlerHandleUSBConnectionEvent); swig_ok {
		swig_g.HandleUSBConnectionEvent(connected, dev)
		return
	}
	panic("call to pure virtual method")
}

//export Swig_DirectorEventHandler_callback_handleUSBConnectionEvent_usb_4f19f1d7d83a7073
func Swig_DirectorEventHandler_callback_handleUSBConnectionEvent_usb_4f19f1d7d83a7073(swig_c int, connected bool, dev uintptr) {
	swig_p := swigDirectorLookup(swig_c).(*_swig_DirectorEventHandler)
	swig_p.HandleUSBConnectionEvent(connected, SwigcptrDevice(dev))
}

type SwigcptrEventHandler uintptr

func (p SwigcptrEventHandler) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrEventHandler) SwigIsEventHandler() {
}

func (p SwigcptrEventHandler) DirectorInterface() interface{} {
	return nil
}

func DeleteEventHandler(arg1 EventHandler) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_EventHandler_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrEventHandler) HandleUSBConnectionEvent(arg2 bool, arg3 Device) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_EventHandler_handleUSBConnectionEvent_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1), C.uintptr_t(_swig_i_2))
}

type EventHandler interface {
	Swigcptr() uintptr
	SwigIsEventHandler()
	DirectorInterface() interface{}
	HandleUSBConnectionEvent(arg2 bool, arg3 Device)
}

type SwigcptrEnumerator uintptr

func (p SwigcptrEnumerator) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrEnumerator) SwigIsEnumerator() {
}

func (arg1 SwigcptrEnumerator) SetEventObject(arg2 EventHandler) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_Enumerator_EventObject_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrEnumerator) GetEventObject() (_swig_ret EventHandler) {
	var swig_r EventHandler
	_swig_i_0 := arg1
	swig_r = (EventHandler)(SwigcptrEventHandler(C._wrap_Enumerator_EventObject_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewEnumerator() (_swig_ret Enumerator) {
	var swig_r Enumerator
	swig_r = (Enumerator)(SwigcptrEnumerator(C._wrap_new_Enumerator_usb_4f19f1d7d83a7073()))
	return swig_r
}

func DeleteEnumerator(arg1 Enumerator) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Enumerator_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrEnumerator) ListenDevices() {
	_swig_i_0 := arg1
	C._wrap_Enumerator_ListenDevices_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrEnumerator) StopListenDevices() {
	_swig_i_0 := arg1
	C._wrap_Enumerator_StopListenDevices_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrEnumerator) HandleEvents() {
	_swig_i_0 := arg1
	C._wrap_Enumerator_HandleEvents_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrEnumerator) SetDevices(arg2 Std_vector_Sl_Device_Sg_) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_Enumerator_Devices_set_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrEnumerator) GetDevices() (_swig_ret Std_vector_Sl_Device_Sg_) {
	var swig_r Std_vector_Sl_Device_Sg_
	_swig_i_0 := arg1
	swig_r = (Std_vector_Sl_Device_Sg_)(SwigcptrStd_vector_Sl_Device_Sg_(C._wrap_Enumerator_Devices_get_usb_4f19f1d7d83a7073(C.uintptr_t(_swig_i_0))))
	return swig_r
}

type Enumerator interface {
	Swigcptr() uintptr
	SwigIsEnumerator()
	SetEventObject(arg2 EventHandler)
	GetEventObject() (_swig_ret EventHandler)
	ListenDevices()
	StopListenDevices()
	HandleEvents()
	SetDevices(arg2 Std_vector_Sl_Device_Sg_)
	GetDevices() (_swig_ret Std_vector_Sl_Device_Sg_)
}


type SwigcptrStd_vector_Sl_Device_Sg_ uintptr
type Std_vector_Sl_Device_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrStd_vector_Sl_Device_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrLibusb_device uintptr
type Libusb_device interface {
	Swigcptr() uintptr;
}
func (p SwigcptrLibusb_device) Swigcptr() uintptr {
	return uintptr(p)
}



var swigDirectorTrack struct {
	sync.Mutex
	m map[int]interface{}
	c int
}

func swigDirectorAdd(v interface{}) int {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m == nil {
		swigDirectorTrack.m = make(map[int]interface{})
	}
	swigDirectorTrack.c++
	ret := swigDirectorTrack.c
	swigDirectorTrack.m[ret] = v
	return ret
}

func swigDirectorLookup(c int) interface{} {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	ret := swigDirectorTrack.m[c]
	if ret == nil {
		panic("C++ director pointer not found (possible	use-after-free)")
	}
	return ret
}

func swigDirectorDelete(c int) {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m[c] == nil {
		if c > swigDirectorTrack.c {
			panic("C++ director pointer invalid (possible memory corruption")
		} else {
			panic("C++ director pointer not found (possible use-after-free)")
		}
	}
	delete(swigDirectorTrack.m, c)
}


