// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_secret_backend_role.html.markdown.
type SecretBackendRole struct {
	s *pulumi.ResourceState
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOpt) (*SecretBackendRole, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.CredentialType == nil {
		return nil, errors.New("missing required argument 'CredentialType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["credentialType"] = nil
		inputs["defaultStsTtl"] = nil
		inputs["maxStsTtl"] = nil
		inputs["name"] = nil
		inputs["policyArns"] = nil
		inputs["policyDocument"] = nil
		inputs["roleArns"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["credentialType"] = args.CredentialType
		inputs["defaultStsTtl"] = args.DefaultStsTtl
		inputs["maxStsTtl"] = args.MaxStsTtl
		inputs["name"] = args.Name
		inputs["policyArns"] = args.PolicyArns
		inputs["policyDocument"] = args.PolicyDocument
		inputs["roleArns"] = args.RoleArns
	}
	s, err := ctx.RegisterResource("vault:aws/secretBackendRole:SecretBackendRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackendRole{s: s}, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretBackendRoleState, opts ...pulumi.ResourceOpt) (*SecretBackendRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backend"] = state.Backend
		inputs["credentialType"] = state.CredentialType
		inputs["defaultStsTtl"] = state.DefaultStsTtl
		inputs["maxStsTtl"] = state.MaxStsTtl
		inputs["name"] = state.Name
		inputs["policyArns"] = state.PolicyArns
		inputs["policyDocument"] = state.PolicyDocument
		inputs["roleArns"] = state.RoleArns
	}
	s, err := ctx.ReadResource("vault:aws/secretBackendRole:SecretBackendRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackendRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecretBackendRole) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecretBackendRole) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The path the AWS secret backend is mounted at,
// with no leading or trailing `/`s.
func (r *SecretBackendRole) Backend() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backend"])
}

// Specifies the type of credential to be used when
// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
// `federationToken`.
func (r *SecretBackendRole) CredentialType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["credentialType"])
}

// The default TTL in seconds for STS credentials.
// When a TTL is not specified when STS credentials are requested,
// and a default TTL is specified on the role,
// then this default TTL will be used. Valid only when `credentialType` is one of
// `assumedRole` or `federationToken`.
func (r *SecretBackendRole) DefaultStsTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["defaultStsTtl"])
}

// The max allowed TTL in seconds for STS credentials
// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
// one of `assumedRole` or `federationToken`.
func (r *SecretBackendRole) MaxStsTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxStsTtl"])
}

// The name to identify this role within the backend.
// Must be unique within the backend.
func (r *SecretBackendRole) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ARN for a pre-existing policy to associate
// with this role. Either `policyDocument` or `policyArns` must be specified.
func (r *SecretBackendRole) PolicyArns() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policyArns"])
}

// The JSON-formatted policy to associate with this
// role. Either `policyDocument` or `policyArns` must be specified.
func (r *SecretBackendRole) PolicyDocument() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyDocument"])
}

// Specifies the ARNs of the AWS roles this Vault role
// is allowed to assume. Required when `credentialType` is `assumedRole` and
// prohibited otherwise.
func (r *SecretBackendRole) RoleArns() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["roleArns"])
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type SecretBackendRoleState struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend interface{}
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType interface{}
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl interface{}
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl interface{}
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name interface{}
	// The ARN for a pre-existing policy to associate
	// with this role. Either `policyDocument` or `policyArns` must be specified.
	PolicyArns interface{}
	// The JSON-formatted policy to associate with this
	// role. Either `policyDocument` or `policyArns` must be specified.
	PolicyDocument interface{}
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns interface{}
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend interface{}
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType interface{}
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl interface{}
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl interface{}
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name interface{}
	// The ARN for a pre-existing policy to associate
	// with this role. Either `policyDocument` or `policyArns` must be specified.
	PolicyArns interface{}
	// The JSON-formatted policy to associate with this
	// role. Either `policyDocument` or `policyArns` must be specified.
	PolicyDocument interface{}
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns interface{}
}
