# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class NetworkInterface(pulumi.CustomResource):
    """
    Provides an Elastic network interface (ENI) resource.
    """
    def __init__(__self__, __name__, __opts__=None, attachments=None, description=None, private_ip=None, private_ips=None, private_ips_count=None, security_groups=None, source_dest_check=None, subnet_id=None, tags=None):
        """Create a NetworkInterface resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['attachments'] = attachments

        __props__['description'] = description

        __props__['private_ip'] = private_ip

        __props__['private_ips'] = private_ips

        __props__['private_ips_count'] = private_ips_count

        __props__['security_groups'] = security_groups

        __props__['source_dest_check'] = source_dest_check

        if not subnet_id:
            raise TypeError('Missing required property subnet_id')
        __props__['subnet_id'] = subnet_id

        __props__['tags'] = tags

        __props__['private_dns_name'] = None

        super(NetworkInterface, __self__).__init__(
            'aws:ec2/networkInterface:NetworkInterface',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

