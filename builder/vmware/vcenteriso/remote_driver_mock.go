package vcenteriso

import (
	vmwcommon "github.com/mitchellh/packer/builder/vmware/common"
)

type RemoteDriverMock struct {
	vmwcommon.DriverMock

	UploadISOCalled bool
	UploadISOPath   string
	UploadISOResult string
	UploadISOErr    error

	RegisterCalled bool
	RegisterPath   string
	RegisterErr    error

	UnregisterCalled bool
	UnregisterPath   string
	UnregisterErr    error

	uploadErr error

	ReloadVMErr error
}

func (d *RemoteDriverMock) UploadISO(path string, checksum string, checksumType string) (string, error) {
	d.UploadISOCalled = true
	d.UploadISOPath = path
	return d.UploadISOResult, d.UploadISOErr
}

func (d *RemoteDriverMock) Register(path string) error {
	d.RegisterCalled = true
	d.RegisterPath = path
	return d.RegisterErr
}

func (d *RemoteDriverMock) Unregister(path string) error {
	d.UnregisterCalled = true
	d.UnregisterPath = path
	return d.UnregisterErr
}

func (d *RemoteDriverMock) upload(dst, src string) error {
	return d.uploadErr
}

func (d *RemoteDriverMock) ReloadVM() error {
	return d.ReloadVMErr
}
