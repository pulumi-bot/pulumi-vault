// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type OidcKeyAllowedClientID struct {
	pulumi.CustomResourceState

	// Client ID to allow usage with the OIDC named key
	AllowedClientId pulumi.StringOutput `pulumi:"allowedClientId"`
	// Name of the OIDC Key allow the Client ID.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
}

// NewOidcKeyAllowedClientID registers a new resource with the given unique name, arguments, and options.
func NewOidcKeyAllowedClientID(ctx *pulumi.Context,
	name string, args *OidcKeyAllowedClientIDArgs, opts ...pulumi.ResourceOption) (*OidcKeyAllowedClientID, error) {
	if args == nil || args.AllowedClientId == nil {
		return nil, errors.New("missing required argument 'AllowedClientId'")
	}
	if args == nil || args.KeyName == nil {
		return nil, errors.New("missing required argument 'KeyName'")
	}
	if args == nil {
		args = &OidcKeyAllowedClientIDArgs{}
	}
	var resource OidcKeyAllowedClientID
	err := ctx.RegisterResource("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcKeyAllowedClientID gets an existing OidcKeyAllowedClientID resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcKeyAllowedClientID(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcKeyAllowedClientIDState, opts ...pulumi.ResourceOption) (*OidcKeyAllowedClientID, error) {
	var resource OidcKeyAllowedClientID
	err := ctx.ReadResource("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcKeyAllowedClientID resources.
type oidcKeyAllowedClientIDState struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId *string `pulumi:"allowedClientId"`
	// Name of the OIDC Key allow the Client ID.
	KeyName *string `pulumi:"keyName"`
}

type OidcKeyAllowedClientIDState struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId pulumi.StringPtrInput
	// Name of the OIDC Key allow the Client ID.
	KeyName pulumi.StringPtrInput
}

func (OidcKeyAllowedClientIDState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcKeyAllowedClientIDState)(nil)).Elem()
}

type oidcKeyAllowedClientIDArgs struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId string `pulumi:"allowedClientId"`
	// Name of the OIDC Key allow the Client ID.
	KeyName string `pulumi:"keyName"`
}

// The set of arguments for constructing a OidcKeyAllowedClientID resource.
type OidcKeyAllowedClientIDArgs struct {
	// Client ID to allow usage with the OIDC named key
	AllowedClientId pulumi.StringInput
	// Name of the OIDC Key allow the Client ID.
	KeyName pulumi.StringInput
}

func (OidcKeyAllowedClientIDArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcKeyAllowedClientIDArgs)(nil)).Elem()
}
