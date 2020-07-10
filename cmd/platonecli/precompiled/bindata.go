package precompile

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _release_linux_conf_contracts_cnsmanager_cpp_abi_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x41\x6b\xb4\x30\x10\xbd\xfb\x2b\x42\xce\x7b\xfa\xbe\xd2\xc3\xde\xba\x85\x42\x2f\x16\xda\xe3\xb2\x87\xa0\xa3\x0c\xd4\x89\x24\x93\x05\x5b\xfa\xdf\x8b\x82\xae\xae\x60\x75\xb5\xed\x36\xe8\x41\xd1\x24\xf3\x98\xf7\x26\x2f\x83\xfb\x40\x08\x21\xde\xab\x7b\x79\x49\x52\x19\xc8\xad\x90\x11\xd9\x67\x48\xd1\x32\x98\x07\xa3\xb3\x47\x42\x96\x9b\xd3\x34\xa4\xdc\xb1\x95\x5b\xb1\x6f\xbe\x75\x03\xf5\x02\x56\xcf\x4d\x7f\x9c\x8b\xbc\x1a\xb7\x6c\x90\x52\xd9\x99\xf0\xb1\x19\x1b\xfd\x08\xc6\xa2\xa6\xc9\x00\xcd\xdb\xa1\x95\x9c\x76\x3c\x35\xbb\x21\x60\x24\xfe\xff\x6f\x0c\x6e\xa4\xc9\xb2\x22\x2e\x17\x25\xea\xd5\xb6\xf9\x6a\xa2\x25\x8e\x22\x2e\x53\x0d\x5a\x0c\x0d\x0b\xe8\xa3\x70\xa3\x01\x54\x1c\x1b\xb0\x76\xad\x8c\x6e\x65\xc4\x68\x20\x5a\xb7\xf4\x9f\x12\x2e\x05\xbe\xd7\xc4\x46\x45\x7c\xd7\xab\xea\x55\xbf\x33\xfc\x65\x80\xdb\x02\xb2\x71\x73\xf5\xab\x2d\x19\xe2\x5a\xc9\x99\x1a\xe6\x2a\x85\xd0\x65\x53\x8b\x75\xb4\x8a\x65\xfc\x17\x7c\x1b\xac\x93\xd1\xbb\xc1\x63\x15\x77\xc5\x93\xc1\x14\x69\x9e\x9a\xfa\x3c\xc6\x8c\x9c\xbd\x20\x1b\x93\x13\xd7\xbb\x22\xec\x1a\xd6\x4f\x59\xde\xd5\x9e\x28\xcb\xb2\xbb\xc8\x99\xe2\x5d\xbb\xf5\x5d\x76\xe1\x3d\xdb\xd7\xe5\xce\x1e\x5b\xc7\xef\x10\x1d\x6a\xc6\xa4\xb8\x8c\xd2\x1a\xcb\x21\xf1\xed\xcd\xd4\xd6\x64\x7a\xde\xf5\x0a\x38\x02\xf1\x97\xa6\x78\xf1\xaf\x8e\x79\x7d\xf3\x32\x69\x05\x87\xe0\x33\x00\x00\xff\xff\x6b\x0c\x17\xf6\xd4\x11\x00\x00")

func release_linux_conf_contracts_cnsmanager_cpp_abi_json() ([]byte, error) {
	return bindata_read(
		_release_linux_conf_contracts_cnsmanager_cpp_abi_json,
		"../../release/linux/conf/contracts/cnsManager.cpp.abi.json",
	)
}

