// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
     * Whether the entity is disabled. Disabled entities' associated tokens cannot be used, but are not revoked.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Metadata to be associated with the entity.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the entity.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Policies to be tied to the entity.
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
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as EntityArgs | undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
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
     * Whether the entity is disabled. Disabled entities' associated tokens cannot be used, but are not revoked.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Metadata to be associated with the entity.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the entity.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Policies to be tied to the entity.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Entity resource.
 */
export interface EntityArgs {
    /**
     * Whether the entity is disabled. Disabled entities' associated tokens cannot be used, but are not revoked.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Metadata to be associated with the entity.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the entity.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Policies to be tied to the entity.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}
