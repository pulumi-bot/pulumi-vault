# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretBackendRootCert(pulumi.CustomResource):
    alt_names: pulumi.Output[list]
    """
    List of alternative names
    """
    backend: pulumi.Output[str]
    """
    The PKI secret backend the resource belongs to.
    """
    certificate: pulumi.Output[str]
    """
    The certificate
    """
    common_name: pulumi.Output[str]
    """
    CN of intermediate to create
    """
    country: pulumi.Output[str]
    """
    The country
    """
    exclude_cn_from_sans: pulumi.Output[bool]
    """
    Flag to exclude CN from SANs
    """
    format: pulumi.Output[str]
    """
    The format of data
    """
    ip_sans: pulumi.Output[list]
    """
    List of alternative IPs
    """
    issuing_ca: pulumi.Output[str]
    """
    The issuing CA
    """
    key_bits: pulumi.Output[float]
    """
    The number of bits to use
    """
    key_type: pulumi.Output[str]
    """
    The desired key type
    """
    locality: pulumi.Output[str]
    """
    The locality
    """
    max_path_length: pulumi.Output[float]
    """
    The maximum path length to encode in the generated certificate
    """
    organization: pulumi.Output[str]
    """
    The organization
    """
    other_sans: pulumi.Output[list]
    """
    List of other SANs
    """
    ou: pulumi.Output[str]
    """
    The organization unit
    """
    permitted_dns_domains: pulumi.Output[list]
    """
    List of domains for which certificates are allowed to be issued
    """
    postal_code: pulumi.Output[str]
    """
    The postal code
    """
    private_key_format: pulumi.Output[str]
    """
    The private key format
    """
    province: pulumi.Output[str]
    """
    The province
    """
    serial: pulumi.Output[str]
    """
    The serial
    """
    street_address: pulumi.Output[str]
    """
    The street address
    """
    ttl: pulumi.Output[str]
    """
    Time to live
    """
    type: pulumi.Output[str]
    """
    Type of intermediate to create. Must be either \"exported\" or \"internal\"
    """
    uri_sans: pulumi.Output[list]
    """
    List of alternative URIs
    """
    def __init__(__self__, resource_name, opts=None, alt_names=None, backend=None, common_name=None, country=None, exclude_cn_from_sans=None, format=None, ip_sans=None, key_bits=None, key_type=None, locality=None, max_path_length=None, organization=None, other_sans=None, ou=None, permitted_dns_domains=None, postal_code=None, private_key_format=None, province=None, street_address=None, ttl=None, type=None, uri_sans=None, __props__=None, __name__=None, __opts__=None):
        """
        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_root_cert.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] alt_names: List of alternative names
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of intermediate to create
        :param pulumi.Input[str] country: The country
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[list] ip_sans: List of alternative IPs
        :param pulumi.Input[float] key_bits: The number of bits to use
        :param pulumi.Input[str] key_type: The desired key type
        :param pulumi.Input[str] locality: The locality
        :param pulumi.Input[float] max_path_length: The maximum path length to encode in the generated certificate
        :param pulumi.Input[str] organization: The organization
        :param pulumi.Input[list] other_sans: List of other SANs
        :param pulumi.Input[str] ou: The organization unit
        :param pulumi.Input[list] permitted_dns_domains: List of domains for which certificates are allowed to be issued
        :param pulumi.Input[str] postal_code: The postal code
        :param pulumi.Input[str] private_key_format: The private key format
        :param pulumi.Input[str] province: The province
        :param pulumi.Input[str] street_address: The street address
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[str] type: Type of intermediate to create. Must be either \"exported\" or \"internal\"
        :param pulumi.Input[list] uri_sans: List of alternative URIs
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

            __props__['alt_names'] = alt_names
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if common_name is None:
                raise TypeError("Missing required property 'common_name'")
            __props__['common_name'] = common_name
            __props__['country'] = country
            __props__['exclude_cn_from_sans'] = exclude_cn_from_sans
            __props__['format'] = format
            __props__['ip_sans'] = ip_sans
            __props__['key_bits'] = key_bits
            __props__['key_type'] = key_type
            __props__['locality'] = locality
            __props__['max_path_length'] = max_path_length
            __props__['organization'] = organization
            __props__['other_sans'] = other_sans
            __props__['ou'] = ou
            __props__['permitted_dns_domains'] = permitted_dns_domains
            __props__['postal_code'] = postal_code
            __props__['private_key_format'] = private_key_format
            __props__['province'] = province
            __props__['street_address'] = street_address
            __props__['ttl'] = ttl
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['uri_sans'] = uri_sans
            __props__['certificate'] = None
            __props__['issuing_ca'] = None
            __props__['serial'] = None
        super(SecretBackendRootCert, __self__).__init__(
            'vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alt_names=None, backend=None, certificate=None, common_name=None, country=None, exclude_cn_from_sans=None, format=None, ip_sans=None, issuing_ca=None, key_bits=None, key_type=None, locality=None, max_path_length=None, organization=None, other_sans=None, ou=None, permitted_dns_domains=None, postal_code=None, private_key_format=None, province=None, serial=None, street_address=None, ttl=None, type=None, uri_sans=None):
        """
        Get an existing SecretBackendRootCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] alt_names: List of alternative names
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] certificate: The certificate
        :param pulumi.Input[str] common_name: CN of intermediate to create
        :param pulumi.Input[str] country: The country
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[list] ip_sans: List of alternative IPs
        :param pulumi.Input[str] issuing_ca: The issuing CA
        :param pulumi.Input[float] key_bits: The number of bits to use
        :param pulumi.Input[str] key_type: The desired key type
        :param pulumi.Input[str] locality: The locality
        :param pulumi.Input[float] max_path_length: The maximum path length to encode in the generated certificate
        :param pulumi.Input[str] organization: The organization
        :param pulumi.Input[list] other_sans: List of other SANs
        :param pulumi.Input[str] ou: The organization unit
        :param pulumi.Input[list] permitted_dns_domains: List of domains for which certificates are allowed to be issued
        :param pulumi.Input[str] postal_code: The postal code
        :param pulumi.Input[str] private_key_format: The private key format
        :param pulumi.Input[str] province: The province
        :param pulumi.Input[str] serial: The serial
        :param pulumi.Input[str] street_address: The street address
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[str] type: Type of intermediate to create. Must be either \"exported\" or \"internal\"
        :param pulumi.Input[list] uri_sans: List of alternative URIs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alt_names"] = alt_names
        __props__["backend"] = backend
        __props__["certificate"] = certificate
        __props__["common_name"] = common_name
        __props__["country"] = country
        __props__["exclude_cn_from_sans"] = exclude_cn_from_sans
        __props__["format"] = format
        __props__["ip_sans"] = ip_sans
        __props__["issuing_ca"] = issuing_ca
        __props__["key_bits"] = key_bits
        __props__["key_type"] = key_type
        __props__["locality"] = locality
        __props__["max_path_length"] = max_path_length
        __props__["organization"] = organization
        __props__["other_sans"] = other_sans
        __props__["ou"] = ou
        __props__["permitted_dns_domains"] = permitted_dns_domains
        __props__["postal_code"] = postal_code
        __props__["private_key_format"] = private_key_format
        __props__["province"] = province
        __props__["serial"] = serial
        __props__["street_address"] = street_address
        __props__["ttl"] = ttl
        __props__["type"] = type
        __props__["uri_sans"] = uri_sans
        return SecretBackendRootCert(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

