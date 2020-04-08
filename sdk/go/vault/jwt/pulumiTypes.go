// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package jwt

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
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
	return o.ApplyT(func(v AuthBackendTune) []string { return v.AllowedResponseHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the request data object.
func (o AuthBackendTunePtrOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.AuditNonHmacRequestKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the response data object.
func (o AuthBackendTunePtrOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.AuditNonHmacResponseKeys }).(pulumi.StringArrayOutput)
}

// Specifies the default time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTunePtrOutput) DefaultLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.DefaultLeaseTtl }).(pulumi.StringPtrOutput)
}

// Specifies whether to show this mount in
// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
func (o AuthBackendTunePtrOutput) ListingVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.ListingVisibility }).(pulumi.StringPtrOutput)
}

// Specifies the maximum time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTunePtrOutput) MaxLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.MaxLeaseTtl }).(pulumi.StringPtrOutput)
}

// List of headers to whitelist and
// pass from the request to the backend.
func (o AuthBackendTunePtrOutput) PassthroughRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTune) []string { return v.PassthroughRequestHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the type of tokens that should be returned by
// the mount. Valid values are "default-service", "default-batch", "service", "batch".
func (o AuthBackendTunePtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTune) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type AuthBackendTuneArgs struct {
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

type AuthBackendTuneArgsInput interface {
	pulumi.Input

	ToAuthBackendTuneArgsOutput() AuthBackendTuneArgsOutput
	ToAuthBackendTuneArgsOutputWithContext(context.Context) AuthBackendTuneArgsOutput
}

type AuthBackendTuneArgsArgs struct {
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

func (AuthBackendTuneArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendTuneArgs)(nil)).Elem()
}

func (i AuthBackendTuneArgsArgs) ToAuthBackendTuneArgsOutput() AuthBackendTuneArgsOutput {
	return i.ToAuthBackendTuneArgsOutputWithContext(context.Background())
}

func (i AuthBackendTuneArgsArgs) ToAuthBackendTuneArgsOutputWithContext(ctx context.Context) AuthBackendTuneArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTuneArgsOutput)
}

func (i AuthBackendTuneArgsArgs) ToAuthBackendTuneArgsPtrOutput() AuthBackendTuneArgsPtrOutput {
	return i.ToAuthBackendTuneArgsPtrOutputWithContext(context.Background())
}

func (i AuthBackendTuneArgsArgs) ToAuthBackendTuneArgsPtrOutputWithContext(ctx context.Context) AuthBackendTuneArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTuneArgsOutput).ToAuthBackendTuneArgsPtrOutputWithContext(ctx)
}

type AuthBackendTuneArgsPtrInput interface {
	pulumi.Input

	ToAuthBackendTuneArgsPtrOutput() AuthBackendTuneArgsPtrOutput
	ToAuthBackendTuneArgsPtrOutputWithContext(context.Context) AuthBackendTuneArgsPtrOutput
}

type authBackendTuneArgsPtrType AuthBackendTuneArgsArgs

func AuthBackendTuneArgsPtr(v *AuthBackendTuneArgsArgs) AuthBackendTuneArgsPtrInput {
	return (*authBackendTuneArgsPtrType)(v)
}

func (*authBackendTuneArgsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendTuneArgs)(nil)).Elem()
}

func (i *authBackendTuneArgsPtrType) ToAuthBackendTuneArgsPtrOutput() AuthBackendTuneArgsPtrOutput {
	return i.ToAuthBackendTuneArgsPtrOutputWithContext(context.Background())
}

func (i *authBackendTuneArgsPtrType) ToAuthBackendTuneArgsPtrOutputWithContext(ctx context.Context) AuthBackendTuneArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTuneArgsPtrOutput)
}

type AuthBackendTuneArgsOutput struct{ *pulumi.OutputState }

func (AuthBackendTuneArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendTuneArgs)(nil)).Elem()
}

func (o AuthBackendTuneArgsOutput) ToAuthBackendTuneArgsOutput() AuthBackendTuneArgsOutput {
	return o
}

func (o AuthBackendTuneArgsOutput) ToAuthBackendTuneArgsOutputWithContext(ctx context.Context) AuthBackendTuneArgsOutput {
	return o
}

func (o AuthBackendTuneArgsOutput) ToAuthBackendTuneArgsPtrOutput() AuthBackendTuneArgsPtrOutput {
	return o.ToAuthBackendTuneArgsPtrOutputWithContext(context.Background())
}

func (o AuthBackendTuneArgsOutput) ToAuthBackendTuneArgsPtrOutputWithContext(ctx context.Context) AuthBackendTuneArgsPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *AuthBackendTuneArgs {
		return &v
	}).(AuthBackendTuneArgsPtrOutput)
}

// List of headers to whitelist and allowing
// a plugin to include them in the response.
func (o AuthBackendTuneArgsOutput) AllowedResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.AllowedResponseHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the request data object.
func (o AuthBackendTuneArgsOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.AuditNonHmacRequestKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the response data object.
func (o AuthBackendTuneArgsOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.AuditNonHmacResponseKeys }).(pulumi.StringArrayOutput)
}

// Specifies the default time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneArgsOutput) DefaultLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.DefaultLeaseTtl }).(pulumi.StringPtrOutput)
}

// Specifies whether to show this mount in
// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
func (o AuthBackendTuneArgsOutput) ListingVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.ListingVisibility }).(pulumi.StringPtrOutput)
}

// Specifies the maximum time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneArgsOutput) MaxLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.MaxLeaseTtl }).(pulumi.StringPtrOutput)
}

