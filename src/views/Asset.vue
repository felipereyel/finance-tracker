<template>
  <div v-if="assetLoading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="asset" class="col ai-s">
    <div class="row jc-sb w-100 ai-c">
      <h2>{{ asset.name }} ({{ formatAssetType(asset.type) }})</h2>
      <button class="btn-sm" v-if="!asset.sellDate" @click="sell">Sell</button>
    </div>
    <div class="col w-100">
      <textarea class="mb-1" cols="30" rows="3" v-model="asset.comment"></textarea>
      <input v-if="asset.sellDate" type="date" v-model="asset.sellDate" />
    </div>
    <br>
    <div class="row jc-sb w-100 ai-c">
      <h3>Prices</h3>
      <button class="btn-sm" @click="addPrice">Add price</button>
    </div>
    <div v-if="pricesLoadig">
      <h4>Loading...</h4>
    </div>
    <div v-else-if="latestPrice">
      <div class="col ai-s w-100">
        <span>Initial price: <b>{{ formatCurrencyBRL(asset.initialPrice) }}</b> (@ {{ formatDate(asset.buyDate) }})</span>
        <span>Current price: <b>{{ formatCurrencyBRL(latestPrice.value) }}</b> (@ {{ formatDate(latestPrice.loggedAt) }})</span>
      </div>
      <Chart type="line" :data="chartData" :options="chartOptions" class="h-30rem" />
    </div>
    <div v-else>Error</div>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import Chart from 'primevue/chart';
import { AssetModel, AssetPriceModel } from '../models';
import { asyncComputed } from '../utils/vue';
import { formatDate } from '../utils/date';
import { formatAssetType } from '../utils/types';
import { formatCurrencyBRL } from '../utils/currency';
import { computed } from 'vue';

const route = useRoute();
const router = useRouter();

const { result: asset, loading: assetLoading } = asyncComputed(() => AssetModel.getAssetById(route.params.id as string));
const { result: prices, loading: pricesLoadig } = asyncComputed(() => AssetPriceModel.getForAsset(route.params.id as string));

const latestPrice = computed(() => {
  if (!prices.value) return;
  return prices.value[prices.value.length - 1];
});

const chartOptions = {
  scales: {
    yAxes: [
      {
        ticks: {
          beginAtZero: true,
        },
      },
    ],
  },
};

const chartData = computed(() => {
  if (!prices.value) return {};
  return {
    labels: prices.value.map((price) => formatDate(price.loggedAt)),
    datasets: [
      {
        label: 'Price',
        data: prices.value.map((price) => price.value),
        fill: false,
        borderColor: 'rgb(75, 192, 192)',
        tension: 0.1,
      },
    ],
  };
});

const addPrice = () => {
  router.push({ name: 'new-asset-price', params: { id: route.params.id } });
};

const sell = () => {
  if (!asset.value) return;
  asset.value.sellDate = new Date().toISOString();
}
</script>

<style scoped></style>
