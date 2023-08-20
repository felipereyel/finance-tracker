import { setTitle } from '../../router';
import { pb } from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { formatCurrencyBRL } from '../../utils/currency';
import { AssetAggregatedModel, AssetType, selectArgs } from '../../models';


export function query() {
  return asyncComputed(async () => {
    setTitle('Assets');
    const data = await pb.collection('assets_agg').getFullList<selectArgs<"assets_agg">>({
        filter: "sell_date = NULL",
        sort: "buy_date",
    });

    const assets = data.map(AssetAggregatedModel.from);
    const total = assets.reduce((acc, cur) => acc + cur.latestPrice, 0);
    
    const _totalByType = assets.reduce((acc, cur) => {
        if (!acc[cur.type]) acc[cur.type] = 0;
        acc[cur.type] += cur.latestPrice;
        return acc;
    }, {} as Record<AssetType, number>)
    
    const totalByType = Object.keys(_totalByType).reduce((acc, cur) => {
        acc[cur as AssetType] = formatCurrencyBRL(_totalByType[cur as AssetType]);
        return acc;
    }, {} as Record<AssetType, string>);

    return { assets, total, totalByType };
  });
}