var _release_linux_conf_contracts_firewall_abi_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\x41\x4b\xc4\x30\x10\x85\xef\xfd\x15\x21\xe7\x9e\xf4\xb6\x37\x51\x04\x4f\x1e\x3c\x2e\x4b\x09\x9b\x59\x09\x64\x27\x21\x33\x41\x16\xd9\xff\x2e\xbb\xc5\x9a\x6a\xd1\x16\x5a\x2c\x21\x3d\x14\xda\x4e\xbf\x47\xdf\xbc\x07\xdd\x56\x42\x08\xf1\x7e\x3d\x5f\x0e\x89\xea\x08\x72\x23\x64\xd3\xd0\x89\x9a\xc7\xb7\x67\x0f\x28\xeb\xaf\xe7\x06\x7d\x64\x92\x1b\xb1\xed\xee\xf5\x09\x3f\x48\x4a\xeb\x00\x44\x09\xa4\x1b\xe1\x93\xbf\x8e\x10\x07\x83\xaf\xb2\x37\x70\xee\xae\x76\x89\xbc\x8b\x3c\x55\xff\x37\x61\x83\x7c\x7b\x33\x46\x77\xef\x90\x58\x21\x5f\x5e\x3a\x28\x4b\x90\x7a\xf2\x49\x3b\x44\xdc\xb3\x71\xd8\x02\xcf\xf5\x5f\xde\xde\x5b\xd7\x07\x15\x73\xe7\x33\xf7\x4e\xeb\xff\xb1\xb6\x1e\x2d\xd0\x7e\xcf\x62\xfc\x10\x2d\x94\x64\x0c\xd5\x0e\x54\xc8\x33\x1b\xf9\x2f\xef\x01\x6c\x9e\xab\x2b\xb5\xfe\x46\x9b\x9a\x8c\x17\xe0\x92\x8c\x92\x8c\xa1\x64\xb0\xe2\x48\xf9\xfe\x68\x8d\x17\x4e\xed\xe5\x10\x67\x70\xf7\xe9\xe8\x5d\x58\x7b\xf1\x56\xb9\xb5\x45\x3b\x51\xed\xaa\x8f\x00\x00\x00\xff\xff\x89\xd7\x9c\x64\xdb\x0d\x00\x00")

func release_linux_conf_contracts_firewall_abi_json() ([]byte, error) {
	return bindata_read(
		_release_linux_conf_contracts_firewall_abi_json,
		"../../release/linux/conf/contracts/fireWall.abi.json",
	)
}

var _release_linux_conf_contracts_groupmanager_cpp_abi_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\xb1\x4e\xc3\x30\x10\xdd\xfb\x15\x96\xe7\x4c\x80\x18\xba\x51\x55\x42\x65\x28\x48\x8c\x55\x07\x2b\xbe\x04\x4b\xc9\xd9\xf2\x9d\x87\x0a\xf5\xdf\x91\x8d\x1a\x52\x01\xa5\x60\x20\x15\x24\x43\xa4\x24\xe7\xe7\x7b\xef\xdd\x8b\x57\x13\x21\x84\x78\x4c\xf7\x78\x49\x54\x2d\xc8\xa9\x90\xa5\x07\xc5\x70\xed\x6d\x70\xb2\x78\xf9\x6c\xd0\x05\x26\x39\x15\xab\xee\xdd\x3e\xc0\x2b\xa0\x3a\x42\x2c\xb0\xb2\x3d\x98\xae\x88\x37\x2e\x15\x11\x7b\x83\xb5\xdc\x2b\xd8\x76\x4f\xeb\x5e\x03\x36\xf0\x67\x3b\x38\xb4\xb1\x41\x3e\x3f\x3b\x66\xdf\xd2\x22\xb1\x42\x8e\x8b\x2a\xd5\x10\xf4\x55\xd9\xa1\x55\x01\x4b\x36\x16\x9f\x01\xb7\xc5\x7b\xe2\xd6\xc0\x57\x4d\x93\xc4\xa5\xb7\xd5\xfd\x41\xca\xc7\x6b\xdd\xe7\xcc\x3e\x64\x51\x7e\x50\x94\xf8\xde\xba\x3b\xf0\xad\x21\x8a\x6b\x7e\x9b\xfa\x97\xdc\xce\x65\x8e\x56\x03\x2d\x43\x9b\x17\xa3\x88\x72\x43\x16\xef\xd9\xff\x9d\x20\xe5\x4a\x5b\x03\xa7\xa1\x9a\x6d\x16\xf3\xef\xf8\x4b\xcd\x0f\x31\x0c\x06\xf9\xf2\x62\x00\x69\x87\x09\x6c\x70\x5a\x31\xcc\xac\xe5\x65\x1c\xe0\x61\xe4\x2d\xfe\x75\x3c\xb2\xcf\x19\xa5\xf5\xce\xc0\xd1\xbf\x0f\x7b\x38\x3d\xff\x34\x34\xa3\x7f\xa7\xef\xdf\x44\xac\x9f\x02\x00\x00\xff\xff\x13\xb9\x10\x95\x48\x0b\x00\x00")

