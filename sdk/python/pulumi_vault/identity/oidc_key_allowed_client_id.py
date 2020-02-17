# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OidcKeyAllowedClientID(pulumi.CustomResource):
    allowed_client_id: pulumi.Output[str]
    """
    Role Client ID allowed to use the key for signing.
    """
    key_name: pulumi.Output[str]
    """
    Name of the key.
    """
    def __init__(__self__, resource_name, opts=None, allowed_client_id=None, key_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a OidcKeyAllowedClientID resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_client_id: Role Client ID allowed to use the key for signing.
        :param pulumi.Input[str] key_name: Name of the key.
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

            if allowed_client_id is None:
                raise TypeError("Missing required property 'allowed_client_id'")
            __props__['allowed_client_id'] = allowed_client_id
            if key_name is None:
                raise TypeError("Missing required property 'key_name'")
            __props__['key_name'] = key_name
        super(OidcKeyAllowedClientID, __self__).__init__(
            'vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allowed_client_id=None, key_name=None):
        """
        Get an existing OidcKeyAllowedClientID resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_client_id: Role Client ID allowed to use the key for signing.
        :param pulumi.Input[str] key_name: Name of the key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_client_id"] = allowed_client_id
        __props__["key_name"] = key_name
        return OidcKeyAllowedClientID(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

