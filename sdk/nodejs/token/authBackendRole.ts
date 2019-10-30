// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
    public static readonly __pulumiType = 'vault:token/authBackendRole:AuthBackendRole';

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
     * List of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     */
    public readonly boundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * List of disallowed policies for given role.
     */
    public readonly disallowedPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * Number of seconds after which issued tokens can no longer be renewed.
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
     * Number of seconds to set the TTL to for issued tokens upon renewal. Makes the token a periodic token, which will
     * never expire as long as it is renewed before the TTL each period.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     */
    public readonly renewable!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime of the generated token
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Period
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token to generate, service or batch
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
     * List of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of disallowed policies for given role.
     */
    readonly disallowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of seconds after which issued tokens can no longer be renewed.
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
     * Number of seconds to set the TTL to for issued tokens upon renewal. Makes the token a periodic token, which will
     * never expire as long as it is renewed before the TTL each period.
     */
    readonly period?: pulumi.Input<string>;
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * Name of the role.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
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
     * List of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of disallowed policies for given role.
     */
    readonly disallowedPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of seconds after which issued tokens can no longer be renewed.
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
     * Number of seconds to set the TTL to for issued tokens upon renewal. Makes the token a periodic token, which will
     * never expire as long as it is renewed before the TTL each period.
     */
    readonly period?: pulumi.Input<string>;
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * Name of the role.
     */
    readonly roleName: pulumi.Input<string>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     */
    readonly tokenType?: pulumi.Input<string>;
}
