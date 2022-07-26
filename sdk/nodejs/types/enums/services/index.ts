// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CustomDomainDomainType = {
    Apex: "apex",
    Subdomain: "subdomain",
} as const;

export type CustomDomainDomainType = (typeof CustomDomainDomainType)[keyof typeof CustomDomainDomainType];

export const CustomDomainVerificationStatus = {
    Verified: "verified",
    Unverified: "unverified",
} as const;

export type CustomDomainVerificationStatus = (typeof CustomDomainVerificationStatus)[keyof typeof CustomDomainVerificationStatus];

export const DeployClearCache = {
    DoNotClear: "do_not_clear",
    Clear: "clear",
} as const;

export type DeployClearCache = (typeof DeployClearCache)[keyof typeof DeployClearCache];

export const StaticSiteServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Whether to auto deploy the service or not upon git push.
 */
export type StaticSiteServiceAutoDeploy = (typeof StaticSiteServiceAutoDeploy)[keyof typeof StaticSiteServiceAutoDeploy];

export const StaticSiteServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

/**
 * The notification setting for this service upon deployment failure.
 */
export type StaticSiteServiceNotifyOnFail = (typeof StaticSiteServiceNotifyOnFail)[keyof typeof StaticSiteServiceNotifyOnFail];

export const StaticSiteServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type StaticSiteServiceSuspended = (typeof StaticSiteServiceSuspended)[keyof typeof StaticSiteServiceSuspended];

export const StaticSiteStaticSiteRouteType = {
    Redirect: "redirect",
    Rewrite: "rewrite",
} as const;

export type StaticSiteStaticSiteRouteType = (typeof StaticSiteStaticSiteRouteType)[keyof typeof StaticSiteStaticSiteRouteType];

export const StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled = (typeof StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled)[keyof typeof StaticSiteStaticSiteServiceDetailsPullRequestPreviewsEnabled];

export const WebServiceServiceAutoDeploy = {
    Yes: "yes",
    No: "no",
} as const;

/**
 * Whether to auto deploy the service or not upon git push.
 */
export type WebServiceServiceAutoDeploy = (typeof WebServiceServiceAutoDeploy)[keyof typeof WebServiceServiceAutoDeploy];

export const WebServiceServiceNotifyOnFail = {
    Default: "default",
    Notify: "notify",
    Ignore: "ignore",
} as const;

/**
 * The notification setting for this service upon deployment failure.
 */
export type WebServiceServiceNotifyOnFail = (typeof WebServiceServiceNotifyOnFail)[keyof typeof WebServiceServiceNotifyOnFail];

export const WebServiceServiceSuspended = {
    Suspended: "suspended",
    NotSuspended: "not_suspended",
} as const;

export type WebServiceServiceSuspended = (typeof WebServiceServiceSuspended)[keyof typeof WebServiceServiceSuspended];

export const WebServiceWebServiceServiceDetailsEnv = {
    Docker: "docker",
    Elixir: "elixir",
    Go: "go",
    Node: "node",
    Python: "python",
    Ruby: "ruby",
    Rust: "rust",
} as const;

export type WebServiceWebServiceServiceDetailsEnv = (typeof WebServiceWebServiceServiceDetailsEnv)[keyof typeof WebServiceWebServiceServiceDetailsEnv];

export const WebServiceWebServiceServiceDetailsPlan = {
    Starter: "starter",
    StarterPlus: "starter_plus",
    Standard: "standard",
    StandardPlus: "standard_plus",
    Pro: "pro",
    ProPlus: "pro_plus",
    ProMax: "pro_max",
    ProUltra: "pro_ultra",
} as const;

export type WebServiceWebServiceServiceDetailsPlan = (typeof WebServiceWebServiceServiceDetailsPlan)[keyof typeof WebServiceWebServiceServiceDetailsPlan];

export const WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled = {
    Yes: "yes",
    No: "no",
} as const;

export type WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled = (typeof WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled)[keyof typeof WebServiceWebServiceServiceDetailsPullRequestPreviewsEnabled];

export const WebServiceWebServiceServiceDetailsRegion = {
    Oregon: "oregon",
    Frankfurt: "frankfurt",
} as const;

export type WebServiceWebServiceServiceDetailsRegion = (typeof WebServiceWebServiceServiceDetailsRegion)[keyof typeof WebServiceWebServiceServiceDetailsRegion];
