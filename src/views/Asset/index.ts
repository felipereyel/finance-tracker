import { setTitle } from '../../router';
import { formatDate } from '../../utils/date';
import { pb } from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetModel, AssetPriceModel, selectArgs } from '../../models';


export function query(assetId: string) {
  return asyncComputed(async () => {
    const [rawAsset, rawPrices] = await Promise.all([
        pb.collection('assets').getOne<selectArgs<"assets"> | null | undefined>(assetId),
        pb.collection('asset_prices').getFullList<selectArgs<"asset_prices">>({ 
            filter: `asset_id = "${assetId}"`, // INJECTION WARNING
            sort: "logged_at" 
        })
    ]);

    if (!rawAsset) return null;

    setTitle(rawAsset.name);
    const asset = AssetModel.from(rawAsset);
    const prices = rawPrices.map(AssetPriceModel.from);
    const latestPrice = prices[prices.length - 1];

    const chart = {
        options: {
            scales: {
                y: {
                    ticks: {
                        beginAtZero: true,
                    },
                },
                x: {
                    type: 'time',
                    time: {
                        unit: 'day',
                    },
                }
            },
        },
        data: {
            labels: prices.map((price) => formatDate(price.loggedAt)),
            datasets: [
              {
                label: 'Price',
                data: prices.map((price) => price.value),
                fill: false,
                borderColor: 'rgb(105, 193, 102)',
              },
            ],
        }
    }
      
    return { asset, prices, latestPrice, chart };
  });
}
