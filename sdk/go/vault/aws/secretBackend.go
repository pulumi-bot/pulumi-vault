// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_secret_backend.html.markdown.
type SecretBackend struct {
	s *pulumi.ResourceState
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOpt) (*SecretBackend, error) {
	if args == nil || args.AccessKey == nil {
		return nil, errors.New("missing required argument 'AccessKey'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessKey"] = nil
		inputs["defaultLeaseTtlSeconds"] = nil
		inputs["description"] = nil
		inputs["maxLeaseTtlSeconds"] = nil
		inputs["path"] = nil
		inputs["region"] = nil
		inputs["secretKey"] = nil
	} else {
		inputs["accessKey"] = args.AccessKey
		inputs["defaultLeaseTtlSeconds"] = args.DefaultLeaseTtlSeconds
		inputs["description"] = args.Description
		inputs["maxLeaseTtlSeconds"] = args.MaxLeaseTtlSeconds
		inputs["path"] = args.Path
		inputs["region"] = args.Region
		inputs["secretKey"] = args.SecretKey
	}
	s, err := ctx.RegisterResource("vault:aws/secretBackend:SecretBackend", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackend{s: s}, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretBackendState, opts ...pulumi.ResourceOpt) (*SecretBackend, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessKey"] = state.AccessKey
		inputs["defaultLeaseTtlSeconds"] = state.DefaultLeaseTtlSeconds
		inputs["description"] = state.Description
		inputs["maxLeaseTtlSeconds"] = state.MaxLeaseTtlSeconds
		inputs["path"] = state.Path
		inputs["region"] = state.Region
		inputs["secretKey"] = state.SecretKey
	}
	s, err := ctx.ReadResource("vault:aws/secretBackend:SecretBackend", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackend{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecretBackend) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecretBackend) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The AWS Access Key ID this backend should use to
// issue new credentials.
func (r *SecretBackend) AccessKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessKey"])
}

// The default TTL for credentials
// issued by this backend.
func (r *SecretBackend) DefaultLeaseTtlSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["defaultLeaseTtlSeconds"])
}

// A human-friendly description for this backend.
func (r *SecretBackend) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The maximum TTL that can be requested
// for credentials issued by this backend.
func (r *SecretBackend) MaxLeaseTtlSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxLeaseTtlSeconds"])
}

// The unique path this backend should be mounted at. Must
// not begin or end with a `/`. Defaults to `aws`.
func (r *SecretBackend) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// The AWS region for API calls. Defaults to `us-east-1`.
func (r *SecretBackend) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The AWS Secret Key this backend should use to
// issue new credentials.
func (r *SecretBackend) SecretKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["secretKey"])
}

// Input properties used for looking up and filtering SecretBackend resources.
type SecretBackendState struct {
	// The AWS Access Key ID this backend should use to
	// issue new credentials.
	AccessKey interface{}
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds interface{}
	// A human-friendly description for this backend.
	Description interface{}
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds interface{}
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path interface{}
	// The AWS region for API calls. Defaults to `us-east-1`.
	Region interface{}
	// The AWS Secret Key this backend should use to
	// issue new credentials.
	SecretKey interface{}
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// The AWS Access Key ID this backend should use to
	// issue new credentials.
	AccessKey interface{}
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds interface{}
	// A human-friendly description for this backend.
	Description interface{}
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds interface{}
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path interface{}
	// The AWS region for API calls. Defaults to `us-east-1`.
	Region interface{}
	// The AWS Secret Key this backend should use to
	// issue new credentials.
	SecretKey interface{}
}
