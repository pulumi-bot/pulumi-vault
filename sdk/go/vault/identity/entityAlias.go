// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity_alias.html.markdown.
type EntityAlias struct {
	pulumi.CustomResourceState

	// Entity ID to which this alias belongs to.
	CanonicalId pulumi.StringOutput `pulumi:"canonicalId"`
	// Accessor of the mount to which the alias should belong to.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEntityAlias registers a new resource with the given unique name, arguments, and options.
func NewEntityAlias(ctx *pulumi.Context,
	name string, args *EntityAliasArgs, opts ...pulumi.ResourceOption) (*EntityAlias, error) {
	if args == nil || args.CanonicalId == nil {
		return nil, errors.New("missing required argument 'CanonicalId'")
	}
	if args == nil || args.MountAccessor == nil {
		return nil, errors.New("missing required argument 'MountAccessor'")
	}
	if args == nil {
		args = &EntityAliasArgs{}
	}
	var resource EntityAlias
	err := ctx.RegisterResource("vault:identity/entityAlias:EntityAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityAlias gets an existing EntityAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityAliasState, opts ...pulumi.ResourceOption) (*EntityAlias, error) {
	var resource EntityAlias
	err := ctx.ReadResource("vault:identity/entityAlias:EntityAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityAlias resources.
type entityAliasState struct {
	// Entity ID to which this alias belongs to.
	CanonicalId *string `pulumi:"canonicalId"`
	// Accessor of the mount to which the alias should belong to.
	MountAccessor *string `pulumi:"mountAccessor"`
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name *string `pulumi:"name"`
}

type EntityAliasState struct {
	// Entity ID to which this alias belongs to.
	CanonicalId pulumi.StringPtrInput
	// Accessor of the mount to which the alias should belong to.
	MountAccessor pulumi.StringPtrInput
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name pulumi.StringPtrInput
}

func (EntityAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAliasState)(nil)).Elem()
}

type entityAliasArgs struct {
	// Entity ID to which this alias belongs to.
	CanonicalId string `pulumi:"canonicalId"`
	// Accessor of the mount to which the alias should belong to.
	MountAccessor string `pulumi:"mountAccessor"`
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a EntityAlias resource.
type EntityAliasArgs struct {
	// Entity ID to which this alias belongs to.
	CanonicalId pulumi.StringInput
	// Accessor of the mount to which the alias should belong to.
	MountAccessor pulumi.StringInput
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name pulumi.StringPtrInput
}

func (EntityAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityAliasArgs)(nil)).Elem()
}

