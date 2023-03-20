<template>
  <div v-if="assetLoading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="asset">
    <h2>New Price: {{ asset.name }}</h2>
    <span>Inital price was R${{ asset.initialPrice }}</span>
    <div class="asset-columns">
      <h3>Price</h3>
      <input type="number" v-model="state.newPrice.value" />
      <br>
      <h3>Logged at</h3>
      <input type="date" v-model="state.newPrice.logged_at" />
      <br>
      <h3>Gain</h3>
      <input type="number" v-model="state.newPrice.gain" />
      <br>
      <h3>Comment</h3>
      <textarea cols="30" rows="3" v-model="state.newPrice.comment"></textarea>
    </div>
    <br>
    <button @click="createPrice">Add price</button>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { AssetModel, AssetPriceModel, AssetPriceCreateDTO } from '../models';
import { asyncComputed } from '../utils/vue';
import { now } from '../utils/date';

const route = useRoute();
const router = useRouter();

const { result: asset, loading: assetLoading } = asyncComputed(() => AssetModel.getAssetById(route.params.id as string));

const state = ref({
  newPrice: { asset_id: route.params.id, gain: 0, logged_at: now() } as AssetPriceCreateDTO,
});

const createPrice = async () => {
  try {
    await AssetPriceModel.create(state.value.newPrice);
    router.push({ name: 'asset', params: { id: route.params.id } });
  } catch (error) {
    console.error(error);
  }
};

watch(() => asset.value?.initialPrice, () => {
  if (!asset.value) return
  state.value.newPrice.value = asset.value.initialPrice;
})

watch(() => state.value.newPrice.value, () => {
  if (!asset.value) return
  state.value.newPrice.gain = state.value.newPrice.value - asset.value.initialPrice;
  state.value.newPrice.gain = Math.round(state.value.newPrice.gain * 100) / 100;
})

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
