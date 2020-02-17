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
    The accessor of the JWT auth backend
    """
    bound_issuer: pulumi.Output[str]
    """
    The value against which to match the iss claim in a JWT
    """
    default_role: pulumi.Output[str]
    """
    The default role to use if none is provided during login
    """
    description: pulumi.Output[str]
    """
    The description of the auth backend
    """
    jwks_ca_pem: pulumi.Output[str]
    """
    The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
    """
    jwks_url: pulumi.Output[str]
    """
    JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
    """
    jwt_supported_algs: pulumi.Output[list]
    """
    A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
    """
    jwt_validation_pubkeys: pulumi.Output[list]
    """
    A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
    """
    oidc_client_id: pulumi.Output[str]
    """
    Client ID used for OIDC backends
    """
    oidc_client_secret: pulumi.Output[str]
    """
    Client Secret used for OIDC backends
    """
    oidc_discovery_ca_pem: pulumi.Output[str]
    """
    The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
    """
    oidc_discovery_url: pulumi.Output[str]
    """
    The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
    """
    path: pulumi.Output[str]
    """
    Path to mount the JWT/OIDC auth backend
    """
    tune: pulumi.Output[dict]
    type: pulumi.Output[str]
    """
    Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
    """
    def __init__(__self__, resource_name, opts=None, bound_issuer=None, default_role=None, description=None, jwks_ca_pem=None, jwks_url=None, jwt_supported_algs=None, jwt_validation_pubkeys=None, oidc_client_id=None, oidc_client_secret=None, oidc_discovery_ca_pem=None, oidc_discovery_url=None, path=None, tune=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource for managing an
        [JWT auth backend within Vault](https://www.vaultproject.io/docs/auth/jwt.html).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/jwt_auth_backend.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bound_issuer: The value against which to match the iss claim in a JWT
        :param pulumi.Input[str] default_role: The default role to use if none is provided during login
        :param pulumi.Input[str] description: The description of the auth backend
        :param pulumi.Input[str] jwks_ca_pem: The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
        :param pulumi.Input[str] jwks_url: JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
        :param pulumi.Input[list] jwt_supported_algs: A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
        :param pulumi.Input[list] jwt_validation_pubkeys: A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
        :param pulumi.Input[str] oidc_client_id: Client ID used for OIDC backends
        :param pulumi.Input[str] oidc_client_secret: Client Secret used for OIDC backends
        :param pulumi.Input[str] oidc_discovery_ca_pem: The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
        :param pulumi.Input[str] oidc_discovery_url: The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
        :param pulumi.Input[str] path: Path to mount the JWT/OIDC auth backend
        :param pulumi.Input[str] type: Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`

        The **tune** object supports the following:

          * `allowedResponseHeaders` (`pulumi.Input[list]`)
          * `auditNonHmacRequestKeys` (`pulumi.Input[list]`)
          * `auditNonHmacResponseKeys` (`pulumi.Input[list]`)
          * `defaultLeaseTtl` (`pulumi.Input[str]`)
          * `listing_visibility` (`pulumi.Input[str]`)
          * `maxLeaseTtl` (`pulumi.Input[str]`)
          * `passthroughRequestHeaders` (`pulumi.Input[list]`)
          * `token_type` (`pulumi.Input[str]`)
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

            __props__['bound_issuer'] = bound_issuer
            __props__['default_role'] = default_role
            __props__['description'] = description
            __props__['jwks_ca_pem'] = jwks_ca_pem
            __props__['jwks_url'] = jwks_url
            __props__['jwt_supported_algs'] = jwt_supported_algs
            __props__['jwt_validation_pubkeys'] = jwt_validation_pubkeys
            __props__['oidc_client_id'] = oidc_client_id
            __props__['oidc_client_secret'] = oidc_client_secret
            __props__['oidc_discovery_ca_pem'] = oidc_discovery_ca_pem
            __props__['oidc_discovery_url'] = oidc_discovery_url
            __props__['path'] = path
            __props__['tune'] = tune
            __props__['type'] = type
            __props__['accessor'] = None
        super(AuthBackend, __self__).__init__(
            'vault:jwt/authBackend:AuthBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessor=None, bound_issuer=None, default_role=None, description=None, jwks_ca_pem=None, jwks_url=None, jwt_supported_algs=None, jwt_validation_pubkeys=None, oidc_client_id=None, oidc_client_secret=None, oidc_discovery_ca_pem=None, oidc_discovery_url=None, path=None, tune=None, type=None):
        """
        Get an existing AuthBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The accessor of the JWT auth backend
        :param pulumi.Input[str] bound_issuer: The value against which to match the iss claim in a JWT
        :param pulumi.Input[str] default_role: The default role to use if none is provided during login
        :param pulumi.Input[str] description: The description of the auth backend
        :param pulumi.Input[str] jwks_ca_pem: The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
        :param pulumi.Input[str] jwks_url: JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
        :param pulumi.Input[list] jwt_supported_algs: A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
        :param pulumi.Input[list] jwt_validation_pubkeys: A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
        :param pulumi.Input[str] oidc_client_id: Client ID used for OIDC backends
        :param pulumi.Input[str] oidc_client_secret: Client Secret used for OIDC backends
        :param pulumi.Input[str] oidc_discovery_ca_pem: The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
        :param pulumi.Input[str] oidc_discovery_url: The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
        :param pulumi.Input[str] path: Path to mount the JWT/OIDC auth backend
        :param pulumi.Input[str] type: Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`

        The **tune** object supports the following:

          * `allowedResponseHeaders` (`pulumi.Input[list]`)
          * `auditNonHmacRequestKeys` (`pulumi.Input[list]`)
          * `auditNonHmacResponseKeys` (`pulumi.Input[list]`)
          * `defaultLeaseTtl` (`pulumi.Input[str]`)
          * `listing_visibility` (`pulumi.Input[str]`)
          * `maxLeaseTtl` (`pulumi.Input[str]`)
          * `passthroughRequestHeaders` (`pulumi.Input[list]`)
          * `token_type` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessor"] = accessor
        __props__["bound_issuer"] = bound_issuer
        __props__["default_role"] = default_role
        __props__["description"] = description
        __props__["jwks_ca_pem"] = jwks_ca_pem
        __props__["jwks_url"] = jwks_url
        __props__["jwt_supported_algs"] = jwt_supported_algs
        __props__["jwt_validation_pubkeys"] = jwt_validation_pubkeys
        __props__["oidc_client_id"] = oidc_client_id
        __props__["oidc_client_secret"] = oidc_client_secret
        __props__["oidc_discovery_ca_pem"] = oidc_discovery_ca_pem
        __props__["oidc_discovery_url"] = oidc_discovery_url
        __props__["path"] = path
        __props__["tune"] = tune
        __props__["type"] = type
        return AuthBackend(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