func release_linux_conf_contracts_groupmanager_cpp_abi_json() ([]byte, error) {
	return bindata_read(
		_release_linux_conf_contracts_groupmanager_cpp_abi_json,
		"../../release/linux/conf/contracts/groupManager.cpp.abi.json",
	)
}

var _release_linux_conf_contracts_nodemanager_cpp_abi_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x4f\x4b\xc3\x30\x18\xc6\xef\xfb\x14\x21\xe7\x9e\x54\x3c\xec\x26\xe8\x65\x42\x2f\x1e\xc7\x0e\xb1\x79\x5b\x02\xe9\x9b\xd2\x3c\x19\x14\xd9\x77\x97\x54\x5b\x3b\xd4\xb1\x5a\x65\xd0\xda\x43\x4b\xc3\xfb\x27\xcf\x2f\x79\x92\xed\x4a\x08\x21\x5e\xda\x77\x7c\x24\xab\x92\xe4\x5a\x48\xa5\xb5\x4c\x3e\x86\x0d\x57\x01\x5e\xae\xc5\xb6\x1f\x3b\x4e\xfc\x54\x80\x9d\xa6\x8d\x77\xfc\x84\x7a\x50\xa8\x0f\x43\x53\xb5\x61\x1e\xb5\xe1\x42\x1e\x05\x1c\xfa\xbf\xdd\x60\x0a\x2e\x60\xec\x1c\x4e\x35\x36\x8c\xeb\xab\x73\xfa\x66\x8e\x3d\x14\x23\x26\xe5\xca\x7a\x1a\x72\xe9\xaa\xe5\x81\x33\x18\xc7\x6f\x05\x0f\xc9\x77\x58\x0b\xc2\x9d\xb5\xa9\xd3\xe4\xbf\xc6\xfb\x87\x8a\xcf\x47\x3d\x94\x8c\x3a\x4c\x52\xbc\x57\xd6\xe8\x8d\x33\x1c\x45\x4f\xdb\x52\x55\x78\xb6\x26\x7b\xa4\x66\x3e\x1b\x6a\x2a\xdd\xe8\x32\x9f\x86\xf2\xdf\xab\xbf\x8e\xb6\x20\x9c\xf0\xe9\x5c\xd0\x5e\xe6\x50\x08\x95\x56\x98\x78\x1a\xb4\xdf\xb1\xca\x92\xb9\xac\xdb\xc5\xae\xaf\xd4\xd5\xa5\xb2\x0f\x91\xcf\x92\xae\xb1\x82\x70\x4f\x96\x40\x7a\x71\xd2\x53\x07\x93\x37\x3f\x33\x6b\xd7\x2b\x18\xc6\xed\xcd\x58\x33\x8e\xd7\xdd\x65\xd0\x9e\x18\xef\xc2\x56\xbb\xd7\x00\x00\x00\xff\xff\x67\x87\x5c\x95\xe6\x0a\x00\x00")

func release_linux_conf_contracts_nodemanager_cpp_abi_json() ([]byte, error) {
	return bindata_read(
		_release_linux_conf_contracts_nodemanager_cpp_abi_json,
		"../../release/linux/conf/contracts/nodeManager.cpp.abi.json",
	)
}

