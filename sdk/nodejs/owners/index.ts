// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as utilities from "../utilities";

// Export members:
export { GetOwnerArgs, GetOwnerResult, GetOwnerOutputArgs } from "./getOwner";
export const getOwner: typeof import("./getOwner").getOwner = null as any;
export const getOwnerOutput: typeof import("./getOwner").getOwnerOutput = null as any;
utilities.lazyLoad(exports, ["getOwner","getOwnerOutput"], () => require("./getOwner"));

export { ListOwnersArgs, ListOwnersResult } from "./listOwners";
export const listOwners: typeof import("./listOwners").listOwners = null as any;
utilities.lazyLoad(exports, ["listOwners"], () => require("./listOwners"));


// Export enums:
export * from "../types/enums/owners";
