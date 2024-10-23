// Code generated by tools/cmd/genoptions/main.go. DO NOT EDIT.

package jwe

import (
	"io/fs"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/option"
)

type Option = option.Interface

// CompactOption describes options that can be passed to `jwe.Compact`
type CompactOption interface {
	Option
	compactOption()
}

type compactOption struct {
	Option
}

func (*compactOption) compactOption() {}

// DecryptOption describes options that can be passed to `jwe.Decrypt`
type DecryptOption interface {
	Option
	decryptOption()
}

type decryptOption struct {
	Option
}

func (*decryptOption) decryptOption() {}

// EncryptDecryptOption describes options that can be passed to either `jwe.Encrypt` or `jwe.Decrypt`
type EncryptDecryptOption interface {
	Option
	encryptOption()
	decryptOption()
}

type encryptDecryptOption struct {
	Option
}

func (*encryptDecryptOption) encryptOption() {}

func (*encryptDecryptOption) decryptOption() {}

// EncryptOption describes options that can be passed to `jwe.Encrypt`
type EncryptOption interface {
	Option
	encryptOption()
}

type encryptOption struct {
	Option
}

func (*encryptOption) encryptOption() {}

// GlobalOption describes options that changes global settings for this package
type GlobalOption interface {
	Option
	globalOption()
}

type globalOption struct {
	Option
}

func (*globalOption) globalOption() {}

// ReadFileOption is a type of `Option` that can be passed to `jwe.Parse`
type ParseOption interface {
	Option
	readFileOption()
}

type parseOption struct {
	Option
}

func (*parseOption) readFileOption() {}

// ReadFileOption is a type of `Option` that can be passed to `jwe.ReadFile`
type ReadFileOption interface {
	Option
	readFileOption()
}

type readFileOption struct {
	Option
}

func (*readFileOption) readFileOption() {}

// JSONSuboption describes suboptions that can be passed to `jwe.WithJSON()` option
type WithJSONSuboption interface {
	Option
	withJSONSuboption()
}

type withJSONSuboption struct {
	Option
}

func (*withJSONSuboption) withJSONSuboption() {}

// WithKeySetSuboption is a suboption passed to the WithKeySet() option
type WithKeySetSuboption interface {
	Option
	withKeySetSuboption()
}

type withKeySetSuboption struct {
	Option
}

func (*withKeySetSuboption) withKeySetSuboption() {}

type identCEK struct{}
type identCompress struct{}
type identContentEncryptionAlgorithm struct{}
type identFS struct{}
type identKey struct{}
type identKeyProvider struct{}
type identKeyUsed struct{}
type identMaxBufferSize struct{}
type identMaxPBES2Count struct{}
type identMergeProtectedHeaders struct{}
type identMessage struct{}
type identPerRecipientHeaders struct{}
type identPretty struct{}
type identProtectedHeaders struct{}
type identRequireKid struct{}
type identSerialization struct{}

func (identCEK) String() string {
	return "WithCEK"
}

func (identCompress) String() string {
	return "WithCompress"
}

func (identContentEncryptionAlgorithm) String() string {
	return "WithContentEncryption"
}

func (identFS) String() string {
	return "WithFS"
}

func (identKey) String() string {
	return "WithKey"
}

func (identKeyProvider) String() string {
	return "WithKeyProvider"
}

func (identKeyUsed) String() string {
	return "WithKeyUsed"
}

func (identMaxBufferSize) String() string {
	return "WithMaxBufferSize"
}

func (identMaxPBES2Count) String() string {
	return "WithMaxPBES2Count"
}

func (identMergeProtectedHeaders) String() string {
	return "WithMergeProtectedHeaders"
}

func (identMessage) String() string {
	return "WithMessage"
}

func (identPerRecipientHeaders) String() string {
	return "WithPerRecipientHeaders"
}

func (identPretty) String() string {
	return "WithPretty"
}

func (identProtectedHeaders) String() string {
	return "WithProtectedHeaders"
}

func (identRequireKid) String() string {
	return "WithRequireKid"
}

func (identSerialization) String() string {
	return "WithSerialization"
}

