// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xyz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticPage struct {
	pulumi.ResourceState

	// The bucket resource.
	Bucket s3.BucketOutput `pulumi:"bucket"`
	// The website URL.
	WebsiteUrl pulumi.StringOutput `pulumi:"websiteUrl"`
}

// NewStaticPage registers a new resource with the given unique name, arguments, and options.
func NewStaticPage(ctx *pulumi.Context,
	name string, args *StaticPageArgs, opts ...pulumi.ResourceOption) (*StaticPage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IndexContent == nil {
		return nil, errors.New("invalid value for required argument 'IndexContent'")
	}
	var resource StaticPage
	err := ctx.RegisterRemoteComponentResource("xyz:index:StaticPage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type staticPageArgs struct {
	Foo *Foo `pulumi:"foo"`
	// The HTML content for index.html.
	IndexContent string `pulumi:"indexContent"`
}

// The set of arguments for constructing a StaticPage resource.
type StaticPageArgs struct {
	Foo *Foo
	// The HTML content for index.html.
	IndexContent pulumi.StringInput
}

func (StaticPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticPageArgs)(nil)).Elem()
}

type StaticPageInput interface {
	pulumi.Input

	ToStaticPageOutput() StaticPageOutput
	ToStaticPageOutputWithContext(ctx context.Context) StaticPageOutput
}

func (*StaticPage) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticPage)(nil)).Elem()
}

func (i *StaticPage) ToStaticPageOutput() StaticPageOutput {
	return i.ToStaticPageOutputWithContext(context.Background())
}

func (i *StaticPage) ToStaticPageOutputWithContext(ctx context.Context) StaticPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticPageOutput)
}

// StaticPageArrayInput is an input type that accepts StaticPageArray and StaticPageArrayOutput values.
// You can construct a concrete instance of `StaticPageArrayInput` via:
//
//          StaticPageArray{ StaticPageArgs{...} }
type StaticPageArrayInput interface {
	pulumi.Input

	ToStaticPageArrayOutput() StaticPageArrayOutput
	ToStaticPageArrayOutputWithContext(context.Context) StaticPageArrayOutput
}

type StaticPageArray []StaticPageInput

func (StaticPageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticPage)(nil)).Elem()
}

func (i StaticPageArray) ToStaticPageArrayOutput() StaticPageArrayOutput {
	return i.ToStaticPageArrayOutputWithContext(context.Background())
}

func (i StaticPageArray) ToStaticPageArrayOutputWithContext(ctx context.Context) StaticPageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticPageArrayOutput)
}

// StaticPageMapInput is an input type that accepts StaticPageMap and StaticPageMapOutput values.
// You can construct a concrete instance of `StaticPageMapInput` via:
//
//          StaticPageMap{ "key": StaticPageArgs{...} }
type StaticPageMapInput interface {
	pulumi.Input

	ToStaticPageMapOutput() StaticPageMapOutput
	ToStaticPageMapOutputWithContext(context.Context) StaticPageMapOutput
}

type StaticPageMap map[string]StaticPageInput

func (StaticPageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticPage)(nil)).Elem()
}

func (i StaticPageMap) ToStaticPageMapOutput() StaticPageMapOutput {
	return i.ToStaticPageMapOutputWithContext(context.Background())
}

func (i StaticPageMap) ToStaticPageMapOutputWithContext(ctx context.Context) StaticPageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticPageMapOutput)
}

type StaticPageOutput struct{ *pulumi.OutputState }

func (StaticPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticPage)(nil)).Elem()
}

func (o StaticPageOutput) ToStaticPageOutput() StaticPageOutput {
	return o
}

func (o StaticPageOutput) ToStaticPageOutputWithContext(ctx context.Context) StaticPageOutput {
	return o
}

// The bucket resource.
func (o StaticPageOutput) Bucket() s3.BucketOutput {
	return o.ApplyT(func(v *StaticPage) s3.BucketOutput { return v.Bucket }).(s3.BucketOutput)
}

// The website URL.
func (o StaticPageOutput) WebsiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticPage) pulumi.StringOutput { return v.WebsiteUrl }).(pulumi.StringOutput)
}

type StaticPageArrayOutput struct{ *pulumi.OutputState }

func (StaticPageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticPage)(nil)).Elem()
}

func (o StaticPageArrayOutput) ToStaticPageArrayOutput() StaticPageArrayOutput {
	return o
}

func (o StaticPageArrayOutput) ToStaticPageArrayOutputWithContext(ctx context.Context) StaticPageArrayOutput {
	return o
}

func (o StaticPageArrayOutput) Index(i pulumi.IntInput) StaticPageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StaticPage {
		return vs[0].([]*StaticPage)[vs[1].(int)]
	}).(StaticPageOutput)
}

type StaticPageMapOutput struct{ *pulumi.OutputState }

func (StaticPageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticPage)(nil)).Elem()
}

func (o StaticPageMapOutput) ToStaticPageMapOutput() StaticPageMapOutput {
	return o
}

func (o StaticPageMapOutput) ToStaticPageMapOutputWithContext(ctx context.Context) StaticPageMapOutput {
	return o
}

func (o StaticPageMapOutput) MapIndex(k pulumi.StringInput) StaticPageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StaticPage {
		return vs[0].(map[string]*StaticPage)[vs[1].(string)]
	}).(StaticPageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticPageInput)(nil)).Elem(), &StaticPage{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticPageArrayInput)(nil)).Elem(), StaticPageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticPageMapInput)(nil)).Elem(), StaticPageMap{})
	pulumi.RegisterOutputType(StaticPageOutput{})
	pulumi.RegisterOutputType(StaticPageArrayOutput{})
	pulumi.RegisterOutputType(StaticPageMapOutput{})
}
