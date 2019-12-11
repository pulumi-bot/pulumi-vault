// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/ldap_auth_backend_group.html.markdown.
type AuthBackendGroup struct {
	s *pulumi.ResourceState
}

// NewAuthBackendGroup registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendGroup(ctx *pulumi.Context,
	name string, args *AuthBackendGroupArgs, opts ...pulumi.ResourceOpt) (*AuthBackendGroup, error) {
	if args == nil || args.Groupname == nil {
		return nil, errors.New("missing required argument 'Groupname'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["groupname"] = nil
		inputs["policies"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["groupname"] = args.Groupname
		inputs["policies"] = args.Policies
	}
	s, err := ctx.RegisterResource("vault:ldap/authBackendGroup:AuthBackendGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendGroup{s: s}, nil
}

// GetAuthBackendGroup gets an existing AuthBackendGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendGroupState, opts ...pulumi.ResourceOpt) (*AuthBackendGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backend"] = state.Backend
		inputs["groupname"] = state.Groupname
		inputs["policies"] = state.Policies
	}
	s, err := ctx.ReadResource("vault:ldap/authBackendGroup:AuthBackendGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Path to the authentication backend
func (r *AuthBackendGroup) Backend() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backend"])
}

// The LDAP groupname
func (r *AuthBackendGroup) Groupname() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["groupname"])
}

// Policies which should be granted to members of the group
func (r *AuthBackendGroup) Policies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["policies"])
}

// Input properties used for looking up and filtering AuthBackendGroup resources.
type AuthBackendGroupState struct {
	// Path to the authentication backend
	Backend interface{}
	// The LDAP groupname
	Groupname interface{}
	// Policies which should be granted to members of the group
	Policies interface{}
}

// The set of arguments for constructing a AuthBackendGroup resource.
type AuthBackendGroupArgs struct {
	// Path to the authentication backend
	Backend interface{}
	// The LDAP groupname
	Groupname interface{}
	// Policies which should be granted to members of the group
	Policies interface{}
}