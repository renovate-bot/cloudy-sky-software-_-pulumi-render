// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebService(ctx *pulumi.Context, args *LookupWebServiceArgs, opts ...pulumi.InvokeOption) (*LookupWebServiceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupWebServiceResult
	err := ctx.Invoke("render:services:getWebService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebServiceArgs struct {
	// (Required) The ID of the service
	Id string `pulumi:"id"`
}

type LookupWebServiceResult struct {
	Items GetWebServiceType `pulumi:"items"`
}

// Defaults sets the appropriate defaults for LookupWebServiceResult
func (val *LookupWebServiceResult) Defaults() *LookupWebServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Items = *tmp.Items.Defaults()

	return &tmp
}

func LookupWebServiceOutput(ctx *pulumi.Context, args LookupWebServiceOutputArgs, opts ...pulumi.InvokeOption) LookupWebServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebServiceResult, error) {
			args := v.(LookupWebServiceArgs)
			r, err := LookupWebService(ctx, &args, opts...)
			var s LookupWebServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebServiceResultOutput)
}

type LookupWebServiceOutputArgs struct {
	// (Required) The ID of the service
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWebServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebServiceArgs)(nil)).Elem()
}

type LookupWebServiceResultOutput struct{ *pulumi.OutputState }

func (LookupWebServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebServiceResult)(nil)).Elem()
}

func (o LookupWebServiceResultOutput) ToLookupWebServiceResultOutput() LookupWebServiceResultOutput {
	return o
}

func (o LookupWebServiceResultOutput) ToLookupWebServiceResultOutputWithContext(ctx context.Context) LookupWebServiceResultOutput {
	return o
}

func (o LookupWebServiceResultOutput) Items() GetWebServiceTypeOutput {
	return o.ApplyT(func(v LookupWebServiceResult) GetWebServiceType { return v.Items }).(GetWebServiceTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebServiceResultOutput{})
}