// List of headers to whitelist and
// pass from the request to the backend.
func (o AuthBackendTuneArgsOutput) PassthroughRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.PassthroughRequestHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the type of tokens that should be returned by
// the mount. Valid values are "default-service", "default-batch", "service", "batch".
func (o AuthBackendTuneArgsOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type AuthBackendTuneArgsPtrOutput struct{ *pulumi.OutputState }

func (AuthBackendTuneArgsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendTuneArgs)(nil)).Elem()
}

func (o AuthBackendTuneArgsPtrOutput) ToAuthBackendTuneArgsPtrOutput() AuthBackendTuneArgsPtrOutput {
	return o
}

func (o AuthBackendTuneArgsPtrOutput) ToAuthBackendTuneArgsPtrOutputWithContext(ctx context.Context) AuthBackendTuneArgsPtrOutput {
	return o
}

func (o AuthBackendTuneArgsPtrOutput) Elem() AuthBackendTuneArgsOutput {
	return o.ApplyT(func(v *AuthBackendTuneArgs) AuthBackendTuneArgs { return *v }).(AuthBackendTuneArgsOutput)
}

// List of headers to whitelist and allowing
// a plugin to include them in the response.
func (o AuthBackendTuneArgsPtrOutput) AllowedResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.AllowedResponseHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the request data object.
func (o AuthBackendTuneArgsPtrOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.AuditNonHmacRequestKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the response data object.
func (o AuthBackendTuneArgsPtrOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.AuditNonHmacResponseKeys }).(pulumi.StringArrayOutput)
}

// Specifies the default time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneArgsPtrOutput) DefaultLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.DefaultLeaseTtl }).(pulumi.StringPtrOutput)
}

// Specifies whether to show this mount in
// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
func (o AuthBackendTuneArgsPtrOutput) ListingVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.ListingVisibility }).(pulumi.StringPtrOutput)
}

// Specifies the maximum time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneArgsPtrOutput) MaxLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.MaxLeaseTtl }).(pulumi.StringPtrOutput)
}

// List of headers to whitelist and
// pass from the request to the backend.
func (o AuthBackendTuneArgsPtrOutput) PassthroughRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) []string { return v.PassthroughRequestHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the type of tokens that should be returned by
// the mount. Valid values are "default-service", "default-batch", "service", "batch".
func (o AuthBackendTuneArgsPtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneArgs) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type AuthBackendTuneState struct {
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

type AuthBackendTuneStateInput interface {
	pulumi.Input

	ToAuthBackendTuneStateOutput() AuthBackendTuneStateOutput
	ToAuthBackendTuneStateOutputWithContext(context.Context) AuthBackendTuneStateOutput
}

type AuthBackendTuneStateArgs struct {
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

func (AuthBackendTuneStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendTuneState)(nil)).Elem()
}

func (i AuthBackendTuneStateArgs) ToAuthBackendTuneStateOutput() AuthBackendTuneStateOutput {
	return i.ToAuthBackendTuneStateOutputWithContext(context.Background())
}

func (i AuthBackendTuneStateArgs) ToAuthBackendTuneStateOutputWithContext(ctx context.Context) AuthBackendTuneStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendTuneStateOutput)
}

type AuthBackendTuneStateOutput struct{ *pulumi.OutputState }

func (AuthBackendTuneStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendTuneState)(nil)).Elem()
}

func (o AuthBackendTuneStateOutput) ToAuthBackendTuneStateOutput() AuthBackendTuneStateOutput {
	return o
}

func (o AuthBackendTuneStateOutput) ToAuthBackendTuneStateOutputWithContext(ctx context.Context) AuthBackendTuneStateOutput {
	return o
}

// List of headers to whitelist and allowing
// a plugin to include them in the response.
func (o AuthBackendTuneStateOutput) AllowedResponseHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneState) []string { return v.AllowedResponseHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the request data object.
func (o AuthBackendTuneStateOutput) AuditNonHmacRequestKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneState) []string { return v.AuditNonHmacRequestKeys }).(pulumi.StringArrayOutput)
}

// Specifies the list of keys that will
// not be HMAC'd by audit devices in the response data object.
func (o AuthBackendTuneStateOutput) AuditNonHmacResponseKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneState) []string { return v.AuditNonHmacResponseKeys }).(pulumi.StringArrayOutput)
}

// Specifies the default time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneStateOutput) DefaultLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneState) *string { return v.DefaultLeaseTtl }).(pulumi.StringPtrOutput)
}

// Specifies whether to show this mount in
// the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
func (o AuthBackendTuneStateOutput) ListingVisibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneState) *string { return v.ListingVisibility }).(pulumi.StringPtrOutput)
}

// Specifies the maximum time-to-live.
// If set, this overrides the global default.
// Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
func (o AuthBackendTuneStateOutput) MaxLeaseTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneState) *string { return v.MaxLeaseTtl }).(pulumi.StringPtrOutput)
}

// List of headers to whitelist and
// pass from the request to the backend.
func (o AuthBackendTuneStateOutput) PassthroughRequestHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthBackendTuneState) []string { return v.PassthroughRequestHeaders }).(pulumi.StringArrayOutput)
}

// Specifies the type of tokens that should be returned by
// the mount. Valid values are "default-service", "default-batch", "service", "batch".
func (o AuthBackendTuneStateOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthBackendTuneState) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthBackendTuneOutput{})
	pulumi.RegisterOutputType(AuthBackendTunePtrOutput{})
	pulumi.RegisterOutputType(AuthBackendTuneArgsOutput{})
	pulumi.RegisterOutputType(AuthBackendTuneArgsPtrOutput{})
	pulumi.RegisterOutputType(AuthBackendTuneStateOutput{})
}
