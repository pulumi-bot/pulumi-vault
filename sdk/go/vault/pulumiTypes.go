// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AuthBackendTune struct {
	// List of headers to whitelist and allowing
	// a plugin to include them in the response.
	AllowedResponseHeaders []string `pulumi:"allowedResponseHeaders"`
	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys []string `pulumi:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys []string `pulumi:"auditNonHmacResponseKeys"`
	// Specifies the default time-to-live.
	// If set, this overrides the global default.
	// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
	DefaultLeaseTtl *string `pulumi:"defaultLeaseTtl"`
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	ListingVisibility *string `pulumi:"listingVisibility"`
	// Specifies the maximum time-to-live.
	// If set, this overrides the global default.
	// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
	MaxLeaseTtl *string `pulumi:"maxLeaseTtl"`
	// List of headers to whitelist and
	// pass from the request to the backend.
	PassthroughRequestHeaders []string `pulumi:"passthroughRequestHeaders"`
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType *string `pulumi:"tokenType"`
}

// AuthBackendTuneInput is an input type that accepts AuthBackendTuneArgs and AuthBackendTuneOutput values.
// You can construct a concrete instance of `AuthBackendTuneInput` via:
//
// 		 AuthBackendTuneArgs{...}
//
type AuthBackendTuneInput interface {
	pulumi.Input

	ToAuthBackendTuneOutput() AuthBackendTuneOutput
	ToAuthBackendTuneOutputWithContext(context.Context) AuthBackendTuneOutput
}

type AuthBackendTuneArgs struct {
	// List of headers to whitelist and allowing
	// a plugin to include them in the response.
	AllowedResponseHeaders pulumi.StringArrayInput `pulumi:"allowedResponseHeaders"`
	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the request data object.
	AuditNonHmacRequestKeys pulumi.StringArrayInput `pulumi:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will
	// not be HMAC'd by audit devices in the response data object.
	AuditNonHmacResponseKeys pulumi.StringArrayInput `pulumi:"auditNonHmacResponseKeys"`
	// Specifies the default time-to-live.
	// If set, this overrides the global default.
	// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
	DefaultLeaseTtl pulumi.StringPtrInput `pulumi:"defaultLeaseTtl"`
	// Specifies whether to show this mount in
	// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
	ListingVisibility pulumi.StringPtrInput `pulumi:"listingVisibility"`
	// Specifies the maximum time-to-live.
	// If set, this overrides the global default.
	// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
	MaxLeaseTtl pulumi.StringPtrInput `pulumi:"maxLeaseTtl"`
	// List of headers to whitelist and
	// pass from the request to the backend.
	PassthroughRequestHeaders pulumi.StringArrayInput `pulumi:"passthroughRequestHeaders"`
	// Specifies the type of tokens that should be returned by
	// the mount. Valid values are "default-service", "default-batch", "service", "batch".
	TokenType pulumi.StringPtrInput `pulumi:"tokenType"`
}

func (AuthBackendTuneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendTune)(nil)).Elem()
}

func (i AuthBackendTuneArgs) ToAuthBackendTuneOutput() AuthBackendTuneOutput {
	return i.ToAuthBackendTuneOutputWithContext(context.Background())
}

func (i AuthBackendTuneArgs) ToAuthBackendTuneOutputWithContext(ctx context.Context) AuthBackendTuneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTuneOutput)
}

func (i AuthBackendTuneArgs) ToAuthBackendTunePtrOutput() AuthBackendTunePtrOutput {
	return i.ToAuthBackendTunePtrOutputWithContext(context.Background())
}

func (i AuthBackendTuneArgs) ToAuthBackendTunePtrOutputWithContext(ctx context.Context) AuthBackendTunePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTuneOutput).ToAuthBackendTunePtrOutputWithContext(ctx)
}

// AuthBackendTunePtrInput is an input type that accepts AuthBackendTuneArgs, AuthBackendTunePtr and AuthBackendTunePtrOutput values.
// You can construct a concrete instance of `AuthBackendTunePtrInput` via:
//
// 		 AuthBackendTuneArgs{...}
//
//  or:
//
// 		 nil
//
type AuthBackendTunePtrInput interface {
	pulumi.Input

	ToAuthBackendTunePtrOutput() AuthBackendTunePtrOutput
	ToAuthBackendTunePtrOutputWithContext(context.Context) AuthBackendTunePtrOutput
}

type authBackendTunePtrType AuthBackendTuneArgs

