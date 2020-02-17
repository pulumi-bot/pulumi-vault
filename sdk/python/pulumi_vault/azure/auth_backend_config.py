# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendConfig(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    The path the Azure auth backend being configured was
    mounted at.  Defaults to `azure`.
    """
    client_id: pulumi.Output[str]
    """
    The client id for credentials to query the Azure APIs.
    Currently read permissions to query compute resources are required.
    """
    client_secret: pulumi.Output[str]
    """
    The client secret for credentials to query the
    Azure APIs.
    """
    environment: pulumi.Output[str]
    """
    The Azure cloud environment. Valid values:
    AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
    AzureGermanCloud.  Defaults to `AzurePublicCloud`.
    """
    resource: pulumi.Output[str]
    """
    The configured URL for the application registered in
    Azure Active Directory.
    """
    tenant_id: pulumi.Output[str]
    """
    The tenant id for the Azure Active Directory
    organization.
    """
    def __init__(__self__, resource_name, opts=None, backend=None, client_id=None, client_secret=None, environment=None, resource=None, tenant_id=None, __props__=None, __name__=None, __opts__=None):
        """
        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/azure_auth_backend_config.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the Azure auth backend being configured was
               mounted at.  Defaults to `azure`.
        :param pulumi.Input[str] client_id: The client id for credentials to query the Azure APIs.
               Currently read permissions to query compute resources are required.
        :param pulumi.Input[str] client_secret: The client secret for credentials to query the
               Azure APIs.
        :param pulumi.Input[str] environment: The Azure cloud environment. Valid values:
               AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
               AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        :param pulumi.Input[str] resource: The configured URL for the application registered in
               Azure Active Directory.
        :param pulumi.Input[str] tenant_id: The tenant id for the Azure Active Directory
               organization.
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

            __props__['backend'] = backend
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['environment'] = environment
            if resource is None:
                raise TypeError("Missing required property 'resource'")
            __props__['resource'] = resource
            if tenant_id is None:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
        super(AuthBackendConfig, __self__).__init__(
            'vault:azure/authBackendConfig:AuthBackendConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, client_id=None, client_secret=None, environment=None, resource=None, tenant_id=None):
        """
        Get an existing AuthBackendConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the Azure auth backend being configured was
               mounted at.  Defaults to `azure`.
        :param pulumi.Input[str] client_id: The client id for credentials to query the Azure APIs.
               Currently read permissions to query compute resources are required.
        :param pulumi.Input[str] client_secret: The client secret for credentials to query the
               Azure APIs.
        :param pulumi.Input[str] environment: The Azure cloud environment. Valid values:
               AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
               AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        :param pulumi.Input[str] resource: The configured URL for the application registered in
               Azure Active Directory.
        :param pulumi.Input[str] tenant_id: The tenant id for the Azure Active Directory
               organization.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["environment"] = environment
        __props__["resource"] = resource
        __props__["tenant_id"] = tenant_id
        return AuthBackendConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

