# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class TopicRule(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, cloudwatch_alarm=None, cloudwatch_metric=None, description=None, dynamodb=None, elasticsearch=None, enabled=None, firehose=None, kinesis=None, lambda_=None, name=None, republish=None, s3=None, sns=None, sql=None, sql_version=None, sqs=None):
        """Create a TopicRule resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cloudwatch_alarm'] = cloudwatch_alarm

        __props__['cloudwatch_metric'] = cloudwatch_metric

        __props__['description'] = description

        __props__['dynamodb'] = dynamodb

        __props__['elasticsearch'] = elasticsearch

        if not enabled:
            raise TypeError('Missing required property enabled')
        __props__['enabled'] = enabled

        __props__['firehose'] = firehose

        __props__['kinesis'] = kinesis

        __props__['lambda_'] = lambda_

        __props__['name'] = name

        __props__['republish'] = republish

        __props__['s3'] = s3

        __props__['sns'] = sns

        if not sql:
            raise TypeError('Missing required property sql')
        __props__['sql'] = sql

        if not sql_version:
            raise TypeError('Missing required property sql_version')
        __props__['sql_version'] = sql_version

        __props__['sqs'] = sqs

        __props__['arn'] = None

        super(TopicRule, __self__).__init__(
            'aws:iot/topicRule:TopicRule',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

