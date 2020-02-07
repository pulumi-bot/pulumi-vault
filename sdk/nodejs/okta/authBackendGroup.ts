// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a group in an
 * [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const example = new vault.okta.AuthBackend("example", {
 *     organization: "dummy",
 *     path: "groupOkta",
 * });
 * const foo = new vault.okta.AuthBackendGroup("foo", {
 *     groupName: "foo",
 *     path: example.path,
 *     policies: [
 *         "one",
 *         "two",
 *     ],
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/okta_auth_backend_group.html.markdown.
 */
export class AuthBackendGroup extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendGroupState, opts?: pulumi.CustomResourceOptions): AuthBackendGroup {
        return new AuthBackendGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:okta/authBackendGroup:AuthBackendGroup';

    /**
     * Returns true if the given object is an instance of AuthBackendGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendGroup.__pulumiType;
    }

    /**
     * Name of the group within the Okta
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * The path where the Okta auth backend is mounted
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Vault policies to associate with this group
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AuthBackendGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendGroupArgs | AuthBackendGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendGroupState | undefined;
            inputs["groupName"] = state ? state.groupName : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as AuthBackendGroupArgs | undefined;
            if (!args || args.groupName === undefined) {
                throw new Error("Missing required property 'groupName'");
            }
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["policies"] = args ? args.policies : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendGroup resources.
 */
export interface AuthBackendGroupState {
    /**
     * Name of the group within the Okta
     */
    readonly groupName?: pulumi.Input<string>;
    /**
     * The path where the Okta auth backend is mounted
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Vault policies to associate with this group
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AuthBackendGroup resource.
 */
export interface AuthBackendGroupArgs {
    /**
     * Name of the group within the Okta
     */
    readonly groupName: pulumi.Input<string>;
    /**
     * The path where the Okta auth backend is mounted
     */
    readonly path: pulumi.Input<string>;
    /**
     * Vault policies to associate with this group
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}
