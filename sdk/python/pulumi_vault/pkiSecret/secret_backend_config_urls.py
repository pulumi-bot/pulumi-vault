# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretBackendConfigUrls(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
    """
    crl_distribution_points: pulumi.Output[list]
    """
    Specifies the URL values for the CRL Distribution Points field.
    """
    issuing_certificates: pulumi.Output[list]
    """
    Specifies the URL values for the Issuing Certificate field.
    """
    ocsp_servers: pulumi.Output[list]
    """
    Specifies the URL values for the OCSP Servers field.
    """
    def __init__(__self__, resource_name, opts=None, backend=None, crl_distribution_points=None, issuing_certificates=None, ocsp_servers=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_config_urls.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[list] crl_distribution_points: Specifies the URL values for the CRL Distribution Points field.
        :param pulumi.Input[list] issuing_certificates: Specifies the URL values for the Issuing Certificate field.
        :param pulumi.Input[list] ocsp_servers: Specifies the URL values for the OCSP Servers field.
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

            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['crl_distribution_points'] = crl_distribution_points
            __props__['issuing_certificates'] = issuing_certificates
            __props__['ocsp_servers'] = ocsp_servers
        super(SecretBackendConfigUrls, __self__).__init__(
            'vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, crl_distribution_points=None, issuing_certificates=None, ocsp_servers=None):
        """
        Get an existing SecretBackendConfigUrls resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[list] crl_distribution_points: Specifies the URL values for the CRL Distribution Points field.
        :param pulumi.Input[list] issuing_certificates: Specifies the URL values for the Issuing Certificate field.
        :param pulumi.Input[list] ocsp_servers: Specifies the URL values for the OCSP Servers field.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["crl_distribution_points"] = crl_distribution_points
        __props__["issuing_certificates"] = issuing_certificates
        __props__["ocsp_servers"] = ocsp_servers
        return SecretBackendConfigUrls(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

