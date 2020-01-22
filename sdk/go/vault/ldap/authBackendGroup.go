// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ldap_auth_backend_group.html.markdown.
type AuthBackendGroup struct {
	pulumi.CustomResourceState

	// Path to the authentication backend
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The LDAP groupname
	Groupname pulumi.StringOutput `pulumi:"groupname"`
	// Policies which should be granted to members of the group
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewAuthBackendGroup registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendGroup(ctx *pulumi.Context,
	name string, args *AuthBackendGroupArgs, opts ...pulumi.ResourceOption) (*AuthBackendGroup, error) {
	if args == nil || args.Groupname == nil {
		return nil, errors.New("missing required argument 'Groupname'")
	}
	if args == nil {
		args = &AuthBackendGroupArgs{}
	}
	var resource AuthBackendGroup
	err := ctx.RegisterResource("vault:ldap/authBackendGroup:AuthBackendGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendGroup gets an existing AuthBackendGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendGroupState, opts ...pulumi.ResourceOption) (*AuthBackendGroup, error) {
	var resource AuthBackendGroup
	err := ctx.ReadResource("vault:ldap/authBackendGroup:AuthBackendGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendGroup resources.
type authBackendGroupState struct {
	// Path to the authentication backend
	Backend *string `pulumi:"backend"`
	// The LDAP groupname
	Groupname *string `pulumi:"groupname"`
	// Policies which should be granted to members of the group
	Policies []string `pulumi:"policies"`
}

type AuthBackendGroupState struct {
	// Path to the authentication backend
	Backend pulumi.StringPtrInput
	// The LDAP groupname
	Groupname pulumi.StringPtrInput
	// Policies which should be granted to members of the group
	Policies pulumi.StringArrayInput
}

func (AuthBackendGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendGroupState)(nil)).Elem()
}

type authBackendGroupArgs struct {
	// Path to the authentication backend
	Backend *string `pulumi:"backend"`
	// The LDAP groupname
	Groupname string `pulumi:"groupname"`
	// Policies which should be granted to members of the group
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a AuthBackendGroup resource.
type AuthBackendGroupArgs struct {
	// Path to the authentication backend
	Backend pulumi.StringPtrInput
	// The LDAP groupname
	Groupname pulumi.StringInput
	// Policies which should be granted to members of the group
	Policies pulumi.StringArrayInput
}

func (AuthBackendGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendGroupArgs)(nil)).Elem()
}

