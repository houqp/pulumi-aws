// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an EC2 Spot Instance Request resource. This allows instances to be
// requested on the spot market.
// 
// By default Terraform creates Spot Instance Requests with a `persistent` type,
// which means that for the duration of their lifetime, AWS will launch an
// instance with the configured details if and when the spot market will accept
// the requested price.
// 
// On destruction, Terraform will make an attempt to terminate the associated Spot
// Instance if there is one present.
// 
// Spot Instances requests with a `one-time` type will close the spot request
// when the instance is terminated either by the request being below the current spot
// price availability or by a user.
// 
// ~> **NOTE:** Because their behavior depends on the live status of the spot
// market, Spot Instance Requests have a unique lifecycle that makes them behave
// differently than other Terraform resources. Most importantly: there is __no
// guarantee__ that a Spot Instance exists to fulfill the request at any given
// point in time. See the [AWS Spot Instance
// documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
// for more information.
// 
type SpotInstanceRequest struct {
	s *pulumi.ResourceState
}

// NewSpotInstanceRequest registers a new resource with the given unique name, arguments, and options.
func NewSpotInstanceRequest(ctx *pulumi.Context,
	name string, args *SpotInstanceRequestArgs, opts ...pulumi.ResourceOpt) (*SpotInstanceRequest, error) {
	if args == nil || args.Ami == nil {
		return nil, errors.New("missing required argument 'Ami'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["ami"] = nil
		inputs["associatePublicIpAddress"] = nil
		inputs["availabilityZone"] = nil
		inputs["blockDurationMinutes"] = nil
		inputs["cpuCoreCount"] = nil
		inputs["cpuThreadsPerCore"] = nil
		inputs["creditSpecification"] = nil
		inputs["disableApiTermination"] = nil
		inputs["ebsBlockDevices"] = nil
		inputs["ebsOptimized"] = nil
		inputs["ephemeralBlockDevices"] = nil
		inputs["getPasswordData"] = nil
		inputs["iamInstanceProfile"] = nil
		inputs["instanceInitiatedShutdownBehavior"] = nil
		inputs["instanceInterruptionBehaviour"] = nil
		inputs["instanceType"] = nil
		inputs["ipv6AddressCount"] = nil
		inputs["ipv6Addresses"] = nil
		inputs["keyName"] = nil
		inputs["launchGroup"] = nil
		inputs["monitoring"] = nil
		inputs["networkInterfaces"] = nil
		inputs["placementGroup"] = nil
		inputs["privateIp"] = nil
		inputs["rootBlockDevice"] = nil
		inputs["securityGroups"] = nil
		inputs["sourceDestCheck"] = nil
		inputs["spotPrice"] = nil
		inputs["spotType"] = nil
		inputs["subnetId"] = nil
		inputs["tags"] = nil
		inputs["tenancy"] = nil
		inputs["userData"] = nil
		inputs["userDataBase64"] = nil
		inputs["validFrom"] = nil
		inputs["validUntil"] = nil
		inputs["volumeTags"] = nil
		inputs["vpcSecurityGroupIds"] = nil
		inputs["waitForFulfillment"] = nil
	} else {
		inputs["ami"] = args.Ami
		inputs["associatePublicIpAddress"] = args.AssociatePublicIpAddress
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["blockDurationMinutes"] = args.BlockDurationMinutes
		inputs["cpuCoreCount"] = args.CpuCoreCount
		inputs["cpuThreadsPerCore"] = args.CpuThreadsPerCore
		inputs["creditSpecification"] = args.CreditSpecification
		inputs["disableApiTermination"] = args.DisableApiTermination
		inputs["ebsBlockDevices"] = args.EbsBlockDevices
		inputs["ebsOptimized"] = args.EbsOptimized
		inputs["ephemeralBlockDevices"] = args.EphemeralBlockDevices
		inputs["getPasswordData"] = args.GetPasswordData
		inputs["iamInstanceProfile"] = args.IamInstanceProfile
		inputs["instanceInitiatedShutdownBehavior"] = args.InstanceInitiatedShutdownBehavior
		inputs["instanceInterruptionBehaviour"] = args.InstanceInterruptionBehaviour
		inputs["instanceType"] = args.InstanceType
		inputs["ipv6AddressCount"] = args.Ipv6AddressCount
		inputs["ipv6Addresses"] = args.Ipv6Addresses
		inputs["keyName"] = args.KeyName
		inputs["launchGroup"] = args.LaunchGroup
		inputs["monitoring"] = args.Monitoring
		inputs["networkInterfaces"] = args.NetworkInterfaces
		inputs["placementGroup"] = args.PlacementGroup
		inputs["privateIp"] = args.PrivateIp
		inputs["rootBlockDevice"] = args.RootBlockDevice
		inputs["securityGroups"] = args.SecurityGroups
		inputs["sourceDestCheck"] = args.SourceDestCheck
		inputs["spotPrice"] = args.SpotPrice
		inputs["spotType"] = args.SpotType
		inputs["subnetId"] = args.SubnetId
		inputs["tags"] = args.Tags
		inputs["tenancy"] = args.Tenancy
		inputs["userData"] = args.UserData
		inputs["userDataBase64"] = args.UserDataBase64
		inputs["validFrom"] = args.ValidFrom
		inputs["validUntil"] = args.ValidUntil
		inputs["volumeTags"] = args.VolumeTags
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
		inputs["waitForFulfillment"] = args.WaitForFulfillment
	}
	inputs["arn"] = nil
	inputs["instanceState"] = nil
	inputs["networkInterfaceId"] = nil
	inputs["passwordData"] = nil
	inputs["primaryNetworkInterfaceId"] = nil
	inputs["privateDns"] = nil
	inputs["publicDns"] = nil
	inputs["publicIp"] = nil
	inputs["spotBidStatus"] = nil
	inputs["spotInstanceId"] = nil
	inputs["spotRequestState"] = nil
	s, err := ctx.RegisterResource("aws:ec2/spotInstanceRequest:SpotInstanceRequest", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SpotInstanceRequest{s: s}, nil
}

// GetSpotInstanceRequest gets an existing SpotInstanceRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpotInstanceRequest(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SpotInstanceRequestState, opts ...pulumi.ResourceOpt) (*SpotInstanceRequest, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["ami"] = state.Ami
		inputs["arn"] = state.Arn
		inputs["associatePublicIpAddress"] = state.AssociatePublicIpAddress
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["blockDurationMinutes"] = state.BlockDurationMinutes
		inputs["cpuCoreCount"] = state.CpuCoreCount
		inputs["cpuThreadsPerCore"] = state.CpuThreadsPerCore
		inputs["creditSpecification"] = state.CreditSpecification
		inputs["disableApiTermination"] = state.DisableApiTermination
		inputs["ebsBlockDevices"] = state.EbsBlockDevices
		inputs["ebsOptimized"] = state.EbsOptimized
		inputs["ephemeralBlockDevices"] = state.EphemeralBlockDevices
		inputs["getPasswordData"] = state.GetPasswordData
		inputs["iamInstanceProfile"] = state.IamInstanceProfile
		inputs["instanceInitiatedShutdownBehavior"] = state.InstanceInitiatedShutdownBehavior
		inputs["instanceInterruptionBehaviour"] = state.InstanceInterruptionBehaviour
		inputs["instanceState"] = state.InstanceState
		inputs["instanceType"] = state.InstanceType
		inputs["ipv6AddressCount"] = state.Ipv6AddressCount
		inputs["ipv6Addresses"] = state.Ipv6Addresses
		inputs["keyName"] = state.KeyName
		inputs["launchGroup"] = state.LaunchGroup
		inputs["monitoring"] = state.Monitoring
		inputs["networkInterfaces"] = state.NetworkInterfaces
		inputs["networkInterfaceId"] = state.NetworkInterfaceId
		inputs["passwordData"] = state.PasswordData
		inputs["placementGroup"] = state.PlacementGroup
		inputs["primaryNetworkInterfaceId"] = state.PrimaryNetworkInterfaceId
		inputs["privateDns"] = state.PrivateDns
		inputs["privateIp"] = state.PrivateIp
		inputs["publicDns"] = state.PublicDns
		inputs["publicIp"] = state.PublicIp
		inputs["rootBlockDevice"] = state.RootBlockDevice
		inputs["securityGroups"] = state.SecurityGroups
		inputs["sourceDestCheck"] = state.SourceDestCheck
		inputs["spotBidStatus"] = state.SpotBidStatus
		inputs["spotInstanceId"] = state.SpotInstanceId
		inputs["spotPrice"] = state.SpotPrice
		inputs["spotRequestState"] = state.SpotRequestState
		inputs["spotType"] = state.SpotType
		inputs["subnetId"] = state.SubnetId
		inputs["tags"] = state.Tags
		inputs["tenancy"] = state.Tenancy
		inputs["userData"] = state.UserData
		inputs["userDataBase64"] = state.UserDataBase64
		inputs["validFrom"] = state.ValidFrom
		inputs["validUntil"] = state.ValidUntil
		inputs["volumeTags"] = state.VolumeTags
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
		inputs["waitForFulfillment"] = state.WaitForFulfillment
	}
	s, err := ctx.ReadResource("aws:ec2/spotInstanceRequest:SpotInstanceRequest", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SpotInstanceRequest{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SpotInstanceRequest) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SpotInstanceRequest) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The AMI to use for the instance.
func (r *SpotInstanceRequest) Ami() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ami"])
}

