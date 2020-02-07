// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/azure_secret_backend.html.markdown.
 */
export class Backend extends pulumi.CustomResource {
    /**
     * Get an existing Backend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendState, opts?: pulumi.CustomResourceOptions): Backend {
        return new Backend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:azure/backend:Backend';

    /**
     * Returns true if the given object is an instance of Backend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backend.__pulumiType;
    }

    /**
     * The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
     * required.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client secret for credentials to query the Azure APIs
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * Path to mount the backend at.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The subscription id for the Azure Active Directory.
     */
    public readonly subscriptionId!: pulumi.Output<string>;
    /**
     * The tenant id for the Azure Active Directory organization.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a Backend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendArgs | BackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BackendState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["subscriptionId"] = state ? state.subscriptionId : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as BackendArgs | undefined;
            if (!args || args.subscriptionId === undefined) {
                throw new Error("Missing required property 'subscriptionId'");
            }
            if (!args || args.tenantId === undefined) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["subscriptionId"] = args ? args.subscriptionId : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Backend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Backend resources.
 */
export interface BackendState {
    /**
     * The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
     * required.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client secret for credentials to query the Azure APIs
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * Path to mount the backend at.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The subscription id for the Azure Active Directory.
     */
    readonly subscriptionId?: pulumi.Input<string>;
    /**
     * The tenant id for the Azure Active Directory organization.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Backend resource.
 */
export interface BackendArgs {
    /**
     * The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are
     * required.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client secret for credentials to query the Azure APIs
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * Path to mount the backend at.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The subscription id for the Azure Active Directory.
     */
    readonly subscriptionId: pulumi.Input<string>;
    /**
     * The tenant id for the Azure Active Directory organization.
     */
    readonly tenantId: pulumi.Input<string>;
}
