// Code generated by go generate via cmd/definitions; DO NOT EDIT.
package dropbox

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	. "go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/pkg/httpclient"
	"go.beyondstorage.io/v5/services"
	. "go.beyondstorage.io/v5/types"
)

var (
	_ Storager
	_ services.ServiceError
	_ httpclient.Options
	_ time.Duration
	_ http.Request
	_ Error
)

// Type is the type for dropbox
const Type = "dropbox"

// ObjectSystemMetadata stores system metadata for object.
type ObjectSystemMetadata struct {
	UploadSessionID string
}

// GetObjectSystemMetadata will get ObjectSystemMetadata from Object.
//
// - This function should not be called by service implementer.
// - The returning ObjectServiceMetadata is read only and should not be modified.
func GetObjectSystemMetadata(o *Object) ObjectSystemMetadata {
	sm, ok := o.GetSystemMetadata()
	if ok {
		return sm.(ObjectSystemMetadata)
	}
	return ObjectSystemMetadata{}
}

// setObjectSystemMetadata will set ObjectSystemMetadata into Object.
//
// - This function should only be called once, please make sure all data has been written before set.
func setObjectSystemMetadata(o *Object, sm ObjectSystemMetadata) {
	o.SetSystemMetadata(sm)
}

// StorageSystemMetadata stores system metadata for object.
type StorageSystemMetadata struct {
	UploadSessionID string
}

// GetStorageSystemMetadata will get StorageSystemMetadata from Storage.
//
// - This function should not be called by service implementer.
// - The returning StorageServiceMetadata is read only and should not be modified.
func GetStorageSystemMetadata(s *StorageMeta) StorageSystemMetadata {
	sm, ok := s.GetSystemMetadata()
	if ok {
		return sm.(StorageSystemMetadata)
	}
	return StorageSystemMetadata{}
}

// setStorageSystemMetadata will set StorageSystemMetadata into Storage.
//
// - This function should only be called once, please make sure all data has been written before set.
func setStorageSystemMetadata(s *StorageMeta, sm StorageSystemMetadata) {
	s.SetSystemMetadata(sm)
}

// WithDefaultStoragePairs will apply default_storage_pairs value to Options.
func WithDefaultStoragePairs(v DefaultStoragePairs) Pair {
	return Pair{Key: "default_storage_pairs", Value: v}
}

// WithStorageFeatures will apply storage_features value to Options.
func WithStorageFeatures(v StorageFeatures) Pair {
	return Pair{Key: "storage_features", Value: v}
}

var pairMap = map[string]string{"content_md5": "string", "content_type": "string", "context": "context.Context", "continuation_token": "string", "credential": "string", "default_content_type": "string", "default_io_callback": "func([]byte)", "default_storage_pairs": "DefaultStoragePairs", "endpoint": "string", "expire": "time.Duration", "http_client_options": "*httpclient.Options", "interceptor": "Interceptor", "io_callback": "func([]byte)", "list_mode": "ListMode", "location": "string", "multipart_id": "string", "name": "string", "object_mode": "ObjectMode", "offset": "int64", "size": "int64", "storage_features": "StorageFeatures", "work_dir": "string"}
var (
	_ Appender = &Storage{}
	_ Direr    = &Storage{}
	_ Storager = &Storage{}
)

type StorageFeatures struct {
}

// pairStorageNew is the parsed struct
type pairStorageNew struct {
	pairs []Pair

	// Required pairs
	HasCredential bool
	Credential    string
	// Optional pairs
	HasDefaultContentType  bool
	DefaultContentType     string
	HasDefaultIoCallback   bool
	DefaultIoCallback      func([]byte)
	HasDefaultStoragePairs bool
	DefaultStoragePairs    DefaultStoragePairs
	HasHTTPClientOptions   bool
	HTTPClientOptions      *httpclient.Options
	HasStorageFeatures     bool
	StorageFeatures        StorageFeatures
	HasWorkDir             bool
	WorkDir                string
	// Enable features
}