func AuthBackendTunePtr(v *AuthBackendTuneArgs) AuthBackendTunePtrInput {
	return (*authBackendTunePtrType)(v)
}

func (*authBackendTunePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendTune)(nil)).Elem()
}

func (i *authBackendTunePtrType) ToAuthBackendTunePtrOutput() AuthBackendTunePtrOutput {
	return i.ToAuthBackendTunePtrOutputWithContext(context.Background())
}

func (i *authBackendTunePtrType) ToAuthBackendTunePtrOutputWithContext(ctx context.Context) AuthBackendTunePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTunePtrOutput)
}

type AuthBackendTuneOutput struct{ *pulumi.OutputState }

func (AuthBackendTuneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendTune)(nil)).Elem()
}

func (o AuthBackendTuneOutput) ToAuthBackendTuneOutput() AuthBackendTuneOutput {
	return o
}

func (o AuthBackendTuneOutput) ToAuthBackendTuneOutputWithContext(ctx context.Context) AuthBackendTuneOutput {
	return o
}

func (o AuthBackendTuneOutput) ToAuthBackendTunePtrOutput() AuthBackendTunePtrOutput {
	return o.ToAuthBackendTunePtrOutputWithContext(context.Background())
}

func (o AuthBackendTuneOutput) ToAuthBackendTunePtrOutputWithContext(ctx context.Context) AuthBackendTunePtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *AuthBackendTune {
		return &v
	}).(AuthBackendTunePtrOutput)
}

// List of headers to whitelist and allowing
// a plugin to include them in the response.
func (o AuthBackendTuneOutput) AllowedResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.AllowedResponseHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the request data object.
func (o AuthBackendTuneOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.AuditNonHmacRequestKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the response data object.
func (o AuthBackendTuneOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.AuditNonHmacResponseKeys }).(pulumi.StringArrayOutput)
}

// Specifies the default time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneOutput) DefaultLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.DefaultLeaseTtl }).(pulumi.StringPtrOutput)
}

// Specifies whether to show this mount in
// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
func (o AuthBackendTuneOutput) ListingVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.ListingVisibility }).(pulumi.StringPtrOutput)
}

// Specifies the maximum time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneOutput) MaxLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.MaxLeaseTtl }).(pulumi.StringPtrOutput)
}

// List of headers to whitelist and
// pass from the request to the backend.
func (o AuthBackendTuneOutput) PassthroughRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.PassthroughRequestHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the type of tokens that should be returned by
// the mount. Valid values are "default-service", "default-batch", "service", "batch".
func (o AuthBackendTuneOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type AuthBackendTunePtrOutput struct{ *pulumi.OutputState }

func (AuthBackendTunePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendTune)(nil)).Elem()
}

func (o AuthBackendTunePtrOutput) ToAuthBackendTunePtrOutput() AuthBackendTunePtrOutput {
	return o
}

func (o AuthBackendTunePtrOutput) ToAuthBackendTunePtrOutputWithContext(ctx context.Context) AuthBackendTunePtrOutput {
	return o
}

func (o AuthBackendTunePtrOutput) Elem() AuthBackendTuneOutput {
	return o.ApplyT(func(v *AuthBackendTune) AuthBackendTune { return *v }).(AuthBackendTuneOutput)
}

// List of headers to whitelist and allowing
// a plugin to include them in the response.
func (o AuthBackendTunePtrOutput) AllowedResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendTune) []string {
		if v == nil {
			return nil
		}
		return v.AllowedResponseHeaders
	}).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the request data object.
func (o AuthBackendTunePtrOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendTune) []string {
		if v == nil {
			return nil
		}
		return v.AuditNonHmacRequestKeys
	}).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the response data object.
func (o AuthBackendTunePtrOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendTune) []string {
		if v == nil {
			return nil
		}
		return v.AuditNonHmacResponseKeys
	}).(pulumi.StringArrayOutput)
}

// Specifies the default time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTunePtrOutput) DefaultLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendTune) *string {
		if v == nil {
			return nil
		}
		return v.DefaultLeaseTtl
	}).(pulumi.StringPtrOutput)
}

// Specifies whether to show this mount in
// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
func (o AuthBackendTunePtrOutput) ListingVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendTune) *string {
		if v == nil {
			return nil
		}
		return v.ListingVisibility
	}).(pulumi.StringPtrOutput)
}

// Specifies the maximum time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTunePtrOutput) MaxLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendTune) *string {
		if v == nil {
			return nil
		}
		return v.MaxLeaseTtl
	}).(pulumi.StringPtrOutput)
}

