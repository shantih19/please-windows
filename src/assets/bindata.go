// Code generated by go-bindata. DO NOT EDIT.
// sources:
// src/assets/pleasew (2.275kB)
// src/assets/plz_complete.sh (837B)

package assets

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len
	return b, nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pleasew = "\x23\x21\x2f\x75\x73\x72\x2f\x62\x69\x6e\x2f\x65\x6e\x76\x20\x62\x61\x73\x68\x0a\x73\x65\x74\x20\x2d\x75\x0a\x0a\x52\x45\x44\x3d\x22\x5c\x78\x31\x42\x5b\x33\x31\x6d\x22\x0a\x47\x52\x45\x45\x4e\x3d\x22\x5c\x78\x31\x42\x5b\x33\x32\x6d\x22\x0a\x59\x45\x4c\x4c\x4f\x57\x3d\x22\x5c\x78\x31\x42\x5b\x33\x33\x6d\x22\x0a\x52\x45\x53\x45\x54\x3d\x22\x5c\x78\x31\x42\x5b\x30\x6d\x22\x0a\x0a\x44\x45\x46\x41\x55\x4c\x54\x5f\x55\x52\x4c\x5f\x42\x41\x53\x45\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x67\x65\x74\x2e\x70\x6c\x65\x61\x73\x65\x2e\x62\x75\x69\x6c\x64\x22\x0a\x23\x20\x57\x65\x20\x6d\x69\x67\x68\x74\x20\x61\x6c\x72\x65\x61\x64\x79\x20\x68\x61\x76\x65\x20\x69\x74\x20\x64\x6f\x77\x6e\x6c\x6f\x61\x64\x65\x64\x2e\x2e\x2e\x0a\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x3d\x60\x67\x72\x65\x70\x20\x2d\x69\x20\x22\x5e\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x22\x20\x2e\x70\x6c\x7a\x63\x6f\x6e\x66\x69\x67\x20\x32\x3e\x2f\x64\x65\x76\x2f\x6e\x75\x6c\x6c\x20\x7c\x20\x63\x75\x74\x20\x2d\x64\x20\x27\x3d\x27\x20\x2d\x66\x20\x32\x20\x7c\x20\x74\x72\x20\x2d\x64\x20\x27\x20\x27\x60\x0a\x69\x66\x20\x5b\x20\x2d\x7a\x20\x22\x24\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x69\x66\x20\x5b\x20\x2d\x7a\x20\x22\x24\x48\x4f\x4d\x45\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x09\x20\x20\x20\x20\x65\x63\x68\x6f\x20\x2d\x65\x20\x3e\x26\x32\x20\x22\x24\x7b\x52\x45\x44\x7d\x5c\x24\x48\x4f\x4d\x45\x20\x6e\x6f\x74\x20\x73\x65\x74\x2c\x20\x6e\x6f\x74\x20\x73\x75\x72\x65\x20\x77\x68\x65\x72\x65\x20\x74\x6f\x20\x6c\x6f\x6f\x6b\x20\x66\x6f\x72\x20\x50\x6c\x65\x61\x73\x65\x2e\x24\x7b\x52\x45\x53\x45\x54\x7d\x22\x0a\x09\x20\x20\x20\x20\x65\x78\x69\x74\x20\x31\x0a\x20\x20\x20\x20\x66\x69\x0a\x20\x20\x20\x20\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x3d\x22\x24\x7b\x48\x4f\x4d\x45\x7d\x2f\x2e\x70\x6c\x65\x61\x73\x65\x22\x0a\x65\x6c\x73\x65\x0a\x20\x20\x20\x20\x23\x20\x49\x74\x20\x63\x61\x6e\x20\x63\x6f\x6e\x74\x61\x69\x6e\x20\x61\x20\x6c\x69\x74\x65\x72\x61\x6c\x20\x7e\x2c\x20\x6e\x65\x65\x64\x20\x74\x6f\x20\x65\x78\x70\x6c\x69\x63\x69\x74\x6c\x79\x20\x68\x61\x6e\x64\x6c\x65\x20\x74\x68\x61\x74\x2e\x0a\x20\x20\x20\x20\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x3d\x22\x24\x7b\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x2f\x5c\x7e\x2f\x24\x48\x4f\x4d\x45\x7d\x22\x0a\x66\x69\x0a\x23\x20\x49\x66\x20\x74\x68\x69\x73\x20\x65\x78\x69\x73\x74\x73\x20\x61\x74\x20\x61\x6e\x79\x20\x76\x65\x72\x73\x69\x6f\x6e\x2c\x20\x6c\x65\x74\x20\x69\x74\x20\x68\x61\x6e\x64\x6c\x65\x20\x61\x6e\x79\x20\x75\x70\x64\x61\x74\x65\x2e\x0a\x54\x41\x52\x47\x45\x54\x3d\x22\x24\x7b\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x7d\x2f\x70\x6c\x65\x61\x73\x65\x22\x0a\x69\x66\x20\x5b\x20\x2d\x66\x20\x22\x24\x54\x41\x52\x47\x45\x54\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x65\x78\x65\x63\x20\x22\x24\x54\x41\x52\x47\x45\x54\x22\x20\x24\x7b\x50\x4c\x5a\x5f\x41\x52\x47\x53\x3a\x2d\x7d\x20\x22\x24\x40\x22\x0a\x66\x69\x0a\x0a\x55\x52\x4c\x5f\x42\x41\x53\x45\x3d\x22\x60\x67\x72\x65\x70\x20\x2d\x69\x20\x22\x5e\x64\x6f\x77\x6e\x6c\x6f\x61\x64\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x22\x20\x2e\x70\x6c\x7a\x63\x6f\x6e\x66\x69\x67\x20\x7c\x20\x63\x75\x74\x20\x2d\x64\x20\x27\x3d\x27\x20\x2d\x66\x20\x32\x20\x7c\x20\x74\x72\x20\x2d\x64\x20\x27\x20\x27\x60\x22\x0a\x69\x66\x20\x5b\x20\x2d\x7a\x20\x22\x24\x55\x52\x4c\x5f\x42\x41\x53\x45\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x55\x52\x4c\x5f\x42\x41\x53\x45\x3d\x24\x44\x45\x46\x41\x55\x4c\x54\x5f\x55\x52\x4c\x5f\x42\x41\x53\x45\x0a\x66\x69\x0a\x55\x52\x4c\x5f\x42\x41\x53\x45\x3d\x22\x24\x7b\x55\x52\x4c\x5f\x42\x41\x53\x45\x25\x2f\x7d\x22\x0a\x0a\x56\x45\x52\x53\x49\x4f\x4e\x3d\x22\x60\x67\x72\x65\x70\x20\x2d\x69\x20\x22\x5e\x76\x65\x72\x73\x69\x6f\x6e\x5b\x5e\x61\x2d\x7a\x5d\x22\x20\x2e\x70\x6c\x7a\x63\x6f\x6e\x66\x69\x67\x60\x22\x0a\x56\x45\x52\x53\x49\x4f\x4e\x3d\x22\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x23\x2a\x3d\x7d\x22\x20\x20\x20\x20\x23\x20\x53\x74\x72\x69\x70\x20\x75\x6e\x74\x69\x6c\x20\x61\x66\x74\x65\x72\x20\x66\x69\x72\x73\x74\x20\x3d\x0a\x56\x45\x52\x53\x49\x4f\x4e\x3d\x22\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x2f\x20\x2f\x7d\x22\x20\x20\x20\x20\x23\x20\x52\x65\x6d\x6f\x76\x65\x20\x61\x6c\x6c\x20\x73\x70\x61\x63\x65\x73\x0a\x56\x45\x52\x53\x49\x4f\x4e\x3d\x22\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x23\x3e\x3d\x7d\x22\x20\x20\x20\x20\x23\x20\x53\x74\x72\x69\x70\x20\x61\x6e\x79\x20\x69\x6e\x69\x74\x69\x61\x6c\x20\x3e\x3d\x0a\x69\x66\x20\x5b\x20\x2d\x7a\x20\x22\x24\x56\x45\x52\x53\x49\x4f\x4e\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x65\x63\x68\x6f\x20\x2d\x65\x20\x3e\x26\x32\x20\x22\x24\x7b\x59\x45\x4c\x4c\x4f\x57\x7d\x43\x61\x6e\x27\x74\x20\x64\x65\x74\x65\x72\x6d\x69\x6e\x65\x20\x76\x65\x72\x73\x69\x6f\x6e\x2c\x20\x77\x69\x6c\x6c\x20\x75\x73\x65\x20\x6c\x61\x74\x65\x73\x74\x2e\x24\x7b\x52\x45\x53\x45\x54\x7d\x22\x0a\x20\x20\x20\x20\x56\x45\x52\x53\x49\x4f\x4e\x3d\x60\x63\x75\x72\x6c\x20\x2d\x66\x73\x53\x4c\x20\x24\x7b\x55\x52\x4c\x5f\x42\x41\x53\x45\x7d\x2f\x6c\x61\x74\x65\x73\x74\x5f\x76\x65\x72\x73\x69\x6f\x6e\x60\x0a\x66\x69\x0a\x0a\x23\x20\x46\x69\x6e\x64\x20\x74\x68\x65\x20\x6f\x73\x20\x2f\x20\x61\x72\x63\x68\x20\x74\x6f\x20\x64\x6f\x77\x6e\x6c\x6f\x61\x64\x2e\x20\x59\x6f\x75\x20\x63\x61\x6e\x20\x64\x6f\x20\x74\x68\x69\x73\x20\x71\x75\x69\x74\x65\x20\x6e\x69\x63\x65\x6c\x79\x20\x77\x69\x74\x68\x20\x67\x6f\x20\x65\x6e\x76\x0a\x23\x20\x62\x75\x74\x20\x77\x65\x20\x75\x73\x65\x20\x74\x68\x69\x73\x20\x73\x63\x72\x69\x70\x74\x20\x6f\x6e\x20\x6d\x61\x63\x68\x69\x6e\x65\x73\x20\x74\x68\x61\x74\x20\x64\x6f\x6e\x27\x74\x20\x6e\x65\x63\x65\x73\x73\x61\x72\x69\x6c\x79\x20\x68\x61\x76\x65\x20\x47\x6f\x20\x69\x74\x73\x65\x6c\x66\x2e\x0a\x4f\x53\x3d\x60\x75\x6e\x61\x6d\x65\x60\x0a\x69\x66\x20\x5b\x20\x22\x24\x4f\x53\x22\x20\x3d\x20\x22\x4c\x69\x6e\x75\x78\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x47\x4f\x4f\x53\x3d\x22\x6c\x69\x6e\x75\x78\x22\x0a\x65\x6c\x69\x66\x20\x5b\x20\x22\x24\x4f\x53\x22\x20\x3d\x20\x22\x44\x61\x72\x77\x69\x6e\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x47\x4f\x4f\x53\x3d\x22\x64\x61\x72\x77\x69\x6e\x22\x0a\x65\x6c\x73\x65\x0a\x20\x20\x20\x20\x65\x63\x68\x6f\x20\x2d\x65\x20\x3e\x26\x32\x20\x22\x24\x7b\x52\x45\x44\x7d\x55\x6e\x6b\x6e\x6f\x77\x6e\x20\x6f\x70\x65\x72\x61\x74\x69\x6e\x67\x20\x73\x79\x73\x74\x65\x6d\x20\x24\x4f\x53\x24\x7b\x52\x45\x53\x45\x54\x7d\x22\x0a\x20\x20\x20\x20\x65\x78\x69\x74\x20\x31\x0a\x66\x69\x0a\x23\x20\x44\x6f\x6e\x27\x74\x20\x68\x61\x76\x65\x20\x61\x6e\x79\x20\x62\x75\x69\x6c\x64\x73\x20\x6f\x74\x68\x65\x72\x20\x74\x68\x61\x6e\x20\x61\x6d\x64\x36\x34\x20\x61\x74\x20\x74\x68\x65\x20\x6d\x6f\x6d\x65\x6e\x74\x2e\x0a\x41\x52\x43\x48\x3d\x22\x61\x6d\x64\x36\x34\x22\x0a\x0a\x50\x4c\x45\x41\x53\x45\x5f\x55\x52\x4c\x3d\x22\x24\x7b\x55\x52\x4c\x5f\x42\x41\x53\x45\x7d\x2f\x24\x7b\x47\x4f\x4f\x53\x7d\x5f\x24\x7b\x41\x52\x43\x48\x7d\x2f\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x7d\x2f\x70\x6c\x65\x61\x73\x65\x5f\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x7d\x2e\x74\x61\x72\x2e\x78\x7a\x22\x0a\x44\x49\x52\x3d\x22\x24\x7b\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x7d\x2f\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x7d\x22\x0a\x23\x20\x50\x6f\x74\x65\x6e\x74\x69\x61\x6c\x6c\x79\x20\x77\x65\x20\x63\x6f\x75\x6c\x64\x20\x72\x65\x75\x73\x65\x20\x74\x68\x69\x73\x20\x62\x75\x74\x20\x69\x74\x27\x73\x20\x65\x61\x73\x69\x65\x72\x20\x6e\x6f\x74\x20\x74\x6f\x20\x72\x65\x61\x6c\x6c\x79\x2e\x0a\x69\x66\x20\x5b\x20\x21\x20\x2d\x64\x20\x22\x24\x44\x49\x52\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x72\x6d\x20\x2d\x72\x66\x20\x22\x24\x44\x49\x52\x22\x0a\x66\x69\x0a\x65\x63\x68\x6f\x20\x2d\x65\x20\x3e\x26\x32\x20\x22\x24\x7b\x47\x52\x45\x45\x4e\x7d\x44\x6f\x77\x6e\x6c\x6f\x61\x64\x69\x6e\x67\x20\x50\x6c\x65\x61\x73\x65\x20\x24\x7b\x56\x45\x52\x53\x49\x4f\x4e\x7d\x20\x74\x6f\x20\x24\x7b\x44\x49\x52\x7d\x2e\x2e\x2e\x24\x7b\x52\x45\x53\x45\x54\x7d\x22\x0a\x6d\x6b\x64\x69\x72\x20\x2d\x70\x20\x22\x24\x44\x49\x52\x22\x0a\x63\x75\x72\x6c\x20\x2d\x66\x73\x53\x4c\x20\x22\x24\x7b\x50\x4c\x45\x41\x53\x45\x5f\x55\x52\x4c\x7d\x22\x20\x7c\x20\x74\x61\x72\x20\x2d\x78\x4a\x70\x66\x2d\x20\x2d\x2d\x73\x74\x72\x69\x70\x2d\x63\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x73\x3d\x31\x20\x2d\x43\x20\x22\x24\x44\x49\x52\x22\x0a\x23\x20\x4c\x69\x6e\x6b\x20\x69\x74\x20\x61\x6c\x6c\x20\x62\x61\x63\x6b\x20\x75\x70\x20\x61\x20\x64\x69\x72\x0a\x66\x6f\x72\x20\x78\x20\x69\x6e\x20\x60\x6c\x73\x20\x22\x24\x44\x49\x52\x22\x60\x3b\x20\x64\x6f\x0a\x20\x20\x20\x20\x6c\x6e\x20\x2d\x73\x66\x20\x22\x24\x7b\x44\x49\x52\x7d\x2f\x24\x7b\x78\x7d\x22\x20\x22\x24\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x22\x0a\x64\x6f\x6e\x65\x0a\x6c\x6e\x20\x2d\x73\x66\x20\x22\x24\x7b\x44\x49\x52\x7d\x2f\x70\x6c\x65\x61\x73\x65\x22\x20\x22\x24\x7b\x4c\x4f\x43\x41\x54\x49\x4f\x4e\x7d\x2f\x70\x6c\x7a\x22\x0a\x65\x63\x68\x6f\x20\x2d\x65\x20\x3e\x26\x32\x20\x22\x24\x7b\x47\x52\x45\x45\x4e\x7d\x53\x68\x6f\x75\x6c\x64\x20\x62\x65\x20\x67\x6f\x6f\x64\x20\x74\x6f\x20\x67\x6f\x20\x6e\x6f\x77\x2c\x20\x72\x75\x6e\x6e\x69\x6e\x67\x20\x70\x6c\x7a\x2e\x2e\x2e\x24\x7b\x52\x45\x53\x45\x54\x7d\x22\x0a\x65\x78\x65\x63\x20\x22\x24\x54\x41\x52\x47\x45\x54\x22\x20\x24\x7b\x50\x4c\x5a\x5f\x41\x52\x47\x53\x3a\x2d\x7d\x20\x22\x24\x40\x22\x0a"

