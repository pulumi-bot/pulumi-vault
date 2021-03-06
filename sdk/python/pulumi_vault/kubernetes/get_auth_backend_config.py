# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetAuthBackendConfigResult:
    """
    A collection of values returned by getAuthBackendConfig.
    """
    def __init__(__self__, backend=None, kubernetes_ca_cert=None, kubernetes_host=None, pem_keys=None, id=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        __self__.backend = backend
        if kubernetes_ca_cert and not isinstance(kubernetes_ca_cert, str):
            raise TypeError("Expected argument 'kubernetes_ca_cert' to be a str")
        __self__.kubernetes_ca_cert = kubernetes_ca_cert
        """
        PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        """
        if kubernetes_host and not isinstance(kubernetes_host, str):
            raise TypeError("Expected argument 'kubernetes_host' to be a str")
        __self__.kubernetes_host = kubernetes_host
        """
        Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        """
        if pem_keys and not isinstance(pem_keys, list):
            raise TypeError("Expected argument 'pem_keys' to be a list")
        __self__.pem_keys = pem_keys
        """
        Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetAuthBackendConfigResult(GetAuthBackendConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthBackendConfigResult(
            backend=self.backend,
            kubernetes_ca_cert=self.kubernetes_ca_cert,
            kubernetes_host=self.kubernetes_host,
            pem_keys=self.pem_keys,
            id=self.id)

def get_auth_backend_config(backend=None,kubernetes_ca_cert=None,kubernetes_host=None,pem_keys=None,opts=None):
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-config) for more
    information.
    
    :param str backend: The unique name for the Kubernetes backend the config to
           retrieve Role attributes for resides in. Defaults to "kubernetes".

    > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/kubernetes_auth_backend_config.html.markdown.
    """
    __args__ = dict()

    __args__['backend'] = backend
    __args__['kubernetesCaCert'] = kubernetes_ca_cert
    __args__['kubernetesHost'] = kubernetes_host
    __args__['pemKeys'] = pem_keys
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig', __args__, opts=opts).value

    return AwaitableGetAuthBackendConfigResult(
        backend=__ret__.get('backend'),
        kubernetes_ca_cert=__ret__.get('kubernetesCaCert'),
        kubernetes_host=__ret__.get('kubernetesHost'),
        pem_keys=__ret__.get('pemKeys'),
        id=__ret__.get('id'))
