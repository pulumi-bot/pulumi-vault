# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EntityAlias(pulumi.CustomResource):
    canonical_id: pulumi.Output[str]
    """
    Entity ID to which this alias belongs to.
    """
    mount_accessor: pulumi.Output[str]
    """
    Accessor of the mount to which the alias should belong to.
    """
    name: pulumi.Output[str]
    """
    Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
    """
    def __init__(__self__, resource_name, opts=None, canonical_id=None, mount_accessor=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a EntityAlias resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] canonical_id: Entity ID to which this alias belongs to.
        :param pulumi.Input[str] mount_accessor: Accessor of the mount to which the alias should belong to.
        :param pulumi.Input[str] name: Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity_alias.html.markdown.
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

            if canonical_id is None:
                raise TypeError("Missing required property 'canonical_id'")
            __props__['canonical_id'] = canonical_id
            if mount_accessor is None:
                raise TypeError("Missing required property 'mount_accessor'")
            __props__['mount_accessor'] = mount_accessor
            __props__['name'] = name
        super(EntityAlias, __self__).__init__(
            'vault:identity/entityAlias:EntityAlias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, canonical_id=None, mount_accessor=None, name=None):
        """
        Get an existing EntityAlias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] canonical_id: Entity ID to which this alias belongs to.
        :param pulumi.Input[str] mount_accessor: Accessor of the mount to which the alias should belong to.
        :param pulumi.Input[str] name: Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity_alias.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["canonical_id"] = canonical_id
        __props__["mount_accessor"] = mount_accessor
        __props__["name"] = name
        return EntityAlias(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

