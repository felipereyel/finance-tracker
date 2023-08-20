<template>
  <div class="col ai-s">
    <h1>Assets</h1>
  </div>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="result">
    <div class="row ai-c jc-sb">
      <h3>Total: {{ formatCurrencyBRL(total) }}</h3>
      <Dropdown v-model="state.seletectedWallet" :options="result.walletOptions" placeholder="Select Type" option-label="label" option-value="value" />
    </div>
    <div>
      <DataTable    
        :value="assets" 
        tableStyle="min-width: 50rem"
        rowGroupMode="subheader" 
        groupRowsBy="type"
        sortMode="single"
        sortField="type" 
        :sortOrder="1"
        @row-click="goToAsset"
      >
        <template #groupheader="{ data }">
          <div class="row jc-sb ai-c mt-2">
            <div class="row ai-c">
              <h2>{{ formatAssetType(data.type) }}</h2>
              <span class="ml-0-5">({{ totalByType[data.type as AssetType] }})</span>
            </div>
            <button class="btn-sm" @click="createAsset(data.type)">Create {{ formatAssetType(data.type) }}</button>
          </div>
        </template>
        <Column field="name" header="Name">
          <template #body="{ data }">
            <a :href="assetUrl(data.id)" >{{ data.name }}</a>
          </template>
        </Column>
        <Column field="buyDate" header="Buy Date">
          <template #body="{ data }">
            {{ formatDate(data.buyDate) }}
          </template>
        </Column>
        <Column field="initialPrice" header="Buy Price">
          <template #body="{ data }">
            {{ formatCurrencyBRL(data.initialPrice) }}
          </template>
        </Column>
        <Column field="latestDate" header="Latest Date">
          <template #body="{ data }">
            {{ formatDate(data.latestDate) }}
          </template>
        </Column>
        <Column field="latestPrice" header="Latest Price">
          <template #body="{ data }">
            {{ formatCurrencyBRL(data.latestPrice) }}
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
  <div v-else>
    <h2>Something went wrong</h2>
  </div>
</template>

<script setup lang="ts">
import Column from 'primevue/column';
import { useRouter } from 'vue-router';
import Dropdown from 'primevue/dropdown';
import DataTable from 'primevue/datatable';
import { computed, nextTick, reactive, watch } from 'vue';

import { query } from "."
import { AssetType } from '../../models';
import { formatDate } from '../../utils/date';
import { shakeDataTable } from '../../utils/vue';
import { formatAssetType } from '../../utils/types';
import { formatCurrencyBRL } from '../../utils/currency';

const router = useRouter();
const { result, loading } = query();

const state = reactive({
  seletectedWallet: "all",
});

const assets = computed(() => {
  if (!result.value) return [];
  if (state.seletectedWallet == "all") return result.value.assets;
  return result.value.assets.filter((asset) => asset.walletId === state.seletectedWallet);
});

const total = computed(() => {
  if (!result.value) return 0;
  return assets.value.reduce((acc, cur) => acc + cur.latestPrice, 0);
});

const totalByType = computed(() => {
  if (!result.value) return {} as Record<AssetType, string>;
  const _totalByType = assets.value.reduce((acc, cur) => {
      if (!acc[cur.type]) acc[cur.type] = 0;
      acc[cur.type] += cur.latestPrice;
      return acc;
  }, {} as Record<AssetType, number>)
  
  return Object.keys(_totalByType).reduce((acc, cur) => {
      acc[cur as AssetType] = formatCurrencyBRL(_totalByType[cur as AssetType]);
      return acc;
  }, {} as Record<AssetType, string>);
});

watch(() => assets.value, () => nextTick().then(shakeDataTable));

const assetUrl = (id: string) => router.resolve({ name: 'asset', params: { id } }).href;
const goToAsset = (event: any) => router.push({ name: 'asset', params: { id: event.data.id } });
const createAsset = async (type?: string) => {
  const query: Record<string, string> = type ? { type } : {};
  if (state.seletectedWallet !== "all") query.wallet = state.seletectedWallet;
  router.push({ name: 'new-asset', query })
};
</script>
