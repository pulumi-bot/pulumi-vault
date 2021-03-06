// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package okta

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a user in an
// [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_user.html.markdown.
type AuthBackendUser struct {
	s *pulumi.ResourceState
}

// NewAuthBackendUser registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendUser(ctx *pulumi.Context,
	name string, args *AuthBackendUserArgs, opts ...pulumi.ResourceOpt) (*AuthBackendUser, error) {
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["groups"] = nil
		inputs["path"] = nil
		inputs["policies"] = nil
		inputs["username"] = nil
	} else {
		inputs["groups"] = args.Groups
		inputs["path"] = args.Path
		inputs["policies"] = args.Policies
		inputs["username"] = args.Username
	}
	s, err := ctx.RegisterResource("vault:okta/authBackendUser:AuthBackendUser", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendUser{s: s}, nil
}

// GetAuthBackendUser gets an existing AuthBackendUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendUserState, opts ...pulumi.ResourceOpt) (*AuthBackendUser, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["groups"] = state.Groups
		inputs["path"] = state.Path
		inputs["policies"] = state.Policies
		inputs["username"] = state.Username
	}
	s, err := ctx.ReadResource("vault:okta/authBackendUser:AuthBackendUser", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendUser{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendUser) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendUser) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// List of Okta groups to associate with this user
func (r *AuthBackendUser) Groups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groups"])
}

// The path where the Okta auth backend is mounted
func (r *AuthBackendUser) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// List of Vault policies to associate with this user
func (r *AuthBackendUser) Policies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policies"])
}

// Name of the user within Okta
func (r *AuthBackendUser) Username() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["username"])
}

// Input properties used for looking up and filtering AuthBackendUser resources.
type AuthBackendUserState struct {
	// List of Okta groups to associate with this user
	Groups interface{}
	// The path where the Okta auth backend is mounted
	Path interface{}
	// List of Vault policies to associate with this user
	Policies interface{}
	// Name of the user within Okta
	Username interface{}
}

// The set of arguments for constructing a AuthBackendUser resource.
type AuthBackendUserArgs struct {
	// List of Okta groups to associate with this user
	Groups interface{}
	// The path where the Okta auth backend is mounted
	Path interface{}
	// List of Vault policies to associate with this user
	Policies interface{}
	// Name of the user within Okta
	Username interface{}
}
