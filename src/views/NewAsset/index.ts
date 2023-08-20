import { WalletModel } from '../../models';
import pbw from '../../services/pocketbase';
import { asyncComputed } from '../../utils/vue';


export function query() {
  return asyncComputed(async () => {
    const rawWallets = await pbw.getFullRecordList('wallets');

    const wallets = rawWallets.map(WalletModel.from);
    const walletOptions = wallets.map(w => ({ value: w.id, label: w.name }));
    
    return { wallets, walletOptions };
  });
}
