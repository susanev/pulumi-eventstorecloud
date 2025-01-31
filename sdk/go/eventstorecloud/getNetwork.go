// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventstorecloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves data for an existing `Network` resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/EventStore/pulumi-eventstorecloud/sdk/go/eventstorecloud"
//	"github.com/pulumi/pulumi-eventstorecloud/sdk/go/eventstorecloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := eventstorecloud.LookupNetwork(ctx, &GetNetworkArgs{
//				Name:      "Example Network",
//				ProjectId: _var.Project_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("networkCidr", example.CidrBlock)
//			return nil
//		})
//	}
//
// ```
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNetworkResult
	err := ctx.Invoke("eventstorecloud:index/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// Address space of the network in CIDR block notation
	CidrBlock string `pulumi:"cidrBlock"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
	// Provider region in which to provision the network
	Region string `pulumi:"region"`
	// Cloud Provider in which to provision the network.
	ResourceProvider string `pulumi:"resourceProvider"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResult, error) {
			args := v.(LookupNetworkArgs)
			r, err := LookupNetwork(ctx, &args, opts...)
			var s LookupNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	Name      pulumi.StringInput `pulumi:"name"`
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// Address space of the network in CIDR block notation
func (o LookupNetworkResultOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.CidrBlock }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Provider region in which to provision the network
func (o LookupNetworkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Region }).(pulumi.StringOutput)
}

// Cloud Provider in which to provision the network.
func (o LookupNetworkResultOutput) ResourceProvider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.ResourceProvider }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
