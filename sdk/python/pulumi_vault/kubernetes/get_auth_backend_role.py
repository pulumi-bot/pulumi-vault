# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAuthBackendRoleResult:
    """
    A collection of values returned by getAuthBackendRole.
    """
    def __init__(__self__, audience=None, backend=None, bound_cidrs=None, bound_service_account_names=None, bound_service_account_namespaces=None, id=None, max_ttl=None, num_uses=None, period=None, policies=None, role_name=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, ttl=None):
        if audience and not isinstance(audience, str):
            raise TypeError("Expected argument 'audience' to be a str")
        __self__.audience = audience
        """
        (Optional) Audience claim to verify in the JWT.
        """
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        __self__.backend = backend
        if bound_cidrs and not isinstance(bound_cidrs, list):
            raise TypeError("Expected argument 'bound_cidrs' to be a list")
        __self__.bound_cidrs = bound_cidrs
        if bound_service_account_names and not isinstance(bound_service_account_names, list):
            raise TypeError("Expected argument 'bound_service_account_names' to be a list")
        __self__.bound_service_account_names = bound_service_account_names
        """
        List of service account names able to access this role. If set to "*" all names are allowed, both this and bound_service_account_namespaces can not be "*".
        """
        if bound_service_account_namespaces and not isinstance(bound_service_account_namespaces, list):
            raise TypeError("Expected argument 'bound_service_account_namespaces' to be a list")
        __self__.bound_service_account_namespaces = bound_service_account_namespaces
        """
        List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if max_ttl and not isinstance(max_ttl, float):
            raise TypeError("Expected argument 'max_ttl' to be a float")
        __self__.max_ttl = max_ttl
        if num_uses and not isinstance(num_uses, float):
            raise TypeError("Expected argument 'num_uses' to be a float")
        __self__.num_uses = num_uses
        if period and not isinstance(period, float):
            raise TypeError("Expected argument 'period' to be a float")
        __self__.period = period
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        __self__.policies = policies
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        __self__.role_name = role_name
        if token_bound_cidrs and not isinstance(token_bound_cidrs, list):
            raise TypeError("Expected argument 'token_bound_cidrs' to be a list")
        __self__.token_bound_cidrs = token_bound_cidrs
        """
        List of CIDR blocks; if set, specifies blocks of IP
        addresses which can authenticate successfully, and ties the resulting token to these blocks
        as well.
        """
        if token_explicit_max_ttl and not isinstance(token_explicit_max_ttl, float):
            raise TypeError("Expected argument 'token_explicit_max_ttl' to be a float")
        __self__.token_explicit_max_ttl = token_explicit_max_ttl
        """
        If set, will encode an
        [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        `token_max_ttl` would otherwise allow a renewal.
        """
        if token_max_ttl and not isinstance(token_max_ttl, float):
            raise TypeError("Expected argument 'token_max_ttl' to be a float")
        __self__.token_max_ttl = token_max_ttl
        """
        The maximum lifetime for generated tokens in number of seconds.
        Its current value will be referenced at renewal time.
        """
        if token_no_default_policy and not isinstance(token_no_default_policy, bool):
            raise TypeError("Expected argument 'token_no_default_policy' to be a bool")
        __self__.token_no_default_policy = token_no_default_policy
        """
        If set, the default policy will not be set on
        generated tokens; otherwise it will be added to the policies set in token_policies.
        """
        if token_num_uses and not isinstance(token_num_uses, float):
            raise TypeError("Expected argument 'token_num_uses' to be a float")
        __self__.token_num_uses = token_num_uses
        """
        The
        [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        if any, in number of seconds to set on the token.
        """
        if token_period and not isinstance(token_period, float):
            raise TypeError("Expected argument 'token_period' to be a float")
        __self__.token_period = token_period
        """
        (Optional) If set, indicates that the
        token generated using this role should never expire. The token should be renewed within the
        duration specified by this value. At each renewal, the token's TTL will be set to the
        value of this field. Specified in seconds.
        """
        if token_policies and not isinstance(token_policies, list):
            raise TypeError("Expected argument 'token_policies' to be a list")
        __self__.token_policies = token_policies
        """
        List of policies to encode onto generated tokens. Depending
        on the auth method, this list may be supplemented by user/group/other values.
        """
        if token_ttl and not isinstance(token_ttl, float):
            raise TypeError("Expected argument 'token_ttl' to be a float")
        __self__.token_ttl = token_ttl
        """
        The incremental lifetime for generated tokens in number of seconds.
        Its current value will be referenced at renewal time.
        """
        if token_type and not isinstance(token_type, str):
            raise TypeError("Expected argument 'token_type' to be a str")
        __self__.token_type = token_type
        """
        The type of token that should be generated. Can be `service`,
        `batch`, or `default` to use the mount's tuned default (which unless changed will be
        `service` tokens). For token store roles, there are two additional possibilities:
        `default-service` and `default-batch` which specify the type to return unless the client
        requests a different type at generation time.
        """
        if ttl and not isinstance(ttl, float):
            raise TypeError("Expected argument 'ttl' to be a float")
        __self__.ttl = ttl
class AwaitableGetAuthBackendRoleResult(GetAuthBackendRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthBackendRoleResult(
            audience=self.audience,
            backend=self.backend,
            bound_cidrs=self.bound_cidrs,
            bound_service_account_names=self.bound_service_account_names,
            bound_service_account_namespaces=self.bound_service_account_namespaces,
            id=self.id,
            max_ttl=self.max_ttl,
            num_uses=self.num_uses,
            period=self.period,
            policies=self.policies,
            role_name=self.role_name,
            token_bound_cidrs=self.token_bound_cidrs,
            token_explicit_max_ttl=self.token_explicit_max_ttl,
            token_max_ttl=self.token_max_ttl,
            token_no_default_policy=self.token_no_default_policy,
            token_num_uses=self.token_num_uses,
            token_period=self.token_period,
            token_policies=self.token_policies,
            token_ttl=self.token_ttl,
            token_type=self.token_type,
            ttl=self.ttl)

def get_auth_backend_role(audience=None,backend=None,bound_cidrs=None,max_ttl=None,num_uses=None,period=None,policies=None,role_name=None,token_bound_cidrs=None,token_explicit_max_ttl=None,token_max_ttl=None,token_no_default_policy=None,token_num_uses=None,token_period=None,token_policies=None,token_ttl=None,token_type=None,ttl=None,opts=None):
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-role) for more
    information.




    :param str audience: (Optional) Audience claim to verify in the JWT.
    :param str backend: The unique name for the Kubernetes backend the role to
           retrieve Role attributes for resides in. Defaults to "kubernetes".
    :param str role_name: The name of the role to retrieve the Role attributes for.
    :param list token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
           addresses which can authenticate successfully, and ties the resulting token to these blocks
           as well.
    :param float token_explicit_max_ttl: If set, will encode an
           [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
           onto the token in number of seconds. This is a hard cap even if `token_ttl` and
           `token_max_ttl` would otherwise allow a renewal.
    :param float token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
           Its current value will be referenced at renewal time.
    :param bool token_no_default_policy: If set, the default policy will not be set on
           generated tokens; otherwise it will be added to the policies set in token_policies.
    :param float token_num_uses: The
           [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
           if any, in number of seconds to set on the token.
    :param float token_period: (Optional) If set, indicates that the
           token generated using this role should never expire. The token should be renewed within the
           duration specified by this value. At each renewal, the token's TTL will be set to the
           value of this field. Specified in seconds.
    :param list token_policies: List of policies to encode onto generated tokens. Depending
           on the auth method, this list may be supplemented by user/group/other values.
    :param float token_ttl: The incremental lifetime for generated tokens in number of seconds.
           Its current value will be referenced at renewal time.
    :param str token_type: The type of token that should be generated. Can be `service`,
           `batch`, or `default` to use the mount's tuned default (which unless changed will be
           `service` tokens). For token store roles, there are two additional possibilities:
           `default-service` and `default-batch` which specify the type to return unless the client
           requests a different type at generation time.
    """
    __args__ = dict()


    __args__['audience'] = audience
    __args__['backend'] = backend
    __args__['boundCidrs'] = bound_cidrs
    __args__['maxTtl'] = max_ttl
    __args__['numUses'] = num_uses
    __args__['period'] = period
    __args__['policies'] = policies
    __args__['roleName'] = role_name
    __args__['tokenBoundCidrs'] = token_bound_cidrs
    __args__['tokenExplicitMaxTtl'] = token_explicit_max_ttl
    __args__['tokenMaxTtl'] = token_max_ttl
    __args__['tokenNoDefaultPolicy'] = token_no_default_policy
    __args__['tokenNumUses'] = token_num_uses
    __args__['tokenPeriod'] = token_period
    __args__['tokenPolicies'] = token_policies
    __args__['tokenTtl'] = token_ttl
    __args__['tokenType'] = token_type
    __args__['ttl'] = ttl
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:kubernetes/getAuthBackendRole:getAuthBackendRole', __args__, opts=opts).value

    return AwaitableGetAuthBackendRoleResult(
        audience=__ret__.get('audience'),
        backend=__ret__.get('backend'),
        bound_cidrs=__ret__.get('boundCidrs'),
        bound_service_account_names=__ret__.get('boundServiceAccountNames'),
        bound_service_account_namespaces=__ret__.get('boundServiceAccountNamespaces'),
        id=__ret__.get('id'),
        max_ttl=__ret__.get('maxTtl'),
        num_uses=__ret__.get('numUses'),
        period=__ret__.get('period'),
        policies=__ret__.get('policies'),
        role_name=__ret__.get('roleName'),
        token_bound_cidrs=__ret__.get('tokenBoundCidrs'),
        token_explicit_max_ttl=__ret__.get('tokenExplicitMaxTtl'),
        token_max_ttl=__ret__.get('tokenMaxTtl'),
        token_no_default_policy=__ret__.get('tokenNoDefaultPolicy'),
        token_num_uses=__ret__.get('tokenNumUses'),
        token_period=__ret__.get('tokenPeriod'),
        token_policies=__ret__.get('tokenPolicies'),
        token_ttl=__ret__.get('tokenTtl'),
        token_type=__ret__.get('tokenType'),
        ttl=__ret__.get('ttl'))
