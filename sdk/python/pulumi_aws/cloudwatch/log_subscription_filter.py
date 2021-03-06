# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class LogSubscriptionFilter(pulumi.CustomResource):
    """
    Provides a CloudWatch Logs subscription filter resource.
    """
    def __init__(__self__, __name__, __opts__=None, destination_arn=None, distribution=None, filter_pattern=None, log_group=None, name=None, role_arn=None):
        """Create a LogSubscriptionFilter resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not destination_arn:
            raise TypeError('Missing required property destination_arn')
        __props__['destination_arn'] = destination_arn

        __props__['distribution'] = distribution

        if not filter_pattern:
            raise TypeError('Missing required property filter_pattern')
        __props__['filter_pattern'] = filter_pattern

        if not log_group:
            raise TypeError('Missing required property log_group')
        __props__['log_group'] = log_group

        __props__['name'] = name

        __props__['role_arn'] = role_arn

        super(LogSubscriptionFilter, __self__).__init__(
            'aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

