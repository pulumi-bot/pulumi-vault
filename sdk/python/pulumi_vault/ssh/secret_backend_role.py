# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretBackendRole(pulumi.CustomResource):
    allow_bare_domains: pulumi.Output[bool]
    """
    Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
    """
    allow_host_certificates: pulumi.Output[bool]
    """
    Specifies if certificates are allowed to be signed for use as a 'host'.
    """
    allow_subdomains: pulumi.Output[bool]
    """
    Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
    """
    allow_user_certificates: pulumi.Output[bool]
    """
    Specifies if certificates are allowed to be signed for use as a 'user'.
    """
    allow_user_key_ids: pulumi.Output[bool]
    """
    Specifies if users can override the key ID for a signed certificate with the `key_id` field.
    """
    allowed_critical_options: pulumi.Output[str]
    """
    Specifies a comma-separated list of critical options that certificates can have when signed.
    """
    allowed_domains: pulumi.Output[str]
    """
    The list of domains for which a client can request a host certificate.
    """
    allowed_extensions: pulumi.Output[str]
    """
    Specifies a comma-separated list of extensions that certificates can have when signed.
    """
    allowed_users: pulumi.Output[str]
    """
    Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
    """
    backend: pulumi.Output[str]
    """
    The path where the SSH secret backend is mounted.
    """
    cidr_list: pulumi.Output[str]
    """
    The comma-separated string of CIDR blocks for which this role is applicable.
    """
    default_critical_options: pulumi.Output[dict]
    """
    Specifies a map of critical options that certificates have when signed.
    """
    default_extensions: pulumi.Output[dict]
    """
    Specifies a map of extensions that certificates have when signed.
    """
    default_user: pulumi.Output[str]
    """
    Specifies the default username for which a credential will be generated.
    """
    key_id_format: pulumi.Output[str]
    """
    Specifies a custom format for the key id of a signed certificate.
    """
    key_type: pulumi.Output[str]
    """
    Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
    """
    max_ttl: pulumi.Output[str]
    """
    Specifies the Time To Live value.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the role to create.
    """
    ttl: pulumi.Output[str]
    """
    Specifies the maximum Time To Live value.
    """
    def __init__(__self__, resource_name, opts=None, allow_bare_domains=None, allow_host_certificates=None, allow_subdomains=None, allow_user_certificates=None, allow_user_key_ids=None, allowed_critical_options=None, allowed_domains=None, allowed_extensions=None, allowed_users=None, backend=None, cidr_list=None, default_critical_options=None, default_extensions=None, default_user=None, key_id_format=None, key_type=None, max_ttl=None, name=None, ttl=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage roles in an SSH secret backend
        [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_bare_domains: Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_host_certificates: Specifies if certificates are allowed to be signed for use as a 'host'.
        :param pulumi.Input[bool] allow_subdomains: Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_user_certificates: Specifies if certificates are allowed to be signed for use as a 'user'.
        :param pulumi.Input[bool] allow_user_key_ids: Specifies if users can override the key ID for a signed certificate with the `key_id` field.
        :param pulumi.Input[str] allowed_critical_options: Specifies a comma-separated list of critical options that certificates can have when signed.
        :param pulumi.Input[str] allowed_domains: The list of domains for which a client can request a host certificate.
        :param pulumi.Input[str] allowed_extensions: Specifies a comma-separated list of extensions that certificates can have when signed.
        :param pulumi.Input[str] allowed_users: Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted.
        :param pulumi.Input[str] cidr_list: The comma-separated string of CIDR blocks for which this role is applicable.
        :param pulumi.Input[dict] default_critical_options: Specifies a map of critical options that certificates have when signed.
        :param pulumi.Input[dict] default_extensions: Specifies a map of extensions that certificates have when signed.
        :param pulumi.Input[str] default_user: Specifies the default username for which a credential will be generated.
        :param pulumi.Input[str] key_id_format: Specifies a custom format for the key id of a signed certificate.
        :param pulumi.Input[str] key_type: Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
        :param pulumi.Input[str] max_ttl: Specifies the Time To Live value.
        :param pulumi.Input[str] name: Specifies the name of the role to create.
        :param pulumi.Input[str] ttl: Specifies the maximum Time To Live value.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ssh_secret_backend_role.html.markdown.
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

            __props__['allow_bare_domains'] = allow_bare_domains
            __props__['allow_host_certificates'] = allow_host_certificates
            __props__['allow_subdomains'] = allow_subdomains
            __props__['allow_user_certificates'] = allow_user_certificates
            __props__['allow_user_key_ids'] = allow_user_key_ids
            __props__['allowed_critical_options'] = allowed_critical_options
            __props__['allowed_domains'] = allowed_domains
            __props__['allowed_extensions'] = allowed_extensions
            __props__['allowed_users'] = allowed_users
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['cidr_list'] = cidr_list
            __props__['default_critical_options'] = default_critical_options
            __props__['default_extensions'] = default_extensions
            __props__['default_user'] = default_user
            __props__['key_id_format'] = key_id_format
            if key_type is None:
                raise TypeError("Missing required property 'key_type'")
            __props__['key_type'] = key_type
            __props__['max_ttl'] = max_ttl
            __props__['name'] = name
            __props__['ttl'] = ttl
        super(SecretBackendRole, __self__).__init__(
            'vault:ssh/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allow_bare_domains=None, allow_host_certificates=None, allow_subdomains=None, allow_user_certificates=None, allow_user_key_ids=None, allowed_critical_options=None, allowed_domains=None, allowed_extensions=None, allowed_users=None, backend=None, cidr_list=None, default_critical_options=None, default_extensions=None, default_user=None, key_id_format=None, key_type=None, max_ttl=None, name=None, ttl=None):
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_bare_domains: Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_host_certificates: Specifies if certificates are allowed to be signed for use as a 'host'.
        :param pulumi.Input[bool] allow_subdomains: Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_user_certificates: Specifies if certificates are allowed to be signed for use as a 'user'.
        :param pulumi.Input[bool] allow_user_key_ids: Specifies if users can override the key ID for a signed certificate with the `key_id` field.
        :param pulumi.Input[str] allowed_critical_options: Specifies a comma-separated list of critical options that certificates can have when signed.
        :param pulumi.Input[str] allowed_domains: The list of domains for which a client can request a host certificate.
        :param pulumi.Input[str] allowed_extensions: Specifies a comma-separated list of extensions that certificates can have when signed.
        :param pulumi.Input[str] allowed_users: Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted.
        :param pulumi.Input[str] cidr_list: The comma-separated string of CIDR blocks for which this role is applicable.
        :param pulumi.Input[dict] default_critical_options: Specifies a map of critical options that certificates have when signed.
        :param pulumi.Input[dict] default_extensions: Specifies a map of extensions that certificates have when signed.
        :param pulumi.Input[str] default_user: Specifies the default username for which a credential will be generated.
        :param pulumi.Input[str] key_id_format: Specifies a custom format for the key id of a signed certificate.
        :param pulumi.Input[str] key_type: Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
        :param pulumi.Input[str] max_ttl: Specifies the Time To Live value.
        :param pulumi.Input[str] name: Specifies the name of the role to create.
        :param pulumi.Input[str] ttl: Specifies the maximum Time To Live value.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ssh_secret_backend_role.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["allow_bare_domains"] = allow_bare_domains
        __props__["allow_host_certificates"] = allow_host_certificates
        __props__["allow_subdomains"] = allow_subdomains
        __props__["allow_user_certificates"] = allow_user_certificates
        __props__["allow_user_key_ids"] = allow_user_key_ids
        __props__["allowed_critical_options"] = allowed_critical_options
        __props__["allowed_domains"] = allowed_domains
        __props__["allowed_extensions"] = allowed_extensions
        __props__["allowed_users"] = allowed_users
        __props__["backend"] = backend
        __props__["cidr_list"] = cidr_list
        __props__["default_critical_options"] = default_critical_options
        __props__["default_extensions"] = default_extensions
        __props__["default_user"] = default_user
        __props__["key_id_format"] = key_id_format
        __props__["key_type"] = key_type
        __props__["max_ttl"] = max_ttl
        __props__["name"] = name
        __props__["ttl"] = ttl
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

