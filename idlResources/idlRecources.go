// Code generated by go-bindata. DO NOT EDIT.
// sources:
// idlResources/0001.idl (20B)
// idlResources/0002.idl (0)
// idlResources/defines.idl (96B)
// idlResources/fitMappings.idl (431B)
// idlResources/fitMessage.idl (26B)
// idlResources/struct.idl (1.588kB)

package idlResources

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

var _idlresources0001Idl = []byte(`#include "0002.idl"
`)

func idlresources0001IdlBytes() ([]byte, error) {
	return _idlresources0001Idl, nil
}

func idlresources0001Idl() (*asset, error) {
	bytes, err := idlresources0001IdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idlResources/0001.idl", size: 20, mode: os.FileMode(0644), modTime: time.Unix(1741119739, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x17, 0xe1, 0x2d, 0xc9, 0xb0, 0xe2, 0xc3, 0x87, 0x1e, 0xe4, 0x1c, 0x60, 0x9b, 0x4c, 0x72, 0xe7, 0xed, 0x57, 0xdc, 0x2a, 0xfa, 0x8, 0x19, 0x28, 0xe2, 0x6b, 0x43, 0x4, 0x54, 0xfe, 0xde, 0x4f}}
	return a, nil
}

var _idlresources0002Idl = []byte("")

func idlresources0002IdlBytes() ([]byte, error) {
	return _idlresources0002Idl, nil
}

