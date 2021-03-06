# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class UsagePlan(pulumi.CustomResource):
    """
    Provides an API Gateway Usage Plan.
    """
    def __init__(__self__, __name__, __opts__=None, api_stages=None, description=None, name=None, product_code=None, quota_settings=None, throttle_settings=None):
        """Create a UsagePlan resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['api_stages'] = api_stages

        __props__['description'] = description

        __props__['name'] = name

        __props__['product_code'] = product_code

        __props__['quota_settings'] = quota_settings

        __props__['throttle_settings'] = throttle_settings

        super(UsagePlan, __self__).__init__(
            'aws:apigateway/usagePlan:UsagePlan',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

