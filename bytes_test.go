package sizefmt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {

	t.Run("ByteSize", func(t *testing.T) {
		t.Run("Prints in the largest possible unit", func(t *testing.T) {
			assert.Equal(t, "10T", ByteSize(10*TERABYTE))
			assert.Equal(t, "10.5T", ByteSize(10.5*TERABYTE))
			assert.Equal(t, "10G", ByteSize(10*GIGABYTE))
			assert.Equal(t, "10.5G", ByteSize(10.5*GIGABYTE))
			assert.Equal(t, "100M", ByteSize(100*MEGABYTE))
			assert.Equal(t, "100.5M", ByteSize(100.5*MEGABYTE))
			assert.Equal(t, "100K", ByteSize(100*KILOBYTE))
			assert.Equal(t, "100.5K", ByteSize(100.5*KILOBYTE))
			assert.Equal(t, "1B", ByteSize(1))
		})

		t.Run("prints '0' for zero bytes", func(t *testing.T) {
			assert.Equal(t, "0", ByteSize(0))
		})
	})

	t.Run("ToMegabytes", func(t *testing.T) {
		t.Run("parses byte amounts with short units (e.g. M, G)", func(t *testing.T) {
			var (
				megabytes int64
				err       error
			)

			megabytes, err = ToMegabytes("5B")
			assert.NoError(t, err)
			assert.Equal(t, int64(0), megabytes)

			megabytes, err = ToMegabytes("5K")
			assert.NoError(t, err)
			assert.Equal(t, int64(0), megabytes)

			megabytes, err = ToMegabytes("5M")
			assert.NoError(t, err)
			assert.Equal(t, int64(5), megabytes)

			megabytes, err = ToMegabytes("5m")
			assert.NoError(t, err)
			assert.Equal(t, int64(5), megabytes)

			megabytes, err = ToMegabytes("2G")
			assert.NoError(t, err)
			assert.Equal(t, int64(2*1024), megabytes)

			megabytes, err = ToMegabytes("3T")
			assert.NoError(t, err)
			assert.Equal(t, int64(3*1024*1024), megabytes)
		})

		t.Run("parses byte amounts with long units (e.g MB, GB)", func(t *testing.T) {
			var (
				megabytes int64
				err       error
			)

			megabytes, err = ToMegabytes("5MB")
			assert.NoError(t, err)
			assert.Equal(t, int64(5), megabytes)

			megabytes, err = ToMegabytes("5mb")
			assert.NoError(t, err)
			assert.Equal(t, int64(5), megabytes)

			megabytes, err = ToMegabytes("2GB")
			assert.NoError(t, err)
			assert.Equal(t, int64(2*1024), megabytes)

			megabytes, err = ToMegabytes("3TB")
			assert.NoError(t, err)
			assert.Equal(t, int64(3*1024*1024), megabytes)
		})

		t.Run("returns an error when the unit is missing", func(t *testing.T) {
			_, err := ToMegabytes("5")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})

		t.Run("returns an error when the unit is unrecognized", func(t *testing.T) {
			_, err := ToMegabytes("5MBB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())

			_, err = ToMegabytes("5BB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})

		t.Run("allows whitespace before and after the value", func(t *testing.T) {
			megabytes, err := ToMegabytes("\t\n\r 5MB ")
			assert.NoError(t, err)
			assert.Equal(t, int64(5), megabytes)
		})

		t.Run("returns an error for negative values", func(t *testing.T) {
			_, err := ToMegabytes("-5MB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})

		t.Run("returns an error for zero values", func(t *testing.T) {
			_, err := ToMegabytes("0TB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})
	})

	t.Run("ToBytes", func(t *testing.T) {
		t.Run("parses byte amounts with short units (e.g. M, G)", func(t *testing.T) {
			var (
				bytes int64
				err   error
			)

			bytes, err = ToBytes("5B")
			assert.NoError(t, err)
			assert.Equal(t, int64(5), bytes)

			bytes, err = ToBytes("5K")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*KILOBYTE), bytes)

			bytes, err = ToBytes("5M")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*MEGABYTE), bytes)

			bytes, err = ToBytes("5m")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*MEGABYTE), bytes)

			bytes, err = ToBytes("5G")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*GIGABYTE), bytes)

			bytes, err = ToBytes("5T")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*TERABYTE), bytes)
		})

		t.Run("parses byte amounts that are float (e.g. 5.3KB)", func(t *testing.T) {
			var (
				bytes int64
				err   error
			)

			bytes, err = ToBytes("13.5KB")
			assert.NoError(t, err)
			assert.Equal(t, int64(13824), bytes)

			bytes, err = ToBytes("4.5KB")
			assert.NoError(t, err)
			assert.Equal(t, int64(4608), bytes)
		})

		t.Run("parses byte amounts with long units (e.g MB, GB)", func(t *testing.T) {
			var (
				bytes int64
				err   error
			)

			bytes, err = ToBytes("5MB")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*MEGABYTE), bytes)

			bytes, err = ToBytes("5mb")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*MEGABYTE), bytes)

			bytes, err = ToBytes("5GB")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*GIGABYTE), bytes)

			bytes, err = ToBytes("5TB")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*TERABYTE), bytes)
		})

		t.Run("returns an error when the unit is missing", func(t *testing.T) {
			_, err := ToBytes("5")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})

		t.Run("returns an error when the unit is unrecognized", func(t *testing.T) {
			_, err := ToBytes("5MBB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())

			_, err = ToBytes("5BB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})

		t.Run("allows whitespace before and after the value", func(t *testing.T) {
			bytes, err := ToBytes("\t\n\r 5MB ")
			assert.NoError(t, err)
			assert.Equal(t, int64(5*MEGABYTE), bytes)
		})

		t.Run("returns an error for negative values", func(t *testing.T) {
			_, err := ToBytes("-5MB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})

		t.Run("returns an error for zero values", func(t *testing.T) {
			_, err := ToBytes("0TB")
			assert.Error(t, err)
			assert.EqualError(t, err, errInvalidByteQuantity.Error())
		})
	})
}
