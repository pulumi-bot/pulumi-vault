// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages Token auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/token.html) for more
 * information.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const example = new vault.tokenauth.AuthBackendRole("example", {
 *     allowedPolicies: [
 *         "dev",
 *         "test",
 *     ],
 *     disallowedPolicies: ["default"],
 *     explicitMaxTtl: "115200",
 *     orphan: true,
 *     pathSuffix: "path-suffix",
 *     period: "86400",
 *     renewable: true,
 *     roleName: "my-role",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/token_auth_backend_role.html.md.
 */
export class AuthBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleState, opts?: pulumi.CustomResourceOptions): AuthBackendRole {
        return new AuthBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:tokenauth/authBackendRole:AuthBackendRole';

    /**
     * Returns true if the given object is an instance of AuthBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRole.__pulumiType;
    }

    /**
     * List of allowed policies for given role.
     */
    public readonly allowedPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * If set, a list of
     * CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     */
    public readonly boundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * List of disallowed policies for given role.
     */
    public readonly disallowedPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * If set, the
     * token will have an explicit max TTL set upon it.
     */
    public readonly explicitMaxTtl!: pulumi.Output<string | undefined>;
    /**
     * If true, tokens created against this policy will be orphan tokens.
     */
    public readonly orphan!: pulumi.Output<boolean | undefined>;
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     */
    public readonly pathSuffix!: pulumi.Output<string | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * Wether to disable the ability of the token to be renewed past its initial TTL.
     */
    public readonly renewable!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleArgs | AuthBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            inputs["allowedPolicies"] = state ? state.allowedPolicies : undefined;
            inputs["boundCidrs"] = state ? state.boundCidrs : undefined;
            inputs["disallowedPolicies"] = state ? state.disallowedPolicies : undefined;
            inputs["explicitMaxTtl"] = state ? state.explicitMaxTtl : undefined;
            inputs["orphan"] = state ? state.orphan : undefined;
            inputs["pathSuffix"] = state ? state.pathSuffix : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["renewable"] = state ? state.renewable : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if (!args || args.roleName === undefined) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["allowedPolicies"] = args ? args.allowedPolicies : undefined;
            inputs["boundCidrs"] = args ? args.boundCidrs : undefined;
            inputs["disallowedPolicies"] = args ? args.disallowedPolicies : undefined;
            inputs["explicitMaxTtl"] = args ? args.explicitMaxTtl : undefined;
            inputs["orphan"] = args ? args.orphan : undefined;
            inputs["pathSuffix"] = args ? args.pathSuffix : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["renewable"] = args ? args.renewable : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRole resources.
 */
export interface AuthBackendRoleState {
    /**
     * List of allowed policies for given role.
     */
    readonly allowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a list of
     * CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     * 
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of disallowed policies for given role.
     */
    readonly disallowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, the
     * token will have an explicit max TTL set upon it.
     * 
     * @deprecated use `token_explicit_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly explicitMaxTtl?: pulumi.Input<string>;
    /**
     * If true, tokens created against this policy will be orphan tokens.
     */
    readonly orphan?: pulumi.Input<boolean>;
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     */
    readonly pathSuffix?: pulumi.Input<string>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<string>;
    /**
     * Wether to disable the ability of the token to be renewed past its initial TTL.
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * The name of the role.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * List of allowed policies for given role.
     */
    readonly allowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, a list of
     * CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     * 
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of disallowed policies for given role.
     */
    readonly disallowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, the
     * token will have an explicit max TTL set upon it.
     * 
     * @deprecated use `token_explicit_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly explicitMaxTtl?: pulumi.Input<string>;
    /**
     * If true, tokens created against this policy will be orphan tokens.
     */
    readonly orphan?: pulumi.Input<boolean>;
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     */
    readonly pathSuffix?: pulumi.Input<string>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<string>;
    /**
     * Wether to disable the ability of the token to be renewed past its initial TTL.
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * The name of the role.
     */
    readonly roleName: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
}
