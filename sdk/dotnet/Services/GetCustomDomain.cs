// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services
{
    public static class GetCustomDomain
    {
        public static Task<GetCustomDomainResult> InvokeAsync(GetCustomDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomDomainResult>("render:services:getCustomDomain", args ?? new GetCustomDomainArgs(), options.WithDefaults());

        public static Output<GetCustomDomainResult> Invoke(GetCustomDomainInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCustomDomainResult>("render:services:getCustomDomain", args ?? new GetCustomDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID or name of the custom domain
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetCustomDomainArgs()
        {
        }
    }

    public sealed class GetCustomDomainInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Required) The ID or name of the custom domain
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// (Required) The ID of the service
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetCustomDomainInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomDomainResult
    {
        public readonly Outputs.CustomDomain Items;

        [OutputConstructor]
        private GetCustomDomainResult(Outputs.CustomDomain items)
        {
            Items = items;
        }
    }
}
