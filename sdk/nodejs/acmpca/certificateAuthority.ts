// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

import {Tags} from "../index";

/**
 * Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).
 * 
 * ~> **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificate_signing_request` attribute and import the signed certificate outside of Terraform. Terraform can support another resource to manage that workflow automatically in the future.
 */
export class CertificateAuthority extends pulumi.CustomResource {
    /**
     * Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateAuthorityState): CertificateAuthority {
        return new CertificateAuthority(name, <any>state, { id });
    }

    /**
     * Amazon Resource Name (ARN) of the certificate authority.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly certificate: pulumi.Output<string>;
    /**
     * Nested argument containing algorithms and certificate subject information. Defined below.
     */
    public readonly certificateAuthorityConfiguration: pulumi.Output<{ keyAlgorithm: string, signingAlgorithm: string, subject: { commonName?: string, country?: string, distinguishedNameQualifier?: string, generationQualifier?: string, givenName?: string, initials?: string, locality?: string, organization?: string, organizationalUnit?: string, pseudonym?: string, state?: string, surname?: string, title?: string } }>;
    /**
     * Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly certificateChain: pulumi.Output<string>;
    /**
     * The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
     */
    public /*out*/ readonly certificateSigningRequest: pulumi.Output<string>;
    /**
     * Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
     */
    public readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly notAfter: pulumi.Output<string>;
    /**
     * Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly notBefore: pulumi.Output<string>;
    /**
     * Nested argument containing revocation configuration. Defined below.
     */
    public readonly revocationConfiguration: pulumi.Output<{ crlConfiguration?: { customCname?: string, enabled?: boolean, expirationInDays: number, s3BucketName?: string } } | undefined>;
    /**
     * Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly serial: pulumi.Output<string>;
    /**
     * Status of the certificate authority.
     */
    public /*out*/ readonly status: pulumi.Output<string>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    public readonly tags: pulumi.Output<Tags | undefined>;
    /**
     * The type of the certificate authority. Currently, this must be `SUBORDINATE`.
     */
    public readonly type: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateAuthority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateAuthorityArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: CertificateAuthorityArgs | CertificateAuthorityState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CertificateAuthorityState = argsOrState as CertificateAuthorityState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateAuthorityConfiguration"] = state ? state.certificateAuthorityConfiguration : undefined;
            inputs["certificateChain"] = state ? state.certificateChain : undefined;
            inputs["certificateSigningRequest"] = state ? state.certificateSigningRequest : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["notAfter"] = state ? state.notAfter : undefined;
            inputs["notBefore"] = state ? state.notBefore : undefined;
            inputs["revocationConfiguration"] = state ? state.revocationConfiguration : undefined;
            inputs["serial"] = state ? state.serial : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as CertificateAuthorityArgs | undefined;
            if (!args || args.certificateAuthorityConfiguration === undefined) {
                throw new Error("Missing required property 'certificateAuthorityConfiguration'");
            }
            inputs["certificateAuthorityConfiguration"] = args ? args.certificateAuthorityConfiguration : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["revocationConfiguration"] = args ? args.revocationConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["certificateChain"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
            inputs["notAfter"] = undefined /*out*/;
            inputs["notBefore"] = undefined /*out*/;
            inputs["serial"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        super("aws:acmpca/certificateAuthority:CertificateAuthority", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateAuthority resources.
 */
export interface CertificateAuthorityState {
    /**
     * Amazon Resource Name (ARN) of the certificate authority.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * Nested argument containing algorithms and certificate subject information. Defined below.
     */
    readonly certificateAuthorityConfiguration?: pulumi.Input<{ keyAlgorithm: pulumi.Input<string>, signingAlgorithm: pulumi.Input<string>, subject: pulumi.Input<{ commonName?: pulumi.Input<string>, country?: pulumi.Input<string>, distinguishedNameQualifier?: pulumi.Input<string>, generationQualifier?: pulumi.Input<string>, givenName?: pulumi.Input<string>, initials?: pulumi.Input<string>, locality?: pulumi.Input<string>, organization?: pulumi.Input<string>, organizationalUnit?: pulumi.Input<string>, pseudonym?: pulumi.Input<string>, state?: pulumi.Input<string>, surname?: pulumi.Input<string>, title?: pulumi.Input<string> }> }>;
    /**
     * Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
     */
    readonly certificateChain?: pulumi.Input<string>;
    /**
     * The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
     */
    readonly certificateSigningRequest?: pulumi.Input<string>;
    /**
     * Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    readonly notAfter?: pulumi.Input<string>;
    /**
     * Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    readonly notBefore?: pulumi.Input<string>;
    /**
     * Nested argument containing revocation configuration. Defined below.
     */
    readonly revocationConfiguration?: pulumi.Input<{ crlConfiguration?: pulumi.Input<{ customCname?: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, expirationInDays: pulumi.Input<number>, s3BucketName?: pulumi.Input<string> }> }>;
    /**
     * Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
     */
    readonly serial?: pulumi.Input<string>;
    /**
     * Status of the certificate authority.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * The type of the certificate authority. Currently, this must be `SUBORDINATE`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertificateAuthority resource.
 */
export interface CertificateAuthorityArgs {
    /**
     * Nested argument containing algorithms and certificate subject information. Defined below.
     */
    readonly certificateAuthorityConfiguration: pulumi.Input<{ keyAlgorithm: pulumi.Input<string>, signingAlgorithm: pulumi.Input<string>, subject: pulumi.Input<{ commonName?: pulumi.Input<string>, country?: pulumi.Input<string>, distinguishedNameQualifier?: pulumi.Input<string>, generationQualifier?: pulumi.Input<string>, givenName?: pulumi.Input<string>, initials?: pulumi.Input<string>, locality?: pulumi.Input<string>, organization?: pulumi.Input<string>, organizationalUnit?: pulumi.Input<string>, pseudonym?: pulumi.Input<string>, state?: pulumi.Input<string>, surname?: pulumi.Input<string>, title?: pulumi.Input<string> }> }>;
    /**
     * Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Nested argument containing revocation configuration. Defined below.
     */
    readonly revocationConfiguration?: pulumi.Input<{ crlConfiguration?: pulumi.Input<{ customCname?: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, expirationInDays: pulumi.Input<number>, s3BucketName?: pulumi.Input<string> }> }>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * The type of the certificate authority. Currently, this must be `SUBORDINATE`.
     */
    readonly type?: pulumi.Input<string>;
}
