// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity_alias.html.markdown.
type EntityAlias struct {
	s *pulumi.ResourceState
}

// NewEntityAlias registers a new resource with the given unique name, arguments, and options.
func NewEntityAlias(ctx *pulumi.Context,
	name string, args *EntityAliasArgs, opts ...pulumi.ResourceOpt) (*EntityAlias, error) {
	if args == nil || args.CanonicalId == nil {
		return nil, errors.New("missing required argument 'CanonicalId'")
	}
	if args == nil || args.MountAccessor == nil {
		return nil, errors.New("missing required argument 'MountAccessor'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["canonicalId"] = nil
		inputs["mountAccessor"] = nil
		inputs["name"] = nil
	} else {
		inputs["canonicalId"] = args.CanonicalId
		inputs["mountAccessor"] = args.MountAccessor
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("vault:identity/entityAlias:EntityAlias", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EntityAlias{s: s}, nil
}

// GetEntityAlias gets an existing EntityAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityAlias(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EntityAliasState, opts ...pulumi.ResourceOpt) (*EntityAlias, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["canonicalId"] = state.CanonicalId
		inputs["mountAccessor"] = state.MountAccessor
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("vault:identity/entityAlias:EntityAlias", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EntityAlias{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EntityAlias) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EntityAlias) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Entity ID to which this alias belongs to.
func (r *EntityAlias) CanonicalId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["canonicalId"])
}

// Accessor of the mount to which the alias should belong to.
func (r *EntityAlias) MountAccessor() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mountAccessor"])
}

// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
func (r *EntityAlias) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering EntityAlias resources.
type EntityAliasState struct {
	// Entity ID to which this alias belongs to.
	CanonicalId interface{}
	// Accessor of the mount to which the alias should belong to.
	MountAccessor interface{}
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name interface{}
}

// The set of arguments for constructing a EntityAlias resource.
type EntityAliasArgs struct {
	// Entity ID to which this alias belongs to.
	CanonicalId interface{}
	// Accessor of the mount to which the alias should belong to.
	MountAccessor interface{}
	// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
	Name interface{}
}