// List of headers to whitelist and
// pass from the request to the backend.
func (o AuthBackendTunePtrOutput) PassthroughRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthBackendTune) []string {
		if v == nil {
			return nil
		}
		return v.PassthroughRequestHeaders
	}).(pulumi.StringArrayOutput)
}

// Specifies the type of tokens that should be returned by
// the mount. Valid values are "default-service", "default-batch", "service", "batch".
func (o AuthBackendTunePtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthBackendTune) *string {
		if v == nil {
			return nil
		}
		return v.TokenType
	}).(pulumi.StringPtrOutput)
}

type GetPolicyDocumentRule struct {
	// Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
	AllowedParameters []GetPolicyDocumentRuleAllowedParameter `pulumi:"allowedParameters"`
	// A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
	Capabilities []string `pulumi:"capabilities"`
	// Blacklists a list of parameter and values. Any values specified here take precedence over `allowedParameter`. See Parameters below.
	DeniedParameters []GetPolicyDocumentRuleDeniedParameter `pulumi:"deniedParameters"`
	// Description of the rule. Will be added as a commend to rendered rule.
	Description *string `pulumi:"description"`
	// The maximum allowed TTL that clients can specify for a wrapped response.
	MaxWrappingTtl *string `pulumi:"maxWrappingTtl"`
	// The minimum allowed TTL that clients can specify for a wrapped response.
	MinWrappingTtl *string `pulumi:"minWrappingTtl"`
	// A path in Vault that this rule applies to.
	Path string `pulumi:"path"`
	// A list of parameters that must be specified.
	RequiredParameters []string `pulumi:"requiredParameters"`
}

// GetPolicyDocumentRuleInput is an input type that accepts GetPolicyDocumentRuleArgs and GetPolicyDocumentRuleOutput values.
// You can construct a concrete instance of `GetPolicyDocumentRuleInput` via:
//
// 		 GetPolicyDocumentRuleArgs{...}
//
type GetPolicyDocumentRuleInput interface {
	pulumi.Input

	ToGetPolicyDocumentRuleOutput() GetPolicyDocumentRuleOutput
	ToGetPolicyDocumentRuleOutputWithContext(context.Context) GetPolicyDocumentRuleOutput
}

type GetPolicyDocumentRuleArgs struct {
	// Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
	AllowedParameters GetPolicyDocumentRuleAllowedParameterArrayInput `pulumi:"allowedParameters"`
	// A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
	Capabilities pulumi.StringArrayInput `pulumi:"capabilities"`
	// Blacklists a list of parameter and values. Any values specified here take precedence over `allowedParameter`. See Parameters below.
	DeniedParameters GetPolicyDocumentRuleDeniedParameterArrayInput `pulumi:"deniedParameters"`
	// Description of the rule. Will be added as a commend to rendered rule.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The maximum allowed TTL that clients can specify for a wrapped response.
	MaxWrappingTtl pulumi.StringPtrInput `pulumi:"maxWrappingTtl"`
	// The minimum allowed TTL that clients can specify for a wrapped response.
	MinWrappingTtl pulumi.StringPtrInput `pulumi:"minWrappingTtl"`
	// A path in Vault that this rule applies to.
	Path pulumi.StringInput `pulumi:"path"`
	// A list of parameters that must be specified.
	RequiredParameters pulumi.StringArrayInput `pulumi:"requiredParameters"`
}

func (GetPolicyDocumentRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentRule)(nil)).Elem()
}

func (i GetPolicyDocumentRuleArgs) ToGetPolicyDocumentRuleOutput() GetPolicyDocumentRuleOutput {
	return i.ToGetPolicyDocumentRuleOutputWithContext(context.Background())
}

func (i GetPolicyDocumentRuleArgs) ToGetPolicyDocumentRuleOutputWithContext(ctx context.Context) GetPolicyDocumentRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentRuleOutput)
}

// GetPolicyDocumentRuleArrayInput is an input type that accepts GetPolicyDocumentRuleArray and GetPolicyDocumentRuleArrayOutput values.
// You can construct a concrete instance of `GetPolicyDocumentRuleArrayInput` via:
//
// 		 GetPolicyDocumentRuleArray{ GetPolicyDocumentRuleArgs{...} }
//
type GetPolicyDocumentRuleArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentRuleArrayOutput() GetPolicyDocumentRuleArrayOutput
	ToGetPolicyDocumentRuleArrayOutputWithContext(context.Context) GetPolicyDocumentRuleArrayOutput
}

