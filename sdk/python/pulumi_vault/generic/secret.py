# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Secret(pulumi.CustomResource):
    data: pulumi.Output[dict]
    """
    Map of strings read from Vault.
    """
    data_json: pulumi.Output[str]
    """
    JSON-encoded secret data to write.
    """
    disable_read: pulumi.Output[bool]
    """
    Don't attempt to read the token from Vault if true; drift won't be detected.
    """
    path: pulumi.Output[str]
    """
    Full path where the generic secret will be written.
    """
    def __init__(__self__, resource_name, opts=None, data_json=None, disable_read=None, path=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Secret resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_json: JSON-encoded secret data to write.
        :param pulumi.Input[bool] disable_read: Don't attempt to read the token from Vault if true; drift won't be detected.
        :param pulumi.Input[str] path: Full path where the generic secret will be written.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if data_json is None:
                raise TypeError("Missing required property 'data_json'")
            __props__['data_json'] = data_json
            __props__['disable_read'] = disable_read
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['data'] = None
        super(Secret, __self__).__init__(
            'vault:generic/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, data=None, data_json=None, disable_read=None, path=None):
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] data: Map of strings read from Vault.
        :param pulumi.Input[str] data_json: JSON-encoded secret data to write.
        :param pulumi.Input[bool] disable_read: Don't attempt to read the token from Vault if true; drift won't be detected.
        :param pulumi.Input[str] path: Full path where the generic secret will be written.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data"] = data
        __props__["data_json"] = data_json
        __props__["disable_read"] = disable_read
        __props__["path"] = path
        return Secret(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

