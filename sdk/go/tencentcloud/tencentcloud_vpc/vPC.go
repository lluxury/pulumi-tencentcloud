// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tencentcloud_vpc

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provide a resource to create a VPC.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-tencentcloud/blob/master/website/docs/r/vpc.html.markdown.
type VPC struct {
	s *pulumi.ResourceState
}

// NewVPC registers a new resource with the given unique name, arguments, and options.
func NewVPC(ctx *pulumi.Context,
	name string, args *VPCArgs, opts ...pulumi.ResourceOpt) (*VPC, error) {
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cidrBlock"] = nil
		inputs["dnsServers"] = nil
		inputs["isMulticast"] = nil
		inputs["name"] = nil
		inputs["tags"] = nil
	} else {
		inputs["cidrBlock"] = args.CidrBlock
		inputs["dnsServers"] = args.DnsServers
		inputs["isMulticast"] = args.IsMulticast
		inputs["name"] = args.Name
		inputs["tags"] = args.Tags
	}
	inputs["createTime"] = nil
	inputs["isDefault"] = nil
	s, err := ctx.RegisterResource("tencentcloud:tencentcloud_vpc/vPC:VPC", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VPC{s: s}, nil
}

// GetVPC gets an existing VPC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPC(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VPCState, opts ...pulumi.ResourceOpt) (*VPC, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cidrBlock"] = state.CidrBlock
		inputs["createTime"] = state.CreateTime
		inputs["dnsServers"] = state.DnsServers
		inputs["isDefault"] = state.IsDefault
		inputs["isMulticast"] = state.IsMulticast
		inputs["name"] = state.Name
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("tencentcloud:tencentcloud_vpc/vPC:VPC", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VPC{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VPC) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VPC) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
func (r *VPC) CidrBlock() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cidrBlock"])
}

// Creation time of VPC.
func (r *VPC) CreateTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["createTime"])
}

// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
func (r *VPC) DnsServers() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["dnsServers"])
}

// Indicates whether it is the default VPC for this region.
func (r *VPC) IsDefault() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isDefault"])
}

// Indicates whether VPC multicast is enabled. The default value is 'true'.
func (r *VPC) IsMulticast() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["isMulticast"])
}

// The name of the VPC.
func (r *VPC) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Tags of the VPC.
func (r *VPC) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering VPC resources.
type VPCState struct {
	// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
	CidrBlock interface{}
	// Creation time of VPC.
	CreateTime interface{}
	// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
	DnsServers interface{}
	// Indicates whether it is the default VPC for this region.
	IsDefault interface{}
	// Indicates whether VPC multicast is enabled. The default value is 'true'.
	IsMulticast interface{}
	// The name of the VPC.
	Name interface{}
	// Tags of the VPC.
	Tags interface{}
}

// The set of arguments for constructing a VPC resource.
type VPCArgs struct {
	// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
	CidrBlock interface{}
	// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
	DnsServers interface{}
	// Indicates whether VPC multicast is enabled. The default value is 'true'.
	IsMulticast interface{}
	// The name of the VPC.
	Name interface{}
	// Tags of the VPC.
	Tags interface{}
}
