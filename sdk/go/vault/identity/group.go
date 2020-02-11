// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// A group can contain multiple entities as its members. A group can also have subgroups. Policies set on the group is granted to all members of the group. During request time, when the token's entity ID is being evaluated for the policies that it has access to; along with the policies on the entity itself, policies that are inherited due to group memberships are also granted.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group.html.markdown.
type Group struct {
	pulumi.CustomResourceState

	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrOutput `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayOutput `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds pulumi.StringArrayOutput `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Name of the identity group to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("vault:identity/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("vault:identity/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity group to create.
	Name *string `pulumi:"name"`
	// A list of policies to apply to the group.
	Policies []string `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type *string `pulumi:"type"`
}

type GroupState struct {
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrInput
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayInput
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds pulumi.StringArrayInput
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapInput
	// Name of the identity group to create.
	Name pulumi.StringPtrInput
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayInput
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies *bool `pulumi:"externalPolicies"`
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds []string `pulumi:"memberEntityIds"`
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds []string `pulumi:"memberGroupIds"`
	// A Map of additional metadata to associate with the group.
	Metadata map[string]string `pulumi:"metadata"`
	// Name of the identity group to create.
	Name *string `pulumi:"name"`
	// A list of policies to apply to the group.
	Policies []string `pulumi:"policies"`
	// Type of the group, internal or external. Defaults to `internal`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.GroupPolicies` to manage policies for this group in a decoupled manner.
	ExternalPolicies pulumi.BoolPtrInput
	// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
	MemberEntityIds pulumi.StringArrayInput
	// A list of Group IDs to be assigned as group members.
	MemberGroupIds pulumi.StringArrayInput
	// A Map of additional metadata to associate with the group.
	Metadata pulumi.StringMapInput
	// Name of the identity group to create.
	Name pulumi.StringPtrInput
	// A list of policies to apply to the group.
	Policies pulumi.StringArrayInput
	// Type of the group, internal or external. Defaults to `internal`.
	Type pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

