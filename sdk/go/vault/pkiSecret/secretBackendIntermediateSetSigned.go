// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecretBackendIntermediateSetSigned struct {
	pulumi.CustomResourceState

	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
}

// NewSecretBackendIntermediateSetSigned registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendIntermediateSetSigned(ctx *pulumi.Context,
	name string, args *SecretBackendIntermediateSetSignedArgs, opts ...pulumi.ResourceOption) (*SecretBackendIntermediateSetSigned, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.Certificate == nil {
		return nil, errors.New("missing required argument 'Certificate'")
	}
	if args == nil {
		args = &SecretBackendIntermediateSetSignedArgs{}
	}
	var resource SecretBackendIntermediateSetSigned
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendIntermediateSetSigned gets an existing SecretBackendIntermediateSetSigned resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendIntermediateSetSigned(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendIntermediateSetSignedState, opts ...pulumi.ResourceOption) (*SecretBackendIntermediateSetSigned, error) {
	var resource SecretBackendIntermediateSetSigned
	err := ctx.ReadResource("vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendIntermediateSetSigned resources.
type secretBackendIntermediateSetSignedState struct {
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The certificate
	Certificate *string `pulumi:"certificate"`
}

type SecretBackendIntermediateSetSignedState struct {
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The certificate
	Certificate pulumi.StringPtrInput
}

func (SecretBackendIntermediateSetSignedState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIntermediateSetSignedState)(nil)).Elem()
}

type secretBackendIntermediateSetSignedArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// The certificate
	Certificate string `pulumi:"certificate"`
}

// The set of arguments for constructing a SecretBackendIntermediateSetSigned resource.
type SecretBackendIntermediateSetSignedArgs struct {
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// The certificate
	Certificate pulumi.StringInput
}

func (SecretBackendIntermediateSetSignedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendIntermediateSetSignedArgs)(nil)).Elem()
}
