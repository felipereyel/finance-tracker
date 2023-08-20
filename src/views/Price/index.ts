import { setTitle } from '../../router';
import { pb } from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetPriceModel, AssetModel, expandedArgs, selectArgs } from '../../models';

type QueryResult = selectArgs<"asset_prices"> & expandedArgs<"asset_prices", ["asset_id"]>;

export function query(priceId: string) {
  return asyncComputed(async () => {
    const data = await pb.collection('asset_prices').getOne<QueryResult | null | undefined>(priceId, {
      expand: "asset_id",
    });
    if (!data) return null;

    setTitle(data.expand.asset_id.name);
    const price = AssetPriceModel.from(data);
    const asset = AssetModel.from(data.expand.asset_id);

    return { asset, price };
  });
}
