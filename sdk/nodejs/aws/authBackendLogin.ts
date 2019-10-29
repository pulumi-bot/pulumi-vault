// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AuthBackendLogin extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendLogin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendLoginState, opts?: pulumi.CustomResourceOptions): AuthBackendLogin {
        return new AuthBackendLogin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendLogin:AuthBackendLogin';

    /**
     * Returns true if the given object is an instance of AuthBackendLogin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendLogin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendLogin.__pulumiType;
    }

    /**
     * The accessor returned from Vault for this token.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * The auth method used to generate this token.
     */
    public /*out*/ readonly authType!: pulumi.Output<string>;
    /**
     * AWS Auth Backend to read the token from.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The token returned by Vault.
     */
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    /**
     * The HTTP method used in the signed request.
     */
    public readonly iamHttpRequestMethod!: pulumi.Output<string | undefined>;
    /**
     * The Base64-encoded body of the signed request.
     */
    public readonly iamRequestBody!: pulumi.Output<string | undefined>;
    /**
     * The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
     */
    public readonly iamRequestHeaders!: pulumi.Output<string | undefined>;
    /**
     * The Base64-encoded HTTP URL used in the signed request.
     */
    public readonly iamRequestUrl!: pulumi.Output<string | undefined>;
    /**
     * Base64-encoded EC2 instance identity document to authenticate with.
     */
    public readonly identity!: pulumi.Output<string | undefined>;
    /**
     * Lease duration in seconds relative to the time in lease_start_time.
     */
    public /*out*/ readonly leaseDuration!: pulumi.Output<number>;
    /**
     * Time at which the lease was read, using the clock of the system where Terraform was running
     */
    public /*out*/ readonly leaseStartTime!: pulumi.Output<string>;
    /**
     * The metadata reported by the Vault server.
     */
    public /*out*/ readonly metadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * The nonce to be used for subsequent login requests.
     */
    public readonly nonce!: pulumi.Output<string>;
    /**
     * PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
     */
    public readonly pkcs7!: pulumi.Output<string | undefined>;
    /**
     * The policies assigned to this token.
     */
    public /*out*/ readonly policies!: pulumi.Output<string[]>;
    /**
     * True if the duration of this lease can be extended through renewal.
     */
    public /*out*/ readonly renewable!: pulumi.Output<boolean>;
    /**
     * AWS Auth Role to read the token from.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
     */
    public readonly signature!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendLogin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthBackendLoginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendLoginArgs | AuthBackendLoginState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendLoginState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["authType"] = state ? state.authType : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["clientToken"] = state ? state.clientToken : undefined;
            inputs["iamHttpRequestMethod"] = state ? state.iamHttpRequestMethod : undefined;
            inputs["iamRequestBody"] = state ? state.iamRequestBody : undefined;
            inputs["iamRequestHeaders"] = state ? state.iamRequestHeaders : undefined;
            inputs["iamRequestUrl"] = state ? state.iamRequestUrl : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["leaseDuration"] = state ? state.leaseDuration : undefined;
            inputs["leaseStartTime"] = state ? state.leaseStartTime : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["nonce"] = state ? state.nonce : undefined;
            inputs["pkcs7"] = state ? state.pkcs7 : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["renewable"] = state ? state.renewable : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["signature"] = state ? state.signature : undefined;
        } else {
            const args = argsOrState as AuthBackendLoginArgs | undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["iamHttpRequestMethod"] = args ? args.iamHttpRequestMethod : undefined;
            inputs["iamRequestBody"] = args ? args.iamRequestBody : undefined;
            inputs["iamRequestHeaders"] = args ? args.iamRequestHeaders : undefined;
            inputs["iamRequestUrl"] = args ? args.iamRequestUrl : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["nonce"] = args ? args.nonce : undefined;
            inputs["pkcs7"] = args ? args.pkcs7 : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["signature"] = args ? args.signature : undefined;
            inputs["accessor"] = undefined /*out*/;
            inputs["authType"] = undefined /*out*/;
            inputs["clientToken"] = undefined /*out*/;
            inputs["leaseDuration"] = undefined /*out*/;
            inputs["leaseStartTime"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["policies"] = undefined /*out*/;
            inputs["renewable"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendLogin.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendLogin resources.
 */
export interface AuthBackendLoginState {
    /**
     * The accessor returned from Vault for this token.
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * The auth method used to generate this token.
     */
    readonly authType?: pulumi.Input<string>;
    /**
     * AWS Auth Backend to read the token from.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The token returned by Vault.
     */
    readonly clientToken?: pulumi.Input<string>;
    /**
     * The HTTP method used in the signed request.
     */
    readonly iamHttpRequestMethod?: pulumi.Input<string>;
    /**
     * The Base64-encoded body of the signed request.
     */
    readonly iamRequestBody?: pulumi.Input<string>;
    /**
     * The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
     */
    readonly iamRequestHeaders?: pulumi.Input<string>;
    /**
     * The Base64-encoded HTTP URL used in the signed request.
     */
    readonly iamRequestUrl?: pulumi.Input<string>;
    /**
     * Base64-encoded EC2 instance identity document to authenticate with.
     */
    readonly identity?: pulumi.Input<string>;
    /**
     * Lease duration in seconds relative to the time in lease_start_time.
     */
    readonly leaseDuration?: pulumi.Input<number>;
    /**
     * Time at which the lease was read, using the clock of the system where Terraform was running
     */
    readonly leaseStartTime?: pulumi.Input<string>;
    /**
     * The metadata reported by the Vault server.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The nonce to be used for subsequent login requests.
     */
    readonly nonce?: pulumi.Input<string>;
    /**
     * PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
     */
    readonly pkcs7?: pulumi.Input<string>;
    /**
     * The policies assigned to this token.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * True if the duration of this lease can be extended through renewal.
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * AWS Auth Role to read the token from.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
     */
    readonly signature?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendLogin resource.
 */
export interface AuthBackendLoginArgs {
    /**
     * AWS Auth Backend to read the token from.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The HTTP method used in the signed request.
     */
    readonly iamHttpRequestMethod?: pulumi.Input<string>;
    /**
     * The Base64-encoded body of the signed request.
     */
    readonly iamRequestBody?: pulumi.Input<string>;
    /**
     * The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
     */
    readonly iamRequestHeaders?: pulumi.Input<string>;
    /**
     * The Base64-encoded HTTP URL used in the signed request.
     */
    readonly iamRequestUrl?: pulumi.Input<string>;
    /**
     * Base64-encoded EC2 instance identity document to authenticate with.
     */
    readonly identity?: pulumi.Input<string>;
    /**
     * The nonce to be used for subsequent login requests.
     */
    readonly nonce?: pulumi.Input<string>;
    /**
     * PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
     */
    readonly pkcs7?: pulumi.Input<string>;
    /**
     * AWS Auth Role to read the token from.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
     */
    readonly signature?: pulumi.Input<string>;
}