func idlresources0002Idl() (*asset, error) {
	bytes, err := idlresources0002IdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idlResources/0002.idl", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1741119739, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var _idlresourcesDefinesIdl = []byte(`#define da (a,a,d) (222)
#define da (a,d) \
    f(\
        ddd(),\
        a,s,d, 222, (222))

`)

func idlresourcesDefinesIdlBytes() ([]byte, error) {
	return _idlresourcesDefinesIdl, nil
}

func idlresourcesDefinesIdl() (*asset, error) {
	bytes, err := idlresourcesDefinesIdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idlResources/defines.idl", size: 96, mode: os.FileMode(0644), modTime: time.Unix(1741119740, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe6, 0x96, 0x78, 0xf2, 0xfa, 0x94, 0xfa, 0x2, 0x5, 0x41, 0x2b, 0xb5, 0x92, 0x30, 0x40, 0xc4, 0x38, 0xbe, 0x32, 0x66, 0x7f, 0xfd, 0x8f, 0xf8, 0x9a, 0x25, 0x7b, 0x26, 0x7f, 0x49, 0x35, 0x33}}
	return a, nil
}

var _idlresourcesFitmappingsIdl = []byte(`#include "0001.idl"

typedef uint64 FitInt64z;
typedef uint32 FitInt32z;
typedef uint16 FitInt16z;
typedef uint8  FitInt08z;

typedef uint64 FitUint64;
typedef uint32 FitUint32;
typedef uint16 FitUint16;
typedef uint8  FitUint08;

typedef int64 FitInt64;
typedef int32 FitInt32;
typedef int16 FitInt16;
typedef int8  FitInt08;




typedef sequence<sequence<FitInt64z>> FitInt64zArray;
typedef sequence<sequence<int32>> Int32Array;
`)

func idlresourcesFitmappingsIdlBytes() ([]byte, error) {
	return _idlresourcesFitmappingsIdl, nil
}

func idlresourcesFitmappingsIdl() (*asset, error) {
	bytes, err := idlresourcesFitmappingsIdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idlResources/fitMappings.idl", size: 431, mode: os.FileMode(0644), modTime: time.Unix(1741119739, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x6, 0x72, 0x89, 0x7b, 0x2e, 0x97, 0x37, 0x8a, 0xd7, 0x3b, 0x5d, 0x1, 0x24, 0x4f, 0xb4, 0xa6, 0x3d, 0xf4, 0x4e, 0xf4, 0x11, 0xc1, 0x51, 0x9, 0xb7, 0xa1, 0xbb, 0x7e, 0x4e, 0x13, 0xd8}}
	return a, nil
}

var _idlresourcesFitmessageIdl = []byte(`typedef int32 dd;



fdsfd`)

func idlresourcesFitmessageIdlBytes() ([]byte, error) {
	return _idlresourcesFitmessageIdl, nil
}

func idlresourcesFitmessageIdl() (*asset, error) {
	bytes, err := idlresourcesFitmessageIdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idlResources/fitMessage.idl", size: 26, mode: os.FileMode(0644), modTime: time.Unix(1741160690, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x50, 0x9c, 0xfc, 0xaf, 0x7, 0x8a, 0x15, 0x51, 0x38, 0x30, 0x2c, 0x9a, 0xa8, 0xa7, 0x56, 0x5f, 0x27, 0xb9, 0xee, 0xe5, 0x2d, 0xc1, 0xa, 0x72, 0xfb, 0x88, 0x68, 0x35, 0x77, 0x95, 0x33}}
	return a, nil
}

var _idlresourcesStructIdl = []byte(`#define DD (233+2)+222+2+2+2
#define multiply( f1, f2 ) ( f1 * f2 )
#define multiply( a1, a2 ) ( a1 * a2 )

#include "fitMappings.idl"


const double D01 = 222+2+2+2+2+2;
const double D02 = 1.3;
const double D03 = .0;
const double D04 = 1.3333e02;
const double D05 = 1.3333e-02;
const double D06 = -1.3333e02;
const double D07 = -1.3333e-02;

const float D011 = 0.0;
const float D022 = 1.3;
const float D033 = .0;
const float D044 = 1.3333e02;
const float D055 = 1.3333e-02;
const float D066 = -1.3333e02;
const float D077 = -1.3333e-02;





const FitInt64z HexNumber = 0xFF;
const string StringNumber = "dddd";
const wstring WideStringNumber = "dddd";


enum NewEnum {a,b,c,e};
const FitInt64z a = ~12;
const int32 b = 12+3;
const int32 c = 12/2;
const int32 d = 12+4*5;
const int32 e = 12-4*5;
struct AA
{
    short Short01;
     short Short02;
    unsigned short Short03;
    unsigned short Short04;
    int32 a;
    int16 b;
    int64 c;
    uint32 d;
    uint16 e;
    uint64 f;
    long g;
    float h;
    double i;
    long long j;
    unsigned long long k;
    sequence<int32> l;
    sequence<int32,3> m;
    short o;
    unsigned short p;
    FitInt64z zz;
    FitInt64zArray zz02;
    NewEnum NewEnum;
    string String01;
    wstring String02;
};

enum Normal
{
    a,b,c,d,e,f
};



struct BB
{
    short Short01;
     short Short02;
    unsigned short Short03;
    unsigned short Short04;
    int32 a;
    int16 b;
    int64 c;
};




enum Abnormal
{
    a=1,b=2,c=3,d=4,e=5,f=6
};



typedef int32 dddd01, dddd02;


typedef struct Proxy
{
    int32 A;
} ProxyA,ProxyB;



`)

func idlresourcesStructIdlBytes() ([]byte, error) {
	return _idlresourcesStructIdl, nil
}

func idlresourcesStructIdl() (*asset, error) {
	bytes, err := idlresourcesStructIdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idlResources/struct.idl", size: 1588, mode: os.FileMode(0644), modTime: time.Unix(1741124568, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0x13, 0xcd, 0x1c, 0x67, 0xad, 0xca, 0xe4, 0x3e, 0x84, 0xa3, 0x67, 0xa4, 0xde, 0x76, 0xdb, 0xbe, 0x87, 0x37, 0xe2, 0xa6, 0x6c, 0xa3, 0x93, 0x4d, 0x64, 0x48, 0x6b, 0x37, 0xbb, 0x2d, 0x1a}}
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
	"idlResources/0001.idl":        idlresources0001Idl,
	"idlResources/0002.idl":        idlresources0002Idl,
	"idlResources/defines.idl":     idlresourcesDefinesIdl,
	"idlResources/fitMappings.idl": idlresourcesFitmappingsIdl,
	"idlResources/fitMessage.idl":  idlresourcesFitmessageIdl,
	"idlResources/struct.idl":      idlresourcesStructIdl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"idlResources": {nil, map[string]*bintree{
		"0001.idl":        {idlresources0001Idl, map[string]*bintree{}},
		"0002.idl":        {idlresources0002Idl, map[string]*bintree{}},
		"defines.idl":     {idlresourcesDefinesIdl, map[string]*bintree{}},
		"fitMappings.idl": {idlresourcesFitmappingsIdl, map[string]*bintree{}},
		"fitMessage.idl":  {idlresourcesFitmessageIdl, map[string]*bintree{}},
		"struct.idl":      {idlresourcesStructIdl, map[string]*bintree{}},
	}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
