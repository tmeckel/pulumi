// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package foo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnumThing int

const (
	EnumThingFour  = EnumThing(4)
	EnumThingSix   = EnumThing(6)
	EnumThingEight = EnumThing(8)
)

func (EnumThing) ElementType() reflect.Type {
	return reflect.TypeOf((*EnumThing)(nil)).Elem()
}

func (e EnumThing) ToEnumThingOutput() EnumThingOutput {
	return pulumi.ToOutput(e).(EnumThingOutput)
}

func (e EnumThing) ToEnumThingOutputWithContext(ctx context.Context) EnumThingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnumThingOutput)
}

func (e EnumThing) ToEnumThingPtrOutput() EnumThingPtrOutput {
	return e.ToEnumThingPtrOutputWithContext(context.Background())
}

func (e EnumThing) ToEnumThingPtrOutputWithContext(ctx context.Context) EnumThingPtrOutput {
	return EnumThing(e).ToEnumThingOutputWithContext(ctx).ToEnumThingPtrOutputWithContext(ctx)
}

func (e EnumThing) ToIntOutput() pulumi.IntOutput {
	return pulumi.ToOutput(pulumi.Int(e)).(pulumi.IntOutput)
}

func (e EnumThing) ToIntOutputWithContext(ctx context.Context) pulumi.IntOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.Int(e)).(pulumi.IntOutput)
}

func (e EnumThing) ToIntPtrOutput() pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntPtrOutputWithContext(context.Background())
}

func (e EnumThing) ToIntPtrOutputWithContext(ctx context.Context) pulumi.IntPtrOutput {
	return pulumi.Int(e).ToIntOutputWithContext(ctx).ToIntPtrOutputWithContext(ctx)
}

type EnumThingOutput struct{ *pulumi.OutputState }

func (EnumThingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnumThing)(nil)).Elem()
}

func (o EnumThingOutput) ToEnumThingOutput() EnumThingOutput {
	return o
}

func (o EnumThingOutput) ToEnumThingOutputWithContext(ctx context.Context) EnumThingOutput {
	return o
}

func (o EnumThingOutput) ToEnumThingPtrOutput() EnumThingPtrOutput {
	return o.ToEnumThingPtrOutputWithContext(context.Background())
}

func (o EnumThingOutput) ToEnumThingPtrOutputWithContext(ctx context.Context) EnumThingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnumThing) *EnumThing {
		return &v
	}).(EnumThingPtrOutput)
}

func (o EnumThingOutput) ToIntOutput() pulumi.IntOutput {
	return o.ToIntOutputWithContext(context.Background())
}

func (o EnumThingOutput) ToIntOutputWithContext(ctx context.Context) pulumi.IntOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnumThing) int {
		return int(e)
	}).(pulumi.IntOutput)
}

func (o EnumThingOutput) ToIntPtrOutput() pulumi.IntPtrOutput {
	return o.ToIntPtrOutputWithContext(context.Background())
}

func (o EnumThingOutput) ToIntPtrOutputWithContext(ctx context.Context) pulumi.IntPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnumThing) *int {
		v := int(e)
		return &v
	}).(pulumi.IntPtrOutput)
}

type EnumThingPtrOutput struct{ *pulumi.OutputState }

func (EnumThingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnumThing)(nil)).Elem()
}

func (o EnumThingPtrOutput) ToEnumThingPtrOutput() EnumThingPtrOutput {
	return o
}

func (o EnumThingPtrOutput) ToEnumThingPtrOutputWithContext(ctx context.Context) EnumThingPtrOutput {
	return o
}

func (o EnumThingPtrOutput) Elem() EnumThingOutput {
	return o.ApplyT(func(v *EnumThing) EnumThing {
		if v != nil {
			return *v
		}
		var ret EnumThing
		return ret
	}).(EnumThingOutput)
}

func (o EnumThingPtrOutput) ToIntPtrOutput() pulumi.IntPtrOutput {
	return o.ToIntPtrOutputWithContext(context.Background())
}

func (o EnumThingPtrOutput) ToIntPtrOutputWithContext(ctx context.Context) pulumi.IntPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnumThing) *int {
		if e == nil {
			return nil
		}
		v := int(*e)
		return &v
	}).(pulumi.IntPtrOutput)
}

// EnumThingInput is an input type that accepts EnumThingArgs and EnumThingOutput values.
// You can construct a concrete instance of `EnumThingInput` via:
//
//	EnumThingArgs{...}
type EnumThingInput interface {
	pulumi.Input

	ToEnumThingOutput() EnumThingOutput
	ToEnumThingOutputWithContext(context.Context) EnumThingOutput
}

var enumThingPtrType = reflect.TypeOf((**EnumThing)(nil)).Elem()

type EnumThingPtrInput interface {
	pulumi.Input

	ToEnumThingPtrOutput() EnumThingPtrOutput
	ToEnumThingPtrOutputWithContext(context.Context) EnumThingPtrOutput
}

type enumThingPtr int

func EnumThingPtr(v int) EnumThingPtrInput {
	return (*enumThingPtr)(&v)
}

func (*enumThingPtr) ElementType() reflect.Type {
	return enumThingPtrType
}

func (in *enumThingPtr) ToEnumThingPtrOutput() EnumThingPtrOutput {
	return pulumi.ToOutput(in).(EnumThingPtrOutput)
}

func (in *enumThingPtr) ToEnumThingPtrOutputWithContext(ctx context.Context) EnumThingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnumThingPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnumThingInput)(nil)).Elem(), EnumThing(4))
	pulumi.RegisterInputType(reflect.TypeOf((*EnumThingPtrInput)(nil)).Elem(), EnumThing(4))
	pulumi.RegisterOutputType(EnumThingOutput{})
	pulumi.RegisterOutputType(EnumThingPtrOutput{})
}
