// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures the periodic tidying operation of the blacklisted role tag entries.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_roletag_blacklist.html.md.
type AuthBackendRoletagBlacklist struct {
	pulumi.CustomResourceState

	// The path the AWS auth backend being configured was
	// mounted at.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy pulumi.BoolPtrOutput `pulumi:"disablePeriodicTidy"`
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer pulumi.IntPtrOutput `pulumi:"safetyBuffer"`
}

// NewAuthBackendRoletagBlacklist registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRoletagBlacklist(ctx *pulumi.Context,
	name string, args *AuthBackendRoletagBlacklistArgs, opts ...pulumi.ResourceOption) (*AuthBackendRoletagBlacklist, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil {
		args = &AuthBackendRoletagBlacklistArgs{}
	}
	var resource AuthBackendRoletagBlacklist
	err := ctx.RegisterResource("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRoletagBlacklist gets an existing AuthBackendRoletagBlacklist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRoletagBlacklist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoletagBlacklistState, opts ...pulumi.ResourceOption) (*AuthBackendRoletagBlacklist, error) {
	var resource AuthBackendRoletagBlacklist
	err := ctx.ReadResource("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRoletagBlacklist resources.
type authBackendRoletagBlacklistState struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend *string `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy *bool `pulumi:"disablePeriodicTidy"`
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer *int `pulumi:"safetyBuffer"`
}

type AuthBackendRoletagBlacklistState struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend pulumi.StringPtrInput
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy pulumi.BoolPtrInput
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer pulumi.IntPtrInput
}

func (AuthBackendRoletagBlacklistState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoletagBlacklistState)(nil)).Elem()
}

type authBackendRoletagBlacklistArgs struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend string `pulumi:"backend"`
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy *bool `pulumi:"disablePeriodicTidy"`
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer *int `pulumi:"safetyBuffer"`
}

// The set of arguments for constructing a AuthBackendRoletagBlacklist resource.
type AuthBackendRoletagBlacklistArgs struct {
	// The path the AWS auth backend being configured was
	// mounted at.
	Backend pulumi.StringInput
	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	DisablePeriodicTidy pulumi.BoolPtrInput
	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	SafetyBuffer pulumi.IntPtrInput
}

func (AuthBackendRoletagBlacklistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoletagBlacklistArgs)(nil)).Elem()
}
