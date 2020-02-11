# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendRole(pulumi.CustomResource):
    allowed_policies: pulumi.Output[list]
    """
    List of allowed policies for given role.
    """
    bound_cidrs: pulumi.Output[list]
    """
    If set, a list of
    CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
    """
    disallowed_policies: pulumi.Output[list]
    """
    List of disallowed policies for given role.
    """
    explicit_max_ttl: pulumi.Output[str]
    """
    Number of seconds after which issued tokens can no longer be renewed.
    """
    orphan: pulumi.Output[bool]
    """
    If true, tokens created against this policy will be orphan tokens.
    """
    path_suffix: pulumi.Output[str]
    """
    Tokens created against this role will have the given suffix as part of their path in addition to the role name.
    """
    period: pulumi.Output[str]
    """
    If set, indicates that the
    token generated using this role should never expire. The token should be renewed within the
    duration specified by this value. At each renewal, the token's TTL will be set to the
    value of this field. Specified in seconds.
    """
    renewable: pulumi.Output[bool]
    """
    Wether to disable the ability of the token to be renewed past its initial TTL.
    """
    role_name: pulumi.Output[str]
    """
    The name of the role.
    """
    token_bound_cidrs: pulumi.Output[list]
    """
    List of CIDR blocks; if set, specifies blocks of IP
    addresses which can authenticate successfully, and ties the resulting token to these blocks
    as well.
    """
    token_explicit_max_ttl: pulumi.Output[float]
    """
    If set, will encode an
    [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
    onto the token in number of seconds. This is a hard cap even if `token_ttl` and
    `token_max_ttl` would otherwise allow a renewal.
    """
    token_max_ttl: pulumi.Output[float]
    """
    The maximum lifetime for generated tokens in number of seconds.
    Its current value will be referenced at renewal time.
    """
    token_no_default_policy: pulumi.Output[bool]
    """
    If set, the default policy will not be set on
    generated tokens; otherwise it will be added to the policies set in token_policies.
    """
    token_num_uses: pulumi.Output[float]
    """
    The
    [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
    if any, in number of seconds to set on the token.
    """
    token_period: pulumi.Output[float]
    """
    If set, indicates that the
    token generated using this role should never expire. The token should be renewed within the
    duration specified by this value. At each renewal, the token's TTL will be set to the
    value of this field. Specified in seconds.
    """
    token_policies: pulumi.Output[list]
    """
    Generated Token's Policies
    """
    token_ttl: pulumi.Output[float]
    """
    The incremental lifetime for generated tokens in number of seconds.
    Its current value will be referenced at renewal time.
    """
    token_type: pulumi.Output[str]
    """
    The type of token that should be generated. Can be `service`,
    `batch`, or `default` to use the mount's tuned default (which unless changed will be
    `service` tokens). For token store roles, there are two additional possibilities:
    `default-service` and `default-batch` which specify the type to return unless the client
    requests a different type at generation time.
    """
    def __init__(__self__, resource_name, opts=None, allowed_policies=None, bound_cidrs=None, disallowed_policies=None, explicit_max_ttl=None, orphan=None, path_suffix=None, period=None, renewable=None, role_name=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages Token auth backend role in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/token.html) for more
        information.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/token_auth_backend_role.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_policies: List of allowed policies for given role.
        :param pulumi.Input[list] bound_cidrs: If set, a list of
               CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        :param pulumi.Input[list] disallowed_policies: List of disallowed policies for given role.
        :param pulumi.Input[str] explicit_max_ttl: Number of seconds after which issued tokens can no longer be renewed.
        :param pulumi.Input[bool] orphan: If true, tokens created against this policy will be orphan tokens.
        :param pulumi.Input[str] path_suffix: Tokens created against this role will have the given suffix as part of their path in addition to the role name.
        :param pulumi.Input[str] period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[bool] renewable: Wether to disable the ability of the token to be renewed past its initial TTL.
        :param pulumi.Input[str] role_name: The name of the role.
        :param pulumi.Input[list] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
               addresses which can authenticate successfully, and ties the resulting token to these blocks
               as well.
        :param pulumi.Input[float] token_explicit_max_ttl: If set, will encode an
               [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
               onto the token in number of seconds. This is a hard cap even if `token_ttl` and
               `token_max_ttl` would otherwise allow a renewal.
        :param pulumi.Input[float] token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[bool] token_no_default_policy: If set, the default policy will not be set on
               generated tokens; otherwise it will be added to the policies set in token_policies.
        :param pulumi.Input[float] token_num_uses: The
               [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
               if any, in number of seconds to set on the token.
        :param pulumi.Input[float] token_period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[list] token_policies: Generated Token's Policies
        :param pulumi.Input[float] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
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

            __props__['allowed_policies'] = allowed_policies
            __props__['bound_cidrs'] = bound_cidrs
            __props__['disallowed_policies'] = disallowed_policies
            __props__['explicit_max_ttl'] = explicit_max_ttl
            __props__['orphan'] = orphan
            __props__['path_suffix'] = path_suffix
            __props__['period'] = period
            __props__['renewable'] = renewable
            if role_name is None:
                raise TypeError("Missing required property 'role_name'")
            __props__['role_name'] = role_name
            __props__['token_bound_cidrs'] = token_bound_cidrs
            __props__['token_explicit_max_ttl'] = token_explicit_max_ttl
            __props__['token_max_ttl'] = token_max_ttl
            __props__['token_no_default_policy'] = token_no_default_policy
            __props__['token_num_uses'] = token_num_uses
            __props__['token_period'] = token_period
            __props__['token_policies'] = token_policies
            __props__['token_ttl'] = token_ttl
            __props__['token_type'] = token_type
        super(AuthBackendRole, __self__).__init__(
            'vault:tokenauth/authBackendRole:AuthBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allowed_policies=None, bound_cidrs=None, disallowed_policies=None, explicit_max_ttl=None, orphan=None, path_suffix=None, period=None, renewable=None, role_name=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None):
        """
        Get an existing AuthBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_policies: List of allowed policies for given role.
        :param pulumi.Input[list] bound_cidrs: If set, a list of
               CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        :param pulumi.Input[list] disallowed_policies: List of disallowed policies for given role.
        :param pulumi.Input[str] explicit_max_ttl: Number of seconds after which issued tokens can no longer be renewed.
        :param pulumi.Input[bool] orphan: If true, tokens created against this policy will be orphan tokens.
        :param pulumi.Input[str] path_suffix: Tokens created against this role will have the given suffix as part of their path in addition to the role name.
        :param pulumi.Input[str] period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[bool] renewable: Wether to disable the ability of the token to be renewed past its initial TTL.
        :param pulumi.Input[str] role_name: The name of the role.
        :param pulumi.Input[list] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
               addresses which can authenticate successfully, and ties the resulting token to these blocks
               as well.
        :param pulumi.Input[float] token_explicit_max_ttl: If set, will encode an
               [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
               onto the token in number of seconds. This is a hard cap even if `token_ttl` and
               `token_max_ttl` would otherwise allow a renewal.
        :param pulumi.Input[float] token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[bool] token_no_default_policy: If set, the default policy will not be set on
               generated tokens; otherwise it will be added to the policies set in token_policies.
        :param pulumi.Input[float] token_num_uses: The
               [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
               if any, in number of seconds to set on the token.
        :param pulumi.Input[float] token_period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[list] token_policies: Generated Token's Policies
        :param pulumi.Input[float] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_policies"] = allowed_policies
        __props__["bound_cidrs"] = bound_cidrs
        __props__["disallowed_policies"] = disallowed_policies
        __props__["explicit_max_ttl"] = explicit_max_ttl
        __props__["orphan"] = orphan
        __props__["path_suffix"] = path_suffix
        __props__["period"] = period
        __props__["renewable"] = renewable
        __props__["role_name"] = role_name
        __props__["token_bound_cidrs"] = token_bound_cidrs
        __props__["token_explicit_max_ttl"] = token_explicit_max_ttl
        __props__["token_max_ttl"] = token_max_ttl
        __props__["token_no_default_policy"] = token_no_default_policy
        __props__["token_num_uses"] = token_num_uses
        __props__["token_period"] = token_period
        __props__["token_policies"] = token_policies
        __props__["token_ttl"] = token_ttl
        __props__["token_type"] = token_type
        return AuthBackendRole(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

