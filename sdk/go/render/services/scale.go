// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Scale struct {
	pulumi.CustomResourceState

	NumInstances pulumi.Float64Output `pulumi:"numInstances"`
}

// NewScale registers a new resource with the given unique name, arguments, and options.
func NewScale(ctx *pulumi.Context,
	name string, args *ScaleArgs, opts ...pulumi.ResourceOption) (*Scale, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NumInstances == nil {
		return nil, errors.New("invalid value for required argument 'NumInstances'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Scale
	err := ctx.RegisterResource("render:services:Scale", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScale gets an existing Scale resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScale(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScaleState, opts ...pulumi.ResourceOption) (*Scale, error) {
	var resource Scale
	err := ctx.ReadResource("render:services:Scale", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Scale resources.
type scaleState struct {
}

type ScaleState struct {
}

func (ScaleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleState)(nil)).Elem()
}

type scaleArgs struct {
	NumInstances float64 `pulumi:"numInstances"`
	// (Required) The ID of the service
	ServiceId *string `pulumi:"serviceId"`
}

// The set of arguments for constructing a Scale resource.
type ScaleArgs struct {
	NumInstances pulumi.Float64Input
	// (Required) The ID of the service
	ServiceId pulumi.StringPtrInput
}

func (ScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scaleArgs)(nil)).Elem()
}

type ScaleInput interface {
	pulumi.Input

	ToScaleOutput() ScaleOutput
	ToScaleOutputWithContext(ctx context.Context) ScaleOutput
}

func (*Scale) ElementType() reflect.Type {
	return reflect.TypeOf((**Scale)(nil)).Elem()
}

func (i *Scale) ToScaleOutput() ScaleOutput {
	return i.ToScaleOutputWithContext(context.Background())
}

func (i *Scale) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleOutput)
}

type ScaleOutput struct{ *pulumi.OutputState }

func (ScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scale)(nil)).Elem()
}

func (o ScaleOutput) ToScaleOutput() ScaleOutput {
	return o
}

func (o ScaleOutput) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return o
}

func (o ScaleOutput) NumInstances() pulumi.Float64Output {
	return o.ApplyT(func(v *Scale) pulumi.Float64Output { return v.NumInstances }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleInput)(nil)).Elem(), &Scale{})
	pulumi.RegisterOutputType(ScaleOutput{})
}
