// Code generated by go-bindata.
// sources:
// template/footer.html
// template/header.html
// template/table2.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
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

var _templateFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\x5d\x6f\xda\x3c\x1c\x47\xef\xf9\x14\x56\xf4\x48\x29\xea\x83\x0d\x84\x42\x83\x00\x89\xb5\xc0\x4a\x06\x23\x20\x28\x74\xda\x85\x93\x38\x89\x43\xfc\x52\xdb\xc9\x60\xd5\xbe\xfb\x44\x61\x2f\x37\xab\x34\xed\xee\xf7\x3f\x17\x7f\x9d\x03\x00\x00\x3d\x14\xd1\x72\x50\x79\x9d\x3a\x54\x54\x1a\xa0\x55\xd8\xb7\x52\x63\xa4\xee\x22\x14\x8a\x88\xc0\xec\xb9\x20\xea\x08\x43\xc1\xd0\x79\xd6\x1c\xd8\x84\x0d\xa8\x73\xca\x20\xa3\x1c\x66\xda\x02\x94\x1b\x92\x28\x6a\x8e\x7d\x4b\xa7\xd8\xb9\x6d\xd5\xbc\xa9\x23\x9a\xf7\x9e\x79\xd8\x97\xbb\x07\xcf\x59\x8f\xe6\x5f\xd9\xac\xe3\xdd\xed\x97\x0a\xa9\x91\x8b\x7c\x99\xb4\xf1\xf0\x69\x32\xfd\x32\xbe\x9f\x6d\xe6\x43\x34\x91\x93\xf1\xd8\x75\xd2\xad\x9c\xdc\x78\xfb\xb9\x05\x42\x25\xb4\x16\x8a\x26\x94\xf7\x2d\xcc\x05\x3f\x32\x51\x68\x6b\xd0\x43\x67\xd7\xb7\xc4\x23\x9e\x69\x18\xe6\xa2\x88\xe2\x1c\x2b\xf2\x6a\x8f\x33\x7c\x40\x39\x0d\x34\x92\x42\x4a\xa2\x60\xa6\x51\x03\x36\x9a\xd0\x45\x05\x8b\x7e\xc0\x3f\x17\x0d\xe5\x3c\x48\x52\xf7\xdd\xf5\xae\xe1\x7b\xa6\x74\x96\xbc\xf3\xe8\xb0\x64\x71\x48\xd7\xae\x87\x56\xa1\xaf\x87\x8b\x4e\xba\xa6\xc1\xd6\x71\xb3\x4e\x8c\xf7\xe3\x85\xde\x97\xdb\x42\x97\x31\xae\x07\x2d\xff\x9f\x8a\x18\x3e\x84\x11\x87\x81\x10\x46\x1b\x85\xe5\xe9\x38\x45\xfd\x04\xa8\x05\xeb\xb0\x8e\x32\xfd\x0b\xbd\xd1\x32\x7d\x5a\xb6\x57\x92\x64\x69\x6b\x5d\x6f\x46\xb7\xd9\x47\xd3\x2e\x3f\x8c\xde\xc7\x04\x4d\xfd\x09\x5d\x2e\x57\xbe\x7f\x58\xc5\xe3\x47\x49\x1b\xb3\xe7\x62\x13\x0d\x8f\xd9\x1a\xab\x9b\xeb\x4e\x7b\xb1\xb9\x63\xbb\xfc\x2f\x5b\xce\x07\x00\xff\x5d\xc5\x05\x0f\x0d\x15\x1c\x5c\x55\xc1\xcb\x85\x9e\xb8\xfd\x29\xc2\x06\xd7\x8c\x48\x92\x9c\xf4\x2d\x23\x44\x6e\xa8\xb4\x3e\xdb\x55\x28\x85\x14\x25\x51\x57\x2f\x46\xd1\x24\x21\xaa\x0b\xec\xf4\x04\xec\xff\x81\xcc\x71\x48\x18\xe1\xa6\x0b\xec\x40\x18\x23\x98\xfd\xad\x7a\xf9\x7a\x19\xbf\xfb\xf4\x50\x20\xa2\xe3\xa0\xd2\x43\xa9\x61\xf9\xa0\x52\xf9\x1e\x00\x00\xff\xff\x86\xb1\xe1\x56\x05\x03\x00\x00")

func templateFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateFooterHtml,
		"template/footer.html",
	)
}

func templateFooterHtml() (*asset, error) {
	bytes, err := templateFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/footer.html", size: 773, mode: os.FileMode(420), modTime: time.Unix(1519499825, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6f\xe2\x30\x10\x85\xef\xf9\x15\xb3\xbe\xae\x88\x03\x04\x36\x62\x93\x48\xd1\x2e\xa0\x72\x68\x69\x41\x22\x3d\x9a\xc4\xc4\xa3\xc6\x76\x1a\xbb\x90\xa8\xea\x7f\xaf\x92\x42\xdb\x43\x7b\xe8\xc9\xe3\x4f\xcf\xef\xf9\x69\xc2\x5f\xff\x6f\xfe\x6d\xef\xd7\x73\x10\x56\x96\xb1\x13\xbe\x1d\x00\xa1\xe0\x2c\xef\x06\x80\xb0\x44\xf5\x00\xa2\xe6\x87\x88\x08\x6b\x2b\x33\xa3\x54\xb2\x26\xcb\x95\xbb\xd7\xda\x1a\x5b\xb3\xaa\xbb\x64\x5a\xd2\x77\x40\x7d\xd7\x73\x3d\x9a\x19\xf3\xc1\x5c\x89\xca\xcd\x8c\x21\x50\xf3\x32\x22\xc6\xb6\x25\x37\x82\x73\x4b\x00\x95\xe5\x45\x8d\xb6\x8d\x88\x11\x6c\x1c\xf8\x83\xa5\x9a\x8c\x03\xbf\x79\xbc\x1d\x32\xbd\x4b\x93\xdf\xde\x24\xb8\x4b\xd7\xcd\xba\x98\x1e\x5a\xff\x6a\x77\xdc\x5e\x0b\x6f\x3e\x9a\x8e\x53\xb9\xc8\x56\xe5\x26\x39\xe1\xb2\x58\x24\x3b\x9a\x27\xb8\x99\xae\x52\x49\x20\xab\xb5\x31\xba\xc6\x02\x55\x44\x98\xd2\xaa\x95\xfa\xc9\x90\x9f\x97\x3a\x68\x65\x07\xec\xc4\x8d\x96\x9c\xfa\xee\x9f\x73\xaf\xcf\xf8\xfb\x6a\xe7\xb8\x9e\xc4\x8e\xe3\x56\x58\xf1\x12\x15\x1f\xb0\xcc\xa2\x56\xf0\xec\x00\x48\x56\x17\xa8\x66\xe0\xc1\xd0\xab\x9a\xbf\xce\xcb\x17\xba\x18\x72\x3c\xf6\x6a\x80\x13\xe6\x56\xcc\x60\x14\xf4\xea\x8e\x5c\x1c\xba\xf7\xe0\xf5\x0e\x7d\x2c\xbd\xe4\x76\xf3\x65\xa5\xe1\x5e\xe7\xed\xf9\x5f\x39\x1e\x63\xe7\x35\x00\x00\xff\xff\xac\x67\x94\xbc\x06\x02\x00\x00")

func templateHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateHeaderHtml,
		"template/header.html",
	)
}

