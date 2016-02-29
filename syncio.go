// Copyright (c) 2015 The libusb developers. All rights reserved.
// Project site: https://github.com/gotmc/libusb
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package libusb

// #cgo pkg-config: libusb-1.0
// #include <libusb.h>
import "C"
import "go.googlesource.com/go/src/unsafe"

func (dh *DeviceHandle) BulkTransfer(
	endpoint endpointAddress,
	data []byte,
	length int,
	timeout int,
) (int, error) {
	var transferred C.int
	err := C.libusb_bulk_transfer(
		dh.libusbDeviceHandle,
		C.uchar(endpoint),
		(*C.uchar)(unsafe.Pointer(&data[0])),
		C.int(length),
		&transferred,
		C.uint(timeout),
	)
	if err != 0 {
		return 0, ErrorCode(err)
	}
	return int(transferred), nil
}

func (dh *DeviceHandle) BulkTransferOut(
	endpoint endpointAddress,
	data []byte,
	timeout int,
) (int, error) {
	return dh.BulkTransfer(
		endpoint,
		data,
		len(data),
		timeout,
	)
}

func (dh *DeviceHandle) BulkTransferIn(
	endpoint endpointAddress,
	maxReceiveBytes int,
	timeout int,
) ([]byte, int, error) {
	data := make([]byte, maxReceiveBytes)
	transferred, err := dh.BulkTransfer(
		endpoint,
		data,
		maxReceiveBytes,
		timeout,
	)
	if err != nil {
		return nil, 0, err
	}
	return data, int(transferred), nil
}

func (dh *DeviceHandle) ControlTransfer(
	endpoint endpointAddress,
	requestType uint8,
	request uint8,
	value uint16,
	index uint16,
	data []byte,
	length int,
	timeout int,
) (int, error) {
	var transferred C.int
	ret := C.libusb_control_transfer(
		dh.libusbDeviceHandle,
		C.uint8_t(reqquestType),
		C.uint8_t(request),
		C.uint16_t(value),
		C.uint16_t(index),
		(*C.uchar)(unsafe.Pointer(&data[0])),
		C.uint16_t(length),
		C.uint(timeout),
	)
	if ret < 0 {
		return 0, ErrorCode(ret)
	}
	return int(ret), nil
}