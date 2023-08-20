import { setTitle } from '../../router';
import pbw from '../../services/pocketbase';
import { formatDate } from '../../utils/date';
import { asyncComputed } from '../../utils/vue';
import { AssetModel, AssetPriceModel } from '../../models';


export function query(assetId: string) {
  return asyncComputed(async () => {
    const [rawAsset, rawPrices] = await Promise.all([
        pbw.getOneRecord('assets', assetId),
        pbw.getFullRecordList('asset_prices', { 
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
