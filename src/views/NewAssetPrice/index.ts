import { setTitle } from '../../router';
import pbw from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetAggregatedModel } from '../../models';


export function query(assetId: string) {
  return asyncComputed(async () => {
    const rawAsset = await pbw.getOneRecord('assets_agg', assetId);
    if (!rawAsset) return null;

    setTitle(rawAsset.name);
    const asset = AssetAggregatedModel.from(rawAsset);

    return { asset };
  });
}
1