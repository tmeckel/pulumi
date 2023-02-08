// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package myedgeorder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The list of configurations.
// API Version: 2020-12-01-preview.
func ListConfigurations(ctx *pulumi.Context, args *ListConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListConfigurationsResult, error) {
	var rv ListConfigurationsResult
	err := ctx.Invoke("myedgeorder::listConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConfigurationsArgs struct {
	// Holds details about product hierarchy information and filterable property.
	ConfigurationFilters []ConfigurationFilters `pulumi:"configurationFilters"`
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
	CustomerSubscriptionDetails *CustomerSubscriptionDetails `pulumi:"customerSubscriptionDetails"`
	// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
	SkipToken *string `pulumi:"skipToken"`
}

// The list of configurations.
type ListConfigurationsResult struct {
	// Link for the next set of configurations.
	NextLink *string `pulumi:"nextLink"`
	// List of configurations.
	Value []ConfigurationResponse `pulumi:"value"`
}

func ListConfigurationsOutput(ctx *pulumi.Context, args ListConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConfigurationsResult, error) {
			args := v.(ListConfigurationsArgs)
			r, err := ListConfigurations(ctx, &args, opts...)
			var s ListConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConfigurationsResultOutput)
}

type ListConfigurationsOutputArgs struct {
	// Holds details about product hierarchy information and filterable property.
	ConfigurationFilters ConfigurationFiltersArrayInput `pulumi:"configurationFilters"`
	// Customer subscription properties. Clients can display available products to unregistered customers by explicitly passing subscription details
	CustomerSubscriptionDetails *CustomerSubscriptionDetailsArgs `pulumi:"customerSubscriptionDetails"`
	// $skipToken is supported on list of configurations, which provides the next page in the list of configurations.
	SkipToken *string `pulumi:"skipToken"`
}

func (ListConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationsArgs)(nil)).Elem()
}

// The list of configurations.
type ListConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConfigurationsResult)(nil)).Elem()
}

func (o ListConfigurationsResultOutput) ToListConfigurationsResultOutput() ListConfigurationsResultOutput {
	return o
}

func (o ListConfigurationsResultOutput) ToListConfigurationsResultOutputWithContext(ctx context.Context) ListConfigurationsResultOutput {
	return o
}

// Link for the next set of configurations.
func (o ListConfigurationsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListConfigurationsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of configurations.
func (o ListConfigurationsResultOutput) Value() ConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListConfigurationsResult) []ConfigurationResponse { return v.Value }).(ConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConfigurationsResultOutput{})
}
