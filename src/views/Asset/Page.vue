<template>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="result" class="col ai-s">
    <div class="row jc-sb w-100 ai-c">
      <h2>{{ result.asset.name }} ({{ formatAssetType(result.asset.type) }})</h2>
      <button class="btn-sm" v-if="!result.asset.sellDate" @click="sell">Sell</button>
    </div>
    <div class="col w-100">
      <textarea class="mb-1" cols="30" rows="3" v-model="result.asset.comment"></textarea>
      <input v-if="result.asset.sellDate" type="date" v-model="result.asset.sellDate" />
    </div>
    <div class="row jc-sb w-100 ai-c">
      <h3>Prices</h3>
      <button class="btn-sm" @click="addPrice">Add price</button>
    </div>
    <div v-if="result.latestPrice">
      <div class="col ai-s w-100">
        <span>Initial price: <b>{{ formatCurrencyBRL(result.asset.initialPrice) }}</b> (@ {{ formatDate(result.asset.buyDate) }})</span>
        <span>Current price: <b>{{ formatCurrencyBRL(result.latestPrice.value) }}</b> (@ {{ formatDate(result.latestPrice.loggedAt) }})</span>
      </div>
      <Chart type="line" :data="result.chart.data" :options="result.chart.options" class="h-30rem" @select="select" />
    </div>
    <div v-else>Error</div>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import 'chartjs-adapter-moment';
import Chart from 'primevue/chart';
import { useRoute, useRouter } from 'vue-router';

import { query } from '.';
import { formatDate } from '../../utils/date';
import { formatAssetType } from '../../utils/types';
import { formatCurrencyBRL } from '../../utils/currency';

const route = useRoute();
const router = useRouter();
const { result, loading } = query(route.params.id as string);

const addPrice = () => {
  router.push({ name: 'new-asset-price', params: { id: route.params.id } });
};

const sell = () => {
  if (!result.value) return;
  result.value.asset.sellDate = new Date().toISOString();
}

const select = ({ element }: any) => {
  if (!element || !result.value) return;

  const price = result.value.prices[element.index]
  if (!price) return;
  
  router.push({ name: 'price', params: { id: price.id } });
}
</script>

<style scoped></style>