func (r *SpotInstanceRequest) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Associate a public ip address with an instance in a VPC.  Boolean value.
func (r *SpotInstanceRequest) AssociatePublicIpAddress() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["associatePublicIpAddress"])
}

// The AZ to start the instance in.
func (r *SpotInstanceRequest) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The required duration for the Spot instances, in minutes. This value must be a multiple of 60 (60, 120, 180, 240, 300, or 360).
// The duration period starts as soon as your Spot instance receives its instance ID. At the end of the duration period, Amazon EC2 marks the Spot instance for termination and provides a Spot instance termination notice, which gives the instance a two-minute warning before it terminates.
// Note that you can't specify an Availability Zone group or a launch group if you specify a duration.
func (r *SpotInstanceRequest) BlockDurationMinutes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["blockDurationMinutes"])
}

// Sets the number of CPU cores for an instance. This option is 
// only supported on creation of instance type that support CPU Options
// [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
func (r *SpotInstanceRequest) CpuCoreCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuCoreCount"])
}

// If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
func (r *SpotInstanceRequest) CpuThreadsPerCore() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuThreadsPerCore"])
}

// Customize the credit specification of the instance. See [Credit Specification](#credit-specification) below for more details.
func (r *SpotInstanceRequest) CreditSpecification() *pulumi.Output {
	return r.s.State["creditSpecification"]
}

