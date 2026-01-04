import { createBlobStorage } from "@nuxthub/core/blob";
import { createDriver } from "@nuxthub/core/blob/drivers/fs";

export { ensureBlob } from "@nuxthub/core/blob";
export const blob = createBlobStorage(createDriver({"dir":"/Users/mojito/Code/sponge/atidraw/.data/blob"}));
