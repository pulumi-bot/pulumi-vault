// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity.html.md.
 */
export class Entity extends pulumi.CustomResource {
    /**
     * Get an existing Entity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntityState, opts?: pulumi.CustomResourceOptions): Entity {
        return new Entity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/entity:Entity';

    /**
     * Returns true if the given object is an instance of Entity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Entity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Entity.__pulumiType;
    }

    /**
     * True/false Is this entity currently disabled. Defaults to `false`
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
     */
    public readonly externalPolicies!: pulumi.Output<boolean | undefined>;
    /**
     * A Map of additional metadata to associate with the user.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the identity entity to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of policies to apply to the entity.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Entity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EntityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntityArgs | EntityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EntityState | undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["externalPolicies"] = state ? state.externalPolicies : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as EntityArgs | undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["externalPolicies"] = args ? args.externalPolicies : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policies"] = args ? args.policies : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Entity.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Entity resources.
 */
export interface EntityState {
    /**
     * True/false Is this entity currently disabled. Defaults to `false`
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
     */
    readonly externalPolicies?: pulumi.Input<boolean>;
    /**
     * A Map of additional metadata to associate with the user.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the identity entity to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of policies to apply to the entity.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Entity resource.
 */
export interface EntityArgs {
    /**
     * True/false Is this entity currently disabled. Defaults to `false`
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
     */
    readonly externalPolicies?: pulumi.Input<boolean>;
    /**
     * A Map of additional metadata to associate with the user.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the identity entity to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of policies to apply to the entity.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}
