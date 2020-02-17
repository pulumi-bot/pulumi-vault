// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_config_ca.html.md.
type SecretBackendConfigCa struct {
	pulumi.CustomResourceState

	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The key and certificate PEM bundle
	PemBundle pulumi.StringOutput `pulumi:"pemBundle"`
}

// NewSecretBackendConfigCa registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendConfigCa(ctx *pulumi.Context,
	name string, args *SecretBackendConfigCaArgs, opts ...pulumi.ResourceOption) (*SecretBackendConfigCa, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.PemBundle == nil {
		return nil, errors.New("missing required argument 'PemBundle'")
	}
	if args == nil {
		args = &SecretBackendConfigCaArgs{}
	}
	var resource SecretBackendConfigCa
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendConfigCa gets an existing SecretBackendConfigCa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendConfigCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendConfigCaState, opts ...pulumi.ResourceOption) (*SecretBackendConfigCa, error) {
	var resource SecretBackendConfigCa
	err := ctx.ReadResource("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendConfigCa resources.
type secretBackendConfigCaState struct {
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The key and certificate PEM bundle
	PemBundle *string `pulumi:"pemBundle"`
}

type SecretBackendConfigCaState struct {
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The key and certificate PEM bundle
	PemBundle pulumi.StringPtrInput
}

func (SecretBackendConfigCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigCaState)(nil)).Elem()
}

type secretBackendConfigCaArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// The key and certificate PEM bundle
	PemBundle string `pulumi:"pemBundle"`
}

// The set of arguments for constructing a SecretBackendConfigCa resource.
type SecretBackendConfigCaArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// The key and certificate PEM bundle
	PemBundle pulumi.StringInput
}

func (SecretBackendConfigCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendConfigCaArgs)(nil)).Elem()
}

