# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class CatalogTable(pulumi.CustomResource):
    """
    Provides a Glue Catalog Table Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality.
    """
    def __init__(__self__, __name__, __opts__=None, catalog_id=None, database_name=None, description=None, name=None, owner=None, parameters=None, partition_keys=None, retention=None, storage_descriptor=None, table_type=None, view_expanded_text=None, view_original_text=None):
        """Create a CatalogTable resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['catalog_id'] = catalog_id

        if not database_name:
            raise TypeError('Missing required property database_name')
        __props__['database_name'] = database_name

        __props__['description'] = description

        __props__['name'] = name

        __props__['owner'] = owner

        __props__['parameters'] = parameters

        __props__['partition_keys'] = partition_keys

        __props__['retention'] = retention

        __props__['storage_descriptor'] = storage_descriptor

        __props__['table_type'] = table_type

        __props__['view_expanded_text'] = view_expanded_text

        __props__['view_original_text'] = view_original_text

        super(CatalogTable, __self__).__init__(
            'aws:glue/catalogTable:CatalogTable',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

