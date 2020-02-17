# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendCert(pulumi.CustomResource):
    aws_public_cert: pulumi.Output[str]
    """
    Base64 encoded AWS Public key required to verify PKCS7 signature of the EC2 instance metadata.
    """
    backend: pulumi.Output[str]
    """
    Unique name of the auth backend to configure.
    """
    cert_name: pulumi.Output[str]
    """
    Name of the certificate to configure.
    """
    type: pulumi.Output[str]
    """
    The type of document that can be verified using the certificate. Must be either "pkcs7" or "identity".
    """
    def __init__(__self__, resource_name, opts=None, aws_public_cert=None, backend=None, cert_name=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a AuthBackendCert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_public_cert: Base64 encoded AWS Public key required to verify PKCS7 signature of the EC2 instance metadata.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[str] cert_name: Name of the certificate to configure.
        :param pulumi.Input[str] type: The type of document that can be verified using the certificate. Must be either "pkcs7" or "identity".
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

            if aws_public_cert is None:
                raise TypeError("Missing required property 'aws_public_cert'")
            __props__['aws_public_cert'] = aws_public_cert
            __props__['backend'] = backend
            if cert_name is None:
                raise TypeError("Missing required property 'cert_name'")
            __props__['cert_name'] = cert_name
            __props__['type'] = type
        super(AuthBackendCert, __self__).__init__(
            'vault:aws/authBackendCert:AuthBackendCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, aws_public_cert=None, backend=None, cert_name=None, type=None):
        """
        Get an existing AuthBackendCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_public_cert: Base64 encoded AWS Public key required to verify PKCS7 signature of the EC2 instance metadata.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[str] cert_name: Name of the certificate to configure.
        :param pulumi.Input[str] type: The type of document that can be verified using the certificate. Must be either "pkcs7" or "identity".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["aws_public_cert"] = aws_public_cert
        __props__["backend"] = backend
        __props__["cert_name"] = cert_name
        __props__["type"] = type
        return AuthBackendCert(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

