package generic

// Code generated by go generate; DO NOT EDIT.

import (
	"github.com/mlange-42/arche/ecs"
)

//////////////////////////////////////////////////////////////////////////

// Filter0 is a helper for building [Query0] query iterators.
type Filter0 filter

// NewFilter0 creates a generic Filter0 for zero components.
//
// See also [ecs.World.Query].
func NewFilter0() *Filter0 {
	f := Filter0(filter{
		include: []Comp{},
	})
	return &f
}

// With adds components that are required, but not accessible via [Query0.Get].
//
// Create the required mask items with [T].
func (q *Filter0) With(mask ...Comp) *Filter0 {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter0) Without(mask ...Comp) *Filter0 {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query0] query for iteration.
func (q *Filter0) Query(w *ecs.World) Query0 {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query0(query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter0.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter0.Query].
func (q *Filter0) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query0 is a generic query iterator for zero components.
//
// Create it with [NewFilter0] and [Filter0.Query]
type Query0 query

//////////////////////////////////////////////////////////////////////////

// Filter1 is a helper for building [Query1] query iterators.
type Filter1[A any] filter

// NewFilter1 creates a generic Filter1 for one components.
//
// See also [ecs.World.Query].
func NewFilter1[A any]() *Filter1[A] {
	f := Filter1[A](filter{
		include: []Comp{typeOf[A]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter1[A]) Optional(mask ...Comp) *Filter1[A] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query1.Get].
//
// Create the required mask items with [T].
func (q *Filter1[A]) With(mask ...Comp) *Filter1[A] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter1[A]) Without(mask ...Comp) *Filter1[A] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query1] query for iteration.
func (q *Filter1[A]) Query(w *ecs.World) Query1[A] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query1[A](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter1.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter1.Query].
func (q *Filter1[A]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query1 is a generic query iterator for one components.
//
// Create it with [NewFilter1] and [Filter1.Query]
type Query1[A any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query1.Entity] to get the current Entity.
func (q *Query1[A]) Get() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

//////////////////////////////////////////////////////////////////////////

// Filter2 is a helper for building [Query2] query iterators.
type Filter2[A any, B any] filter

// NewFilter2 creates a generic Filter2 for two components.
//
// See also [ecs.World.Query].
func NewFilter2[A any, B any]() *Filter2[A, B] {
	f := Filter2[A, B](filter{
		include: []Comp{typeOf[A](), typeOf[B]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter2[A, B]) Optional(mask ...Comp) *Filter2[A, B] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query2.Get].
//
// Create the required mask items with [T].
func (q *Filter2[A, B]) With(mask ...Comp) *Filter2[A, B] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter2[A, B]) Without(mask ...Comp) *Filter2[A, B] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query2] query for iteration.
func (q *Filter2[A, B]) Query(w *ecs.World) Query2[A, B] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query2[A, B](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter2.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter2.Query].
func (q *Filter2[A, B]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query2 is a generic query iterator for two components.
//
// Create it with [NewFilter2] and [Filter2.Query]
type Query2[A any, B any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query2.Entity] to get the current Entity.
func (q *Query2[A, B]) Get() (*A, *B) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1]))
}

//////////////////////////////////////////////////////////////////////////

// Filter3 is a helper for building [Query3] query iterators.
type Filter3[A any, B any, C any] filter

// NewFilter3 creates a generic Filter3 for three components.
//
// See also [ecs.World.Query].
func NewFilter3[A any, B any, C any]() *Filter3[A, B, C] {
	f := Filter3[A, B, C](filter{
		include: []Comp{typeOf[A](), typeOf[B](), typeOf[C]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter3[A, B, C]) Optional(mask ...Comp) *Filter3[A, B, C] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query3.Get].
//
// Create the required mask items with [T].
func (q *Filter3[A, B, C]) With(mask ...Comp) *Filter3[A, B, C] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter3[A, B, C]) Without(mask ...Comp) *Filter3[A, B, C] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query3] query for iteration.
func (q *Filter3[A, B, C]) Query(w *ecs.World) Query3[A, B, C] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query3[A, B, C](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter3.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter3.Query].
func (q *Filter3[A, B, C]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query3 is a generic query iterator for three components.
//
// Create it with [NewFilter3] and [Filter3.Query]
type Query3[A any, B any, C any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query3.Entity] to get the current Entity.
func (q *Query3[A, B, C]) Get() (*A, *B, *C) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2]))
}

