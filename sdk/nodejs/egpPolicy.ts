// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
 * 
 * **Note** this feature is available only with Vault Enterprise.
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const allowAll = new vault.EgpPolicy("allow-all", {
 *     enforcementLevel: "soft-mandatory",
 *     paths: ["*"],
 *     policy: `main = rule {
 *   true
 * }
 * `,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/egp_policy.html.md.
 */
export class EgpPolicy extends pulumi.CustomResource {
    /**
     * Get an existing EgpPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EgpPolicyState, opts?: pulumi.CustomResourceOptions): EgpPolicy {
        return new EgpPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/egpPolicy:EgpPolicy';

    /**
     * Returns true if the given object is an instance of EgpPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EgpPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EgpPolicy.__pulumiType;
    }

    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     */
    public readonly enforcementLevel!: pulumi.Output<string>;
    /**
     * The name of the policy
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of paths to which the policy will be applied to
     */
    public readonly paths!: pulumi.Output<string[]>;
    /**
     * String containing a Sentinel policy
     */
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a EgpPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EgpPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EgpPolicyArgs | EgpPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EgpPolicyState | undefined;
            inputs["enforcementLevel"] = state ? state.enforcementLevel : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["paths"] = state ? state.paths : undefined;
            inputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as EgpPolicyArgs | undefined;
            if (!args || args.enforcementLevel === undefined) {
                throw new Error("Missing required property 'enforcementLevel'");
            }
            if (!args || args.paths === undefined) {
                throw new Error("Missing required property 'paths'");
            }
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            inputs["enforcementLevel"] = args ? args.enforcementLevel : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["paths"] = args ? args.paths : undefined;
            inputs["policy"] = args ? args.policy : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EgpPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EgpPolicy resources.
 */
export interface EgpPolicyState {
    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     */
    readonly enforcementLevel?: pulumi.Input<string>;
    /**
     * The name of the policy
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of paths to which the policy will be applied to
     */
    readonly paths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String containing a Sentinel policy
     */
    readonly policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EgpPolicy resource.
 */
export interface EgpPolicyArgs {
    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     */
    readonly enforcementLevel: pulumi.Input<string>;
    /**
     * The name of the policy
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of paths to which the policy will be applied to
     */
    readonly paths: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String containing a Sentinel policy
     */
    readonly policy: pulumi.Input<string>;
}
