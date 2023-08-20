import { setTitle } from '../../router';
import pbw from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetPriceModel, AssetModel } from '../../models';

export function query(priceId: string) {
  return asyncComputed(async () => {
    const data = await pbw.getOneExpandedRecord("asset_prices", priceId, ["asset_id"]);
    if (!data) return null;

    setTitle(data.expand.asset_id.name);
    const price = AssetPriceModel.from(data);
    const asset = AssetModel.from(data.expand.asset_id);

    return { asset, price };
  });
}
