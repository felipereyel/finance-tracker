import { setTitle } from '../../router';
import { pb } from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetAggregatedModel, selectArgs } from '../../models';


export function query(assetId: string) {
  return asyncComputed(async () => {
    const data = await pb.collection('assets_agg').getOne<selectArgs<"assets_agg"> | null | undefined>(assetId);
    if (!data) return null;

    setTitle(data.name);
    const asset = AssetAggregatedModel.from(data);

    return { asset };
  });
}
1