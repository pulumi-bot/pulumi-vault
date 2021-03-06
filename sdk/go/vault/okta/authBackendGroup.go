// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package okta

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a group in an
// [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_group.html.markdown.
type AuthBackendGroup struct {
	s *pulumi.ResourceState
}

// NewAuthBackendGroup registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendGroup(ctx *pulumi.Context,
	name string, args *AuthBackendGroupArgs, opts ...pulumi.ResourceOpt) (*AuthBackendGroup, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["groupName"] = nil
		inputs["path"] = nil
		inputs["policies"] = nil
	} else {
		inputs["groupName"] = args.GroupName
		inputs["path"] = args.Path
		inputs["policies"] = args.Policies
	}
	s, err := ctx.RegisterResource("vault:okta/authBackendGroup:AuthBackendGroup", name, true, inputs, opts...)
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
		inputs["groupName"] = state.GroupName
		inputs["path"] = state.Path
		inputs["policies"] = state.Policies
	}
	s, err := ctx.ReadResource("vault:okta/authBackendGroup:AuthBackendGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackendGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackendGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackendGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Name of the group within the Okta
func (r *AuthBackendGroup) GroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupName"])
}

// The path where the Okta auth backend is mounted
func (r *AuthBackendGroup) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// Vault policies to associate with this group
func (r *AuthBackendGroup) Policies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policies"])
}

// Input properties used for looking up and filtering AuthBackendGroup resources.
type AuthBackendGroupState struct {
	// Name of the group within the Okta
	GroupName interface{}
	// The path where the Okta auth backend is mounted
	Path interface{}
	// Vault policies to associate with this group
	Policies interface{}
}

// The set of arguments for constructing a AuthBackendGroup resource.
type AuthBackendGroupArgs struct {
	// Name of the group within the Okta
	GroupName interface{}
	// The path where the Okta auth backend is mounted
	Path interface{}
	// Vault policies to associate with this group
	Policies interface{}
}