type GetPolicyDocumentRuleArray []GetPolicyDocumentRuleInput

func (GetPolicyDocumentRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentRule)(nil)).Elem()
}

func (i GetPolicyDocumentRuleArray) ToGetPolicyDocumentRuleArrayOutput() GetPolicyDocumentRuleArrayOutput {
	return i.ToGetPolicyDocumentRuleArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentRuleArray) ToGetPolicyDocumentRuleArrayOutputWithContext(ctx context.Context) GetPolicyDocumentRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentRuleArrayOutput)
}

type GetPolicyDocumentRuleOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentRule)(nil)).Elem()
}

func (o GetPolicyDocumentRuleOutput) ToGetPolicyDocumentRuleOutput() GetPolicyDocumentRuleOutput {
	return o
}

func (o GetPolicyDocumentRuleOutput) ToGetPolicyDocumentRuleOutputWithContext(ctx context.Context) GetPolicyDocumentRuleOutput {
	return o
}

// Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
func (o GetPolicyDocumentRuleOutput) AllowedParameters() GetPolicyDocumentRuleAllowedParameterArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) []GetPolicyDocumentRuleAllowedParameter { return v.AllowedParameters }).(GetPolicyDocumentRuleAllowedParameterArrayOutput)
}

// A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
func (o GetPolicyDocumentRuleOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) []string { return v.Capabilities }).(pulumi.StringArrayOutput)
}

// Blacklists a list of parameter and values. Any values specified here take precedence over `allowedParameter`. See Parameters below.
func (o GetPolicyDocumentRuleOutput) DeniedParameters() GetPolicyDocumentRuleDeniedParameterArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) []GetPolicyDocumentRuleDeniedParameter { return v.DeniedParameters }).(GetPolicyDocumentRuleDeniedParameterArrayOutput)
}

// Description of the rule. Will be added as a commend to rendered rule.
func (o GetPolicyDocumentRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The maximum allowed TTL that clients can specify for a wrapped response.
func (o GetPolicyDocumentRuleOutput) MaxWrappingTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) *string { return v.MaxWrappingTtl }).(pulumi.StringPtrOutput)
}

// The minimum allowed TTL that clients can specify for a wrapped response.
func (o GetPolicyDocumentRuleOutput) MinWrappingTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) *string { return v.MinWrappingTtl }).(pulumi.StringPtrOutput)
}

// A path in Vault that this rule applies to.
func (o GetPolicyDocumentRuleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) string { return v.Path }).(pulumi.StringOutput)
}

// A list of parameters that must be specified.
func (o GetPolicyDocumentRuleOutput) RequiredParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentRule) []string { return v.RequiredParameters }).(pulumi.StringArrayOutput)
}

type GetPolicyDocumentRuleArrayOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentRule)(nil)).Elem()
}

func (o GetPolicyDocumentRuleArrayOutput) ToGetPolicyDocumentRuleArrayOutput() GetPolicyDocumentRuleArrayOutput {
	return o
}

func (o GetPolicyDocumentRuleArrayOutput) ToGetPolicyDocumentRuleArrayOutputWithContext(ctx context.Context) GetPolicyDocumentRuleArrayOutput {
	return o
}

func (o GetPolicyDocumentRuleArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetPolicyDocumentRule {
		return vs[0].([]GetPolicyDocumentRule)[vs[1].(int)]
	}).(GetPolicyDocumentRuleOutput)
}

type GetPolicyDocumentRuleAllowedParameter struct {
	// name of permitted or denied parameter.
	Key string `pulumi:"key"`
	// list of values what are permitted or denied by policy rule.
	Values []string `pulumi:"values"`
}

// GetPolicyDocumentRuleAllowedParameterInput is an input type that accepts GetPolicyDocumentRuleAllowedParameterArgs and GetPolicyDocumentRuleAllowedParameterOutput values.
// You can construct a concrete instance of `GetPolicyDocumentRuleAllowedParameterInput` via:
//
// 		 GetPolicyDocumentRuleAllowedParameterArgs{...}
//
type GetPolicyDocumentRuleAllowedParameterInput interface {
	pulumi.Input

	ToGetPolicyDocumentRuleAllowedParameterOutput() GetPolicyDocumentRuleAllowedParameterOutput
	ToGetPolicyDocumentRuleAllowedParameterOutputWithContext(context.Context) GetPolicyDocumentRuleAllowedParameterOutput
}

