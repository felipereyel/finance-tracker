import { setTitle } from '../../router';
import pbw from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetAggregatedModel } from '../../models';


export function query(assetId: string) {
  return asyncComputed(async () => {
    const data = await pbw.getOneRecord('assets_agg', assetId);
    if (!data) return null;

    setTitle(data.name);
    const asset = AssetAggregatedModel.from(data);

    return { asset };
  });
}
1