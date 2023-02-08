// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pet struct {
	pulumi.CustomResourceState

	Name pulumi.StringPtrOutput `pulumi:"name"`
}

// NewPet registers a new resource with the given unique name, arguments, and options.
func NewPet(ctx *pulumi.Context,
	name string, args *PetArgs, opts ...pulumi.ResourceOption) (*Pet, error) {
	if args == nil {
		args = &PetArgs{}
	}

	var resource Pet
	err := ctx.RegisterResource("example::Pet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPet gets an existing Pet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PetState, opts ...pulumi.ResourceOption) (*Pet, error) {
	var resource Pet
	err := ctx.ReadResource("example::Pet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pet resources.
type petState struct {
}

type PetState struct {
}

func (PetState) ElementType() reflect.Type {
	return reflect.TypeOf((*petState)(nil)).Elem()
}

type petArgs struct {
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Pet resource.
type PetArgs struct {
	Name *string
}

func (PetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*petArgs)(nil)).Elem()
}

type PetInput interface {
	pulumi.Input

	ToPetOutput() PetOutput
	ToPetOutputWithContext(ctx context.Context) PetOutput
}

func (*Pet) ElementType() reflect.Type {
	return reflect.TypeOf((**Pet)(nil)).Elem()
}

func (i *Pet) ToPetOutput() PetOutput {
	return i.ToPetOutputWithContext(context.Background())
}

func (i *Pet) ToPetOutputWithContext(ctx context.Context) PetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PetOutput)
}

// PetArrayInput is an input type that accepts PetArray and PetArrayOutput values.
// You can construct a concrete instance of `PetArrayInput` via:
//
//	PetArray{ PetArgs{...} }
type PetArrayInput interface {
	pulumi.Input

	ToPetArrayOutput() PetArrayOutput
	ToPetArrayOutputWithContext(context.Context) PetArrayOutput
}

type PetArray []PetInput

func (PetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pet)(nil)).Elem()
}

func (i PetArray) ToPetArrayOutput() PetArrayOutput {
	return i.ToPetArrayOutputWithContext(context.Background())
}

func (i PetArray) ToPetArrayOutputWithContext(ctx context.Context) PetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PetArrayOutput)
}

// PetMapInput is an input type that accepts PetMap and PetMapOutput values.
// You can construct a concrete instance of `PetMapInput` via:
//
//	PetMap{ "key": PetArgs{...} }
type PetMapInput interface {
	pulumi.Input

	ToPetMapOutput() PetMapOutput
	ToPetMapOutputWithContext(context.Context) PetMapOutput
}

type PetMap map[string]PetInput

func (PetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pet)(nil)).Elem()
}

func (i PetMap) ToPetMapOutput() PetMapOutput {
	return i.ToPetMapOutputWithContext(context.Background())
}

func (i PetMap) ToPetMapOutputWithContext(ctx context.Context) PetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PetMapOutput)
}

type PetOutput struct{ *pulumi.OutputState }

func (PetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pet)(nil)).Elem()
}

func (o PetOutput) ToPetOutput() PetOutput {
	return o
}

func (o PetOutput) ToPetOutputWithContext(ctx context.Context) PetOutput {
	return o
}

func (o PetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pet) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

type PetArrayOutput struct{ *pulumi.OutputState }

func (PetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pet)(nil)).Elem()
}

func (o PetArrayOutput) ToPetArrayOutput() PetArrayOutput {
	return o
}

func (o PetArrayOutput) ToPetArrayOutputWithContext(ctx context.Context) PetArrayOutput {
	return o
}

func (o PetArrayOutput) Index(i pulumi.IntInput) PetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pet {
		return vs[0].([]*Pet)[vs[1].(int)]
	}).(PetOutput)
}

type PetMapOutput struct{ *pulumi.OutputState }

func (PetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pet)(nil)).Elem()
}

func (o PetMapOutput) ToPetMapOutput() PetMapOutput {
	return o
}

func (o PetMapOutput) ToPetMapOutputWithContext(ctx context.Context) PetMapOutput {
	return o
}

func (o PetMapOutput) MapIndex(k pulumi.StringInput) PetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pet {
		return vs[0].(map[string]*Pet)[vs[1].(string)]
	}).(PetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PetInput)(nil)).Elem(), &Pet{})
	pulumi.RegisterInputType(reflect.TypeOf((*PetArrayInput)(nil)).Elem(), PetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PetMapInput)(nil)).Elem(), PetMap{})
	pulumi.RegisterOutputType(PetOutput{})
	pulumi.RegisterOutputType(PetArrayOutput{})
	pulumi.RegisterOutputType(PetMapOutput{})
}
