// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an PKI certificate.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_root_sign_intermediate.html.md.
 */
export class SecretBackendRootSignIntermediate extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRootSignIntermediate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRootSignIntermediateState, opts?: pulumi.CustomResourceOptions): SecretBackendRootSignIntermediate {
        return new SecretBackendRootSignIntermediate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate';

    /**
     * Returns true if the given object is an instance of SecretBackendRootSignIntermediate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRootSignIntermediate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRootSignIntermediate.__pulumiType;
    }

    /**
     * List of alternative names
     */
    public readonly altNames!: pulumi.Output<string[] | undefined>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The CA chain
     */
    public /*out*/ readonly caChain!: pulumi.Output<string>;
    /**
     * The certificate
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * CN of intermediate to create
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * The country
     */
    public readonly country!: pulumi.Output<string | undefined>;
    /**
     * The CSR
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * Flag to exclude CN from SANs
     */
    public readonly excludeCnFromSans!: pulumi.Output<boolean | undefined>;
    /**
     * The format of data
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * List of alternative IPs
     */
    public readonly ipSans!: pulumi.Output<string[] | undefined>;
    /**
     * The issuing CA
     */
    public /*out*/ readonly issuingCa!: pulumi.Output<string>;
    /**
     * The locality
     */
    public readonly locality!: pulumi.Output<string | undefined>;
    /**
     * The maximum path length to encode in the generated certificate
     */
    public readonly maxPathLength!: pulumi.Output<number | undefined>;
    /**
     * The organization
     */
    public readonly organization!: pulumi.Output<string | undefined>;
    /**
     * List of other SANs
     */
    public readonly otherSans!: pulumi.Output<string[] | undefined>;
    /**
     * The organization unit
     */
    public readonly ou!: pulumi.Output<string | undefined>;
    /**
     * List of domains for which certificates are allowed to be issued
     */
    public readonly permittedDnsDomains!: pulumi.Output<string[] | undefined>;
    /**
     * The postal code
     */
    public readonly postalCode!: pulumi.Output<string | undefined>;
    /**
     * The province
     */
    public readonly province!: pulumi.Output<string | undefined>;
    /**
     * The serial
     */
    public /*out*/ readonly serial!: pulumi.Output<string>;
    /**
     * The street address
     */
    public readonly streetAddress!: pulumi.Output<string | undefined>;
    /**
     * Time to live
     */
    public readonly ttl!: pulumi.Output<string | undefined>;
    /**
     * List of alternative URIs
     */
    public readonly uriSans!: pulumi.Output<string[] | undefined>;
    /**
     * Preserve CSR values
     */
    public readonly useCsrValues!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SecretBackendRootSignIntermediate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRootSignIntermediateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRootSignIntermediateArgs | SecretBackendRootSignIntermediateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendRootSignIntermediateState | undefined;
            inputs["altNames"] = state ? state.altNames : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["caChain"] = state ? state.caChain : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["commonName"] = state ? state.commonName : undefined;
            inputs["country"] = state ? state.country : undefined;
            inputs["csr"] = state ? state.csr : undefined;
            inputs["excludeCnFromSans"] = state ? state.excludeCnFromSans : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["ipSans"] = state ? state.ipSans : undefined;
            inputs["issuingCa"] = state ? state.issuingCa : undefined;
            inputs["locality"] = state ? state.locality : undefined;
            inputs["maxPathLength"] = state ? state.maxPathLength : undefined;
            inputs["organization"] = state ? state.organization : undefined;
            inputs["otherSans"] = state ? state.otherSans : undefined;
            inputs["ou"] = state ? state.ou : undefined;
            inputs["permittedDnsDomains"] = state ? state.permittedDnsDomains : undefined;
            inputs["postalCode"] = state ? state.postalCode : undefined;
            inputs["province"] = state ? state.province : undefined;
            inputs["serial"] = state ? state.serial : undefined;
            inputs["streetAddress"] = state ? state.streetAddress : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["uriSans"] = state ? state.uriSans : undefined;
            inputs["useCsrValues"] = state ? state.useCsrValues : undefined;
        } else {
            const args = argsOrState as SecretBackendRootSignIntermediateArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.commonName === undefined) {
                throw new Error("Missing required property 'commonName'");
            }
            if (!args || args.csr === undefined) {
                throw new Error("Missing required property 'csr'");
            }
            inputs["altNames"] = args ? args.altNames : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["commonName"] = args ? args.commonName : undefined;
            inputs["country"] = args ? args.country : undefined;
            inputs["csr"] = args ? args.csr : undefined;
            inputs["excludeCnFromSans"] = args ? args.excludeCnFromSans : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["ipSans"] = args ? args.ipSans : undefined;
            inputs["locality"] = args ? args.locality : undefined;
            inputs["maxPathLength"] = args ? args.maxPathLength : undefined;
            inputs["organization"] = args ? args.organization : undefined;
            inputs["otherSans"] = args ? args.otherSans : undefined;
            inputs["ou"] = args ? args.ou : undefined;
            inputs["permittedDnsDomains"] = args ? args.permittedDnsDomains : undefined;
            inputs["postalCode"] = args ? args.postalCode : undefined;
            inputs["province"] = args ? args.province : undefined;
            inputs["streetAddress"] = args ? args.streetAddress : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["uriSans"] = args ? args.uriSans : undefined;
            inputs["useCsrValues"] = args ? args.useCsrValues : undefined;
            inputs["caChain"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["issuingCa"] = undefined /*out*/;
            inputs["serial"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendRootSignIntermediate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRootSignIntermediate resources.
 */
export interface SecretBackendRootSignIntermediateState {
    /**
     * List of alternative names
     */
    readonly altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The CA chain
     */
    readonly caChain?: pulumi.Input<string>;
    /**
     * The certificate
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * CN of intermediate to create
     */
    readonly commonName?: pulumi.Input<string>;
    /**
     * The country
     */
    readonly country?: pulumi.Input<string>;
    /**
     * The CSR
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    readonly excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The format of data
     */
    readonly format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    readonly ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The issuing CA
     */
    readonly issuingCa?: pulumi.Input<string>;
    /**
     * The locality
     */
    readonly locality?: pulumi.Input<string>;
    /**
     * The maximum path length to encode in the generated certificate
     */
    readonly maxPathLength?: pulumi.Input<number>;
    /**
     * The organization
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    readonly otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization unit
     */
    readonly ou?: pulumi.Input<string>;
    /**
     * List of domains for which certificates are allowed to be issued
     */
    readonly permittedDnsDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code
     */
    readonly postalCode?: pulumi.Input<string>;
    /**
     * The province
     */
    readonly province?: pulumi.Input<string>;
    /**
     * The serial
     */
    readonly serial?: pulumi.Input<string>;
    /**
     * The street address
     */
    readonly streetAddress?: pulumi.Input<string>;
    /**
     * Time to live
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    readonly uriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Preserve CSR values
     */
    readonly useCsrValues?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SecretBackendRootSignIntermediate resource.
 */
export interface SecretBackendRootSignIntermediateArgs {
    /**
     * List of alternative names
     */
    readonly altNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The PKI secret backend the resource belongs to.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * CN of intermediate to create
     */
    readonly commonName: pulumi.Input<string>;
    /**
     * The country
     */
    readonly country?: pulumi.Input<string>;
    /**
     * The CSR
     */
    readonly csr: pulumi.Input<string>;
    /**
     * Flag to exclude CN from SANs
     */
    readonly excludeCnFromSans?: pulumi.Input<boolean>;
    /**
     * The format of data
     */
    readonly format?: pulumi.Input<string>;
    /**
     * List of alternative IPs
     */
    readonly ipSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The locality
     */
    readonly locality?: pulumi.Input<string>;
    /**
     * The maximum path length to encode in the generated certificate
     */
    readonly maxPathLength?: pulumi.Input<number>;
    /**
     * The organization
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * List of other SANs
     */
    readonly otherSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The organization unit
     */
    readonly ou?: pulumi.Input<string>;
    /**
     * List of domains for which certificates are allowed to be issued
     */
    readonly permittedDnsDomains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The postal code
     */
    readonly postalCode?: pulumi.Input<string>;
    /**
     * The province
     */
    readonly province?: pulumi.Input<string>;
    /**
     * The street address
     */
    readonly streetAddress?: pulumi.Input<string>;
    /**
     * Time to live
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * List of alternative URIs
     */
    readonly uriSans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Preserve CSR values
     */
    readonly useCsrValues?: pulumi.Input<boolean>;
}