var _release_linux_conf_contracts_parammanager_cpp_abi_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\x41\x4f\xc2\x30\x14\xbe\xef\x57\x2c\x3b\x73\x52\xe3\x81\xdb\x44\x43\x4c\x8c\xe1\x80\x27\xc2\xa1\x96\x82\x0d\x5b\xbb\xf4\xbd\x2a\x8b\xe1\xbf\x9b\xb2\x0c\x21\x2b\xae\xd2\x61\xea\x38\x90\xac\xe9\xbe\x7e\xef\xfb\xbe\xee\xbd\x59\x14\xc7\x71\xfc\xb9\xfb\x37\xbf\x44\x90\x9c\x25\xc3\x38\x01\x86\x63\x02\x23\x29\x50\x11\x8a\xcf\x66\x75\xf0\xbd\x8b\x8b\x42\x23\x24\xc3\x78\xb6\x5f\x3b\xc6\x69\xe0\x51\x3b\xd2\x7e\x1f\x96\x45\x75\x2e\x2a\x2e\x56\xc9\xd1\x86\xed\xfe\x69\x7e\xc0\x41\x6a\xac\x49\x1c\x2e\x53\x29\x00\x89\x40\x03\xb6\x24\x19\x1c\xf1\xae\x4f\x59\x6a\x41\x91\x4b\x51\x9d\xb3\x1d\x9c\x52\x61\xe5\xa8\xc2\x09\x5e\xae\xe2\x74\x23\xc8\x61\xe5\xa8\xb4\x57\xe1\xc0\xf0\x11\x26\x4a\x2e\x34\x65\x0f\x79\x81\xe5\x5d\x26\xe9\xda\x2f\x02\xfc\x47\xc0\x06\x55\xcd\x05\x5e\x5f\x05\x94\x04\x67\x41\x2e\x98\x06\x67\x4d\x3a\x0e\xc3\x74\x33\x26\xf0\xc4\x73\x8e\x7e\x19\x40\x1b\x4e\x83\x98\xb1\xfe\xf6\x26\x20\xeb\xdb\xca\xbf\xa0\xe3\xee\x5a\x74\x6c\xf9\x2e\xe0\xdd\xb8\xfe\x7a\x02\xca\xa3\xd8\x3f\x32\xde\x41\x84\x3e\x7a\x9f\x66\x99\xfc\x48\x45\x99\x52\x2a\xb5\xc0\x7b\x56\x64\xb2\xac\xdb\xa0\x6f\x17\x70\x05\xb7\xca\x11\x4a\x36\x80\xe1\xe8\x8d\xd1\x75\xcd\xbb\xaa\x62\xc2\x54\xce\x01\xcc\xfb\x7e\xd3\x92\x81\xb6\x82\x59\x45\x09\xaa\x4d\x9e\x21\x4b\xdf\x1a\xe6\xea\xac\x1b\xd4\x37\x15\x76\x33\x64\x5a\x14\x4a\xbe\xb3\xaa\x7e\xb6\xe8\xec\x1b\xd2\x0a\xdb\xa0\x1d\xdc\x45\xf9\xa5\x38\xbd\x8c\xc7\x74\xf3\x02\x6c\x4c\xc0\x37\x0f\x16\x9c\x06\xb1\xe0\x86\x8b\xb6\xf2\xff\xeb\x64\x11\xcd\xa3\xe8\x2b\x00\x00\xff\xff\xdd\x9e\xe4\x2c\x59\x10\x00\x00")

func release_linux_conf_contracts_parammanager_cpp_abi_json() ([]byte, error) {
	return bindata_read(
		_release_linux_conf_contracts_parammanager_cpp_abi_json,
		"../../release/linux/conf/contracts/paramManager.cpp.abi.json",
	)
}

var _release_linux_conf_contracts_usermanager_cpp_abi_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd9\x41\x6b\x2a\x31\x10\x00\xe0\xbb\xbf\x22\xe4\xec\xe9\xbd\x9b\x37\xdf\xf3\x52\x28\x16\x5a\x7a\x12\x0f\x61\x33\xbb\x06\x62\xb2\x24\xb3\x07\x29\xfe\xf7\x92\x85\xda\x5d\xd7\xec\xa6\x4d\xa4\x45\xe3\x41\x30\x2e\x93\x99\x2f\xc3\x08\x71\x33\x23\x84\x90\xb7\xf6\xdd\xbd\xa8\x62\x7b\xa0\x0b\x42\x2d\xe0\x4b\x53\x83\x59\xf2\xbd\x50\x74\xfe\xf9\x80\x50\x75\x83\x96\x2e\xc8\x66\xdb\x59\xd5\x0d\x5e\x5a\x2e\xb4\xb2\xc8\x14\xba\x88\x25\x93\x16\xba\x91\xf0\x50\xb7\x5b\x95\x8d\x2a\x50\x68\x45\xdb\xaf\x8e\x73\x5f\x4a\x15\xe0\x92\x73\xf3\x28\x2c\x3e\x95\xcf\x5a\xc2\xe5\xb4\x4e\x6b\xfd\x30\x83\x70\xc8\x4c\x05\x78\x16\x67\x90\x9c\x45\x23\x54\x45\x7b\x0f\x1c\x4f\x9f\x3c\x04\xa1\x29\xa4\xd9\xb8\x8b\x8c\xa6\x89\x35\x76\x22\xf6\xdf\xc1\x51\x83\xb5\x71\xc6\x8c\x73\x73\x87\xba\x84\x4c\xfb\xae\xdd\x4a\x14\xae\xea\x47\xb8\x1b\x5c\x9f\x2d\xe3\xfc\xff\x8e\x09\xd5\x8e\xac\xdc\xc0\x97\x8c\x83\x67\x70\x28\x72\x7c\x17\x67\xe1\x21\xc8\x5a\x73\x48\x06\x7c\x63\x63\x22\x1e\x98\x83\xcc\x73\x22\x15\x72\x60\x1b\x67\xe4\xab\x20\x73\x90\x19\x39\x0d\xf2\xd8\x0f\x9e\x56\x68\x58\x81\x99\xf9\xca\xbd\x9c\xa1\x53\x42\x8f\xce\xe6\x0f\xea\x15\xd4\x52\x1f\xc0\x64\xed\x48\xed\x80\xb6\xce\xd6\x49\xac\x47\xba\xfa\xd5\x82\x89\x43\x6d\x2c\x98\x07\x55\xea\x0c\xdb\x31\xa9\x39\x43\x70\xb6\x2b\xb0\xc5\x99\xce\x37\x1b\xb7\xdf\xfd\x61\x95\xce\x43\x37\xe0\xc3\x34\x23\x2c\x6f\xe3\x10\x2b\xc0\xa5\x94\xee\x10\x3d\x63\xe7\xd7\x95\x9c\xe0\x2e\xd3\x95\x3b\x31\x6d\xfd\xe5\x4d\xf4\xab\xb7\xba\xe3\xed\x4a\xfa\x6f\x23\x32\xe3\x34\xe3\x8e\x59\xff\xdf\x16\xe9\x01\xc9\x17\x66\x26\x35\x5a\xc2\x7a\xf4\xa6\xe8\xe7\x26\xa6\x50\xf8\xf7\xcf\x95\x0e\x6b\xb6\x7d\x0f\x00\x00\xff\xff\x7e\x86\x6e\xc1\x03\x1b\x00\x00")