// If true, enables [EC2 Instance
// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
func (r *SpotInstanceRequest) DisableApiTermination() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disableApiTermination"])
}

// Additional EBS block devices to attach to the
// instance.  See [Block Devices](#block-devices) below for details.
func (r *SpotInstanceRequest) EbsBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ebsBlockDevices"])
}

// If true, the launched EC2 instance will be EBS-optimized.
// Note that if this is not set on an instance type that is optimized by default then
// this will show as disabled but if the instance type is optimized by default then
// there is no need to set this and there is no effect to disabling it.
// See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
func (r *SpotInstanceRequest) EbsOptimized() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ebsOptimized"])
}

// Customize Ephemeral (also known as
// "Instance Store") volumes on the instance. See [Block Devices](#block-devices) below for details.
func (r *SpotInstanceRequest) EphemeralBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ephemeralBlockDevices"])
}

// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
func (r *SpotInstanceRequest) GetPasswordData() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["getPasswordData"])
}

// The IAM Instance Profile to
// launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
// * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
func (r *SpotInstanceRequest) IamInstanceProfile() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["iamInstanceProfile"])
}

// Shutdown behavior for the
// instance. Amazon defaults this to `stop` for EBS-backed instances and
// `terminate` for instance-store instances. Cannot be set on instance-store
// instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
func (r *SpotInstanceRequest) InstanceInitiatedShutdownBehavior() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceInitiatedShutdownBehavior"])
}

// Indicates whether a Spot instance stops or terminates when it is interrupted. Default is `terminate` as this is the current AWS behaviour.
func (r *SpotInstanceRequest) InstanceInterruptionBehaviour() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceInterruptionBehaviour"])
}

func (r *SpotInstanceRequest) InstanceState() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceState"])
}

// The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
func (r *SpotInstanceRequest) InstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceType"])
}

func (r *SpotInstanceRequest) Ipv6AddressCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ipv6AddressCount"])
}

// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
func (r *SpotInstanceRequest) Ipv6Addresses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipv6Addresses"])
}

