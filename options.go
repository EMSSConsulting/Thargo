package thargo

// Options allows you to configure various parameters relating
// to the way in which Thargo accesses and writes your archives.
type Options struct {
	GZip            bool
	GZipLevel       int
	CreateIfMissing bool
}

// DefaultOptions retrieves the default Thargo configuration which
// can then be modified to suit your needs.
var DefaultOptions = &Options{
	GZip:            true,
	GZipLevel:       9,
	CreateIfMissing: false,
}
