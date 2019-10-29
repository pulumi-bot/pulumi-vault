// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class GroupAlias extends pulumi.CustomResource {
    /**
     * Get an existing GroupAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupAliasState, opts?: pulumi.CustomResourceOptions): GroupAlias {
        return new GroupAlias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/groupAlias:GroupAlias';

    /**
     * Returns true if the given object is an instance of GroupAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupAlias.__pulumiType;
    }

    /**
     * ID of the group to which this is an alias.
     */
    public readonly canonicalId!: pulumi.Output<string>;
    /**
     * Mount accessor to which this alias belongs to.
     */
    public readonly mountAccessor!: pulumi.Output<string>;
    /**
     * Name of the group alias.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a GroupAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupAliasArgs | GroupAliasState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupAliasState | undefined;
            inputs["canonicalId"] = state ? state.canonicalId : undefined;
            inputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as GroupAliasArgs | undefined;
            if (!args || args.canonicalId === undefined) {
                throw new Error("Missing required property 'canonicalId'");
            }
            if (!args || args.mountAccessor === undefined) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            inputs["canonicalId"] = args ? args.canonicalId : undefined;
            inputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupAlias.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupAlias resources.
 */
export interface GroupAliasState {
    /**
     * ID of the group to which this is an alias.
     */
    readonly canonicalId?: pulumi.Input<string>;
    /**
     * Mount accessor to which this alias belongs to.
     */
    readonly mountAccessor?: pulumi.Input<string>;
    /**
     * Name of the group alias.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupAlias resource.
 */
export interface GroupAliasArgs {
    /**
     * ID of the group to which this is an alias.
     */
    readonly canonicalId: pulumi.Input<string>;
    /**
     * Mount accessor to which this alias belongs to.
     */
    readonly mountAccessor: pulumi.Input<string>;
    /**
     * Name of the group alias.
     */
    readonly name?: pulumi.Input<string>;
}
