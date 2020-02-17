// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/jwt_auth_backend.html.md.
 */
export class AuthBackend extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendState, opts?: pulumi.CustomResourceOptions): AuthBackend {
        return new AuthBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:jwt/authBackend:AuthBackend';

    /**
     * Returns true if the given object is an instance of AuthBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackend.__pulumiType;
    }

    /**
     * The accessor of the JWT auth backend
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The value against which to match the iss claim in a JWT
     */
    public readonly boundIssuer!: pulumi.Output<string | undefined>;
    /**
     * The default role to use if none is provided during login
     */
    public readonly defaultRole!: pulumi.Output<string | undefined>;
    /**
     * The description of the auth backend
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     */
    public readonly jwksCaPem!: pulumi.Output<string | undefined>;
    /**
     * JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
     */
    public readonly jwksUrl!: pulumi.Output<string | undefined>;
    /**
     * A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     */
    public readonly jwtSupportedAlgs!: pulumi.Output<string[] | undefined>;
    /**
     * A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
     */
    public readonly jwtValidationPubkeys!: pulumi.Output<string[] | undefined>;
    /**
     * Client ID used for OIDC backends
     */
    public readonly oidcClientId!: pulumi.Output<string | undefined>;
    /**
     * Client Secret used for OIDC backends
     */
    public readonly oidcClientSecret!: pulumi.Output<string | undefined>;
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     */
    public readonly oidcDiscoveryCaPem!: pulumi.Output<string | undefined>;
    /**
     * The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
     */
    public readonly oidcDiscoveryUrl!: pulumi.Output<string | undefined>;
    /**
     * Path to mount the JWT/OIDC auth backend
     */
    public readonly path!: pulumi.Output<string | undefined>;
    public readonly tune!: pulumi.Output<outputs.jwt.AuthBackendTune>;
    /**
     * Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendArgs | AuthBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["boundIssuer"] = state ? state.boundIssuer : undefined;
            inputs["defaultRole"] = state ? state.defaultRole : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["jwksCaPem"] = state ? state.jwksCaPem : undefined;
            inputs["jwksUrl"] = state ? state.jwksUrl : undefined;
            inputs["jwtSupportedAlgs"] = state ? state.jwtSupportedAlgs : undefined;
            inputs["jwtValidationPubkeys"] = state ? state.jwtValidationPubkeys : undefined;
            inputs["oidcClientId"] = state ? state.oidcClientId : undefined;
            inputs["oidcClientSecret"] = state ? state.oidcClientSecret : undefined;
            inputs["oidcDiscoveryCaPem"] = state ? state.oidcDiscoveryCaPem : undefined;
            inputs["oidcDiscoveryUrl"] = state ? state.oidcDiscoveryUrl : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["tune"] = state ? state.tune : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuthBackendArgs | undefined;
            inputs["boundIssuer"] = args ? args.boundIssuer : undefined;
            inputs["defaultRole"] = args ? args.defaultRole : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["jwksCaPem"] = args ? args.jwksCaPem : undefined;
            inputs["jwksUrl"] = args ? args.jwksUrl : undefined;
            inputs["jwtSupportedAlgs"] = args ? args.jwtSupportedAlgs : undefined;
            inputs["jwtValidationPubkeys"] = args ? args.jwtValidationPubkeys : undefined;
            inputs["oidcClientId"] = args ? args.oidcClientId : undefined;
            inputs["oidcClientSecret"] = args ? args.oidcClientSecret : undefined;
            inputs["oidcDiscoveryCaPem"] = args ? args.oidcDiscoveryCaPem : undefined;
            inputs["oidcDiscoveryUrl"] = args ? args.oidcDiscoveryUrl : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["tune"] = args ? args.tune : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["accessor"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackend resources.
 */
export interface AuthBackendState {
    /**
     * The accessor of the JWT auth backend
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * The value against which to match the iss claim in a JWT
     */
    readonly boundIssuer?: pulumi.Input<string>;
    /**
     * The default role to use if none is provided during login
     */
    readonly defaultRole?: pulumi.Input<string>;
    /**
     * The description of the auth backend
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     */
    readonly jwksCaPem?: pulumi.Input<string>;
    /**
     * JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
     */
    readonly jwksUrl?: pulumi.Input<string>;
    /**
     * A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     */
    readonly jwtSupportedAlgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
     */
    readonly jwtValidationPubkeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Client ID used for OIDC backends
     */
    readonly oidcClientId?: pulumi.Input<string>;
    /**
     * Client Secret used for OIDC backends
     */
    readonly oidcClientSecret?: pulumi.Input<string>;
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     */
    readonly oidcDiscoveryCaPem?: pulumi.Input<string>;
    /**
     * The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
     */
    readonly oidcDiscoveryUrl?: pulumi.Input<string>;
    /**
     * Path to mount the JWT/OIDC auth backend
     */
    readonly path?: pulumi.Input<string>;
    readonly tune?: pulumi.Input<inputs.jwt.AuthBackendTune>;
    /**
     * Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackend resource.
 */
export interface AuthBackendArgs {
    /**
     * The value against which to match the iss claim in a JWT
     */
    readonly boundIssuer?: pulumi.Input<string>;
    /**
     * The default role to use if none is provided during login
     */
    readonly defaultRole?: pulumi.Input<string>;
    /**
     * The description of the auth backend
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     */
    readonly jwksCaPem?: pulumi.Input<string>;
    /**
     * JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
     */
    readonly jwksUrl?: pulumi.Input<string>;
    /**
     * A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     */
    readonly jwtSupportedAlgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
     */
    readonly jwtValidationPubkeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Client ID used for OIDC backends
     */
    readonly oidcClientId?: pulumi.Input<string>;
    /**
     * Client Secret used for OIDC backends
     */
    readonly oidcClientSecret?: pulumi.Input<string>;
    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     */
    readonly oidcDiscoveryCaPem?: pulumi.Input<string>;
    /**
     * The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
     */
    readonly oidcDiscoveryUrl?: pulumi.Input<string>;
    /**
     * Path to mount the JWT/OIDC auth backend
     */
    readonly path?: pulumi.Input<string>;
    readonly tune?: pulumi.Input<inputs.jwt.AuthBackendTune>;
    /**
     * Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     */
    readonly type?: pulumi.Input<string>;
}
