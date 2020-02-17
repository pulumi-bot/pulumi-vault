# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendStsRole(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    AWS account ID to be associated with STS role.
    """
    backend: pulumi.Output[str]
    """
    Unique name of the auth backend to configure.
    """
    sts_role: pulumi.Output[str]
    """
    AWS ARN for STS role to be assumed when interacting with the account specified.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, backend=None, sts_role=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a AuthBackendStsRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID to be associated with STS role.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[str] sts_role: AWS ARN for STS role to be assumed when interacting with the account specified.
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

            if account_id is None:
                raise TypeError("Missing required property 'account_id'")
            __props__['account_id'] = account_id
            __props__['backend'] = backend
            if sts_role is None:
                raise TypeError("Missing required property 'sts_role'")
            __props__['sts_role'] = sts_role
        super(AuthBackendStsRole, __self__).__init__(
            'vault:aws/authBackendStsRole:AuthBackendStsRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, backend=None, sts_role=None):
        """
        Get an existing AuthBackendStsRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID to be associated with STS role.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[str] sts_role: AWS ARN for STS role to be assumed when interacting with the account specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["backend"] = backend
        __props__["sts_role"] = sts_role
        return AuthBackendStsRole(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