//////////////////////////////////////////////////////////////////////////

// Filter4 is a helper for building [Query4] query iterators.
type Filter4[A any, B any, C any, D any] filter

// NewFilter4 creates a generic Filter4 for four components.
//
// See also [ecs.World.Query].
func NewFilter4[A any, B any, C any, D any]() *Filter4[A, B, C, D] {
	f := Filter4[A, B, C, D](filter{
		include: []Comp{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter4[A, B, C, D]) Optional(mask ...Comp) *Filter4[A, B, C, D] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query4.Get].
//
// Create the required mask items with [T].
func (q *Filter4[A, B, C, D]) With(mask ...Comp) *Filter4[A, B, C, D] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter4[A, B, C, D]) Without(mask ...Comp) *Filter4[A, B, C, D] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query4] query for iteration.
func (q *Filter4[A, B, C, D]) Query(w *ecs.World) Query4[A, B, C, D] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query4[A, B, C, D](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter4.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter4.Query].
func (q *Filter4[A, B, C, D]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query4 is a generic query iterator for four components.
//
// Create it with [NewFilter4] and [Filter4.Query]
type Query4[A any, B any, C any, D any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query4.Entity] to get the current Entity.
func (q *Query4[A, B, C, D]) Get() (*A, *B, *C, *D) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3]))
}

//////////////////////////////////////////////////////////////////////////

// Filter5 is a helper for building [Query5] query iterators.
type Filter5[A any, B any, C any, D any, E any] filter

// NewFilter5 creates a generic Filter5 for five components.
//
// See also [ecs.World.Query].
func NewFilter5[A any, B any, C any, D any, E any]() *Filter5[A, B, C, D, E] {
	f := Filter5[A, B, C, D, E](filter{
		include: []Comp{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter5[A, B, C, D, E]) Optional(mask ...Comp) *Filter5[A, B, C, D, E] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query5.Get].
//
// Create the required mask items with [T].
func (q *Filter5[A, B, C, D, E]) With(mask ...Comp) *Filter5[A, B, C, D, E] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter5[A, B, C, D, E]) Without(mask ...Comp) *Filter5[A, B, C, D, E] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query5] query for iteration.
func (q *Filter5[A, B, C, D, E]) Query(w *ecs.World) Query5[A, B, C, D, E] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query5[A, B, C, D, E](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter5.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter5.Query].
func (q *Filter5[A, B, C, D, E]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query5 is a generic query iterator for five components.
//
// Create it with [NewFilter5] and [Filter5.Query]
type Query5[A any, B any, C any, D any, E any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query5.Entity] to get the current Entity.
func (q *Query5[A, B, C, D, E]) Get() (*A, *B, *C, *D, *E) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4]))
}

//////////////////////////////////////////////////////////////////////////

// Filter6 is a helper for building [Query6] query iterators.
type Filter6[A any, B any, C any, D any, E any, F any] filter

// NewFilter6 creates a generic Filter6 for six components.
//
// See also [ecs.World.Query].
func NewFilter6[A any, B any, C any, D any, E any, F any]() *Filter6[A, B, C, D, E, F] {
	f := Filter6[A, B, C, D, E, F](filter{
		include: []Comp{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E](), typeOf[F]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter6[A, B, C, D, E, F]) Optional(mask ...Comp) *Filter6[A, B, C, D, E, F] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query6.Get].
//
// Create the required mask items with [T].
func (q *Filter6[A, B, C, D, E, F]) With(mask ...Comp) *Filter6[A, B, C, D, E, F] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter6[A, B, C, D, E, F]) Without(mask ...Comp) *Filter6[A, B, C, D, E, F] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query6] query for iteration.
func (q *Filter6[A, B, C, D, E, F]) Query(w *ecs.World) Query6[A, B, C, D, E, F] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query6[A, B, C, D, E, F](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter6.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter6.Query].
func (q *Filter6[A, B, C, D, E, F]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query6 is a generic query iterator for six components.
//
// Create it with [NewFilter6] and [Filter6.Query]
type Query6[A any, B any, C any, D any, E any, F any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query6.Entity] to get the current Entity.
func (q *Query6[A, B, C, D, E, F]) Get() (*A, *B, *C, *D, *E, *F) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4])), (*F)(q.Query.Get(q.ids[5]))
}

