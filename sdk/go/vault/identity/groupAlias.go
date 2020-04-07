// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Identity Group Alias for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// Group aliases allows entity membership in external groups to be managed semi-automatically. External group serves as a mapping to a group that is outside of the identity store. External groups can have one (and only one) alias. This alias should map to a notion of group that is outside of the identity store. For example, groups in LDAP, and teams in GitHub. A username in LDAP, belonging to a group in LDAP, can get its entity ID added as a member of a group in Vault automatically during logins and token renewals. This works only if the group in Vault is an external group and has an alias that maps to the group in LDAP. If the user is removed from the group in LDAP, that change gets reflected in Vault only upon the subsequent login or renewal operation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group_alias.html.md.
type GroupAlias struct {
	pulumi.CustomResourceState

	// ID of the group to which this is an alias.
	CanonicalId pulumi.StringOutput `pulumi:"canonicalId"`
	// Mount accessor of the authentication backend to which this alias belongs to.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// Name of the group alias to create.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGroupAlias registers a new resource with the given unique name, arguments, and options.
func NewGroupAlias(ctx *pulumi.Context,
	name string, args *GroupAliasArgs, opts ...pulumi.ResourceOption) (*GroupAlias, error) {
	if args == nil || args.CanonicalId == nil {
		return nil, errors.New("missing required argument 'CanonicalId'")
	}
	if args == nil || args.MountAccessor == nil {
		return nil, errors.New("missing required argument 'MountAccessor'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &GroupAliasArgs{}
	}
	var resource GroupAlias
	err := ctx.RegisterResource("vault:identity/groupAlias:GroupAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupAlias gets an existing GroupAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupAliasState, opts ...pulumi.ResourceOption) (*GroupAlias, error) {
	var resource GroupAlias
	err := ctx.ReadResource("vault:identity/groupAlias:GroupAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupAlias resources.
type groupAliasState struct {
	// ID of the group to which this is an alias.
	CanonicalId *string `pulumi:"canonicalId"`
	// Mount accessor of the authentication backend to which this alias belongs to.
	MountAccessor *string `pulumi:"mountAccessor"`
	// Name of the group alias to create.
	Name *string `pulumi:"name"`
}

type GroupAliasState struct {
	// ID of the group to which this is an alias.
	CanonicalId pulumi.StringPtrInput
	// Mount accessor of the authentication backend to which this alias belongs to.
	MountAccessor pulumi.StringPtrInput
	// Name of the group alias to create.
	Name pulumi.StringPtrInput
}

func (GroupAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupAliasState)(nil)).Elem()
}

type groupAliasArgs struct {
	// ID of the group to which this is an alias.
	CanonicalId string `pulumi:"canonicalId"`
	// Mount accessor of the authentication backend to which this alias belongs to.
	MountAccessor string `pulumi:"mountAccessor"`
	// Name of the group alias to create.
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a GroupAlias resource.
type GroupAliasArgs struct {
	// ID of the group to which this is an alias.
	CanonicalId pulumi.StringInput
	// Mount accessor of the authentication backend to which this alias belongs to.
	MountAccessor pulumi.StringInput
	// Name of the group alias to create.
	Name pulumi.StringInput
}

func (GroupAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupAliasArgs)(nil)).Elem()
}