// parsePairStorageNew will parse Pair slice into *pairStorageNew
func parsePairStorageNew(opts []Pair) (pairStorageNew, error) {
	result :=
		pairStorageNew{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "credential":
			if result.HasCredential {
				continue
			}
			result.HasCredential = true
			result.Credential = v.Value.(string)
		case "default_content_type":
			if result.HasDefaultContentType {
				continue
			}
			result.HasDefaultContentType = true
			result.DefaultContentType = v.Value.(string)
		case "default_io_callback":
			if result.HasDefaultIoCallback {
				continue
			}
			result.HasDefaultIoCallback = true
			result.DefaultIoCallback = v.Value.(func([]byte))
		case "default_storage_pairs":
			if result.HasDefaultStoragePairs {
				continue
			}
			result.HasDefaultStoragePairs = true
			result.DefaultStoragePairs = v.Value.(DefaultStoragePairs)
		case "http_client_options":
			if result.HasHTTPClientOptions {
				continue
			}
			result.HasHTTPClientOptions = true
			result.HTTPClientOptions = v.Value.(*httpclient.Options)
		case "storage_features":
			if result.HasStorageFeatures {
				continue
			}
			result.HasStorageFeatures = true
			result.StorageFeatures = v.Value.(StorageFeatures)
		case "work_dir":
			if result.HasWorkDir {
				continue
			}
			result.HasWorkDir = true
			result.WorkDir = v.Value.(string)
		}
	}
	// Enable features

	// Default pairs
	if result.HasDefaultContentType {
		result.HasDefaultStoragePairs = true
		result.DefaultStoragePairs.Write = append(result.DefaultStoragePairs.Write, WithContentType(result.DefaultContentType))
	}
	if result.HasDefaultIoCallback {
		result.HasDefaultStoragePairs = true
		result.DefaultStoragePairs.Read = append(result.DefaultStoragePairs.Read, WithIoCallback(result.DefaultIoCallback))
		result.DefaultStoragePairs.Write = append(result.DefaultStoragePairs.Write, WithIoCallback(result.DefaultIoCallback))
	}
	if !result.HasCredential {
		return pairStorageNew{}, services.PairRequiredError{Keys: []string{"credential"}}
	}
	return result, nil
}

