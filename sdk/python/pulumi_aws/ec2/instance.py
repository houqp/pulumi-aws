# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    """
    Provides an EC2 instance resource. This allows instances to be created, updated,
    and deleted. Instances also support [provisioning](https://www.terraform.io/docs/provisioners/index.html).
    """
    def __init__(__self__, __name__, __opts__=None, ami=None, associate_public_ip_address=None, availability_zone=None, cpu_core_count=None, cpu_threads_per_core=None, credit_specification=None, disable_api_termination=None, ebs_block_devices=None, ebs_optimized=None, ephemeral_block_devices=None, get_password_data=None, iam_instance_profile=None, instance_initiated_shutdown_behavior=None, instance_type=None, ipv6_address_count=None, ipv6_addresses=None, key_name=None, monitoring=None, network_interfaces=None, placement_group=None, private_ip=None, root_block_device=None, security_groups=None, source_dest_check=None, subnet_id=None, tags=None, tenancy=None, user_data=None, user_data_base64=None, volume_tags=None, vpc_security_group_ids=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not ami:
            raise TypeError('Missing required property ami')
        __props__['ami'] = ami

        __props__['associate_public_ip_address'] = associate_public_ip_address

        __props__['availability_zone'] = availability_zone

        __props__['cpu_core_count'] = cpu_core_count

        __props__['cpu_threads_per_core'] = cpu_threads_per_core

        __props__['credit_specification'] = credit_specification

        __props__['disable_api_termination'] = disable_api_termination

        __props__['ebs_block_devices'] = ebs_block_devices

        __props__['ebs_optimized'] = ebs_optimized

        __props__['ephemeral_block_devices'] = ephemeral_block_devices

        __props__['get_password_data'] = get_password_data

        __props__['iam_instance_profile'] = iam_instance_profile

        __props__['instance_initiated_shutdown_behavior'] = instance_initiated_shutdown_behavior

        if not instance_type:
            raise TypeError('Missing required property instance_type')
        __props__['instance_type'] = instance_type

        __props__['ipv6_address_count'] = ipv6_address_count

        __props__['ipv6_addresses'] = ipv6_addresses

        __props__['key_name'] = key_name

        __props__['monitoring'] = monitoring

        __props__['network_interfaces'] = network_interfaces

        __props__['placement_group'] = placement_group

        __props__['private_ip'] = private_ip

        __props__['root_block_device'] = root_block_device

        __props__['security_groups'] = security_groups

        __props__['source_dest_check'] = source_dest_check

        __props__['subnet_id'] = subnet_id

        __props__['tags'] = tags

        __props__['tenancy'] = tenancy

        __props__['user_data'] = user_data

        __props__['user_data_base64'] = user_data_base64

        __props__['volume_tags'] = volume_tags

        __props__['vpc_security_group_ids'] = vpc_security_group_ids

        __props__['arn'] = None
        __props__['instance_state'] = None
        __props__['network_interface_id'] = None
        __props__['password_data'] = None
        __props__['primary_network_interface_id'] = None
        __props__['private_dns'] = None
        __props__['public_dns'] = None
        __props__['public_ip'] = None

        super(Instance, __self__).__init__(
            'aws:ec2/instance:Instance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