func templateHeaderHtml() (*asset, error) {
	bytes, err := templateHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/header.html", size: 518, mode: os.FileMode(420), modTime: time.Unix(1519503784, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateTable2Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xcf\x6e\xdb\x30\x0c\xc6\xef\x79\x0a\x42\x77\x3b\x1d\x7a\x0b\x1c\x03\xbd\x0c\x18\x30\x0c\xc5\xb2\x3d\x80\x22\xd3\x36\x51\x5b\xf2\x24\x66\x6d\x20\xf8\xdd\x07\xc9\x7f\x12\xc7\x4b\x93\x8b\x14\xea\xe3\x8f\x9f\x28\x3a\xab\x9f\xc1\xf1\xb9\xc1\xbd\x68\xa5\xad\x48\xef\xe0\x09\xbe\x3c\x75\x1f\x22\xf7\x1e\xd2\x1f\xb2\x45\xe8\xfb\x6c\x5b\x3f\xe7\x9b\xac\xa0\xbf\x93\xb8\x20\xd7\x35\xf2\xbc\x83\xb2\xc1\x0f\x01\xaa\x91\xce\xed\x45\x47\x1d\x36\xa4\x51\xe4\x1b\x00\xef\xc1\x4a\x5d\x21\xa4\xaf\x63\xd8\x41\xdf\x6f\x00\x22\xe7\x26\x23\x91\x8a\xc9\xe8\x98\x38\x28\x86\x1d\x40\xe6\xd8\x1a\x5d\x4d\x85\xdf\x6b\x62\x4c\x5c\x27\x15\xee\x40\x9b\x77\x2b\xbb\x1b\xab\x83\x7e\x04\x6d\x23\x29\xee\x2f\x7e\x5e\x62\xad\xd1\xcd\x4d\x35\x39\x39\x3b\xb2\x86\x23\xeb\xc4\xb5\x71\x39\x36\x46\xbd\x6d\x46\x55\x80\x25\x40\x25\xe0\x1f\x10\x87\x93\x52\x88\x05\x16\x02\xd2\x03\x4b\x3e\xcd\xe0\xf0\x0b\xb9\xe6\xc4\xf1\x8e\x2e\x28\x9d\x5b\x40\xb0\x71\x38\x91\xbe\x4a\x6a\x3e\xc1\x14\xc1\xbd\xbd\x9b\xfd\x4d\xbf\x5a\x53\x59\x74\xee\x2e\xa1\xb3\xd4\x4a\x7b\x5e\x21\xee\xf9\x45\x65\x74\xb1\x4a\xd0\x45\xdf\x5f\x3a\x21\xa0\xb6\x58\xee\x85\xf7\x54\x42\xfa\xfb\xe7\xf7\xbe\xf7\x7e\x5a\xa3\x56\x40\x21\x59\x26\x6c\xaa\x2a\xbc\x20\x1b\xd3\x30\x75\x02\x98\x38\xfc\xf7\x7e\x74\x3b\x2b\x95\xd1\x8c\x9a\xf7\x62\x51\x37\x7d\x39\x71\x6d\xec\xf5\x9d\x62\xf8\x70\x6a\xc3\xa5\xae\xe3\x22\xbf\x7e\xa9\xd0\xa0\xf4\x17\xc5\xf1\x98\xc3\xe1\x20\xc6\xe2\xc9\x82\x39\x74\xf5\x46\x3b\x0f\xd8\x42\xa8\x0b\xb8\xee\xc5\x50\xeb\xf1\x50\x00\x64\x34\xcd\x59\x29\xa1\x94\x89\xaa\x51\xbd\x89\x3c\xdb\x52\xbe\x72\xf2\x78\x3a\xd6\x3c\xa6\x16\xdd\xe7\xbc\x07\xf3\xb2\x66\x5a\x2c\x2d\xba\x3a\x6c\x5d\x47\xfa\x3f\xf4\x45\x3b\xb2\xad\x5c\x7c\x84\x2b\xd1\x1c\xbf\x44\xc7\xd0\xbf\x00\x00\x00\xff\xff\x15\xaf\x68\x7e\x96\x04\x00\x00")

func templateTable2HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateTable2Html,
		"template/table2.html",
	)
}

func templateTable2Html() (*asset, error) {
	bytes, err := templateTable2HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/table2.html", size: 1174, mode: os.FileMode(420), modTime: time.Unix(1519504151, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"template/footer.html": templateFooterHtml,
	"template/header.html": templateHeaderHtml,
	"template/table2.html": templateTable2Html,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"template": &bintree{nil, map[string]*bintree{
		"footer.html": &bintree{templateFooterHtml, map[string]*bintree{}},
		"header.html": &bintree{templateHeaderHtml, map[string]*bintree{}},
		"table2.html": &bintree{templateTable2Html, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