// DefaultStoragePairs is default pairs for specific action
type DefaultStoragePairs struct {
	CommitAppend []Pair
	Create       []Pair
	CreateAppend []Pair
	CreateDir    []Pair
	Delete       []Pair
	List         []Pair
	Metadata     []Pair
	Read         []Pair
	Stat         []Pair
	Write        []Pair
	WriteAppend  []Pair
}
type pairStorageCommitAppend struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageCommitAppend(opts []Pair) (pairStorageCommitAppend, error) {
	result :=
		pairStorageCommitAppend{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageCommitAppend{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageCreate struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasObjectMode bool
	ObjectMode    ObjectMode
}

func (s *Storage) parsePairStorageCreate(opts []Pair) (pairStorageCreate, error) {
	result :=
		pairStorageCreate{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
		default:
			return pairStorageCreate{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageCreateAppend struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageCreateAppend(opts []Pair) (pairStorageCreateAppend, error) {
	result :=
		pairStorageCreateAppend{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageCreateAppend{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageCreateDir struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageCreateDir(opts []Pair) (pairStorageCreateDir, error) {
	result :=
		pairStorageCreateDir{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageCreateDir{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageDelete struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasObjectMode bool
	ObjectMode    ObjectMode
}

func (s *Storage) parsePairStorageDelete(opts []Pair) (pairStorageDelete, error) {
	result :=
		pairStorageDelete{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
		default:
			return pairStorageDelete{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageList struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasListMode bool
	ListMode    ListMode
}

func (s *Storage) parsePairStorageList(opts []Pair) (pairStorageList, error) {
	result :=
		pairStorageList{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "list_mode":
			if result.HasListMode {
				continue
			}
			result.HasListMode = true
			result.ListMode = v.Value.(ListMode)
		default:
			return pairStorageList{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageMetadata struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageMetadata(opts []Pair) (pairStorageMetadata, error) {
	result :=
		pairStorageMetadata{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageMetadata{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageRead struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasIoCallback bool
	IoCallback    func([]byte)
	HasOffset     bool
	Offset        int64
	HasSize       bool
	Size          int64
}

func (s *Storage) parsePairStorageRead(opts []Pair) (pairStorageRead, error) {
	result :=
		pairStorageRead{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "io_callback":
			if result.HasIoCallback {
				continue
			}
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
		case "offset":
			if result.HasOffset {
				continue
			}
			result.HasOffset = true
			result.Offset = v.Value.(int64)
		case "size":
			if result.HasSize {
				continue
			}
			result.HasSize = true
			result.Size = v.Value.(int64)
		default:
			return pairStorageRead{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageStat struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasObjectMode bool
	ObjectMode    ObjectMode
}

func (s *Storage) parsePairStorageStat(opts []Pair) (pairStorageStat, error) {
	result :=
		pairStorageStat{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
		default:
			return pairStorageStat{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageWrite struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
	HasContentMd5  bool
	ContentMd5     string
	HasContentType bool
	ContentType    string
	HasIoCallback  bool
	IoCallback     func([]byte)
}

func (s *Storage) parsePairStorageWrite(opts []Pair) (pairStorageWrite, error) {
	result :=
		pairStorageWrite{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		case "content_md5":
			if result.HasContentMd5 {
				continue
			}
			result.HasContentMd5 = true
			result.ContentMd5 = v.Value.(string)
		case "content_type":
			if result.HasContentType {
				continue
			}
			result.HasContentType = true
			result.ContentType = v.Value.(string)
		case "io_callback":
			if result.HasIoCallback {
				continue
			}
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
		default:
			return pairStorageWrite{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}

type pairStorageWriteAppend struct {
	pairs []Pair
	// Required pairs
	// Optional pairs
}

func (s *Storage) parsePairStorageWriteAppend(opts []Pair) (pairStorageWriteAppend, error) {
	result :=
		pairStorageWriteAppend{pairs: opts}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageWriteAppend{}, services.PairUnsupportedError{Pair: v}
		}
	}

	return result, nil
}
func (s *Storage) CommitAppend(o *Object, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.CommitAppendWithContext(ctx, o, pairs...)
}
func (s *Storage) CommitAppendWithContext(ctx context.Context, o *Object, pairs ...Pair) (err error) {
	defer func() {
		err =
			s.formatError("commit_append", err)
	}()
	if !o.Mode.IsAppend() {
		err = services.ObjectModeInvalidError{Expected: ModeAppend, Actual: o.Mode}
		return
	}
	pairs = append(pairs, s.defaultPairs.CommitAppend...)
	var opt pairStorageCommitAppend

	opt, err = s.parsePairStorageCommitAppend(pairs)
	if err != nil {
		return
	}
	return s.commitAppend(ctx, o, opt)
}
func (s *Storage) Create(path string, pairs ...Pair) (o *Object) {
	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairStorageCreate

	// Ignore error while handling local functions.
	opt, _ = s.parsePairStorageCreate(pairs)
	return s.create(path, opt)
}
func (s *Storage) CreateAppend(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.CreateAppendWithContext(ctx, path, pairs...)
}
func (s *Storage) CreateAppendWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err =
			s.formatError("create_append", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.CreateAppend...)
	var opt pairStorageCreateAppend

	opt, err = s.parsePairStorageCreateAppend(pairs)
	if err != nil {
		return
	}
	return s.createAppend(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) CreateDir(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.CreateDirWithContext(ctx, path, pairs...)
}
func (s *Storage) CreateDirWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err =
			s.formatError("create_dir", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.CreateDir...)
	var opt pairStorageCreateDir

	opt, err = s.parsePairStorageCreateDir(pairs)
	if err != nil {
		return
	}
	return s.createDir(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) Delete(path string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, path, pairs...)
}
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...Pair) (err error) {
	defer func() {
		err =
			s.formatError("delete", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairStorageDelete

	opt, err = s.parsePairStorageDelete(pairs)
	if err != nil {
		return
	}
	return s.delete(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) List(path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, path, pairs...)
}
func (s *Storage) ListWithContext(ctx context.Context, path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	defer func() {
		err =
			s.formatError("list", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairStorageList

	opt, err = s.parsePairStorageList(pairs)
	if err != nil {
		return
	}
	return s.list(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) Metadata(pairs ...Pair) (meta *StorageMeta) {
	pairs = append(pairs, s.defaultPairs.Metadata...)
	var opt pairStorageMetadata

	// Ignore error while handling local functions.
	opt, _ = s.parsePairStorageMetadata(pairs)
	return s.metadata(opt)
}
func (s *Storage) Read(path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.ReadWithContext(ctx, path, w, pairs...)
}
func (s *Storage) ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	defer func() {
		err =
			s.formatError("read", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Read...)
	var opt pairStorageRead

	opt, err = s.parsePairStorageRead(pairs)
	if err != nil {
		return
	}
	return s.read(ctx, strings.ReplaceAll(path, "\\", "/"), w, opt)
}
func (s *Storage) Stat(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.StatWithContext(ctx, path, pairs...)
}
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err =
			s.formatError("stat", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Stat...)
	var opt pairStorageStat

	opt, err = s.parsePairStorageStat(pairs)
	if err != nil {
		return
	}
	return s.stat(ctx, strings.ReplaceAll(path, "\\", "/"), opt)
}
func (s *Storage) Write(path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.WriteWithContext(ctx, path, r, size, pairs...)
}
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	defer func() {
		err =
			s.formatError("write", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Write...)
	var opt pairStorageWrite

	opt, err = s.parsePairStorageWrite(pairs)
	if err != nil {
		return
	}
	return s.write(ctx, strings.ReplaceAll(path, "\\", "/"), r, size, opt)
}
func (s *Storage) WriteAppend(o *Object, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.WriteAppendWithContext(ctx, o, r, size, pairs...)
}
func (s *Storage) WriteAppendWithContext(ctx context.Context, o *Object, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	defer func() {
		err =
			s.formatError("write_append", err)
	}()
	if !o.Mode.IsAppend() {
		err = services.ObjectModeInvalidError{Expected: ModeAppend, Actual: o.Mode}
		return
	}
	pairs = append(pairs, s.defaultPairs.WriteAppend...)
	var opt pairStorageWriteAppend

	opt, err = s.parsePairStorageWriteAppend(pairs)
	if err != nil {
		return
	}
	return s.writeAppend(ctx, o, r, size, opt)
}
func init() {
	services.RegisterStorager(Type, NewStorager)
	services.RegisterSchema(Type, pairMap)
}
