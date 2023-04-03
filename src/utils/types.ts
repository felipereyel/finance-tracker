import { AssetType, AssetTypeMap } from "../models/types";

export const formatAssetType = (assetType: AssetType) => {
  return AssetTypeMap[assetType];
};
