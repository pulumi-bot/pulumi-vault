# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackend(pulumi.CustomResource):
    accessor: pulumi.Output[str]
    """
    The accessor for this auth mount.
    """
    binddn: pulumi.Output[str]
    """
    DN of object to bind when performing user search
    """
    bindpass: pulumi.Output[str]
    """
    Password to use with `binddn` when performing user search
    """
    certificate: pulumi.Output[str]
    """
    Trusted CA to validate TLS certificate
    """
    deny_null_bind: pulumi.Output[bool]
    description: pulumi.Output[str]
    """
    Description for the LDAP auth backend mount
    """
    discoverdn: pulumi.Output[bool]
    groupattr: pulumi.Output[str]
    """
    LDAP attribute to follow on objects returned by groupfilter
    """
    groupdn: pulumi.Output[str]
    """
    Base DN under which to perform group search
    """
    groupfilter: pulumi.Output[str]
    """
    Go template used to construct group membership query
    """
    insecure_tls: pulumi.Output[bool]
    """
    Control whether or TLS certificates must be validated
    """
    path: pulumi.Output[str]
    """
    Path to mount the LDAP auth backend under
    """
    starttls: pulumi.Output[bool]
    """
    Control use of TLS when conecting to LDAP
    """
    tls_max_version: pulumi.Output[str]
    """
    Maximum acceptable version of TLS
    """
    tls_min_version: pulumi.Output[str]
    """
    Minimum acceptable version of TLS
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
    List of policies to encode onto generated tokens. Depending
    on the auth method, this list may be supplemented by user/group/other values.
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
    upndomain: pulumi.Output[str]
    """
    The userPrincipalDomain used to construct UPN string
    """
    url: pulumi.Output[str]
    """
    The URL of the LDAP server
    """
    use_token_groups: pulumi.Output[bool]
    """
    Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
    """
    userattr: pulumi.Output[str]
    """
    Attribute on user object matching username passed in
    """
    userdn: pulumi.Output[str]
    """
    Base DN under which to perform user search
    """
    def __init__(__self__, resource_name, opts=None, binddn=None, bindpass=None, certificate=None, deny_null_bind=None, description=None, discoverdn=None, groupattr=None, groupdn=None, groupfilter=None, insecure_tls=None, path=None, starttls=None, tls_max_version=None, tls_min_version=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, upndomain=None, url=None, use_token_groups=None, userattr=None, userdn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource for managing an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ldap_auth_backend.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] binddn: DN of object to bind when performing user search
        :param pulumi.Input[str] bindpass: Password to use with `binddn` when performing user search
        :param pulumi.Input[str] certificate: Trusted CA to validate TLS certificate
        :param pulumi.Input[str] description: Description for the LDAP auth backend mount
        :param pulumi.Input[str] groupattr: LDAP attribute to follow on objects returned by groupfilter
        :param pulumi.Input[str] groupdn: Base DN under which to perform group search
        :param pulumi.Input[str] groupfilter: Go template used to construct group membership query
        :param pulumi.Input[bool] insecure_tls: Control whether or TLS certificates must be validated
        :param pulumi.Input[str] path: Path to mount the LDAP auth backend under
        :param pulumi.Input[bool] starttls: Control use of TLS when conecting to LDAP
        :param pulumi.Input[str] tls_max_version: Maximum acceptable version of TLS
        :param pulumi.Input[str] tls_min_version: Minimum acceptable version of TLS
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
        :param pulumi.Input[list] token_policies: List of policies to encode onto generated tokens. Depending
               on the auth method, this list may be supplemented by user/group/other values.
        :param pulumi.Input[float] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        :param pulumi.Input[str] upndomain: The userPrincipalDomain used to construct UPN string
        :param pulumi.Input[str] url: The URL of the LDAP server
        :param pulumi.Input[bool] use_token_groups: Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
        :param pulumi.Input[str] userattr: Attribute on user object matching username passed in
        :param pulumi.Input[str] userdn: Base DN under which to perform user search
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

            __props__['binddn'] = binddn
            __props__['bindpass'] = bindpass
            __props__['certificate'] = certificate
            __props__['deny_null_bind'] = deny_null_bind
            __props__['description'] = description
            __props__['discoverdn'] = discoverdn
            __props__['groupattr'] = groupattr
            __props__['groupdn'] = groupdn
            __props__['groupfilter'] = groupfilter
            __props__['insecure_tls'] = insecure_tls
            __props__['path'] = path
            __props__['starttls'] = starttls
            __props__['tls_max_version'] = tls_max_version
            __props__['tls_min_version'] = tls_min_version
            __props__['token_bound_cidrs'] = token_bound_cidrs
            __props__['token_explicit_max_ttl'] = token_explicit_max_ttl
            __props__['token_max_ttl'] = token_max_ttl
            __props__['token_no_default_policy'] = token_no_default_policy
            __props__['token_num_uses'] = token_num_uses
            __props__['token_period'] = token_period
            __props__['token_policies'] = token_policies
            __props__['token_ttl'] = token_ttl
            __props__['token_type'] = token_type
            __props__['upndomain'] = upndomain
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['use_token_groups'] = use_token_groups
            __props__['userattr'] = userattr
            __props__['userdn'] = userdn
            __props__['accessor'] = None
        super(AuthBackend, __self__).__init__(
            'vault:ldap/authBackend:AuthBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessor=None, binddn=None, bindpass=None, certificate=None, deny_null_bind=None, description=None, discoverdn=None, groupattr=None, groupdn=None, groupfilter=None, insecure_tls=None, path=None, starttls=None, tls_max_version=None, tls_min_version=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, upndomain=None, url=None, use_token_groups=None, userattr=None, userdn=None):
        """
        Get an existing AuthBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The accessor for this auth mount.
        :param pulumi.Input[str] binddn: DN of object to bind when performing user search
        :param pulumi.Input[str] bindpass: Password to use with `binddn` when performing user search
        :param pulumi.Input[str] certificate: Trusted CA to validate TLS certificate
        :param pulumi.Input[str] description: Description for the LDAP auth backend mount
        :param pulumi.Input[str] groupattr: LDAP attribute to follow on objects returned by groupfilter
        :param pulumi.Input[str] groupdn: Base DN under which to perform group search
        :param pulumi.Input[str] groupfilter: Go template used to construct group membership query
        :param pulumi.Input[bool] insecure_tls: Control whether or TLS certificates must be validated
        :param pulumi.Input[str] path: Path to mount the LDAP auth backend under
        :param pulumi.Input[bool] starttls: Control use of TLS when conecting to LDAP
        :param pulumi.Input[str] tls_max_version: Maximum acceptable version of TLS
        :param pulumi.Input[str] tls_min_version: Minimum acceptable version of TLS
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
        :param pulumi.Input[list] token_policies: List of policies to encode onto generated tokens. Depending
               on the auth method, this list may be supplemented by user/group/other values.
        :param pulumi.Input[float] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        :param pulumi.Input[str] upndomain: The userPrincipalDomain used to construct UPN string
        :param pulumi.Input[str] url: The URL of the LDAP server
        :param pulumi.Input[bool] use_token_groups: Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
        :param pulumi.Input[str] userattr: Attribute on user object matching username passed in
        :param pulumi.Input[str] userdn: Base DN under which to perform user search
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessor"] = accessor
        __props__["binddn"] = binddn
        __props__["bindpass"] = bindpass
        __props__["certificate"] = certificate
        __props__["deny_null_bind"] = deny_null_bind
        __props__["description"] = description
        __props__["discoverdn"] = discoverdn
        __props__["groupattr"] = groupattr
        __props__["groupdn"] = groupdn
        __props__["groupfilter"] = groupfilter
        __props__["insecure_tls"] = insecure_tls
        __props__["path"] = path
        __props__["starttls"] = starttls
        __props__["tls_max_version"] = tls_max_version
        __props__["tls_min_version"] = tls_min_version
        __props__["token_bound_cidrs"] = token_bound_cidrs
        __props__["token_explicit_max_ttl"] = token_explicit_max_ttl
        __props__["token_max_ttl"] = token_max_ttl
        __props__["token_no_default_policy"] = token_no_default_policy
        __props__["token_num_uses"] = token_num_uses
        __props__["token_period"] = token_period
        __props__["token_policies"] = token_policies
        __props__["token_ttl"] = token_ttl
        __props__["token_type"] = token_type
        __props__["upndomain"] = upndomain
        __props__["url"] = url
        __props__["use_token_groups"] = use_token_groups
        __props__["userattr"] = userattr
        __props__["userdn"] = userdn
        return AuthBackend(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

