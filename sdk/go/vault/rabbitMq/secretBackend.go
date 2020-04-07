// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitMq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecretBackend struct {
	pulumi.CustomResourceState

	// Specifies the RabbitMQ connection URI.
	ConnectionUri pulumi.StringOutput `pulumi:"connectionUri"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// Specifies the RabbitMQ management administrator password.
	Password pulumi.StringOutput `pulumi:"password"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Specifies the RabbitMQ management administrator username.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection pulumi.BoolPtrOutput `pulumi:"verifyConnection"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil || args.ConnectionUri == nil {
		return nil, errors.New("missing required argument 'ConnectionUri'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &SecretBackendArgs{}
	}
	var resource SecretBackend
	err := ctx.RegisterResource("vault:rabbitMq/secretBackend:SecretBackend", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:rabbitMq/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri *string `pulumi:"connectionUri"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Specifies the RabbitMQ management administrator password.
	Password *string `pulumi:"password"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path *string `pulumi:"path"`
	// Specifies the RabbitMQ management administrator username.
	Username *string `pulumi:"username"`
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

type SecretBackendState struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri pulumi.StringPtrInput
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Specifies the RabbitMQ management administrator password.
	Password pulumi.StringPtrInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path pulumi.StringPtrInput
	// Specifies the RabbitMQ management administrator username.
	Username pulumi.StringPtrInput
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri string `pulumi:"connectionUri"`
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Specifies the RabbitMQ management administrator password.
	Password string `pulumi:"password"`
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path *string `pulumi:"path"`
	// Specifies the RabbitMQ management administrator username.
	Username string `pulumi:"username"`
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection *bool `pulumi:"verifyConnection"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Specifies the RabbitMQ connection URI.
	ConnectionUri pulumi.StringInput
	// The default TTL for credentials
	// issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Specifies the RabbitMQ management administrator password.
	Password pulumi.StringInput
	// The unique path this backend should be mounted at. Must
	// not begin or end with a `/`. Defaults to `aws`.
	Path pulumi.StringPtrInput
	// Specifies the RabbitMQ management administrator username.
	Username pulumi.StringInput
	// Specifies whether to verify connection URI, username, and password.
	// Defaults to `true`.
	VerifyConnection pulumi.BoolPtrInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}
