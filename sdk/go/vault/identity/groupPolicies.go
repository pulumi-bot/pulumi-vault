// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group_policies.html.markdown.
type GroupPolicies struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	Exclusive pulumi.BoolPtrOutput `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the group that are assigned the policies.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// List of policies to assign to the group
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewGroupPolicies registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicies(ctx *pulumi.Context,
	name string, args *GroupPoliciesArgs, opts ...pulumi.ResourceOption) (*GroupPolicies, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil || args.Policies == nil {
		return nil, errors.New("missing required argument 'Policies'")
	}
	if args == nil {
		args = &GroupPoliciesArgs{}
	}
	var resource GroupPolicies
	err := ctx.RegisterResource("vault:identity/groupPolicies:GroupPolicies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicies gets an existing GroupPolicies resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPoliciesState, opts ...pulumi.ResourceOption) (*GroupPolicies, error) {
	var resource GroupPolicies
	err := ctx.ReadResource("vault:identity/groupPolicies:GroupPolicies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicies resources.
type groupPoliciesState struct {
	// Defaults to `true`.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId *string `pulumi:"groupId"`
	// The name of the group that are assigned the policies.
	GroupName *string `pulumi:"groupName"`
	// List of policies to assign to the group
	Policies []string `pulumi:"policies"`
}

type GroupPoliciesState struct {
	// Defaults to `true`.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign policies to.
	GroupId pulumi.StringPtrInput
	// The name of the group that are assigned the policies.
	GroupName pulumi.StringPtrInput
	// List of policies to assign to the group
	Policies pulumi.StringArrayInput
}

func (GroupPoliciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPoliciesState)(nil)).Elem()
}

type groupPoliciesArgs struct {
	// Defaults to `true`.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId string `pulumi:"groupId"`
	// List of policies to assign to the group
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a GroupPolicies resource.
type GroupPoliciesArgs struct {
	// Defaults to `true`.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign policies to.
	GroupId pulumi.StringInput
	// List of policies to assign to the group
	Policies pulumi.StringArrayInput
}

func (GroupPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPoliciesArgs)(nil)).Elem()
}

