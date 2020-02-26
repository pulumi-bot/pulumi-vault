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
    The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
    """
    base_url: pulumi.Output[str]
    """
    The Okta url. Examples: oktapreview.com, okta.com
    """
    bypass_okta_mfa: pulumi.Output[bool]
    """
    When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
    """
    description: pulumi.Output[str]
    """
    The description of the auth backend
    """
    groups: pulumi.Output[list]
    """
    Associate Okta groups with policies within Vault.
    See below for more details.

      * `group_name` (`str`) - Name of the group within the Okta
      * `policies` (`list`) - List of Vault policies to associate with this user
    """
    max_ttl: pulumi.Output[str]
    """
    Maximum duration after which authentication will be expired
    [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
    """
    organization: pulumi.Output[str]
    """
    The Okta organization. This will be the first part of the url `https://XXX.okta.com`
    """
    path: pulumi.Output[str]
    """
    Path to mount the Okta auth backend
    """
    token: pulumi.Output[str]
    """
    The Okta API token. This is required to query Okta for user group membership.
    If this is not supplied only locally configured groups will be enabled.
    """
    ttl: pulumi.Output[str]
    """
    Duration after which authentication will be expired.
    [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
    """
    users: pulumi.Output[list]
    """
    Associate Okta users with groups or policies within Vault.
    See below for more details.

      * `groups` (`list`) - List of Okta groups to associate with this user
      * `policies` (`list`) - List of Vault policies to associate with this user
      * `username` (`str`) - Name of the user within Okta
    """
    def __init__(__self__, resource_name, opts=None, base_url=None, bypass_okta_mfa=None, description=None, groups=None, max_ttl=None, organization=None, path=None, token=None, ttl=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource for managing an
        [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base_url: The Okta url. Examples: oktapreview.com, okta.com
        :param pulumi.Input[bool] bypass_okta_mfa: When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
        :param pulumi.Input[str] description: The description of the auth backend
        :param pulumi.Input[list] groups: Associate Okta groups with policies within Vault.
               See below for more details.
        :param pulumi.Input[str] max_ttl: Maximum duration after which authentication will be expired
               [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
        :param pulumi.Input[str] organization: The Okta organization. This will be the first part of the url `https://XXX.okta.com`
        :param pulumi.Input[str] path: Path to mount the Okta auth backend
        :param pulumi.Input[str] token: The Okta API token. This is required to query Okta for user group membership.
               If this is not supplied only locally configured groups will be enabled.
        :param pulumi.Input[str] ttl: Duration after which authentication will be expired.
               [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
        :param pulumi.Input[list] users: Associate Okta users with groups or policies within Vault.
               See below for more details.

        The **groups** object supports the following:

          * `group_name` (`pulumi.Input[str]`) - Name of the group within the Okta
          * `policies` (`pulumi.Input[list]`) - List of Vault policies to associate with this user

        The **users** object supports the following:

          * `groups` (`pulumi.Input[list]`) - List of Okta groups to associate with this user
          * `policies` (`pulumi.Input[list]`) - List of Vault policies to associate with this user
          * `username` (`pulumi.Input[str]`) - Name of the user within Okta
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

            __props__['base_url'] = base_url
            __props__['bypass_okta_mfa'] = bypass_okta_mfa
            __props__['description'] = description
            __props__['groups'] = groups
            __props__['max_ttl'] = max_ttl
            if organization is None:
                raise TypeError("Missing required property 'organization'")
            __props__['organization'] = organization
            __props__['path'] = path
            __props__['token'] = token
            __props__['ttl'] = ttl
            __props__['users'] = users
            __props__['accessor'] = None
        super(AuthBackend, __self__).__init__(
            'vault:okta/authBackend:AuthBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessor=None, base_url=None, bypass_okta_mfa=None, description=None, groups=None, max_ttl=None, organization=None, path=None, token=None, ttl=None, users=None):
        """
        Get an existing AuthBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
        :param pulumi.Input[str] base_url: The Okta url. Examples: oktapreview.com, okta.com
        :param pulumi.Input[bool] bypass_okta_mfa: When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
        :param pulumi.Input[str] description: The description of the auth backend
        :param pulumi.Input[list] groups: Associate Okta groups with policies within Vault.
               See below for more details.
        :param pulumi.Input[str] max_ttl: Maximum duration after which authentication will be expired
               [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
        :param pulumi.Input[str] organization: The Okta organization. This will be the first part of the url `https://XXX.okta.com`
        :param pulumi.Input[str] path: Path to mount the Okta auth backend
        :param pulumi.Input[str] token: The Okta API token. This is required to query Okta for user group membership.
               If this is not supplied only locally configured groups will be enabled.
        :param pulumi.Input[str] ttl: Duration after which authentication will be expired.
               [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
        :param pulumi.Input[list] users: Associate Okta users with groups or policies within Vault.
               See below for more details.

        The **groups** object supports the following:

          * `group_name` (`pulumi.Input[str]`) - Name of the group within the Okta
          * `policies` (`pulumi.Input[list]`) - List of Vault policies to associate with this user

        The **users** object supports the following:

          * `groups` (`pulumi.Input[list]`) - List of Okta groups to associate with this user
          * `policies` (`pulumi.Input[list]`) - List of Vault policies to associate with this user
          * `username` (`pulumi.Input[str]`) - Name of the user within Okta
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessor"] = accessor
        __props__["base_url"] = base_url
        __props__["bypass_okta_mfa"] = bypass_okta_mfa
        __props__["description"] = description
        __props__["groups"] = groups
        __props__["max_ttl"] = max_ttl
        __props__["organization"] = organization
        __props__["path"] = path
        __props__["token"] = token
        __props__["ttl"] = ttl
        __props__["users"] = users
        return AuthBackend(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

