package filter

import (
	"internal/base"
)

// Mask is a mask for a combination of components.
type Mask = base.Mask

// MaskFilter is the interface for logic filters
type MaskFilter interface {
	Matches(mask base.BitMask) bool
}

// All matches all the given components.
//
// Like [And] for combining individual components.
func All(comps ...base.ID) Mask {
	return base.NewMask(comps...)
}

// OneOf matches any of the two components.
//
// Like [Or] for combining individual components.
func OneOf(compA base.ID, compB base.ID) *OR {
	return &OR{base.NewMask(compA), base.NewMask(compB)}
}

// AND is a filter for ANDing together components
type AND struct {
	L MaskFilter
	R MaskFilter
}

// And constructs a pointer to a AND filter
func And(l, r MaskFilter) *AND {
	return &AND{L: l, R: r}
}

// Matches matches a filter against a mask
func (f *AND) Matches(mask base.BitMask) bool {
	return f.L.Matches(mask) && f.R.Matches(mask)
}

// OR is a filter for ORing together components
type OR struct {
	L MaskFilter
	R MaskFilter
}

// Or constructs a pointer to a OR filter
func Or(l, r MaskFilter) *OR {
	return &OR{L: l, R: r}
}

// Matches matches a filter against a mask
func (f *OR) Matches(mask base.BitMask) bool {
	return f.L.Matches(mask) || f.R.Matches(mask)
}

// XOR is a filter for XORing together components
type XOR struct {
	L MaskFilter
	R MaskFilter
}

// XOr constructs a pointer to a XOR filter
func XOr(l, r MaskFilter) *XOR {
	return &XOR{L: l, R: r}
}

// Matches matches a filter against a mask
func (f *XOR) Matches(mask base.BitMask) bool {
	return f.L.Matches(mask) != f.R.Matches(mask)
}

// NOT is a filter for excluding entities with all given components
type NOT Mask

// Not constructs a NOT filter
func Not(comps ...base.ID) NOT {
	return NOT(base.NewMask(comps...))
}

// Matches matches a filter against a mask
func (f NOT) Matches(mask base.BitMask) bool {
	return !mask.Contains(f.BitMask)
}

// NotANY is a filter for excluding entities with any of the the given components
type NotANY Mask

// NotAny constructs a NotANY filter
func NotAny(comps ...base.ID) NotANY {
	return NotANY(base.NewMask(comps...))
}

// Matches matches a filter against a mask
func (f NotANY) Matches(mask base.BitMask) bool {
	return !mask.ContainsAny(f.BitMask)
}
