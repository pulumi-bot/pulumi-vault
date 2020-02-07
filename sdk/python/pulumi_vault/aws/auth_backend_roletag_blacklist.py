# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendRoletagBlacklist(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    The path the AWS auth backend being configured was
    mounted at.
    """
    disable_periodic_tidy: pulumi.Output[bool]
    """
    If set to true, disables the periodic
    tidying of the roletag blacklist entries. Defaults to false.
    """
    safety_buffer: pulumi.Output[float]
    """
    The amount of extra time that must have passed
    beyond the roletag expiration, before it is removed from the backend storage.
    Defaults to 259,200 seconds, or 72 hours.
    """
    def __init__(__self__, resource_name, opts=None, backend=None, disable_periodic_tidy=None, safety_buffer=None, __props__=None, __name__=None, __opts__=None):
        """
        Configures the periodic tidying operation of the blacklisted role tag entries.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_roletag_blacklist.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.
        :param pulumi.Input[bool] disable_periodic_tidy: If set to true, disables the periodic
               tidying of the roletag blacklist entries. Defaults to false.
        :param pulumi.Input[float] safety_buffer: The amount of extra time that must have passed
               beyond the roletag expiration, before it is removed from the backend storage.
               Defaults to 259,200 seconds, or 72 hours.
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

            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['disable_periodic_tidy'] = disable_periodic_tidy
            __props__['safety_buffer'] = safety_buffer
        super(AuthBackendRoletagBlacklist, __self__).__init__(
            'vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, disable_periodic_tidy=None, safety_buffer=None):
        """
        Get an existing AuthBackendRoletagBlacklist resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.
        :param pulumi.Input[bool] disable_periodic_tidy: If set to true, disables the periodic
               tidying of the roletag blacklist entries. Defaults to false.
        :param pulumi.Input[float] safety_buffer: The amount of extra time that must have passed
               beyond the roletag expiration, before it is removed from the backend storage.
               Defaults to 259,200 seconds, or 72 hours.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["disable_periodic_tidy"] = disable_periodic_tidy
        __props__["safety_buffer"] = safety_buffer
        return AuthBackendRoletagBlacklist(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

