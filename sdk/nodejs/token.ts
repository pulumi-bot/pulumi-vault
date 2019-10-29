// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Token extends pulumi.CustomResource {
    /**
     * Get an existing Token resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TokenState, opts?: pulumi.CustomResourceOptions): Token {
        return new Token(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/token:Token';

    /**
     * Returns true if the given object is an instance of Token.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Token {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Token.__pulumiType;
    }

    /**
     * The client token.
     */
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    /**
     * The display name of the token.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The explicit max TTL of the token.
     */
    public readonly explicitMaxTtl!: pulumi.Output<string | undefined>;
    /**
     * The token lease duration.
     */
    public /*out*/ readonly leaseDuration!: pulumi.Output<number>;
    /**
     * The token lease started on.
     */
    public /*out*/ readonly leaseStarted!: pulumi.Output<string>;
    /**
     * Flag to disable the default policy.
     */
    public readonly noDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to create a token without parent.
     */
    public readonly noParent!: pulumi.Output<boolean>;
    /**
     * The number of allowed uses of the token.
     */
    public readonly numUses!: pulumi.Output<number>;
    /**
     * The period of the token.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * List of policies.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * The renew increment.
     */
    public readonly renewIncrement!: pulumi.Output<number | undefined>;
    /**
     * The minimum lease to renew token.
     */
    public readonly renewMinLease!: pulumi.Output<number | undefined>;
    /**
     * Flag to allow the token to be renewed
     */
    public readonly renewable!: pulumi.Output<boolean>;
    /**
     * The token role name.
     */
    public readonly roleName!: pulumi.Output<string | undefined>;
    /**
     * The TTL period of the token.
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
    /**
     * The client wrapped token.
     */
    public /*out*/ readonly wrappedToken!: pulumi.Output<string>;
    /**
     * The client wrapping accessor.
     */
    public /*out*/ readonly wrappingAccessor!: pulumi.Output<string>;
    /**
     * The TTL period of the wrapped token.
     */
    public readonly wrappingTtl!: pulumi.Output<string | undefined>;

    /**
     * Create a Token resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TokenArgs | TokenState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TokenState | undefined;
            inputs["clientToken"] = state ? state.clientToken : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["explicitMaxTtl"] = state ? state.explicitMaxTtl : undefined;
            inputs["leaseDuration"] = state ? state.leaseDuration : undefined;
            inputs["leaseStarted"] = state ? state.leaseStarted : undefined;
            inputs["noDefaultPolicy"] = state ? state.noDefaultPolicy : undefined;
            inputs["noParent"] = state ? state.noParent : undefined;
            inputs["numUses"] = state ? state.numUses : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["renewIncrement"] = state ? state.renewIncrement : undefined;
            inputs["renewMinLease"] = state ? state.renewMinLease : undefined;
            inputs["renewable"] = state ? state.renewable : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["wrappedToken"] = state ? state.wrappedToken : undefined;
            inputs["wrappingAccessor"] = state ? state.wrappingAccessor : undefined;
            inputs["wrappingTtl"] = state ? state.wrappingTtl : undefined;
        } else {
            const args = argsOrState as TokenArgs | undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["explicitMaxTtl"] = args ? args.explicitMaxTtl : undefined;
            inputs["noDefaultPolicy"] = args ? args.noDefaultPolicy : undefined;
            inputs["noParent"] = args ? args.noParent : undefined;
            inputs["numUses"] = args ? args.numUses : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["renewIncrement"] = args ? args.renewIncrement : undefined;
            inputs["renewMinLease"] = args ? args.renewMinLease : undefined;
            inputs["renewable"] = args ? args.renewable : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["wrappingTtl"] = args ? args.wrappingTtl : undefined;
            inputs["clientToken"] = undefined /*out*/;
            inputs["leaseDuration"] = undefined /*out*/;
            inputs["leaseStarted"] = undefined /*out*/;
            inputs["wrappedToken"] = undefined /*out*/;
            inputs["wrappingAccessor"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Token.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Token resources.
 */
export interface TokenState {
    /**
     * The client token.
     */
    readonly clientToken?: pulumi.Input<string>;
    /**
     * The display name of the token.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The explicit max TTL of the token.
     */
    readonly explicitMaxTtl?: pulumi.Input<string>;
    /**
     * The token lease duration.
     */
    readonly leaseDuration?: pulumi.Input<number>;
    /**
     * The token lease started on.
     */
    readonly leaseStarted?: pulumi.Input<string>;
    /**
     * Flag to disable the default policy.
     */
    readonly noDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * Flag to create a token without parent.
     */
    readonly noParent?: pulumi.Input<boolean>;
    /**
     * The number of allowed uses of the token.
     */
    readonly numUses?: pulumi.Input<number>;
    /**
     * The period of the token.
     */
    readonly period?: pulumi.Input<string>;
    /**
     * List of policies.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The renew increment.
     */
    readonly renewIncrement?: pulumi.Input<number>;
    /**
     * The minimum lease to renew token.
     */
    readonly renewMinLease?: pulumi.Input<number>;
    /**
     * Flag to allow the token to be renewed
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * The token role name.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The TTL period of the token.
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * The client wrapped token.
     */
    readonly wrappedToken?: pulumi.Input<string>;
    /**
     * The client wrapping accessor.
     */
    readonly wrappingAccessor?: pulumi.Input<string>;
    /**
     * The TTL period of the wrapped token.
     */
    readonly wrappingTtl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Token resource.
 */
export interface TokenArgs {
    /**
     * The display name of the token.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The explicit max TTL of the token.
     */
    readonly explicitMaxTtl?: pulumi.Input<string>;
    /**
     * Flag to disable the default policy.
     */
    readonly noDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * Flag to create a token without parent.
     */
    readonly noParent?: pulumi.Input<boolean>;
    /**
     * The number of allowed uses of the token.
     */
    readonly numUses?: pulumi.Input<number>;
    /**
     * The period of the token.
     */
    readonly period?: pulumi.Input<string>;
    /**
     * List of policies.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The renew increment.
     */
    readonly renewIncrement?: pulumi.Input<number>;
    /**
     * The minimum lease to renew token.
     */
    readonly renewMinLease?: pulumi.Input<number>;
    /**
     * Flag to allow the token to be renewed
     */
    readonly renewable?: pulumi.Input<boolean>;
    /**
     * The token role name.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The TTL period of the token.
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * The TTL period of the wrapped token.
     */
    readonly wrappingTtl?: pulumi.Input<string>;
}
