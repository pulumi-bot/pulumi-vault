# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, address=None, auth_logins=None, ca_cert_dir=None, ca_cert_file=None, client_auths=None, max_lease_ttl_seconds=None, max_retries=None, namespace=None, skip_tls_verify=None, token=None, token_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the vault package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **auth_logins** object supports the following:
        
          * `namespace` (`pulumi.Input[str]`)
          * `parameters` (`pulumi.Input[dict]`)
          * `path` (`pulumi.Input[str]`)
        
        The **client_auths** object supports the following:
        
          * `certFile` (`pulumi.Input[str]`)
          * `keyFile` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/index.html.markdown.
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

            if address is None:
                address = utilities.get_env('VAULT_ADDR')
            __props__['address'] = address
            __props__['auth_logins'] = pulumi.Output.from_input(auth_logins).apply(json.dumps) if auth_logins is not None else None
            if ca_cert_dir is None:
                ca_cert_dir = utilities.get_env('VAULT_CAPATH')
            __props__['ca_cert_dir'] = ca_cert_dir
            if ca_cert_file is None:
                ca_cert_file = utilities.get_env('VAULT_CACERT')
            __props__['ca_cert_file'] = ca_cert_file
            __props__['client_auths'] = pulumi.Output.from_input(client_auths).apply(json.dumps) if client_auths is not None else None
            if max_lease_ttl_seconds is None:
                max_lease_ttl_seconds = (utilities.get_env_int('TERRAFORM_VAULT_MAX_TTL') or 20)
            __props__['max_lease_ttl_seconds'] = pulumi.Output.from_input(max_lease_ttl_seconds).apply(json.dumps) if max_lease_ttl_seconds is not None else None
            if max_retries is None:
                max_retries = (utilities.get_env_int('VAULT_MAX_RETRIES') or 2)
            __props__['max_retries'] = pulumi.Output.from_input(max_retries).apply(json.dumps) if max_retries is not None else None
            if namespace is None:
                namespace = utilities.get_env('VAULT_NAMESPACE')
            __props__['namespace'] = namespace
            if skip_tls_verify is None:
                skip_tls_verify = utilities.get_env_bool('VAULT_SKIP_VERIFY')
            __props__['skip_tls_verify'] = pulumi.Output.from_input(skip_tls_verify).apply(json.dumps) if skip_tls_verify is not None else None
            if token is None:
                token = utilities.get_env('VAULT_TOKEN')
            __props__['token'] = token
            if token_name is None:
                token_name = utilities.get_env('VAULT_TOKEN_NAME')
            __props__['token_name'] = token_name
        super(Provider, __self__).__init__(
            'vault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Provider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/index.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        return Provider(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