type GetPolicyDocumentRuleAllowedParameterArgs struct {
	// name of permitted or denied parameter.
	Key pulumi.StringInput `pulumi:"key"`
	// list of values what are permitted or denied by policy rule.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GetPolicyDocumentRuleAllowedParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentRuleAllowedParameter)(nil)).Elem()
}

func (i GetPolicyDocumentRuleAllowedParameterArgs) ToGetPolicyDocumentRuleAllowedParameterOutput() GetPolicyDocumentRuleAllowedParameterOutput {
	return i.ToGetPolicyDocumentRuleAllowedParameterOutputWithContext(context.Background())
}

func (i GetPolicyDocumentRuleAllowedParameterArgs) ToGetPolicyDocumentRuleAllowedParameterOutputWithContext(ctx context.Context) GetPolicyDocumentRuleAllowedParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentRuleAllowedParameterOutput)
}

// GetPolicyDocumentRuleAllowedParameterArrayInput is an input type that accepts GetPolicyDocumentRuleAllowedParameterArray and GetPolicyDocumentRuleAllowedParameterArrayOutput values.
// You can construct a concrete instance of `GetPolicyDocumentRuleAllowedParameterArrayInput` via:
//
// 		 GetPolicyDocumentRuleAllowedParameterArray{ GetPolicyDocumentRuleAllowedParameterArgs{...} }
//
type GetPolicyDocumentRuleAllowedParameterArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentRuleAllowedParameterArrayOutput() GetPolicyDocumentRuleAllowedParameterArrayOutput
	ToGetPolicyDocumentRuleAllowedParameterArrayOutputWithContext(context.Context) GetPolicyDocumentRuleAllowedParameterArrayOutput
}

type GetPolicyDocumentRuleAllowedParameterArray []GetPolicyDocumentRuleAllowedParameterInput

func (GetPolicyDocumentRuleAllowedParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentRuleAllowedParameter)(nil)).Elem()
}

func (i GetPolicyDocumentRuleAllowedParameterArray) ToGetPolicyDocumentRuleAllowedParameterArrayOutput() GetPolicyDocumentRuleAllowedParameterArrayOutput {
	return i.ToGetPolicyDocumentRuleAllowedParameterArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentRuleAllowedParameterArray) ToGetPolicyDocumentRuleAllowedParameterArrayOutputWithContext(ctx context.Context) GetPolicyDocumentRuleAllowedParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentRuleAllowedParameterArrayOutput)
}

type GetPolicyDocumentRuleAllowedParameterOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentRuleAllowedParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentRuleAllowedParameter)(nil)).Elem()
}

func (o GetPolicyDocumentRuleAllowedParameterOutput) ToGetPolicyDocumentRuleAllowedParameterOutput() GetPolicyDocumentRuleAllowedParameterOutput {
	return o
}

func (o GetPolicyDocumentRuleAllowedParameterOutput) ToGetPolicyDocumentRuleAllowedParameterOutputWithContext(ctx context.Context) GetPolicyDocumentRuleAllowedParameterOutput {
	return o
}

// name of permitted or denied parameter.
func (o GetPolicyDocumentRuleAllowedParameterOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentRuleAllowedParameter) string { return v.Key }).(pulumi.StringOutput)
}

// list of values what are permitted or denied by policy rule.
func (o GetPolicyDocumentRuleAllowedParameterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentRuleAllowedParameter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GetPolicyDocumentRuleAllowedParameterArrayOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentRuleAllowedParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentRuleAllowedParameter)(nil)).Elem()
}

func (o GetPolicyDocumentRuleAllowedParameterArrayOutput) ToGetPolicyDocumentRuleAllowedParameterArrayOutput() GetPolicyDocumentRuleAllowedParameterArrayOutput {
	return o
}

func (o GetPolicyDocumentRuleAllowedParameterArrayOutput) ToGetPolicyDocumentRuleAllowedParameterArrayOutputWithContext(ctx context.Context) GetPolicyDocumentRuleAllowedParameterArrayOutput {
	return o
}

func (o GetPolicyDocumentRuleAllowedParameterArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentRuleAllowedParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetPolicyDocumentRuleAllowedParameter {
		return vs[0].([]GetPolicyDocumentRuleAllowedParameter)[vs[1].(int)]
	}).(GetPolicyDocumentRuleAllowedParameterOutput)
}

type GetPolicyDocumentRuleDeniedParameter struct {
	// name of permitted or denied parameter.
	Key string `pulumi:"key"`
	// list of values what are permitted or denied by policy rule.
	Values []string `pulumi:"values"`
}

