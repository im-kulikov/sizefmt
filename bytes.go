package sizefmt

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	_ = 1.0 << (10 * iota) // ignore first value by assigning to blank identifier
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

var (
	bytesPattern           = regexp.MustCompile(`(?i)^(-?\d+(?:\.\d+)?)([KMGT]B?|B)$`)
	errInvalidByteQuantity = errors.New("byte quantity must be a positive integer with a unit of measurement like M, MB, G, or GB")
)

// ByteSize returns a human-readable byte string of the form 10M, 12.5K, and so forth.  The following units are available:
//	T: Terabyte
//	G: Gigabyte
//	M: Megabyte
//	K: Kilobyte
//	B: Byte
// The unit that results in the smallest number greater than or equal to 1 is always chosen.
func ByteSize(b float64) string {
	var (
		unit string
		del  float64 = 1
	)

	switch {
	case b >= YB:
		unit = "Y"
		del = YB
	case b >= ZB:
		unit = "Z"
		del = ZB
	case b >= EB:
		unit = "E"
		del = EB
	case b >= PB:
		unit = "P"
		del = PB
	case b >= TB:
		unit = "T"
		del = TB
	case b >= GB:
		unit = "G"
		del = GB
	case b >= MB:
		unit = "M"
		del = MB
	case b >= KB:
		unit = "K"
		del = KB
	case b == 0:
		return "0"
	default:
		unit = "B"
	}
	return strings.TrimSuffix(
		strconv.FormatFloat(b/del, 'f', 1, 32),
		".0",
	) + unit
}

// ToMegabytes parses a string formatted by ByteSize as megabytes.
func ToMegabytes(s string) (int64, error) {
	bytes, err := ToBytes(s)
	if err != nil {
		return 0, err
	}
	return bytes / MB, nil
}

// ToBytes parses a string formatted by ByteSize as bytes.
func ToBytes(s string) (int64, error) {
	parts := bytesPattern.FindStringSubmatch(strings.TrimSpace(s))
	if len(parts) < 3 {
		return 0, errInvalidByteQuantity
	}

	value, err := strconv.ParseFloat(parts[1], 64)
	if err != nil || value <= 0 {
		return 0, errInvalidByteQuantity
	}

	var bytes int64
	unit := strings.ToUpper(parts[2])
	switch unit[:1] {
	case "T":
		bytes = int64(value * TB)
	case "G":
		bytes = int64(value * GB)
	case "M":
		bytes = int64(value * MB)
	case "K":
		bytes = int64(value * KB)
	case "B":
		bytes = int64(value)
	}

	return bytes, nil
}
