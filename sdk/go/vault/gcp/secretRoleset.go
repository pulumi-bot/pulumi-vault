// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
// 
// Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/gcp_secret_roleset.html.markdown.
type SecretRoleset struct {
	s *pulumi.ResourceState
}

// NewSecretRoleset registers a new resource with the given unique name, arguments, and options.
func NewSecretRoleset(ctx *pulumi.Context,
	name string, args *SecretRolesetArgs, opts ...pulumi.ResourceOpt) (*SecretRoleset, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.Bindings == nil {
		return nil, errors.New("missing required argument 'Bindings'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil || args.Roleset == nil {
		return nil, errors.New("missing required argument 'Roleset'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["backend"] = nil
		inputs["bindings"] = nil
		inputs["project"] = nil
		inputs["roleset"] = nil
		inputs["secretType"] = nil
		inputs["tokenScopes"] = nil
	} else {
		inputs["backend"] = args.Backend
		inputs["bindings"] = args.Bindings
		inputs["project"] = args.Project
		inputs["roleset"] = args.Roleset
		inputs["secretType"] = args.SecretType
		inputs["tokenScopes"] = args.TokenScopes
	}
	inputs["serviceAccountEmail"] = nil
	s, err := ctx.RegisterResource("vault:gcp/secretRoleset:SecretRoleset", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretRoleset{s: s}, nil
}

// GetSecretRoleset gets an existing SecretRoleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretRoleset(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretRolesetState, opts ...pulumi.ResourceOpt) (*SecretRoleset, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["backend"] = state.Backend
		inputs["bindings"] = state.Bindings
		inputs["project"] = state.Project
		inputs["roleset"] = state.Roleset
		inputs["secretType"] = state.SecretType
		inputs["serviceAccountEmail"] = state.ServiceAccountEmail
		inputs["tokenScopes"] = state.TokenScopes
	}
	s, err := ctx.ReadResource("vault:gcp/secretRoleset:SecretRoleset", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretRoleset{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecretRoleset) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecretRoleset) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Path where the GCP Secrets Engine is mounted
func (r *SecretRoleset) Backend() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backend"])
}

// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
func (r *SecretRoleset) Bindings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["bindings"])
}

// Name of the GCP project that this roleset's service account will belong to.
func (r *SecretRoleset) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// Name of the Roleset to create
func (r *SecretRoleset) Roleset() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleset"])
}

// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
func (r *SecretRoleset) SecretType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["secretType"])
}

// Email of the service account created by Vault for this Roleset
func (r *SecretRoleset) ServiceAccountEmail() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceAccountEmail"])
}

// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
func (r *SecretRoleset) TokenScopes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tokenScopes"])
}

// Input properties used for looking up and filtering SecretRoleset resources.
type SecretRolesetState struct {
	// Path where the GCP Secrets Engine is mounted
	Backend interface{}
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings interface{}
	// Name of the GCP project that this roleset's service account will belong to.
	Project interface{}
	// Name of the Roleset to create
	Roleset interface{}
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType interface{}
	// Email of the service account created by Vault for this Roleset
	ServiceAccountEmail interface{}
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes interface{}
}

// The set of arguments for constructing a SecretRoleset resource.
type SecretRolesetArgs struct {
	// Path where the GCP Secrets Engine is mounted
	Backend interface{}
	// Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
	Bindings interface{}
	// Name of the GCP project that this roleset's service account will belong to.
	Project interface{}
	// Name of the Roleset to create
	Roleset interface{}
	// Type of secret generated for this role set. Accepted values: `accessToken`, `serviceAccountKey`. Defaults to `accessToken`.
	SecretType interface{}
	// List of OAuth scopes to assign to `accessToken` secrets generated under this role set (`accessToken` role sets only).
	TokenScopes interface{}
}
