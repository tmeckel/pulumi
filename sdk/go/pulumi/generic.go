package pulumi

import (
	"context"
	"fmt"
	"reflect"
)

func typeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

type InputT[T any] interface {
	Input

	ToOutputT() OutputT[T]
	// TODO: Can we enforce that T is assignable to ElementType?
}

type OutputT[T any] struct {
	*OutputState // TODO: hide me
}

func (OutputT[T]) isInputT() {} // TODO: hack; fix

func T[T any](v T) OutputT[T] {
	state := newOutputState(nil /* joinGroup */, typeOf[T]())
	state.resolve(v, true, false, nil /* deps */)
	return OutputT[T]{OutputState: state}
}

func Cast[T any](o Output) OutputT[T] {
	typ := typeOf[T]()
	if o.ElementType().AssignableTo(typ) {
		return OutputT[T]{o.getState()}
	}

	// TODO: should this return an error instead?
	// With a MustCast[T] function that panics?
	state := newOutputState(nil /* joinGroup */, typ)
	state.reject(fmt.Errorf("cannot cast %v to %v", o.ElementType(), typ))
	return OutputT[T]{state}
}

var (
	_ Output      = OutputT[any]{}
	_ Input       = OutputT[any]{}
	_ InputT[int] = OutputT[int]{}
)

func (o OutputT[T]) ElementType() reflect.Type {
	return typeOf[T]()
}

func (o OutputT[T]) ToOutputT() OutputT[T] {
	return o
}

func (o OutputT[T]) ToAnyOutput() AnyOutput {
	return AnyOutput(o)
}

// awaitT is a type-safe variant of OutputState.await.
func awaitT[I InputT[T], T any](ctx context.Context, o InputT[T]) (v T, known, secret bool, deps []Resource, err error) {
	var value any
	// TODO: make OutputState type-safe internally.
	value, known, secret, deps, err = ToOutput(o).getState().await(ctx)
	if err == nil {
		// TODO: should this turn into an error?
		var ok bool
		v, ok = value.(T)
		if !ok && value != nil {
			err = fmt.Errorf("awaited value of type %T but got %T", v, value)
		}
	}
	return v, known, secret, deps, err
}

type ArrayInputT[T any] interface {
	InputT[[]T]
}

type ArrayOutputT[T any] struct{ *OutputState }

var (
	_ Output           = ArrayOutputT[any]{}
	_ Input            = ArrayOutputT[any]{}
	_ InputT[[]int]    = ArrayOutputT[int]{}
	_ ArrayInputT[int] = ArrayOutputT[int]{}
)

func ArrayOf[T any](items ...InputT[T]) ArrayOutputT[T] {
	return ArrayOutputT[T]{
		// TODO: is this right?
		OutputState: ToOutput(items).getState(),
	}
}

func ArrayFrom[T any](items InputT[[]T]) ArrayOutputT[T] {
	return ArrayOutputT[T]{
		// TODO: is this right?
		OutputState: ToOutput(items).getState(),
	}
}

func (ArrayOutputT[T]) isInputT() {} // TODO: hack; fix

func (ArrayOutputT[T]) ElementType() reflect.Type {
	return reflect.SliceOf(typeOf[T]())
}

func (o ArrayOutputT[T]) ToOutputT() OutputT[[]T] {
	return OutputT[[]T](o)
}

func (o ArrayOutputT[T]) Index(i InputT[int]) OutputT[T] {
	return ApplyT2(o, i, func(items []T, idx int) T {
		if idx < 0 || idx >= len(items) {
			var zero T
			return zero
		}
		return items[idx]
	})
}

type PtrOutputT[T any] struct{ *OutputState }

type PtrInputT[T any] interface {
	InputT[*T]
}

var (
	_ Output         = PtrOutputT[any]{}
	_ Input          = PtrOutputT[any]{}
	_ InputT[*int]   = PtrOutputT[int]{}
	_ PtrInputT[int] = PtrOutputT[int]{}
)

func Ptr[T any](v T) PtrOutputT[T] {
	os := newOutputState(nil /* joinGroup */, typeOf[T]())
	os.resolve(&v, true, false, nil /* deps */)
	return PtrOutputT[T]{OutputState: os}
}

