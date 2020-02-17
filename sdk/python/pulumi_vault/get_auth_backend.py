# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetAuthBackendResult:
    """
    A collection of values returned by getAuthBackend.
    """
    def __init__(__self__, accessor=None, default_lease_ttl_seconds=None, description=None, id=None, listing_visibility=None, local=None, max_lease_ttl_seconds=None, path=None, type=None):
        if accessor and not isinstance(accessor, str):
            raise TypeError("Expected argument 'accessor' to be a str")
        __self__.accessor = accessor
        """
        The accessor for this auth method
        """
        if default_lease_ttl_seconds and not isinstance(default_lease_ttl_seconds, float):
            raise TypeError("Expected argument 'default_lease_ttl_seconds' to be a float")
        __self__.default_lease_ttl_seconds = default_lease_ttl_seconds
        """
        The default lease duration in seconds.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A description of the auth method.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if listing_visibility and not isinstance(listing_visibility, str):
            raise TypeError("Expected argument 'listing_visibility' to be a str")
        __self__.listing_visibility = listing_visibility
        """
        Speficies whether to show this mount in the UI-specific listing endpoint.
        """
        if local and not isinstance(local, bool):
            raise TypeError("Expected argument 'local' to be a bool")
        __self__.local = local
        """
        Specifies if the auth method is local only.
        """
        if max_lease_ttl_seconds and not isinstance(max_lease_ttl_seconds, float):
            raise TypeError("Expected argument 'max_lease_ttl_seconds' to be a float")
        __self__.max_lease_ttl_seconds = max_lease_ttl_seconds
        """
        The maximum lease duration in seconds.
        """
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The name of the auth method type.
        """
class AwaitableGetAuthBackendResult(GetAuthBackendResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthBackendResult(
            accessor=self.accessor,
            default_lease_ttl_seconds=self.default_lease_ttl_seconds,
            description=self.description,
            id=self.id,
            listing_visibility=self.listing_visibility,
            local=self.local,
            max_lease_ttl_seconds=self.max_lease_ttl_seconds,
            path=self.path,
            type=self.type)

def get_auth_backend(path=None,opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str path: The auth backend mount point.
    """
    __args__ = dict()


    __args__['path'] = path
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:index/getAuthBackend:getAuthBackend', __args__, opts=opts).value

    return AwaitableGetAuthBackendResult(
        accessor=__ret__.get('accessor'),
        default_lease_ttl_seconds=__ret__.get('defaultLeaseTtlSeconds'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        listing_visibility=__ret__.get('listingVisibility'),
        local=__ret__.get('local'),
        max_lease_ttl_seconds=__ret__.get('maxLeaseTtlSeconds'),
        path=__ret__.get('path'),
        type=__ret__.get('type'))