func release_linux_conf_contracts_usermanager_cpp_abi_json() ([]byte, error) {
	return bindata_read(
		_release_linux_conf_contracts_usermanager_cpp_abi_json,
		"../../release/linux/conf/contracts/userManager.cpp.abi.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"../../release/linux/conf/contracts/cnsManager.cpp.abi.json": release_linux_conf_contracts_cnsmanager_cpp_abi_json,
	"../../release/linux/conf/contracts/fireWall.abi.json": release_linux_conf_contracts_firewall_abi_json,
	"../../release/linux/conf/contracts/groupManager.cpp.abi.json": release_linux_conf_contracts_groupmanager_cpp_abi_json,
	"../../release/linux/conf/contracts/nodeManager.cpp.abi.json": release_linux_conf_contracts_nodemanager_cpp_abi_json,
	"../../release/linux/conf/contracts/paramManager.cpp.abi.json": release_linux_conf_contracts_parammanager_cpp_abi_json,
	"../../release/linux/conf/contracts/userManager.cpp.abi.json": release_linux_conf_contracts_usermanager_cpp_abi_json,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"..": &_bintree_t{nil, map[string]*_bintree_t{
		"..": &_bintree_t{nil, map[string]*_bintree_t{
			"release": &_bintree_t{nil, map[string]*_bintree_t{
				"linux": &_bintree_t{nil, map[string]*_bintree_t{
					"conf": &_bintree_t{nil, map[string]*_bintree_t{
						"contracts": &_bintree_t{nil, map[string]*_bintree_t{
							"cnsManager.cpp.abi.json": &_bintree_t{release_linux_conf_contracts_cnsmanager_cpp_abi_json, map[string]*_bintree_t{
							}},
							"fireWall.abi.json": &_bintree_t{release_linux_conf_contracts_firewall_abi_json, map[string]*_bintree_t{
							}},
							"groupManager.cpp.abi.json": &_bintree_t{release_linux_conf_contracts_groupmanager_cpp_abi_json, map[string]*_bintree_t{
							}},
							"nodeManager.cpp.abi.json": &_bintree_t{release_linux_conf_contracts_nodemanager_cpp_abi_json, map[string]*_bintree_t{
							}},
							"paramManager.cpp.abi.json": &_bintree_t{release_linux_conf_contracts_parammanager_cpp_abi_json, map[string]*_bintree_t{
							}},
							"userManager.cpp.abi.json": &_bintree_t{release_linux_conf_contracts_usermanager_cpp_abi_json, map[string]*_bintree_t{
							}},
						}},
					}},
				}},
			}},
		}},
	}},
}}