<template>
  <h2>New Asset</h2>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div class="asset-columns" v-else-if="result">
    <h3>Name</h3>
    <input v-model="state.newAsset.name" />
    <h3>Wallet</h3>
    <Dropdown v-model="state.newAsset.wallet" :options="result.walletOptions" placeholder="Select Type" option-label="label" option-value="value" />
    <h3>Type</h3>
    <Dropdown v-model="state.newAsset.type" :options="assetTypeOptions" placeholder="Select Type" />
    <h3>Initial Price</h3>
    <input type="number" v-model="state.newAsset.initial_price" />
    <h3>Buy at</h3>
    <input type="date" v-model="state.newAsset.buy_date" />
    <h3>Comment</h3>
    <textarea cols="30" rows="3" v-model="state.newAsset.comment"></textarea>
    <button @click="createAsset">Add Asset</button>
  </div>
  <div v-else>
    <h2>Something went wrong</h2>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Dropdown from 'primevue/dropdown';
import { useRouter, useRoute } from 'vue-router';

import { query } from '.';
import { now } from '../../utils/date';
import { setTitle } from '../../router';
import { AssetModel, AssetCreateDTO, AssetPriceModel, assetTypeOptions } from '../../models';

setTitle();
const route = useRoute();
const router = useRouter();

const { result, loading } = query();

const state = ref({
  newAsset: { 
    name: "New Asset", 
    initial_price: 0, 
    buy_date: now() ,
    type: route.query.type as string | undefined,
    wallet: route.query.wallet as string | undefined,
  } as AssetCreateDTO,
});

const createAsset = async () => {
  try {
    const asset = await AssetModel.create(state.value.newAsset);
    await AssetPriceModel.create({
      asset_id: asset.id,
      value: asset.initialPrice,
      logged_at: asset.buyDate,
      comment: 'Initial Price',
      gain: 0,
    });
    router.push({ name: 'asset', params: { id: asset.id } });
  } catch (error) {
    console.error(error);
  }
};

</script>

<style scoped>
.asset-columns {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.asset-columns > *:not(:last-child) {
  margin-bottom: 1rem;
}
</style>
