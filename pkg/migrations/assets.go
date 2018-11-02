// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/20180519220506_add_events_table.down.sql (142B)
// sql/20180519220506_add_events_table.up.sql (312B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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

var __20180519220506_add_events_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2d\x4b\xcd\x2b\x29\x8e\x2f\x28\x4d\xca\xc9\x4c\x8e\xcf\x4e\xad\x8c\xcf\x4c\xa9\x50\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\xc2\xa7\xa7\x28\x35\x39\xbf\x28\x25\x35\x25\x3e\xb1\x04\x55\x13\x44\x57\x88\xa3\x93\x8f\x2b\x86\x2e\xb8\x2a\x40\x00\x00\x00\xff\xff\x33\x48\xa7\x81\x8e\x00\x00\x00")

func _20180519220506_add_events_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180519220506_add_events_tableDownSql,
		"20180519220506_add_events_table.down.sql",
	)
}

func _20180519220506_add_events_tableDownSql() (*asset, error) {
	bytes, err := _20180519220506_add_events_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180519220506_add_events_table.down.sql", size: 142, mode: os.FileMode(420), modTime: time.Unix(1541167761, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0x58, 0x29, 0x23, 0xf6, 0x14, 0xa8, 0x5, 0x65, 0xb8, 0x2c, 0xae, 0x74, 0x15, 0x8b, 0x36, 0x94, 0x51, 0x44, 0x8d, 0x83, 0x66, 0x9d, 0x58, 0x68, 0x3c, 0x75, 0xc5, 0xb9, 0x5e, 0xa4, 0x5e}}
	return a, nil
}

var __20180519220506_add_events_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xc1\x4a\x86\x40\x18\x45\xf7\xf3\x14\x77\xa9\xd0\x1b\xb8\x1a\xf3\x93\x86\xc6\x51\x66\x3e\x51\xdb\x0c\xe6\xcc\x42\x8a\x0a\xb3\xa8\xb7\x0f\x25\x52\xe8\x5f\xfc\xcb\xcb\xe1\x5c\xce\xad\x25\xc9\x04\x96\xb9\x26\xa8\x12\xa6\x66\x50\xaf\x1c\x3b\xc4\xcf\xf8\xb2\xbe\x23\x11\xc0\x1c\xe0\xc8\x2a\xa9\xd1\x58\x55\x49\x3b\xe0\x9e\x86\x1b\x01\xbc\x7d\x3c\x3e\xcf\x93\x7f\x8a\xdf\x60\xea\x79\xd7\x4d\xab\xf5\xc6\x96\x38\xbd\x2e\x21\x06\x3f\xae\x60\x55\x91\x63\x59\x35\xe8\x14\xdf\xed\x13\x0f\xb5\x21\x14\x54\xca\x56\x6f\x62\x97\xa4\x9b\x15\xc6\x75\x44\x3e\x30\x49\x91\x66\x42\xfc\xf6\x29\x53\x50\x7f\xb1\xcf\x1f\x09\x7e\x0e\x5f\x02\xa8\xcd\x5f\xfa\xc1\xae\xfb\x3a\x25\xff\x3f\x3b\xc1\x34\xfb\x09\x00\x00\xff\xff\x4c\xd7\xeb\x51\x38\x01\x00\x00")

func _20180519220506_add_events_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180519220506_add_events_tableUpSql,
		"20180519220506_add_events_table.up.sql",
	)
}

func _20180519220506_add_events_tableUpSql() (*asset, error) {
	bytes, err := _20180519220506_add_events_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180519220506_add_events_table.up.sql", size: 312, mode: os.FileMode(420), modTime: time.Unix(1541167761, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x94, 0x5, 0xb6, 0x0, 0x5, 0x4, 0xd3, 0xb, 0x2f, 0xda, 0xc9, 0xab, 0xfb, 0x32, 0x8f, 0xc2, 0x3, 0x46, 0x60, 0xa3, 0xf2, 0x8f, 0x31, 0x3, 0xd2, 0x2f, 0xbc, 0xbe, 0xd2, 0x17, 0x7, 0x66}}
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
	"20180519220506_add_events_table.down.sql": _20180519220506_add_events_tableDownSql,

	"20180519220506_add_events_table.up.sql": _20180519220506_add_events_tableUpSql,
}

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
	"20180519220506_add_events_table.down.sql": &bintree{_20180519220506_add_events_tableDownSql, map[string]*bintree{}},
	"20180519220506_add_events_table.up.sql":   &bintree{_20180519220506_add_events_tableUpSql, map[string]*bintree{}},
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