func pleasewBytes() ([]byte, error) {
	return bindataRead(
		_pleasew,
		"pleasew",
	)
}

func pleasew() (*asset, error) {
	bytes, err := pleasewBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pleasew", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf7, 0x29, 0x83, 0xe, 0x78, 0xc5, 0x6d, 0x7c, 0x2b, 0xe3, 0xf7, 0xc3, 0x52, 0x97, 0xd5, 0xc9, 0x6f, 0xaa, 0xc0, 0xdc, 0x7d, 0xa2, 0x3, 0x62, 0x76, 0xac, 0xb1, 0xa7, 0x46, 0xae, 0x16}}
	return a, nil
}

var _plz_completeSh = "\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x0a\x23\x20\x70\x6c\x7a\x20\x63\x6f\x6d\x70\x6c\x65\x74\x69\x6f\x6e\x0a\x23\x0a\x23\x20\x61\x64\x64\x0a\x23\x20\x73\x6f\x75\x72\x63\x65\x20\x3c\x28\x70\x6c\x7a\x20\x2d\x2d\x63\x6f\x6d\x70\x6c\x65\x74\x69\x6f\x6e\x5f\x73\x63\x72\x69\x70\x74\x29\x0a\x23\x20\x74\x6f\x20\x79\x6f\x75\x72\x20\x2e\x62\x61\x73\x68\x72\x63\x20\x2f\x2e\x7a\x73\x68\x72\x63\x20\x74\x6f\x20\x61\x63\x74\x69\x76\x61\x74\x65\x20\x74\x68\x69\x73\x2e\x0a\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x23\x0a\x0a\x5f\x70\x6c\x7a\x5f\x63\x6f\x6d\x70\x6c\x65\x74\x65\x5f\x62\x61\x73\x68\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x43\x4f\x4d\x50\x5f\x57\x4f\x52\x44\x42\x52\x45\x41\x4b\x53\x3d\x24\x7b\x43\x4f\x4d\x50\x5f\x57\x4f\x52\x44\x42\x52\x45\x41\x4b\x53\x2f\x2f\x3a\x7d\x0a\x20\x20\x20\x20\x61\x72\x67\x73\x3d\x28\x22\x24\x7b\x43\x4f\x4d\x50\x5f\x57\x4f\x52\x44\x53\x5b\x40\x5d\x3a\x31\x3a\x24\x43\x4f\x4d\x50\x5f\x43\x57\x4f\x52\x44\x7d\x22\x29\x0a\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x20\x49\x46\x53\x3d\x24\x27\x5c\x6e\x27\x0a\x20\x20\x20\x20\x43\x4f\x4d\x50\x52\x45\x50\x4c\x59\x3d\x28\x24\x28\x47\x4f\x5f\x46\x4c\x41\x47\x53\x5f\x43\x4f\x4d\x50\x4c\x45\x54\x49\x4f\x4e\x3d\x31\x20\x24\x7b\x43\x4f\x4d\x50\x5f\x57\x4f\x52\x44\x53\x5b\x30\x5d\x7d\x20\x2d\x70\x20\x2d\x76\x20\x30\x20\x2d\x2d\x6e\x6f\x75\x70\x64\x61\x74\x65\x20\x22\x24\x7b\x61\x72\x67\x73\x5b\x40\x5d\x7d\x22\x29\x29\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x30\x0a\x7d\x0a\x0a\x5f\x70\x6c\x7a\x5f\x63\x6f\x6d\x70\x6c\x65\x74\x65\x5f\x7a\x73\x68\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x20\x61\x72\x67\x73\x3d\x28\x22\x24\x7b\x77\x6f\x72\x64\x73\x5b\x40\x5d\x3a\x31\x3a\x24\x43\x55\x52\x52\x45\x4e\x54\x7d\x22\x29\x0a\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x20\x49\x46\x53\x3d\x24\x27\x5c\x6e\x27\x0a\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x20\x63\x6f\x6d\x70\x6c\x65\x74\x69\x6f\x6e\x73\x3d\x28\x24\x28\x47\x4f\x5f\x46\x4c\x41\x47\x53\x5f\x43\x4f\x4d\x50\x4c\x45\x54\x49\x4f\x4e\x3d\x31\x20\x24\x7b\x77\x6f\x72\x64\x73\x5b\x31\x5d\x7d\x20\x2d\x70\x20\x2d\x76\x20\x30\x20\x2d\x2d\x6e\x6f\x75\x70\x64\x61\x74\x65\x20\x22\x24\x7b\x61\x72\x67\x73\x5b\x40\x5d\x7d\x22\x29\x29\x0a\x20\x20\x20\x20\x66\x6f\x72\x20\x63\x6f\x6d\x70\x6c\x65\x74\x69\x6f\x6e\x20\x69\x6e\x20\x24\x63\x6f\x6d\x70\x6c\x65\x74\x69\x6f\x6e\x73\x3b\x20\x64\x6f\x0a\x09\x63\x6f\x6d\x70\x61\x64\x64\x20\x24\x63\x6f\x6d\x70\x6c\x65\x74\x69\x6f\x6e\x0a\x20\x20\x20\x20\x64\x6f\x6e\x65\x0a\x7d\x0a\x0a\x69\x66\x20\x5b\x20\x2d\x6e\x20\x22\x24\x42\x41\x53\x48\x5f\x56\x45\x52\x53\x49\x4f\x4e\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x63\x6f\x6d\x70\x6c\x65\x74\x65\x20\x2d\x46\x20\x5f\x70\x6c\x7a\x5f\x63\x6f\x6d\x70\x6c\x65\x74\x65\x5f\x62\x61\x73\x68\x20\x70\x6c\x7a\x0a\x65\x6c\x69\x66\x20\x5b\x20\x2d\x6e\x20\x22\x24\x5a\x53\x48\x5f\x56\x45\x52\x53\x49\x4f\x4e\x22\x20\x5d\x3b\x20\x74\x68\x65\x6e\x0a\x20\x20\x20\x20\x63\x6f\x6d\x70\x64\x65\x66\x20\x5f\x70\x6c\x7a\x5f\x63\x6f\x6d\x70\x6c\x65\x74\x65\x5f\x7a\x73\x68\x20\x70\x6c\x7a\x0a\x66\x69\x0a"

func plz_completeShBytes() ([]byte, error) {
	return bindataRead(
		_plz_completeSh,
		"plz_complete.sh",
	)
}

func plz_completeSh() (*asset, error) {
	bytes, err := plz_completeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plz_complete.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2d, 0xf, 0x58, 0x9c, 0xf, 0x6e, 0x52, 0xcd, 0xf8, 0x7c, 0xbf, 0xd1, 0x50, 0x69, 0xd5, 0x91, 0x69, 0xa2, 0x51, 0x4c, 0x9c, 0x57, 0xbc, 0x77, 0xc, 0x59, 0x98, 0xd0, 0xbc, 0xa0, 0x4c, 0x77}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"pleasew":         pleasew,
	"plz_complete.sh": plz_completeSh,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"pleasew":         {pleasew, map[string]*bintree{}},
	"plz_complete.sh": {plz_completeSh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
