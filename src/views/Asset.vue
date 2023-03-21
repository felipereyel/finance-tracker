<template>
  <div v-if="assetLoading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="asset">
    <h2>{{ asset.name }}</h2>
    <div class="asset-columns">
      <textarea cols="30" rows="3" v-model="asset.comment"></textarea>
      <input type="date" v-model="asset.sellDate" />
    </div>
    <br>
    <h3>Prices</h3>
    <div v-if="pricesLoadig">
      <h4>Loading...</h4>
    </div>
    <div v-else>
      <ul>
        <li v-for="price in prices" :key="price.id">
          {{ formatDate(price.loggedAt) }}: <b>R${{ price.value }}</b>
        </li>
      </ul>
    </div>
    <button @click="addPrice">Add price</button>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { AssetModel, AssetPriceModel } from '../models';
import { asyncComputed } from '../utils/vue';
import { formatDate } from '../utils/date';

const route = useRoute();
const router = useRouter();

const { result: asset, loading: assetLoading } = asyncComputed(() => AssetModel.getAssetById(route.params.id as string));
const { result: prices, loading: pricesLoadig } = asyncComputed(() => AssetPriceModel.getForAsset(route.params.id as string));

const clearSellDate = () => {
  if (asset.value) {
    asset.value.sellDate = null;
  }
};

const addPrice = () => {
  router.push({ name: 'new-asset-price', params: { id: route.params.id } });
};
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

.asset-columns > *:not(:last-child) {
  margin-bottom: 1rem;
}
</style>