// GetPolicyDocumentRuleDeniedParameterInput is an input type that accepts GetPolicyDocumentRuleDeniedParameterArgs and GetPolicyDocumentRuleDeniedParameterOutput values.
// You can construct a concrete instance of `GetPolicyDocumentRuleDeniedParameterInput` via:
//
// 		 GetPolicyDocumentRuleDeniedParameterArgs{...}
//
type GetPolicyDocumentRuleDeniedParameterInput interface {
	pulumi.Input

	ToGetPolicyDocumentRuleDeniedParameterOutput() GetPolicyDocumentRuleDeniedParameterOutput
	ToGetPolicyDocumentRuleDeniedParameterOutputWithContext(context.Context) GetPolicyDocumentRuleDeniedParameterOutput
}

type GetPolicyDocumentRuleDeniedParameterArgs struct {
	// name of permitted or denied parameter.
	Key pulumi.StringInput `pulumi:"key"`
	// list of values what are permitted or denied by policy rule.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GetPolicyDocumentRuleDeniedParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentRuleDeniedParameter)(nil)).Elem()
}

func (i GetPolicyDocumentRuleDeniedParameterArgs) ToGetPolicyDocumentRuleDeniedParameterOutput() GetPolicyDocumentRuleDeniedParameterOutput {
	return i.ToGetPolicyDocumentRuleDeniedParameterOutputWithContext(context.Background())
}

func (i GetPolicyDocumentRuleDeniedParameterArgs) ToGetPolicyDocumentRuleDeniedParameterOutputWithContext(ctx context.Context) GetPolicyDocumentRuleDeniedParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentRuleDeniedParameterOutput)
}

// GetPolicyDocumentRuleDeniedParameterArrayInput is an input type that accepts GetPolicyDocumentRuleDeniedParameterArray and GetPolicyDocumentRuleDeniedParameterArrayOutput values.
// You can construct a concrete instance of `GetPolicyDocumentRuleDeniedParameterArrayInput` via:
//
// 		 GetPolicyDocumentRuleDeniedParameterArray{ GetPolicyDocumentRuleDeniedParameterArgs{...} }
//
type GetPolicyDocumentRuleDeniedParameterArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentRuleDeniedParameterArrayOutput() GetPolicyDocumentRuleDeniedParameterArrayOutput
	ToGetPolicyDocumentRuleDeniedParameterArrayOutputWithContext(context.Context) GetPolicyDocumentRuleDeniedParameterArrayOutput
}

type GetPolicyDocumentRuleDeniedParameterArray []GetPolicyDocumentRuleDeniedParameterInput

func (GetPolicyDocumentRuleDeniedParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentRuleDeniedParameter)(nil)).Elem()
}

func (i GetPolicyDocumentRuleDeniedParameterArray) ToGetPolicyDocumentRuleDeniedParameterArrayOutput() GetPolicyDocumentRuleDeniedParameterArrayOutput {
	return i.ToGetPolicyDocumentRuleDeniedParameterArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentRuleDeniedParameterArray) ToGetPolicyDocumentRuleDeniedParameterArrayOutputWithContext(ctx context.Context) GetPolicyDocumentRuleDeniedParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentRuleDeniedParameterArrayOutput)
}

type GetPolicyDocumentRuleDeniedParameterOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentRuleDeniedParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentRuleDeniedParameter)(nil)).Elem()
}

func (o GetPolicyDocumentRuleDeniedParameterOutput) ToGetPolicyDocumentRuleDeniedParameterOutput() GetPolicyDocumentRuleDeniedParameterOutput {
	return o
}

func (o GetPolicyDocumentRuleDeniedParameterOutput) ToGetPolicyDocumentRuleDeniedParameterOutputWithContext(ctx context.Context) GetPolicyDocumentRuleDeniedParameterOutput {
	return o
}

// name of permitted or denied parameter.
func (o GetPolicyDocumentRuleDeniedParameterOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyDocumentRuleDeniedParameter) string { return v.Key }).(pulumi.StringOutput)
}

// list of values what are permitted or denied by policy rule.
func (o GetPolicyDocumentRuleDeniedParameterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyDocumentRuleDeniedParameter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GetPolicyDocumentRuleDeniedParameterArrayOutput struct{ *pulumi.OutputState }

func (GetPolicyDocumentRuleDeniedParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentRuleDeniedParameter)(nil)).Elem()
}

func (o GetPolicyDocumentRuleDeniedParameterArrayOutput) ToGetPolicyDocumentRuleDeniedParameterArrayOutput() GetPolicyDocumentRuleDeniedParameterArrayOutput {
	return o
}

