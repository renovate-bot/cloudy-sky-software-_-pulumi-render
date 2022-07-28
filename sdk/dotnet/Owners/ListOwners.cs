// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Render.Owners
{
    public static class ListOwners
    {
        public static Task<ListOwnersResult> InvokeAsync(ListOwnersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListOwnersResult>("render:owners:listOwners", args ?? new ListOwnersArgs(), options.WithDefaults());
    }


    public sealed class ListOwnersArgs : Pulumi.InvokeArgs
    {
        public ListOwnersArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListOwnersResult
    {
        public readonly ImmutableArray<Outputs.ListOwnersResponse> Items;

        [OutputConstructor]
        private ListOwnersResult(ImmutableArray<Outputs.ListOwnersResponse> items)
        {
            Items = items;
        }
    }
}
