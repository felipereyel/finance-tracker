import { setTitle } from '../../router';
import pbw from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';
import { AssetAggregatedModel, WalletModel } from '../../models';


export function query() {
  return asyncComputed(async () => {
    setTitle('Assets');
    const [rawAssets, rawWallets] = await Promise.all([
      pbw.getFullRecordList('assets_agg', { sort: 'buy_date', filter: 'sell_date = NULL' }),
      pbw.getFullRecordList('wallets'),
    ]);

    const wallets = rawWallets.map(WalletModel.from);
    const assets = rawAssets.map(AssetAggregatedModel.from);
    const walletOptions = [{value: "all", label: "All"}, ...wallets.map(w => ({ value: w.id, label: w.name }))];
    
    return { assets, wallets, walletOptions };
  });
}
