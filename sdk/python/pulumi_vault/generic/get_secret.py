# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, data=None, data_json=None, id=None, lease_duration=None, lease_id=None, lease_renewable=None, lease_start_time=None, path=None, version=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        __self__.data = data
        if data_json and not isinstance(data_json, str):
            raise TypeError("Expected argument 'data_json' to be a str")
        __self__.data_json = data_json
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if lease_duration and not isinstance(lease_duration, float):
            raise TypeError("Expected argument 'lease_duration' to be a float")
        __self__.lease_duration = lease_duration
        if lease_id and not isinstance(lease_id, str):
            raise TypeError("Expected argument 'lease_id' to be a str")
        __self__.lease_id = lease_id
        if lease_renewable and not isinstance(lease_renewable, bool):
            raise TypeError("Expected argument 'lease_renewable' to be a bool")
        __self__.lease_renewable = lease_renewable
        if lease_start_time and not isinstance(lease_start_time, str):
            raise TypeError("Expected argument 'lease_start_time' to be a str")
        __self__.lease_start_time = lease_start_time
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        __self__.path = path
        if version and not isinstance(version, float):
            raise TypeError("Expected argument 'version' to be a float")
        __self__.version = version
class AwaitableGetSecretResult(GetSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretResult(
            data=self.data,
            data_json=self.data_json,
            id=self.id,
            lease_duration=self.lease_duration,
            lease_id=self.lease_id,
            lease_renewable=self.lease_renewable,
            lease_start_time=self.lease_start_time,
            path=self.path,
            version=self.version)

def get_secret(path=None,version=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['path'] = path
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:generic/getSecret:getSecret', __args__, opts=opts).value

    return AwaitableGetSecretResult(
        data=__ret__.get('data'),
        data_json=__ret__.get('dataJson'),
        id=__ret__.get('id'),
        lease_duration=__ret__.get('leaseDuration'),
        lease_id=__ret__.get('leaseId'),
        lease_renewable=__ret__.get('leaseRenewable'),
        lease_start_time=__ret__.get('leaseStartTime'),
        path=__ret__.get('path'),
        version=__ret__.get('version'))
