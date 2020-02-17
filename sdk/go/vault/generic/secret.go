// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package generic

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Secret struct {
	pulumi.CustomResourceState

	// Map of strings read from Vault.
	Data pulumi.MapOutput `pulumi:"data"`
	// JSON-encoded secret data to write.
	DataJson pulumi.StringOutput `pulumi:"dataJson"`
	// Don't attempt to read the token from Vault if true; drift won't be detected.
	DisableRead pulumi.BoolPtrOutput `pulumi:"disableRead"`
	// Full path where the generic secret will be written.
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil || args.DataJson == nil {
		return nil, errors.New("missing required argument 'DataJson'")
	}
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &SecretArgs{}
	}
	var resource Secret
	err := ctx.RegisterResource("vault:generic/secret:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("vault:generic/secret:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	// Map of strings read from Vault.
	Data map[string]interface{} `pulumi:"data"`
	// JSON-encoded secret data to write.
	DataJson *string `pulumi:"dataJson"`
	// Don't attempt to read the token from Vault if true; drift won't be detected.
	DisableRead *bool `pulumi:"disableRead"`
	// Full path where the generic secret will be written.
	Path *string `pulumi:"path"`
}

type SecretState struct {
	// Map of strings read from Vault.
	Data pulumi.MapInput
	// JSON-encoded secret data to write.
	DataJson pulumi.StringPtrInput
	// Don't attempt to read the token from Vault if true; drift won't be detected.
	DisableRead pulumi.BoolPtrInput
	// Full path where the generic secret will be written.
	Path pulumi.StringPtrInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// JSON-encoded secret data to write.
	DataJson string `pulumi:"dataJson"`
	// Don't attempt to read the token from Vault if true; drift won't be detected.
	DisableRead *bool `pulumi:"disableRead"`
	// Full path where the generic secret will be written.
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// JSON-encoded secret data to write.
	DataJson pulumi.StringInput
	// Don't attempt to read the token from Vault if true; drift won't be detected.
	DisableRead pulumi.BoolPtrInput
	// Full path where the generic secret will be written.
	Path pulumi.StringInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