// The key name of the Key Pair to use for the instance; which can be managed using [the `aws_key_pair` resource](key_pair.html).
func (r *SpotInstanceRequest) KeyName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyName"])
}

// A launch group is a group of spot instances that launch together and terminate together.
// If left empty instances are launched and terminated individually.
func (r *SpotInstanceRequest) LaunchGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["launchGroup"])
}

// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
func (r *SpotInstanceRequest) Monitoring() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["monitoring"])
}

// Customize network interfaces to be attached at instance boot time. See [Network Interfaces](#network-interfaces) below for more details.
func (r *SpotInstanceRequest) NetworkInterfaces() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkInterfaces"])
}

// The ID of the network interface to attach.
func (r *SpotInstanceRequest) NetworkInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkInterfaceId"])
}

func (r *SpotInstanceRequest) PasswordData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordData"])
}

// The Placement Group to start the instance in.
func (r *SpotInstanceRequest) PlacementGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["placementGroup"])
}

func (r *SpotInstanceRequest) PrimaryNetworkInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["primaryNetworkInterfaceId"])
}

// The private DNS name assigned to the instance. Can only be
// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
// for your VPC
func (r *SpotInstanceRequest) PrivateDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateDns"])
}

// Private IP address to associate with the
// instance in a VPC.
func (r *SpotInstanceRequest) PrivateIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIp"])
}

// The public DNS name assigned to the instance. For EC2-VPC, this
// is only available if you've enabled DNS hostnames for your VPC
func (r *SpotInstanceRequest) PublicDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicDns"])
}

// The public IP address assigned to the instance, if applicable.
func (r *SpotInstanceRequest) PublicIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicIp"])
}

// Customize details about the root block
// device of the instance. See [Block Devices](#block-devices) below for details.
func (r *SpotInstanceRequest) RootBlockDevice() *pulumi.Output {
	return r.s.State["rootBlockDevice"]
}

// A list of security group names to associate with.
func (r *SpotInstanceRequest) SecurityGroups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroups"])
}

// Controls if traffic is routed to the instance when
// the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
func (r *SpotInstanceRequest) SourceDestCheck() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["sourceDestCheck"])
}

// The current [bid
// status](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-bid-status.html)
// of the Spot Instance Request.
// * `spot_request_state` The current [request
// state](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html#creating-spot-request-status)
// of the Spot Instance Request.
func (r *SpotInstanceRequest) SpotBidStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["spotBidStatus"])
}

// The Instance ID (if any) that is currently fulfilling
// the Spot Instance request.
func (r *SpotInstanceRequest) SpotInstanceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["spotInstanceId"])
}

// The maximum price to request on the spot market.
func (r *SpotInstanceRequest) SpotPrice() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["spotPrice"])
}

func (r *SpotInstanceRequest) SpotRequestState() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["spotRequestState"])
}

// If set to `one-time`, after
// the instance is terminated, the spot request will be closed.
func (r *SpotInstanceRequest) SpotType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["spotType"])
}

// The VPC Subnet ID to launch in.
func (r *SpotInstanceRequest) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// A mapping of tags to assign to the resource.
func (r *SpotInstanceRequest) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
func (r *SpotInstanceRequest) Tenancy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenancy"])
}

// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
func (r *SpotInstanceRequest) UserData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userData"])
}

// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
func (r *SpotInstanceRequest) UserDataBase64() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userDataBase64"])
}

// The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
func (r *SpotInstanceRequest) ValidFrom() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["validFrom"])
}

// The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request. The default end date is 7 days from the current date.
func (r *SpotInstanceRequest) ValidUntil() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["validUntil"])
}

// A mapping of tags to assign to the devices created by the instance at launch time.
func (r *SpotInstanceRequest) VolumeTags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["volumeTags"])
}

// A list of security group IDs to associate with.
func (r *SpotInstanceRequest) VpcSecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vpcSecurityGroupIds"])
}

// If set, Terraform will
// wait for the Spot Request to be fulfilled, and will throw an error if the
// timeout of 10m is reached.
func (r *SpotInstanceRequest) WaitForFulfillment() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["waitForFulfillment"])
}

