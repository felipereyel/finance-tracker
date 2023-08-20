import { setTitle } from '../../router';
import pbw from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetPriceModel, AssetModel } from '../../models';

export function query(priceId: string) {
  return asyncComputed(async () => {
    const rawPrice = await pbw.getOneExpandedRecord("asset_prices", priceId, ["asset_id"]);
    if (!rawPrice) return null;

    setTitle(rawPrice.expand.asset_id.name);
    const price = AssetPriceModel.from(rawPrice);
    const asset = AssetModel.from(rawPrice.expand.asset_id);

    return { asset, price };
  });
}
