<template>
    <h2>New Asset</h2>
    <div class="asset-columns">
      <h3>Name</h3>
      <input v-model="state.newAsset.name" />
      <br>
      <h3>Type</h3>
      <Dropdown v-model="state.newAsset.type" :options="assetTypeOption" placeholder="Select Type" />
      <br>
      <h3>Initial Price</h3>
      <input type="number" v-model="state.newAsset.initial_price" />
      <br>
      <h3>Buy at</h3>
      <input type="date" v-model="state.newAsset.buy_date" />
      <br>
      <h3>Comment</h3>
      <textarea cols="30" rows="3" v-model="state.newAsset.comment"></textarea>
    </div>
    <br>
    <button @click="createAsset">Add Asset</button>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { AssetModel, AssetCreateDTO, AssetPriceModel } from '../models';
import { now } from '../utils/date';
import Dropdown from 'primevue/dropdown';

const router = useRouter();

const state = ref({
  newAsset: { 
    name: "New Asset", 
    initial_price: 0, 
    buy_date: now() 
  } as AssetCreateDTO,
});

const createAsset = async () => {
  try {
    const asset = await AssetModel.create(state.value.newAsset);
    await AssetPriceModel.create({
      asset_id: asset.id,
      value: asset.initialPrice,
      logged_at: asset.buy_date,
      comment: 'Initial Price',
      gain: 0,
    });
    router.push({ name: 'asset', params: { id: asset.id } });
  } catch (error) {
    console.error(error);
  }
};

const assetTypeOption = ['fii', 'federal_bond', 'cdb', 'hedge_fund', 'stock', 'other'];
</script>

<style scoped>
.asset-columns {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.asset-sell-date {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>
