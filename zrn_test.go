package zrn_test

import (
	"testing"

	"github.com/zeiss/zrn"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		zrn.New("zrn", "vision", "microscopy", "de", "foo", "bar")
	}
}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		zrn.Parse("zrn:vision:microscopy:de:foo:bar")
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		desc        string
		namespace   zrn.Match
		partition   zrn.Match
		product     zrn.Match
		region      zrn.Match
		identifier  zrn.Match
		resource    zrn.Match
		expected    *zrn.ZRN
		expectedErr bool
	}{
		{
			desc:       "valid ZRN",
			namespace:  "zrn",
			partition:  "vision",
			product:    "microscopy",
			region:     "de",
			identifier: "foo",
			resource:   "bar",
			expected: &zrn.ZRN{
				Namespace:  "zrn",
				Partition:  "vision",
				Product:    "microscopy",
				Region:     "de",
				Identifier: "foo",
				Resource:   "bar",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			zrn, err := zrn.New(test.namespace, test.partition, test.product, test.region, test.identifier, test.resource)

			if test.expectedErr {
				require.Error(t, err)
			}

			assert.Equal(t, test.expected, zrn)
		})
	}
}

func TestMust(t *testing.T) {
	tests := []struct {
		desc        string
		namespace   zrn.Match
		partition   zrn.Match
		product     zrn.Match
		region      zrn.Match
		identifier  zrn.Match
		resource    zrn.Match
		expected    *zrn.ZRN
		expectedErr bool
	}{
		{
			desc:       "invalid ZRN",
			namespace:  "",
			partition:  "vision",
			product:    "microscopy",
			region:     "de",
			identifier: "foo",
			resource:   "bar",
			expected:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			assert.Panics(t, func() {
				zrn.Must(test.namespace, test.partition, test.product, test.region, test.identifier, test.resource)
			})
		})
	}
}