//////////////////////////////////////////////////////////////////////////

// Filter7 is a helper for building [Query7] query iterators.
type Filter7[A any, B any, C any, D any, E any, F any, G any] filter

// NewFilter7 creates a generic Filter7 for seven components.
//
// See also [ecs.World.Query].
func NewFilter7[A any, B any, C any, D any, E any, F any, G any]() *Filter7[A, B, C, D, E, F, G] {
	f := Filter7[A, B, C, D, E, F, G](filter{
		include: []Comp{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E](), typeOf[F](), typeOf[G]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter7[A, B, C, D, E, F, G]) Optional(mask ...Comp) *Filter7[A, B, C, D, E, F, G] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query7.Get].
//
// Create the required mask items with [T].
func (q *Filter7[A, B, C, D, E, F, G]) With(mask ...Comp) *Filter7[A, B, C, D, E, F, G] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter7[A, B, C, D, E, F, G]) Without(mask ...Comp) *Filter7[A, B, C, D, E, F, G] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query7] query for iteration.
func (q *Filter7[A, B, C, D, E, F, G]) Query(w *ecs.World) Query7[A, B, C, D, E, F, G] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query7[A, B, C, D, E, F, G](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter7.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter7.Query].
func (q *Filter7[A, B, C, D, E, F, G]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query7 is a generic query iterator for seven components.
//
// Create it with [NewFilter7] and [Filter7.Query]
type Query7[A any, B any, C any, D any, E any, F any, G any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query7.Entity] to get the current Entity.
func (q *Query7[A, B, C, D, E, F, G]) Get() (*A, *B, *C, *D, *E, *F, *G) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4])), (*F)(q.Query.Get(q.ids[5])), (*G)(q.Query.Get(q.ids[6]))
}

//////////////////////////////////////////////////////////////////////////

// Filter8 is a helper for building [Query8] query iterators.
type Filter8[A any, B any, C any, D any, E any, F any, G any, H any] filter

// NewFilter8 creates a generic Filter8 for eight components.
//
// See also [ecs.World.Query].
func NewFilter8[A any, B any, C any, D any, E any, F any, G any, H any]() *Filter8[A, B, C, D, E, F, G, H] {
	f := Filter8[A, B, C, D, E, F, G, H](filter{
		include: []Comp{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E](), typeOf[F](), typeOf[G](), typeOf[H]()},
	})
	return &f
}

// Optional makes some of the query's components optional.
//
// Create the required mask items with [T].
//
// Only affects component types that were specified in the query.
func (q *Filter8[A, B, C, D, E, F, G, H]) Optional(mask ...Comp) *Filter8[A, B, C, D, E, F, G, H] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds components that are required, but not accessible via [Query8.Get].
//
// Create the required mask items with [T].
func (q *Filter8[A, B, C, D, E, F, G, H]) With(mask ...Comp) *Filter8[A, B, C, D, E, F, G, H] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask items with [T].
func (q *Filter8[A, B, C, D, E, F, G, H]) Without(mask ...Comp) *Filter8[A, B, C, D, E, F, G, H] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query8] query for iteration.
func (q *Filter8[A, B, C, D, E, F, G, H]) Query(w *ecs.World) Query8[A, B, C, D, E, F, G, H] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query8[A, B, C, D, E, F, G, H](query{
		w.Query(&q.compiled.filter),
		q.compiled.Ids,
	})
}

// Filter generates and return the [ecs.Filter] used after [Filter8.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter8.Query].
func (q *Filter8[A, B, C, D, E, F, G, H]) Filter(w *ecs.World) ecs.MaskFilter {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.filter
}

// Query8 is a generic query iterator for eight components.
//
// Create it with [NewFilter8] and [Filter8.Query]
type Query8[A any, B any, C any, D any, E any, F any, G any, H any] query

// Get returns all queried components for the current query iterator position.
//
// Use [Query8.Entity] to get the current Entity.
func (q *Query8[A, B, C, D, E, F, G, H]) Get() (*A, *B, *C, *D, *E, *F, *G, *H) {
	return (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4])), (*F)(q.Query.Get(q.ids[5])), (*G)(q.Query.Get(q.ids[6])), (*H)(q.Query.Get(q.ids[7]))
}
