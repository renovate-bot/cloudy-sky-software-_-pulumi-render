// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Services.Outputs
{

    [OutputType]
    public sealed class ListServiceResponse
    {
        public readonly string? Cursor;
        public readonly Union<Outputs.StaticSite, Outputs.WebService>? Service;

        [OutputConstructor]
        private ListServiceResponse(
            string? cursor,

            Union<Outputs.StaticSite, Outputs.WebService>? service)
        {
            Cursor = cursor;
            Service = service;
        }
    }
}
