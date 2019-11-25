// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appRole

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Logs into Vault using the AppRole auth backend. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/approle.html) for more
// information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/approle_auth_backend_login.html.markdown.
type AuthBackendLogin struct {
	s *pulumi.ResourceState
}

// NewAuthBackendLogin registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendLogin(ctx *pulumi.Context,
	name string, args *AuthBackendLoginArgs, opts ...pulumi.ResourceOpt) (*AuthBackendLogin, error) {
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["roleId"] = nil
		inputs["secretId"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["roleId"] = args.RoleId
		inputs["secretId"] = args.SecretId
	}
	inputs["accessor"] = nil
	inputs["clientToken"] = nil
	inputs["leaseDuration"] = nil
	inputs["leaseStarted"] = nil
	inputs["metadata"] = nil
	inputs["policies"] = nil
	inputs["renewable"] = nil
	s, err := ctx.RegisterResource("vault:appRole/authBackendLogin:AuthBackendLogin", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendLogin{s: s}, nil
}

// GetAuthBackendLogin gets an existing AuthBackendLogin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendLogin(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendLoginState, opts ...pulumi.ResourceOpt) (*AuthBackendLogin, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessor"] = state.Accessor
		inputs["backend"] = state.Backend
		inputs["clientToken"] = state.ClientToken
		inputs["leaseDuration"] = state.LeaseDuration
		inputs["leaseStarted"] = state.LeaseStarted
		inputs["metadata"] = state.Metadata
		inputs["policies"] = state.Policies
		inputs["renewable"] = state.Renewable
		inputs["roleId"] = state.RoleId
		inputs["secretId"] = state.SecretId
	}
	s, err := ctx.ReadResource("vault:appRole/authBackendLogin:AuthBackendLogin", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendLogin{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendLogin) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendLogin) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The accessor for the token.
func (r *AuthBackendLogin) Accessor() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessor"])
}

// The unique path of the Vault backend to log in with.
func (r *AuthBackendLogin) Backend() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backend"])
}

// The Vault token created.
func (r *AuthBackendLogin) ClientToken() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientToken"])
}

// How long the token is valid for, in seconds.
func (r *AuthBackendLogin) LeaseDuration() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["leaseDuration"])
}

// The date and time the lease started, in RFC 3339 format.
func (r *AuthBackendLogin) LeaseStarted() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["leaseStarted"])
}

// The metadata associated with the token.
func (r *AuthBackendLogin) Metadata() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["metadata"])
}

// A list of policies applied to the token.
func (r *AuthBackendLogin) Policies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["policies"])
}

// Whether the token is renewable or not.
func (r *AuthBackendLogin) Renewable() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["renewable"])
}

// The ID of the role to log in with.
func (r *AuthBackendLogin) RoleId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["roleId"])
}

// The secret ID of the role to log in with. Required
// unless `bindSecretId` is set to false on the role.
func (r *AuthBackendLogin) SecretId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secretId"])
}

// Input properties used for looking up and filtering AuthBackendLogin resources.
type AuthBackendLoginState struct {
	// The accessor for the token.
	Accessor interface{}
	// The unique path of the Vault backend to log in with.
	Backend interface{}
	// The Vault token created.
	ClientToken interface{}
	// How long the token is valid for, in seconds.
	LeaseDuration interface{}
	// The date and time the lease started, in RFC 3339 format.
	LeaseStarted interface{}
	// The metadata associated with the token.
	Metadata interface{}
	// A list of policies applied to the token.
	Policies interface{}
	// Whether the token is renewable or not.
	Renewable interface{}
	// The ID of the role to log in with.
	RoleId interface{}
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId interface{}
}

// The set of arguments for constructing a AuthBackendLogin resource.
type AuthBackendLoginArgs struct {
	// The unique path of the Vault backend to log in with.
	Backend interface{}
	// The ID of the role to log in with.
	RoleId interface{}
	// The secret ID of the role to log in with. Required
	// unless `bindSecretId` is set to false on the role.
	SecretId interface{}
}