// WithCEK allows users to specify a variable to store the CEK used in the
// message upon successful decryption. The variable must be a pointer to
// a byte slice, and it will only be populated if the decryption is successful.
//
// This option is currently considered EXPERIMENTAL, and is subject to
// future changes across minor/micro versions.
func WithCEK(v *[]byte) DecryptOption {
	return &decryptOption{option.New(identCEK{}, v)}
}

// WithCompress specifies the compression algorithm to use when encrypting
// a payload using `jwe.Encrypt` (Yes, we know it can only be "" or "DEF",
// but the way the specification is written it could allow for more options,
// and therefore this option takes an argument)
func WithCompress(v jwa.CompressionAlgorithm) EncryptOption {
	return &encryptOption{option.New(identCompress{}, v)}
}

// WithContentEncryptionAlgorithm specifies the algorithm to encrypt the
// JWE message content with. If not provided, `jwa.A256GCM` is used.
func WithContentEncryption(v jwa.ContentEncryptionAlgorithm) EncryptOption {
	return &encryptOption{option.New(identContentEncryptionAlgorithm{}, v)}
}

// WithFS specifies the source `fs.FS` object to read the file from.
func WithFS(v fs.FS) ReadFileOption {
	return &readFileOption{option.New(identFS{}, v)}
}

func WithKeyProvider(v KeyProvider) DecryptOption {
	return &decryptOption{option.New(identKeyProvider{}, v)}
}

// WithKeyUsed allows you to specify the `jwe.Decrypt()` function to
// return the key used for decryption. This may be useful when
// you specify multiple key sources or if you pass a `jwk.Set`
// and you want to know which key was successful at decrypting the
// signature.
//
// `v` must be a pointer to an empty `interface{}`. Do not use
// `jwk.Key` here unless you are 100% sure that all keys that you
// have provided are instances of `jwk.Key` (remember that the
// jwx API allows users to specify a raw key such as *rsa.PublicKey)
func WithKeyUsed(v interface{}) DecryptOption {
	return &decryptOption{option.New(identKeyUsed{}, v)}
}

// WithMaxBufferSize specifies the maximum buffer size for internal
// calculations, such as when AES-CBC is performed. The default value is 256MB.
// If set to an invalid value, the default value is used.
//
// This option has a global effect.
func WithMaxBufferSize(v int64) GlobalOption {
	return &globalOption{option.New(identMaxBufferSize{}, v)}
}

// WithMaxPBES2Count specifies the maximum number of PBES2 iterations
// to use when decrypting a message. If not specified, the default
// value of 10,000 is used.
//
// This option has a global effect.
func WithMaxPBES2Count(v int) GlobalOption {
	return &globalOption{option.New(identMaxPBES2Count{}, v)}
}

// WithMergeProtectedHeaders specify that when given multiple headers
// as options to `jwe.Encrypt`, these headers should be merged instead
// of overwritten
func WithMergeProtectedHeaders(v bool) EncryptOption {
	return &encryptOption{option.New(identMergeProtectedHeaders{}, v)}
}

// WithMessage provides a message object to be populated by `jwe.Decrpt`
// Using this option allows you to decrypt AND obtain the `jwe.Message`
// in one go.
//
// Note that you should NOT be using the message object for anything other
// than inspecting its contents. Particularly, do not expect the message
// reliable when you call `Decrypt` on it. `(jwe.Message).Decrypt` is
// slated to be deprecated in the next major version.
func WithMessage(v *Message) DecryptOption {
	return &decryptOption{option.New(identMessage{}, v)}
}

// WithPretty specifies whether the JSON output should be formatted and
// indented
func WithPretty(v bool) WithJSONSuboption {
	return &withJSONSuboption{option.New(identPretty{}, v)}
}

// WithrequiredKid specifies whether the keys in the jwk.Set should
// only be matched if the target JWE message's Key ID and the Key ID
// in the given key matches.
func WithRequireKid(v bool) WithKeySetSuboption {
	return &withKeySetSuboption{option.New(identRequireKid{}, v)}
}

// WithCompact specifies that the result of `jwe.Encrypt()` is serialized in
// compact format.
//
// By default `jwe.Encrypt()` will opt to use compact format, so you usually
// do not need to specify this option other than to be explicit about it
func WithCompact() EncryptOption {
	return &encryptOption{option.New(identSerialization{}, fmtCompact)}
}
