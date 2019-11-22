// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AuthBackend struct {
	s *pulumi.ResourceState
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOpt) (*AuthBackend, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultLeaseTtlSeconds"] = nil
		inputs["description"] = nil
		inputs["listingVisibility"] = nil
		inputs["local"] = nil
		inputs["maxLeaseTtlSeconds"] = nil
		inputs["path"] = nil
		inputs["type"] = nil
	} else {
		inputs["defaultLeaseTtlSeconds"] = args.DefaultLeaseTtlSeconds
		inputs["description"] = args.Description
		inputs["listingVisibility"] = args.ListingVisibility
		inputs["local"] = args.Local
		inputs["maxLeaseTtlSeconds"] = args.MaxLeaseTtlSeconds
		inputs["path"] = args.Path
		inputs["type"] = args.Type
	}
	inputs["accessor"] = nil
	s, err := ctx.RegisterResource("vault:index/authBackend:AuthBackend", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackend{s: s}, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AuthBackendState, opts ...pulumi.ResourceOpt) (*AuthBackend, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessor"] = state.Accessor
		inputs["defaultLeaseTtlSeconds"] = state.DefaultLeaseTtlSeconds
		inputs["description"] = state.Description
		inputs["listingVisibility"] = state.ListingVisibility
		inputs["local"] = state.Local
		inputs["maxLeaseTtlSeconds"] = state.MaxLeaseTtlSeconds
		inputs["path"] = state.Path
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("vault:index/authBackend:AuthBackend", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AuthBackend{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AuthBackend) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AuthBackend) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The accessor for this auth method
func (r *AuthBackend) Accessor() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accessor"])
}

// The default lease duration in seconds.
func (r *AuthBackend) DefaultLeaseTtlSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["defaultLeaseTtlSeconds"])
}

// A description of the auth method
func (r *AuthBackend) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Speficies whether to show this mount in the UI-specific listing endpoint.
func (r *AuthBackend) ListingVisibility() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["listingVisibility"])
}

// Specifies if the auth method is local only.
func (r *AuthBackend) Local() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["local"])
}

// The maximum lease duration in seconds.
func (r *AuthBackend) MaxLeaseTtlSeconds() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxLeaseTtlSeconds"])
}

// The path to mount the auth method — this defaults to the name of the type
func (r *AuthBackend) Path() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["path"])
}

// The name of the auth method type
func (r *AuthBackend) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering AuthBackend resources.
type AuthBackendState struct {
	// The accessor for this auth method
	Accessor interface{}
	// The default lease duration in seconds.
	DefaultLeaseTtlSeconds interface{}
	// A description of the auth method
	Description interface{}
	// Speficies whether to show this mount in the UI-specific listing endpoint.
	ListingVisibility interface{}
	// Specifies if the auth method is local only.
	Local interface{}
	// The maximum lease duration in seconds.
	MaxLeaseTtlSeconds interface{}
	// The path to mount the auth method — this defaults to the name of the type
	Path interface{}
	// The name of the auth method type
	Type interface{}
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// The default lease duration in seconds.
	DefaultLeaseTtlSeconds interface{}
	// A description of the auth method
	Description interface{}
	// Speficies whether to show this mount in the UI-specific listing endpoint.
	ListingVisibility interface{}
	// Specifies if the auth method is local only.
	Local interface{}
	// The maximum lease duration in seconds.
	MaxLeaseTtlSeconds interface{}
	// The path to mount the auth method — this defaults to the name of the type
	Path interface{}
	// The name of the auth method type
	Type interface{}
}
