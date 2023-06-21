// Code generated by go generate; DO NOT EDIT.
package images

import (
	"io"
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *PushOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *PushOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAll set field All to given value
func (o *PushOptions) WithAll(value bool) *PushOptions {
	o.All = &value
	return o
}

// GetAll returns value of field All
func (o *PushOptions) GetAll() bool {
	if o.All == nil {
		var z bool
		return z
	}
	return *o.All
}

// WithAuthfile set field Authfile to given value
func (o *PushOptions) WithAuthfile(value string) *PushOptions {
	o.Authfile = &value
	return o
}

// GetAuthfile returns value of field Authfile
func (o *PushOptions) GetAuthfile() string {
	if o.Authfile == nil {
		var z string
		return z
	}
	return *o.Authfile
}

// WithCompress set field Compress to given value
func (o *PushOptions) WithCompress(value bool) *PushOptions {
	o.Compress = &value
	return o
}

// GetCompress returns value of field Compress
func (o *PushOptions) GetCompress() bool {
	if o.Compress == nil {
		var z bool
		return z
	}
	return *o.Compress
}

// WithCompressionFormat set field CompressionFormat to given value
func (o *PushOptions) WithCompressionFormat(value string) *PushOptions {
	o.CompressionFormat = &value
	return o
}

// GetCompressionFormat returns value of field CompressionFormat
func (o *PushOptions) GetCompressionFormat() string {
	if o.CompressionFormat == nil {
		var z string
		return z
	}
	return *o.CompressionFormat
}

// WithCompressionLevel set field CompressionLevel to given value
func (o *PushOptions) WithCompressionLevel(value int) *PushOptions {
	o.CompressionLevel = &value
	return o
}

// GetCompressionLevel returns value of field CompressionLevel
func (o *PushOptions) GetCompressionLevel() int {
	if o.CompressionLevel == nil {
		var z int
		return z
	}
	return *o.CompressionLevel
}

// WithFormat set field Format to given value
func (o *PushOptions) WithFormat(value string) *PushOptions {
	o.Format = &value
	return o
}

// GetFormat returns value of field Format
func (o *PushOptions) GetFormat() string {
	if o.Format == nil {
		var z string
		return z
	}
	return *o.Format
}

// WithPassword set field Password to given value
func (o *PushOptions) WithPassword(value string) *PushOptions {
	o.Password = &value
	return o
}

// GetPassword returns value of field Password
func (o *PushOptions) GetPassword() string {
	if o.Password == nil {
		var z string
		return z
	}
	return *o.Password
}

// WithProgressWriter set field ProgressWriter to given value
func (o *PushOptions) WithProgressWriter(value io.Writer) *PushOptions {
	o.ProgressWriter = &value
	return o
}

// GetProgressWriter returns value of field ProgressWriter
func (o *PushOptions) GetProgressWriter() io.Writer {
	if o.ProgressWriter == nil {
		var z io.Writer
		return z
	}
	return *o.ProgressWriter
}

// WithSkipTLSVerify set field SkipTLSVerify to given value
func (o *PushOptions) WithSkipTLSVerify(value bool) *PushOptions {
	o.SkipTLSVerify = &value
	return o
}

// GetSkipTLSVerify returns value of field SkipTLSVerify
func (o *PushOptions) GetSkipTLSVerify() bool {
	if o.SkipTLSVerify == nil {
		var z bool
		return z
	}
	return *o.SkipTLSVerify
}

// WithRemoveSignatures set field RemoveSignatures to given value
func (o *PushOptions) WithRemoveSignatures(value bool) *PushOptions {
	o.RemoveSignatures = &value
	return o
}

// GetRemoveSignatures returns value of field RemoveSignatures
func (o *PushOptions) GetRemoveSignatures() bool {
	if o.RemoveSignatures == nil {
		var z bool
		return z
	}
	return *o.RemoveSignatures
}

// WithUsername set field Username to given value
func (o *PushOptions) WithUsername(value string) *PushOptions {
	o.Username = &value
	return o
}

// GetUsername returns value of field Username
func (o *PushOptions) GetUsername() string {
	if o.Username == nil {
		var z string
		return z
	}
	return *o.Username
}

// WithQuiet set field Quiet to given value
func (o *PushOptions) WithQuiet(value bool) *PushOptions {
	o.Quiet = &value
	return o
}

// GetQuiet returns value of field Quiet
func (o *PushOptions) GetQuiet() bool {
	if o.Quiet == nil {
		var z bool
		return z
	}
	return *o.Quiet
}

// WithManifestDigest set field ManifestDigest to given value
func (o *PushOptions) WithManifestDigest(value string) *PushOptions {
	o.ManifestDigest = &value
	return o
}

// GetManifestDigest returns value of field ManifestDigest
func (o *PushOptions) GetManifestDigest() string {
	if o.ManifestDigest == nil {
		var z string
		return z
	}
	return *o.ManifestDigest
}
