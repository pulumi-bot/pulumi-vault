# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendClient(pulumi.CustomResource):
    access_key: pulumi.Output[str]
    """
    The AWS access key that Vault should use for the
    auth backend.
    """
    backend: pulumi.Output[str]
    """
    The path the AWS auth backend being configured was
    mounted at.  Defaults to `aws`.
    """
    ec2_endpoint: pulumi.Output[str]
    """
    Override the URL Vault uses when making EC2 API
    calls.
    """
    iam_endpoint: pulumi.Output[str]
    """
    Override the URL Vault uses when making IAM API
    calls.
    """
    iam_server_id_header_value: pulumi.Output[str]
    """
    The value to require in the
    `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
    that are used in the IAM auth method.
    """
    secret_key: pulumi.Output[str]
    """
    The AWS secret key that Vault should use for the
    auth backend.
    """
    sts_endpoint: pulumi.Output[str]
    """
    Override the URL Vault uses when making STS API
    calls.
    """
    def __init__(__self__, resource_name, opts=None, access_key=None, backend=None, ec2_endpoint=None, iam_endpoint=None, iam_server_id_header_value=None, secret_key=None, sts_endpoint=None, __props__=None, __name__=None, __opts__=None):
        """
        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_client.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS access key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.  Defaults to `aws`.
        :param pulumi.Input[str] ec2_endpoint: Override the URL Vault uses when making EC2 API
               calls.
        :param pulumi.Input[str] iam_endpoint: Override the URL Vault uses when making IAM API
               calls.
        :param pulumi.Input[str] iam_server_id_header_value: The value to require in the
               `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
               that are used in the IAM auth method.
        :param pulumi.Input[str] secret_key: The AWS secret key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] sts_endpoint: Override the URL Vault uses when making STS API
               calls.
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

            __props__['access_key'] = access_key
            __props__['backend'] = backend
            __props__['ec2_endpoint'] = ec2_endpoint
            __props__['iam_endpoint'] = iam_endpoint
            __props__['iam_server_id_header_value'] = iam_server_id_header_value
            __props__['secret_key'] = secret_key
            __props__['sts_endpoint'] = sts_endpoint
        super(AuthBackendClient, __self__).__init__(
            'vault:aws/authBackendClient:AuthBackendClient',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_key=None, backend=None, ec2_endpoint=None, iam_endpoint=None, iam_server_id_header_value=None, secret_key=None, sts_endpoint=None):
        """
        Get an existing AuthBackendClient resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS access key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.  Defaults to `aws`.
        :param pulumi.Input[str] ec2_endpoint: Override the URL Vault uses when making EC2 API
               calls.
        :param pulumi.Input[str] iam_endpoint: Override the URL Vault uses when making IAM API
               calls.
        :param pulumi.Input[str] iam_server_id_header_value: The value to require in the
               `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
               that are used in the IAM auth method.
        :param pulumi.Input[str] secret_key: The AWS secret key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] sts_endpoint: Override the URL Vault uses when making STS API
               calls.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_key"] = access_key
        __props__["backend"] = backend
        __props__["ec2_endpoint"] = ec2_endpoint
        __props__["iam_endpoint"] = iam_endpoint
        __props__["iam_server_id_header_value"] = iam_server_id_header_value
        __props__["secret_key"] = secret_key
        __props__["sts_endpoint"] = sts_endpoint
        return AuthBackendClient(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

