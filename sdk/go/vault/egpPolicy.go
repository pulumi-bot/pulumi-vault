// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vault

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
// 
// **Note** this feature is available only with Vault Enterprise.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/egp_policy.html.markdown.
type EgpPolicy struct {
	pulumi.CustomResourceState

	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringOutput `pulumi:"enforcementLevel"`
	// The name of the policy
	Name pulumi.StringOutput `pulumi:"name"`
	// List of paths to which the policy will be applied to
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// String containing a Sentinel policy
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewEgpPolicy registers a new resource with the given unique name, arguments, and options.
func NewEgpPolicy(ctx *pulumi.Context,
	name string, args *EgpPolicyArgs, opts ...pulumi.ResourceOption) (*EgpPolicy, error) {
	if args == nil || args.EnforcementLevel == nil {
		return nil, errors.New("missing required argument 'EnforcementLevel'")
	}
	if args == nil || args.Paths == nil {
		return nil, errors.New("missing required argument 'Paths'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil {
		args = &EgpPolicyArgs{}
	}
	var resource EgpPolicy
	err := ctx.RegisterResource("vault:index/egpPolicy:EgpPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEgpPolicy gets an existing EgpPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEgpPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EgpPolicyState, opts ...pulumi.ResourceOption) (*EgpPolicy, error) {
	var resource EgpPolicy
	err := ctx.ReadResource("vault:index/egpPolicy:EgpPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EgpPolicy resources.
type egpPolicyState struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel *string `pulumi:"enforcementLevel"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// List of paths to which the policy will be applied to
	Paths []string `pulumi:"paths"`
	// String containing a Sentinel policy
	Policy *string `pulumi:"policy"`
}

type EgpPolicyState struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringPtrInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// List of paths to which the policy will be applied to
	Paths pulumi.StringArrayInput
	// String containing a Sentinel policy
	Policy pulumi.StringPtrInput
}

func (EgpPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*egpPolicyState)(nil)).Elem()
}

type egpPolicyArgs struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel string `pulumi:"enforcementLevel"`
	// The name of the policy
	Name *string `pulumi:"name"`
	// List of paths to which the policy will be applied to
	Paths []string `pulumi:"paths"`
	// String containing a Sentinel policy
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a EgpPolicy resource.
type EgpPolicyArgs struct {
	// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
	EnforcementLevel pulumi.StringInput
	// The name of the policy
	Name pulumi.StringPtrInput
	// List of paths to which the policy will be applied to
	Paths pulumi.StringArrayInput
	// String containing a Sentinel policy
	Policy pulumi.StringInput
}

func (EgpPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*egpPolicyArgs)(nil)).Elem()
}

