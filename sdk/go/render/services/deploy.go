// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Deploy struct {
	pulumi.CustomResourceState

	ClearCache DeployClearCachePtrOutput `pulumi:"clearCache"`
	Commit     DeployCommitPtrOutput     `pulumi:"commit"`
}

// NewDeploy registers a new resource with the given unique name, arguments, and options.
func NewDeploy(ctx *pulumi.Context,
	name string, args *DeployArgs, opts ...pulumi.ResourceOption) (*Deploy, error) {
	if args == nil {
		args = &DeployArgs{}
	}

	if isZero(args.ClearCache) {
		args.ClearCache = DeployClearCache("do_not_clear")
	}
	var resource Deploy
	err := ctx.RegisterResource("render:services:Deploy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploy gets an existing Deploy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployState, opts ...pulumi.ResourceOption) (*Deploy, error) {
	var resource Deploy
	err := ctx.ReadResource("render:services:Deploy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deploy resources.
type deployState struct {
}

type DeployState struct {
}

func (DeployState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployState)(nil)).Elem()
}

type deployArgs struct {
	ClearCache *DeployClearCache `pulumi:"clearCache"`
}

// The set of arguments for constructing a Deploy resource.
type DeployArgs struct {
	ClearCache DeployClearCachePtrInput
}

func (DeployArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployArgs)(nil)).Elem()
}

type DeployInput interface {
	pulumi.Input

	ToDeployOutput() DeployOutput
	ToDeployOutputWithContext(ctx context.Context) DeployOutput
}

func (*Deploy) ElementType() reflect.Type {
	return reflect.TypeOf((**Deploy)(nil)).Elem()
}

func (i *Deploy) ToDeployOutput() DeployOutput {
	return i.ToDeployOutputWithContext(context.Background())
}

func (i *Deploy) ToDeployOutputWithContext(ctx context.Context) DeployOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployOutput)
}

type DeployOutput struct{ *pulumi.OutputState }

func (DeployOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deploy)(nil)).Elem()
}

func (o DeployOutput) ToDeployOutput() DeployOutput {
	return o
}

func (o DeployOutput) ToDeployOutputWithContext(ctx context.Context) DeployOutput {
	return o
}

func (o DeployOutput) ClearCache() DeployClearCachePtrOutput {
	return o.ApplyT(func(v *Deploy) DeployClearCachePtrOutput { return v.ClearCache }).(DeployClearCachePtrOutput)
}

func (o DeployOutput) Commit() DeployCommitPtrOutput {
	return o.ApplyT(func(v *Deploy) DeployCommitPtrOutput { return v.Commit }).(DeployCommitPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployInput)(nil)).Elem(), &Deploy{})
	pulumi.RegisterOutputType(DeployOutput{})
}