// TODO: Should this take an InputT
func PtrOf[T any](o InputT[T]) PtrOutputT[T] {
	// As of Go 1.20, Output[T] cannot refer to Output[*T] directly.
	// This refers to the following limitation at
	// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#generic-types:
	//
	//  A generic type can refer to itself in cases
	//  where a type can ordinarily refer to itself,
	//  but when it does so the type arguments must be the type parameters,
	//  listed in the same order.
	//  This restriction prevents infinite recursion of type instantiation.
	//
	// In other words, Output[T]'s methods can only refer to Output[T],
	// or concrete instances of Output[T] (e.g. Output[int]).
	// They cannot refer to variants of T (e.g. Output[*T] or Output[[]T]).
	// Doing so will result in a compile error.
	//
	//   func (o Output[T]) Ptr() Output[*T] {
	//     var _ = OutputT[*T]{}
	//     // error: instantiation cycle
	//
	// This restriction applies for both, direct and indirect references
	// so a method on Output[T] also cannot call a function
	// that instantiates Output[*T].
	//
	//   func (o Output[T]) Ptr() Output[*T] {
	//     return PtrOf(o)
	//     // error: instantiation cycle
	//
	// This restriction may be lifted in the future,
	// but meanwhile it means that PtrOf must be a top-level function.
	p := ApplyT(o, func(v T) *T { return &v })
	return PtrFrom(p)
}

func PtrFrom[T any](o OutputT[*T]) PtrOutputT[T] {
	// No need to check if o.ElementType() is assignable.
	// It's already a pointer type.
	return PtrOutputT[T]{o.getState()}
}

func (PtrOutputT[T]) isInputT() {} // TODO: hack; fix

func (PtrOutputT[T]) ElementType() reflect.Type {
	return reflect.PtrTo(typeOf[T]())
}

func (o PtrOutputT[T]) ToOutputT() OutputT[*T] {
	return OutputT[*T](o)
}

func (o PtrOutputT[T]) Elem() OutputT[T] {
	return ApplyT(o, func(v *T) T {
		if v == nil {
			var zero T
			return zero
		}
		return *v
	})
}

type MapOutputT[V any] struct{ *OutputState }

type MapInputT[V any] interface {
	InputT[map[string]V]
}

var (
	_ Output                 = MapOutputT[any]{}
	_ Input                  = MapOutputT[any]{}
	_ InputT[map[string]int] = MapOutputT[int]{}
	_ MapInputT[any]         = MapOutputT[any]{}
)

func MapOf[V any](items map[string]InputT[V]) MapOutputT[V] {
	return MapOutputT[V]{
		ToOutput(items).getState(),
	}
}

func MapFrom[V any](items InputT[map[string]V]) MapOutputT[V] {
	return MapOutputT[V]{
		ToOutput(items).getState(),
	}
}

func (MapOutputT[V]) isInputT() {} // TODO: hack; fix

func (MapOutputT[V]) ElementType() reflect.Type {
	return reflect.MapOf(typeOf[string](), typeOf[V]())
}

func (o MapOutputT[V]) ToOutputT() OutputT[map[string]V] {
	return OutputT[map[string]V](o)
}

func (o MapOutputT[V]) MapIndex(i InputT[string]) OutputT[V] {
	return ApplyT2(o, i, func(items map[string]V, idx string) V {
		return items[idx]
	})
}

// TODO: Should we parameterize the applier so context, etc. can be optionally
// passed in?
func ApplyT[O any, I InputT[T], T any](o I, f func(T) O) OutputT[O] {
	state := newOutputState(nil, typeOf[O]())
	go func() {
		v, known, secret, deps, err := awaitT[I, T](context.Background(), o)
		if err != nil || !known {
			var zero O
			state.fulfill(zero, known, secret, deps, err)
			return
		}
		state.fulfill(f(v), known, secret, deps, err)
	}()
	return OutputT[O]{state}
}

func ApplyT2[O any, I1 InputT[T1], I2 InputT[T2], T1, T2 any](o1 I1, o2 I2, f func(T1, T2) O) OutputT[O] {
	// TODO: context
	state := newOutputState(nil, typeOf[O]())
	go func() {
		v1, known, secret, deps, err := awaitT[I1, T1](context.Background(), o1)
		if err != nil || !known {
			var zero O
			state.fulfill(zero, known, secret, deps, err)
			return
		}

		v2, known2, secret2, deps2, err := awaitT[I2, T2](context.Background(), o2)
		known = known && known2
		secret = secret || secret2
		deps = mergeDependencies(deps, deps2)
		if err != nil || !known {
			var zero O
			state.fulfill(zero, known, secret, deps, err)
			return
		}

		state.fulfill(f(v1, v2), known, secret, deps, nil)
	}()
	return OutputT[O]{state}
}

/*
Apply2(a, b, f) -> Output[U]
Apply3(a, b, c, f) -> Output[U]
// ...
Apply8(a, b, c, d, e, f, g, h, f) -> Output[U]
*/

// TODO Make a typed outputState[T].
// TODO Make the embeds above private.
