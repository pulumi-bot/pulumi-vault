// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transit

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configure the cache for the Transit Secret Backend in Vault.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/transit_secret_cache_config.html.markdown.
type SecretCacheConfig struct {
	s *pulumi.ResourceState
}

// NewSecretCacheConfig registers a new resource with the given unique name, arguments, and options.
func NewSecretCacheConfig(ctx *pulumi.Context,
	name string, args *SecretCacheConfigArgs, opts ...pulumi.ResourceOpt) (*SecretCacheConfig, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["size"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["size"] = args.Size
	}
	s, err := ctx.RegisterResource("vault:transit/secretCacheConfig:SecretCacheConfig", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretCacheConfig{s: s}, nil
}

// GetSecretCacheConfig gets an existing SecretCacheConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretCacheConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretCacheConfigState, opts ...pulumi.ResourceOpt) (*SecretCacheConfig, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backend"] = state.Backend
		inputs["size"] = state.Size
	}
	s, err := ctx.ReadResource("vault:transit/secretCacheConfig:SecretCacheConfig", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretCacheConfig{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecretCacheConfig) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecretCacheConfig) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
func (r *SecretCacheConfig) Backend() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["backend"])
}

// The number of cache entries. 0 means unlimited.
func (r *SecretCacheConfig) Size() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["size"])
}

// Input properties used for looking up and filtering SecretCacheConfig resources.
type SecretCacheConfigState struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend interface{}
	// The number of cache entries. 0 means unlimited.
	Size interface{}
}

// The set of arguments for constructing a SecretCacheConfig resource.
type SecretCacheConfigArgs struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend interface{}
	// The number of cache entries. 0 means unlimited.
	Size interface{}
}