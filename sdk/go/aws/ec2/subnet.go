// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an VPC subnet resource.
type Subnet struct {
	s *pulumi.ResourceState
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOpt) (*Subnet, error) {
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["assignIpv6AddressOnCreation"] = nil
		inputs["availabilityZone"] = nil
		inputs["cidrBlock"] = nil
		inputs["ipv6CidrBlock"] = nil
		inputs["mapPublicIpOnLaunch"] = nil
		inputs["tags"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["assignIpv6AddressOnCreation"] = args.AssignIpv6AddressOnCreation
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["cidrBlock"] = args.CidrBlock
		inputs["ipv6CidrBlock"] = args.Ipv6CidrBlock
		inputs["mapPublicIpOnLaunch"] = args.MapPublicIpOnLaunch
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	inputs["arn"] = nil
	inputs["ipv6CidrBlockAssociationId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/subnet:Subnet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subnet{s: s}, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetState, opts ...pulumi.ResourceOpt) (*Subnet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["assignIpv6AddressOnCreation"] = state.AssignIpv6AddressOnCreation
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["cidrBlock"] = state.CidrBlock
		inputs["ipv6CidrBlock"] = state.Ipv6CidrBlock
		inputs["ipv6CidrBlockAssociationId"] = state.Ipv6CidrBlockAssociationId
		inputs["mapPublicIpOnLaunch"] = state.MapPublicIpOnLaunch
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/subnet:Subnet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subnet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Subnet) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Subnet) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The ARN of the subnet.
// * `availability_zone`- The AZ for the subnet.
func (r *Subnet) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Specify true to indicate
// that network interfaces created in the specified subnet should be
// assigned an IPv6 address. Default is `false`
func (r *Subnet) AssignIpv6AddressOnCreation() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["assignIpv6AddressOnCreation"])
}

func (r *Subnet) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The CIDR block for the subnet.
func (r *Subnet) CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cidrBlock"])
}

// The IPv6 network range for the subnet,
// in CIDR notation. The subnet size must use a /64 prefix length.
func (r *Subnet) Ipv6CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6CidrBlock"])
}

// The association ID for the IPv6 CIDR block.
func (r *Subnet) Ipv6CidrBlockAssociationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6CidrBlockAssociationId"])
}

// Specify true to indicate
// that instances launched into the subnet should be assigned
// a public IP address. Default is `false`.
func (r *Subnet) MapPublicIpOnLaunch() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["mapPublicIpOnLaunch"])
}

// A mapping of tags to assign to the resource.
func (r *Subnet) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The VPC ID.
func (r *Subnet) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering Subnet resources.
type SubnetState struct {
	// The ARN of the subnet.
	// * `availability_zone`- The AZ for the subnet.
	Arn interface{}
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation interface{}
	AvailabilityZone interface{}
	// The CIDR block for the subnet.
	CidrBlock interface{}
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock interface{}
	// The association ID for the IPv6 CIDR block.
	Ipv6CidrBlockAssociationId interface{}
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VPC ID.
	VpcId interface{}
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation interface{}
	AvailabilityZone interface{}
	// The CIDR block for the subnet.
	CidrBlock interface{}
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock interface{}
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VPC ID.
	VpcId interface{}
}
