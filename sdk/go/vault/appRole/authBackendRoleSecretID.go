// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appRole

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/approle.html) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/approle_auth_backend_role_secret_id.html.markdown.
type AuthBackendRoleSecretID struct {
	s *pulumi.ResourceState
}

// NewAuthBackendRoleSecretID registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRoleSecretID(ctx *pulumi.Context,
	name string, args *AuthBackendRoleSecretIDArgs, opts ...pulumi.ResourceOpt) (*AuthBackendRoleSecretID, error) {
	if args == nil || args.RoleName == nil {
		return nil, errors.New("missing required argument 'RoleName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["cidrLists"] = nil
		inputs["metadata"] = nil
		inputs["roleName"] = nil
		inputs["secretId"] = nil
		inputs["wrappingTtl"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["cidrLists"] = args.CidrLists
		inputs["metadata"] = args.Metadata
		inputs["roleName"] = args.RoleName
		inputs["secretId"] = args.SecretId
		inputs["wrappingTtl"] = args.WrappingTtl
	}
	inputs["accessor"] = nil
	inputs["wrappingAccessor"] = nil
	inputs["wrappingToken"] = nil
	s, err := ctx.RegisterResource("vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRoleSecretID{s: s}, nil
}

// GetAuthBackendRoleSecretID gets an existing AuthBackendRoleSecretID resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRoleSecretID(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendRoleSecretIDState, opts ...pulumi.ResourceOpt) (*AuthBackendRoleSecretID, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessor"] = state.Accessor
		inputs["backend"] = state.Backend
		inputs["cidrLists"] = state.CidrLists
		inputs["metadata"] = state.Metadata
		inputs["roleName"] = state.RoleName
		inputs["secretId"] = state.SecretId
		inputs["wrappingAccessor"] = state.WrappingAccessor
		inputs["wrappingToken"] = state.WrappingToken
		inputs["wrappingTtl"] = state.WrappingTtl
	}
	s, err := ctx.ReadResource("vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendRoleSecretID{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendRoleSecretID) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendRoleSecretID) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The unique ID for this SecretID that can be safely logged.
func (r *AuthBackendRoleSecretID) Accessor() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessor"])
}

// Unique name of the auth backend to configure.
func (r *AuthBackendRoleSecretID) Backend() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backend"])
}

// If set, specifies blocks of IP addresses which can
// perform the login operation using this SecretID.
func (r *AuthBackendRoleSecretID) CidrLists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["cidrLists"])
}

// A JSON-encoded string containing metadata in
// key-value pairs to be set on tokens issued with this SecretID.
func (r *AuthBackendRoleSecretID) Metadata() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["metadata"])
}

// The name of the role to create the SecretID for.
func (r *AuthBackendRoleSecretID) RoleName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleName"])
}

// The SecretID to be created. If set, uses "Push"
// mode.  Defaults to Vault auto-generating SecretIDs.
func (r *AuthBackendRoleSecretID) SecretId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secretId"])
}

// The unique ID for the response-wrapped SecretID that can
// be safely logged.
func (r *AuthBackendRoleSecretID) WrappingAccessor() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["wrappingAccessor"])
}

// The token used to retrieve a response-wrapped SecretID.
func (r *AuthBackendRoleSecretID) WrappingToken() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["wrappingToken"])
}

// If set, the SecretID response will be
// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping.html)
// and available for the duration specified. Only a single unwrapping of the
// token is allowed.
func (r *AuthBackendRoleSecretID) WrappingTtl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["wrappingTtl"])
}

// Input properties used for looking up and filtering AuthBackendRoleSecretID resources.
type AuthBackendRoleSecretIDState struct {
	// The unique ID for this SecretID that can be safely logged.
	Accessor interface{}
	// Unique name of the auth backend to configure.
	Backend interface{}
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists interface{}
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata interface{}
	// The name of the role to create the SecretID for.
	RoleName interface{}
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId interface{}
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor interface{}
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken interface{}
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping.html)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl interface{}
}

// The set of arguments for constructing a AuthBackendRoleSecretID resource.
type AuthBackendRoleSecretIDArgs struct {
	// Unique name of the auth backend to configure.
	Backend interface{}
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists interface{}
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata interface{}
	// The name of the role to create the SecretID for.
	RoleName interface{}
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId interface{}
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping.html)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl interface{}
}
