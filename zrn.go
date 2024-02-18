package zrn

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	// DefaultSeparator is the default separator used to join the URN segments.
	DefaultSeparator = ":"
	// DefaultNamespace is the default namespace used for ZEISS URNs.
	DefaultNamespace = "zeiss"
)

// ErrorInvalid is returned when parsing an URN with an invalid format.
var ErrorInvalid = errors.New("zrn: invalid format")

// use a single instance of Validate, it caches struct info.
var validate = validator.New()

// Match is a string that can be used to match a URN segment.
type Match string

var (
	// Wildcard is the wildcard used to match any value.
	Wildcard Match = "*"
	// Empty is the empty string
	Empty Match
)

// String returns the string representation of the match.
func (m Match) String() string {
	return string(m)
}

// ZRN is a Uniform Resource Name (URN) as defined in RFC 2141.
type ZRN struct {
	// Namespace is the URN namespace.
	Namespace Match `json:"namespace" validate:"required"` // e.g. "zeiss"
	// Partition is the URN partition.
	Partition Match `json:"partition" validate:"max=256"` // e.g. "smt"
	// Product/Service is the URN product or service.
	Product Match `json:"product" validate:"max=256"` // e.g. "microscope"
	// Region is the URN region.
	Region Match `json:"region" validate:"max=256"` // e.g. "us"
	// Identifier is the URN identifier.
	Identifier Match `json:"identifier" validate:"max=64"` // e.g. "1234"
	// Resource is the URN resource.
	Resource Match `json:"resource" validate:"required,max=256"`
}

// String returns the string representation of the ZRN.
func (z *ZRN) String() string {
	return strings.Join([]string{z.Namespace.String(), z.Partition.String(), z.Product.String(), z.Region.String(), z.Identifier.String(), z.Resource.String()}, DefaultSeparator)
}

// Match returns true if the ZRN matches the given pattern.
//
// nolint: gocyclo
func (z *ZRN) Match(zrn *ZRN) bool {
	return (z.Namespace == zrn.Namespace || (z.Namespace == Wildcard && zrn.Namespace == Wildcard) || (z.Namespace == Empty && zrn.Namespace == Empty) || zrn.Namespace == Wildcard || zrn.Namespace == Empty) &&
		(z.Partition == zrn.Partition || (z.Partition == Wildcard && zrn.Partition == Wildcard) || (z.Partition == Empty && zrn.Partition == Empty) || zrn.Partition == Wildcard || zrn.Partition == Empty) &&
		(z.Product == zrn.Product || (z.Product == Wildcard && zrn.Product == Wildcard) || (z.Product == Empty && zrn.Product == Empty) || zrn.Product == Wildcard || zrn.Product == Empty) &&
		(z.Region == zrn.Region || (z.Region == Wildcard && zrn.Region == Wildcard) || (z.Region == Empty && zrn.Region == Empty) || zrn.Region == Wildcard || zrn.Region == Empty) &&
		(z.Identifier == zrn.Identifier || (z.Identifier == Wildcard && zrn.Identifier == Wildcard) || (z.Identifier == Empty && zrn.Identifier == Empty) || zrn.Identifier == Wildcard || zrn.Identifier == Empty) &&
		(z.Resource == zrn.Resource || (z.Resource == Wildcard && zrn.Resource == Wildcard) || (z.Resource == Empty && zrn.Resource == Empty) || zrn.Resource == Wildcard || zrn.Resource == Empty)
}

// ExactMatch returns true if the ZRN matches the given pattern exactly.
func (z *ZRN) ExactMatch(zrn *ZRN) bool {
	return z.Namespace == zrn.Namespace &&
		z.Partition == zrn.Partition &&
		z.Product == zrn.Product &&
		z.Region == zrn.Region &&
		z.Identifier == zrn.Identifier &&
		z.Resource == zrn.Resource
}

// New returns a new ZRN.
func New(namespace, partition, product, region, identifier, resource Match) (*ZRN, error) {
	zrn := &ZRN{
		Namespace:  namespace,
		Partition:  partition,
		Product:    product,
		Region:     region,
		Identifier: identifier,
		Resource:   resource,
	}

	validate = validator.New()

	if err := validate.Struct(zrn); err != nil {
		return nil, err
	}

	return zrn, nil
}

// Must returns a new ZRN. It panics if the ZRN is invalid.
func Must(namespace, partition, product, region, identifier, resource Match) *ZRN {
	zrn, err := New(namespace, partition, product, region, identifier, resource)
	if err != nil {
		panic(err)
	}

	return zrn
}

// Parse returns a new ZRN from the given string.
func Parse(s string) (*ZRN, error) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	if Match(s) == Wildcard {
		return &ZRN{
			Namespace:  Wildcard,
			Partition:  Wildcard,
			Product:    Wildcard,
			Region:     Wildcard,
			Identifier: Wildcard,
			Resource:   Wildcard,
		}, nil
	}

	segments := strings.SplitN(s, DefaultSeparator, 6)
	if len(segments) < 5 {
		return nil, ErrorInvalid
	}

	mm := make([]Match, len(segments))
	for i, segment := range segments {
		mm[i] = Match(segment)
		if Match(segment) == Wildcard || Match(segment) == Empty {
			mm[i] = Wildcard
		}
	}

	urn := &ZRN{
		Namespace:  mm[0],
		Partition:  mm[1],
		Product:    mm[2],
		Region:     mm[3],
		Identifier: mm[4],
		Resource:   mm[5],
	}

	validate = validator.New()

	if err := validate.Struct(urn); err != nil {
		return nil, err
	}

	return urn, nil
}

// MustParse returns a new ZRN from the given string. It panics if the string is invalid.
func MustParse(s string) *ZRN {
	urn, err := Parse(s)
	if err != nil {
		panic(err)
	}

	return urn
}