func (o GetPolicyDocumentRuleDeniedParameterArrayOutput) ToGetPolicyDocumentRuleDeniedParameterArrayOutputWithContext(ctx context.Context) GetPolicyDocumentRuleDeniedParameterArrayOutput {
	return o
}

func (o GetPolicyDocumentRuleDeniedParameterArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentRuleDeniedParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetPolicyDocumentRuleDeniedParameter {
		return vs[0].([]GetPolicyDocumentRuleDeniedParameter)[vs[1].(int)]
	}).(GetPolicyDocumentRuleDeniedParameterOutput)
}

type ProviderAuthLogin struct {
	Namespace  *string           `pulumi:"namespace"`
	Parameters map[string]string `pulumi:"parameters"`
	Path       string            `pulumi:"path"`
}

// ProviderAuthLoginInput is an input type that accepts ProviderAuthLoginArgs and ProviderAuthLoginOutput values.
// You can construct a concrete instance of `ProviderAuthLoginInput` via:
//
// 		 ProviderAuthLoginArgs{...}
//
type ProviderAuthLoginInput interface {
	pulumi.Input

	ToProviderAuthLoginOutput() ProviderAuthLoginOutput
	ToProviderAuthLoginOutputWithContext(context.Context) ProviderAuthLoginOutput
}

type ProviderAuthLoginArgs struct {
	Namespace  pulumi.StringPtrInput `pulumi:"namespace"`
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
	Path       pulumi.StringInput    `pulumi:"path"`
}

func (ProviderAuthLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAuthLogin)(nil)).Elem()
}

func (i ProviderAuthLoginArgs) ToProviderAuthLoginOutput() ProviderAuthLoginOutput {
	return i.ToProviderAuthLoginOutputWithContext(context.Background())
}

func (i ProviderAuthLoginArgs) ToProviderAuthLoginOutputWithContext(ctx context.Context) ProviderAuthLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAuthLoginOutput)
}

// ProviderAuthLoginArrayInput is an input type that accepts ProviderAuthLoginArray and ProviderAuthLoginArrayOutput values.
// You can construct a concrete instance of `ProviderAuthLoginArrayInput` via:
//
// 		 ProviderAuthLoginArray{ ProviderAuthLoginArgs{...} }
//
type ProviderAuthLoginArrayInput interface {
	pulumi.Input

	ToProviderAuthLoginArrayOutput() ProviderAuthLoginArrayOutput
	ToProviderAuthLoginArrayOutputWithContext(context.Context) ProviderAuthLoginArrayOutput
}

type ProviderAuthLoginArray []ProviderAuthLoginInput

func (ProviderAuthLoginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderAuthLogin)(nil)).Elem()
}

func (i ProviderAuthLoginArray) ToProviderAuthLoginArrayOutput() ProviderAuthLoginArrayOutput {
	return i.ToProviderAuthLoginArrayOutputWithContext(context.Background())
}

func (i ProviderAuthLoginArray) ToProviderAuthLoginArrayOutputWithContext(ctx context.Context) ProviderAuthLoginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAuthLoginArrayOutput)
}

type ProviderAuthLoginOutput struct{ *pulumi.OutputState }

func (ProviderAuthLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAuthLogin)(nil)).Elem()
}

func (o ProviderAuthLoginOutput) ToProviderAuthLoginOutput() ProviderAuthLoginOutput {
	return o
}

func (o ProviderAuthLoginOutput) ToProviderAuthLoginOutputWithContext(ctx context.Context) ProviderAuthLoginOutput {
	return o
}

func (o ProviderAuthLoginOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAuthLogin) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ProviderAuthLoginOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProviderAuthLogin) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o ProviderAuthLoginOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderAuthLogin) string { return v.Path }).(pulumi.StringOutput)
}

type ProviderAuthLoginArrayOutput struct{ *pulumi.OutputState }

func (ProviderAuthLoginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderAuthLogin)(nil)).Elem()
}

func (o ProviderAuthLoginArrayOutput) ToProviderAuthLoginArrayOutput() ProviderAuthLoginArrayOutput {
	return o
}

func (o ProviderAuthLoginArrayOutput) ToProviderAuthLoginArrayOutputWithContext(ctx context.Context) ProviderAuthLoginArrayOutput {
	return o
}

func (o ProviderAuthLoginArrayOutput) Index(i pulumi.IntInput) ProviderAuthLoginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderAuthLogin {
		return vs[0].([]ProviderAuthLogin)[vs[1].(int)]
	}).(ProviderAuthLoginOutput)
}

