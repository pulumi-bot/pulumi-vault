// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gcp

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/gcp_secret_backend.html.md.
type SecretBackend struct {
	pulumi.CustomResourceState

	// The GCP service account credentials in JSON format.
	Credentials pulumi.StringPtrOutput `pulumi:"credentials"`
	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	DefaultLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	MaxLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"maxLeaseTtlSeconds"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `gcp`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		args = &SecretBackendArgs{}
	}
	var resource SecretBackend
	err := ctx.RegisterResource("vault:gcp/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:gcp/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// The GCP service account credentials in JSON format.
	Credentials *string `pulumi:"credentials"`
	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `gcp`.
	Path *string `pulumi:"path"`
}

type SecretBackendState struct {
	// The GCP service account credentials in JSON format.
	Credentials pulumi.StringPtrInput
	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `gcp`.
	Path pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// The GCP service account credentials in JSON format.
	Credentials *string `pulumi:"credentials"`
	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `gcp`.
	Path *string `pulumi:"path"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// The GCP service account credentials in JSON format.
	Credentials pulumi.StringPtrInput
	// The default TTL for credentials
	// issued by this backend. Defaults to '0'.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend. Defaults to '0'.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `gcp`.
	Path pulumi.StringPtrInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

