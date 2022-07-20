// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package services

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeployCommit struct {
	CreatedAt *string `pulumi:"createdAt"`
	Id        *string `pulumi:"id"`
	Message   *string `pulumi:"message"`
}

type DeployCommitOutput struct{ *pulumi.OutputState }

func (DeployCommitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeployCommit)(nil)).Elem()
}

func (o DeployCommitOutput) ToDeployCommitOutput() DeployCommitOutput {
	return o
}

func (o DeployCommitOutput) ToDeployCommitOutputWithContext(ctx context.Context) DeployCommitOutput {
	return o
}

func (o DeployCommitOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeployCommit) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DeployCommitOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeployCommit) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o DeployCommitOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeployCommit) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type DeployCommitPtrOutput struct{ *pulumi.OutputState }

func (DeployCommitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployCommit)(nil)).Elem()
}

func (o DeployCommitPtrOutput) ToDeployCommitPtrOutput() DeployCommitPtrOutput {
	return o
}

func (o DeployCommitPtrOutput) ToDeployCommitPtrOutputWithContext(ctx context.Context) DeployCommitPtrOutput {
	return o
}

func (o DeployCommitPtrOutput) Elem() DeployCommitOutput {
	return o.ApplyT(func(v *DeployCommit) DeployCommit {
		if v != nil {
			return *v
		}
		var ret DeployCommit
		return ret
	}).(DeployCommitOutput)
}

func (o DeployCommitPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeployCommit) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o DeployCommitPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeployCommit) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DeployCommitPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeployCommit) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type ServerProperties struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

// ServerPropertiesInput is an input type that accepts ServerPropertiesArgs and ServerPropertiesOutput values.
// You can construct a concrete instance of `ServerPropertiesInput` via:
//
//          ServerPropertiesArgs{...}
type ServerPropertiesInput interface {
	pulumi.Input

	ToServerPropertiesOutput() ServerPropertiesOutput
	ToServerPropertiesOutputWithContext(context.Context) ServerPropertiesOutput
}

type ServerPropertiesArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ServerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerProperties)(nil)).Elem()
}

func (i ServerPropertiesArgs) ToServerPropertiesOutput() ServerPropertiesOutput {
	return i.ToServerPropertiesOutputWithContext(context.Background())
}

func (i ServerPropertiesArgs) ToServerPropertiesOutputWithContext(ctx context.Context) ServerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesOutput)
}

type ServerPropertiesOutput struct{ *pulumi.OutputState }

func (ServerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerProperties)(nil)).Elem()
}

func (o ServerPropertiesOutput) ToServerPropertiesOutput() ServerPropertiesOutput {
	return o
}

func (o ServerPropertiesOutput) ToServerPropertiesOutputWithContext(ctx context.Context) ServerPropertiesOutput {
	return o
}

func (o ServerPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ServiceDisk struct {
	MountPath string   `pulumi:"mountPath"`
	Name      string   `pulumi:"name"`
	SizeGB    *float64 `pulumi:"sizeGB"`
}

// Defaults sets the appropriate defaults for ServiceDisk
func (val *ServiceDisk) Defaults() *ServiceDisk {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SizeGB) {
		sizeGB_ := 1.0
		tmp.SizeGB = &sizeGB_
	}
	return &tmp
}

type ServiceDockerDetails struct {
	DockerCommand  *string `pulumi:"dockerCommand"`
	DockerContext  *string `pulumi:"dockerContext"`
	DockerfilePath *string `pulumi:"dockerfilePath"`
}

// Defaults sets the appropriate defaults for ServiceDockerDetails
func (val *ServiceDockerDetails) Defaults() *ServiceDockerDetails {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DockerfilePath) {
		dockerfilePath_ := "./Dockerfile"
		tmp.DockerfilePath = &dockerfilePath_
	}
	return &tmp
}

type ServiceNativeEnvironmentDetails struct {
	BuildCommand string `pulumi:"buildCommand"`
	StartCommand string `pulumi:"startCommand"`
}

// A service header object
type ServiceServiceHeader struct {
	Name  string `pulumi:"name"`
	Path  string `pulumi:"path"`
	Value string `pulumi:"value"`
}

// A static website service
type ServiceStaticSite struct {
	BuildCommand               *string                                      `pulumi:"buildCommand"`
	Headers                    []ServiceServiceHeader                       `pulumi:"headers"`
	ParentServer               *ServiceStaticSiteParentServerProperties     `pulumi:"parentServer"`
	PublishPath                *string                                      `pulumi:"publishPath"`
	PullRequestPreviewsEnabled *ServiceStaticSitePullRequestPreviewsEnabled `pulumi:"pullRequestPreviewsEnabled"`
	Routes                     []ServiceStaticSiteRoute                     `pulumi:"routes"`
	// The HTTPS service URL. A subdomain of onrender.com, by default.
	Url *string `pulumi:"url"`
}

// Defaults sets the appropriate defaults for ServiceStaticSite
func (val *ServiceStaticSite) Defaults() *ServiceStaticSite {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PublishPath) {
		publishPath_ := "public"
		tmp.PublishPath = &publishPath_
	}
	if isZero(tmp.PullRequestPreviewsEnabled) {
		pullRequestPreviewsEnabled_ := ServiceStaticSitePullRequestPreviewsEnabled("no")
		tmp.PullRequestPreviewsEnabled = &pullRequestPreviewsEnabled_
	}
	return &tmp
}

type ServiceStaticSiteParentServerProperties struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

// A route object for a static site
type ServiceStaticSiteRoute struct {
	Destination string                     `pulumi:"destination"`
	Source      string                     `pulumi:"source"`
	Type        ServiceStaticSiteRouteType `pulumi:"type"`
}

type ServiceWebService struct {
	Disk                       *ServiceDisk                                 `pulumi:"disk"`
	Env                        ServiceWebServiceEnv                         `pulumi:"env"`
	EnvSpecificDetails         interface{}                                  `pulumi:"envSpecificDetails"`
	HealthCheckPath            *string                                      `pulumi:"healthCheckPath"`
	NumInstances               *float64                                     `pulumi:"numInstances"`
	Plan                       *ServiceWebServicePlan                       `pulumi:"plan"`
	PullRequestPreviewsEnabled *ServiceWebServicePullRequestPreviewsEnabled `pulumi:"pullRequestPreviewsEnabled"`
	Region                     *ServiceWebServiceRegion                     `pulumi:"region"`
}

// Defaults sets the appropriate defaults for ServiceWebService
func (val *ServiceWebService) Defaults() *ServiceWebService {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Disk = tmp.Disk.Defaults()

	if isZero(tmp.NumInstances) {
		numInstances_ := 1.0
		tmp.NumInstances = &numInstances_
	}
	if isZero(tmp.Plan) {
		plan_ := ServiceWebServicePlan("starter")
		tmp.Plan = &plan_
	}
	if isZero(tmp.PullRequestPreviewsEnabled) {
		pullRequestPreviewsEnabled_ := ServiceWebServicePullRequestPreviewsEnabled("no")
		tmp.PullRequestPreviewsEnabled = &pullRequestPreviewsEnabled_
	}
	if isZero(tmp.Region) {
		region_ := ServiceWebServiceRegion("oregon")
		tmp.Region = &region_
	}
	return &tmp
}
func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPropertiesInput)(nil)).Elem(), ServerPropertiesArgs{})
	pulumi.RegisterOutputType(DeployCommitOutput{})
	pulumi.RegisterOutputType(DeployCommitPtrOutput{})
	pulumi.RegisterOutputType(ServerPropertiesOutput{})
}