type ProviderClientAuth struct {
	CertFile string `pulumi:"certFile"`
	KeyFile  string `pulumi:"keyFile"`
}

// ProviderClientAuthInput is an input type that accepts ProviderClientAuthArgs and ProviderClientAuthOutput values.
// You can construct a concrete instance of `ProviderClientAuthInput` via:
//
// 		 ProviderClientAuthArgs{...}
//
type ProviderClientAuthInput interface {
	pulumi.Input

	ToProviderClientAuthOutput() ProviderClientAuthOutput
	ToProviderClientAuthOutputWithContext(context.Context) ProviderClientAuthOutput
}

type ProviderClientAuthArgs struct {
	CertFile pulumi.StringInput `pulumi:"certFile"`
	KeyFile  pulumi.StringInput `pulumi:"keyFile"`
}

func (ProviderClientAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderClientAuth)(nil)).Elem()
}

func (i ProviderClientAuthArgs) ToProviderClientAuthOutput() ProviderClientAuthOutput {
	return i.ToProviderClientAuthOutputWithContext(context.Background())
}

func (i ProviderClientAuthArgs) ToProviderClientAuthOutputWithContext(ctx context.Context) ProviderClientAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderClientAuthOutput)
}

// ProviderClientAuthArrayInput is an input type that accepts ProviderClientAuthArray and ProviderClientAuthArrayOutput values.
// You can construct a concrete instance of `ProviderClientAuthArrayInput` via:
//
// 		 ProviderClientAuthArray{ ProviderClientAuthArgs{...} }
//
type ProviderClientAuthArrayInput interface {
	pulumi.Input

	ToProviderClientAuthArrayOutput() ProviderClientAuthArrayOutput
	ToProviderClientAuthArrayOutputWithContext(context.Context) ProviderClientAuthArrayOutput
}

type ProviderClientAuthArray []ProviderClientAuthInput

func (ProviderClientAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderClientAuth)(nil)).Elem()
}

func (i ProviderClientAuthArray) ToProviderClientAuthArrayOutput() ProviderClientAuthArrayOutput {
	return i.ToProviderClientAuthArrayOutputWithContext(context.Background())
}

func (i ProviderClientAuthArray) ToProviderClientAuthArrayOutputWithContext(ctx context.Context) ProviderClientAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderClientAuthArrayOutput)
}

type ProviderClientAuthOutput struct{ *pulumi.OutputState }

func (ProviderClientAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderClientAuth)(nil)).Elem()
}

func (o ProviderClientAuthOutput) ToProviderClientAuthOutput() ProviderClientAuthOutput {
	return o
}

func (o ProviderClientAuthOutput) ToProviderClientAuthOutputWithContext(ctx context.Context) ProviderClientAuthOutput {
	return o
}

func (o ProviderClientAuthOutput) CertFile() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderClientAuth) string { return v.CertFile }).(pulumi.StringOutput)
}

func (o ProviderClientAuthOutput) KeyFile() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderClientAuth) string { return v.KeyFile }).(pulumi.StringOutput)
}

type ProviderClientAuthArrayOutput struct{ *pulumi.OutputState }

func (ProviderClientAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderClientAuth)(nil)).Elem()
}

func (o ProviderClientAuthArrayOutput) ToProviderClientAuthArrayOutput() ProviderClientAuthArrayOutput {
	return o
}

func (o ProviderClientAuthArrayOutput) ToProviderClientAuthArrayOutputWithContext(ctx context.Context) ProviderClientAuthArrayOutput {
	return o
}

func (o ProviderClientAuthArrayOutput) Index(i pulumi.IntInput) ProviderClientAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderClientAuth {
		return vs[0].([]ProviderClientAuth)[vs[1].(int)]
	}).(ProviderClientAuthOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthBackendTuneOutput{})
	pulumi.RegisterOutputType(AuthBackendTunePtrOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentRuleOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentRuleArrayOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentRuleAllowedParameterOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentRuleAllowedParameterArrayOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentRuleDeniedParameterOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentRuleDeniedParameterArrayOutput{})
	pulumi.RegisterOutputType(ProviderAuthLoginOutput{})
	pulumi.RegisterOutputType(ProviderAuthLoginArrayOutput{})
	pulumi.RegisterOutputType(ProviderClientAuthOutput{})
	pulumi.RegisterOutputType(ProviderClientAuthArrayOutput{})
}
