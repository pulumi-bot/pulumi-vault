// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type OidcRole struct {
	pulumi.CustomResourceState

	// The value that will be included in the `aud` field of all the OIDC identity tokens issued by this role
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// A configured named key, the key must already exist.
	Key pulumi.StringOutput `pulumi:"key"`
	// Name of the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
}

// NewOidcRole registers a new resource with the given unique name, arguments, and options.
func NewOidcRole(ctx *pulumi.Context,
	name string, args *OidcRoleArgs, opts ...pulumi.ResourceOption) (*OidcRole, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil {
		args = &OidcRoleArgs{}
	}
	var resource OidcRole
	err := ctx.RegisterResource("vault:identity/oidcRole:OidcRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcRole gets an existing OidcRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcRoleState, opts ...pulumi.ResourceOption) (*OidcRole, error) {
	var resource OidcRole
	err := ctx.ReadResource("vault:identity/oidcRole:OidcRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcRole resources.
type oidcRoleState struct {
	// The value that will be included in the `aud` field of all the OIDC identity tokens issued by this role
	ClientId *string `pulumi:"clientId"`
	// A configured named key, the key must already exist.
	Key *string `pulumi:"key"`
	// Name of the role.
	Name *string `pulumi:"name"`
	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	Template *string `pulumi:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	Ttl *int `pulumi:"ttl"`
}

type OidcRoleState struct {
	// The value that will be included in the `aud` field of all the OIDC identity tokens issued by this role
	ClientId pulumi.StringPtrInput
	// A configured named key, the key must already exist.
	Key pulumi.StringPtrInput
	// Name of the role.
	Name pulumi.StringPtrInput
	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	Template pulumi.StringPtrInput
	// TTL of the tokens generated against the role in number of seconds.
	Ttl pulumi.IntPtrInput
}

func (OidcRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcRoleState)(nil)).Elem()
}

type oidcRoleArgs struct {
	// A configured named key, the key must already exist.
	Key string `pulumi:"key"`
	// Name of the role.
	Name *string `pulumi:"name"`
	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	Template *string `pulumi:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a OidcRole resource.
type OidcRoleArgs struct {
	// A configured named key, the key must already exist.
	Key pulumi.StringInput
	// Name of the role.
	Name pulumi.StringPtrInput
	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	Template pulumi.StringPtrInput
	// TTL of the tokens generated against the role in number of seconds.
	Ttl pulumi.IntPtrInput
}

func (OidcRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcRoleArgs)(nil)).Elem()
}