// Input properties used for looking up and filtering SpotInstanceRequest resources.
type SpotInstanceRequestState struct {
	// The AMI to use for the instance.
	Ami interface{}
	Arn interface{}
	// Associate a public ip address with an instance in a VPC.  Boolean value.
	AssociatePublicIpAddress interface{}
	// The AZ to start the instance in.
	AvailabilityZone interface{}
	// The required duration for the Spot instances, in minutes. This value must be a multiple of 60 (60, 120, 180, 240, 300, or 360).
	// The duration period starts as soon as your Spot instance receives its instance ID. At the end of the duration period, Amazon EC2 marks the Spot instance for termination and provides a Spot instance termination notice, which gives the instance a two-minute warning before it terminates.
	// Note that you can't specify an Availability Zone group or a launch group if you specify a duration.
	BlockDurationMinutes interface{}
	// Sets the number of CPU cores for an instance. This option is 
	// only supported on creation of instance type that support CPU Options
	// [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
	CpuCoreCount interface{}
	// If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
	CpuThreadsPerCore interface{}
	// Customize the credit specification of the instance. See [Credit Specification](#credit-specification) below for more details.
	CreditSpecification interface{}
	// If true, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination interface{}
	// Additional EBS block devices to attach to the
	// instance.  See [Block Devices](#block-devices) below for details.
	EbsBlockDevices interface{}
	// If true, the launched EC2 instance will be EBS-optimized.
	// Note that if this is not set on an instance type that is optimized by default then
	// this will show as disabled but if the instance type is optimized by default then
	// there is no need to set this and there is no effect to disabling it.
	// See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
	EbsOptimized interface{}
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See [Block Devices](#block-devices) below for details.
	EphemeralBlockDevices interface{}
	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData interface{}
	// The IAM Instance Profile to
	// launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
	// * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	IamInstanceProfile interface{}
	// Shutdown behavior for the
	// instance. Amazon defaults this to `stop` for EBS-backed instances and
	// `terminate` for instance-store instances. Cannot be set on instance-store
	// instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
	InstanceInitiatedShutdownBehavior interface{}
	// Indicates whether a Spot instance stops or terminates when it is interrupted. Default is `terminate` as this is the current AWS behaviour.
	InstanceInterruptionBehaviour interface{}
	InstanceState interface{}
	// The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
	InstanceType interface{}
	Ipv6AddressCount interface{}
	// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
	Ipv6Addresses interface{}
	// The key name of the Key Pair to use for the instance; which can be managed using [the `aws_key_pair` resource](key_pair.html).
	KeyName interface{}
	// A launch group is a group of spot instances that launch together and terminate together.
	// If left empty instances are launched and terminated individually.
	LaunchGroup interface{}
	// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
	Monitoring interface{}
	// Customize network interfaces to be attached at instance boot time. See [Network Interfaces](#network-interfaces) below for more details.
	NetworkInterfaces interface{}
	// The ID of the network interface to attach.
	NetworkInterfaceId interface{}
	PasswordData interface{}
	// The Placement Group to start the instance in.
	PlacementGroup interface{}
	PrimaryNetworkInterfaceId interface{}
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns interface{}
	// Private IP address to associate with the
	// instance in a VPC.
	PrivateIp interface{}
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns interface{}
	// The public IP address assigned to the instance, if applicable.
	PublicIp interface{}
	// Customize details about the root block
	// device of the instance. See [Block Devices](#block-devices) below for details.
	RootBlockDevice interface{}
	// A list of security group names to associate with.
	SecurityGroups interface{}
	// Controls if traffic is routed to the instance when
	// the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
	SourceDestCheck interface{}
	// The current [bid
	// status](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-bid-status.html)
	// of the Spot Instance Request.
	// * `spot_request_state` The current [request
	// state](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html#creating-spot-request-status)
	// of the Spot Instance Request.
	SpotBidStatus interface{}
	// The Instance ID (if any) that is currently fulfilling
	// the Spot Instance request.
	SpotInstanceId interface{}
	// The maximum price to request on the spot market.
	SpotPrice interface{}
	SpotRequestState interface{}
	// If set to `one-time`, after
	// the instance is terminated, the spot request will be closed.
	SpotType interface{}
	// The VPC Subnet ID to launch in.
	SubnetId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
	Tenancy interface{}
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
	UserData interface{}
	// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 interface{}
	// The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
	ValidFrom interface{}
	// The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request. The default end date is 7 days from the current date.
	ValidUntil interface{}
	// A mapping of tags to assign to the devices created by the instance at launch time.
	VolumeTags interface{}
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds interface{}
	// If set, Terraform will
	// wait for the Spot Request to be fulfilled, and will throw an error if the
	// timeout of 10m is reached.
	WaitForFulfillment interface{}
}

// The set of arguments for constructing a SpotInstanceRequest resource.
type SpotInstanceRequestArgs struct {
	// The AMI to use for the instance.
	Ami interface{}
	// Associate a public ip address with an instance in a VPC.  Boolean value.
	AssociatePublicIpAddress interface{}
	// The AZ to start the instance in.
	AvailabilityZone interface{}
	// The required duration for the Spot instances, in minutes. This value must be a multiple of 60 (60, 120, 180, 240, 300, or 360).
	// The duration period starts as soon as your Spot instance receives its instance ID. At the end of the duration period, Amazon EC2 marks the Spot instance for termination and provides a Spot instance termination notice, which gives the instance a two-minute warning before it terminates.
	// Note that you can't specify an Availability Zone group or a launch group if you specify a duration.
	BlockDurationMinutes interface{}
	// Sets the number of CPU cores for an instance. This option is 
	// only supported on creation of instance type that support CPU Options
	// [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
	CpuCoreCount interface{}
	// If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
	CpuThreadsPerCore interface{}
	// Customize the credit specification of the instance. See [Credit Specification](#credit-specification) below for more details.
	CreditSpecification interface{}
	// If true, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination interface{}
	// Additional EBS block devices to attach to the
	// instance.  See [Block Devices](#block-devices) below for details.
	EbsBlockDevices interface{}
	// If true, the launched EC2 instance will be EBS-optimized.
	// Note that if this is not set on an instance type that is optimized by default then
	// this will show as disabled but if the instance type is optimized by default then
	// there is no need to set this and there is no effect to disabling it.
	// See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
	EbsOptimized interface{}
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See [Block Devices](#block-devices) below for details.
	EphemeralBlockDevices interface{}
	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData interface{}
	// The IAM Instance Profile to
	// launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
	// * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	IamInstanceProfile interface{}
	// Shutdown behavior for the
	// instance. Amazon defaults this to `stop` for EBS-backed instances and
	// `terminate` for instance-store instances. Cannot be set on instance-store
	// instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
	InstanceInitiatedShutdownBehavior interface{}
	// Indicates whether a Spot instance stops or terminates when it is interrupted. Default is `terminate` as this is the current AWS behaviour.
	InstanceInterruptionBehaviour interface{}
	// The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
	InstanceType interface{}
	Ipv6AddressCount interface{}
	// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
	Ipv6Addresses interface{}
	// The key name of the Key Pair to use for the instance; which can be managed using [the `aws_key_pair` resource](key_pair.html).
	KeyName interface{}
	// A launch group is a group of spot instances that launch together and terminate together.
	// If left empty instances are launched and terminated individually.
	LaunchGroup interface{}
	// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
	Monitoring interface{}
	// Customize network interfaces to be attached at instance boot time. See [Network Interfaces](#network-interfaces) below for more details.
	NetworkInterfaces interface{}
	// The Placement Group to start the instance in.
	PlacementGroup interface{}
	// Private IP address to associate with the
	// instance in a VPC.
	PrivateIp interface{}
	// Customize details about the root block
	// device of the instance. See [Block Devices](#block-devices) below for details.
	RootBlockDevice interface{}
	// A list of security group names to associate with.
	SecurityGroups interface{}
	// Controls if traffic is routed to the instance when
	// the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
	SourceDestCheck interface{}
	// The maximum price to request on the spot market.
	SpotPrice interface{}
	// If set to `one-time`, after
	// the instance is terminated, the spot request will be closed.
	SpotType interface{}
	// The VPC Subnet ID to launch in.
	SubnetId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
	Tenancy interface{}
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
	UserData interface{}
	// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 interface{}
	// The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
	ValidFrom interface{}
	// The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request. The default end date is 7 days from the current date.
	ValidUntil interface{}
	// A mapping of tags to assign to the devices created by the instance at launch time.
	VolumeTags interface{}
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds interface{}
	// If set, Terraform will
	// wait for the Spot Request to be fulfilled, and will throw an error if the
	// timeout of 10m is reached.
	WaitForFulfillment interface{}
}
