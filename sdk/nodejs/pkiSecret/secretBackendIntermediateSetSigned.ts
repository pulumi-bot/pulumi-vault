// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_intermediate_set_signed.html.md.
 */
export class SecretBackendIntermediateSetSigned extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendIntermediateSetSigned resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendIntermediateSetSignedState, opts?: pulumi.CustomResourceOptions): SecretBackendIntermediateSetSigned {
        return new SecretBackendIntermediateSetSigned(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned';

    /**
     * Returns true if the given object is an instance of SecretBackendIntermediateSetSigned.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendIntermediateSetSigned {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendIntermediateSetSigned.__pulumiType;
    }

    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The certificate
     */
    public readonly certificate!: pulumi.Output<string>;

    /**
     * Create a SecretBackendIntermediateSetSigned resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendIntermediateSetSignedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendIntermediateSetSignedArgs | SecretBackendIntermediateSetSignedState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendIntermediateSetSignedState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
        } else {
            const args = argsOrState as SecretBackendIntermediateSetSignedArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.certificate === undefined) {
                throw new Error("Missing required property 'certificate'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendIntermediateSetSigned.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendIntermediateSetSigned resources.
 */
export interface SecretBackendIntermediateSetSignedState {
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The certificate
     */
    readonly certificate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendIntermediateSetSigned resource.
 */
export interface SecretBackendIntermediateSetSignedArgs {
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * The certificate
     */
    readonly certificate: pulumi.Input<string>;
}
