<template>
  <div class="col ai-s">
    <h1>Assets</h1>
  </div>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else>
    <div class="row ai-s">
      <h3>Total: {{ formatCurrencyBRL(total) }}</h3>
    </div>
    <div v-if="assets">
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
              <span class="ml-0-5">({{ calculateTypeTotal(data.type) }})</span>
            </div>
            <button class="btn-sm" @click="createAsset(data.type)">Create {{ formatAssetType(data.type) }}</button>
          </div>
        </template>
        <Column field="name" header="Name"></Column>
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
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { computed, nextTick, watch } from 'vue';
import Column from 'primevue/column';
import DataTable from 'primevue/datatable';

import { setTitle } from '../utils/title';
import { formatDate } from '../utils/date';
import { asyncComputed } from '../utils/vue';
import { formatAssetType } from '../utils/types';
import { AssetAggregatedModel } from '../models';
import { formatCurrencyBRL } from '../utils/currency';

const route = useRoute();
const router = useRouter();
setTitle(route, 'Assets');

const { result: assets, loading } = asyncComputed(() => AssetAggregatedModel.getLiveAssets());
const total = computed(() => assets.value?.reduce((acc, cur) => acc + cur.latestPrice, 0) ?? 0);

const createAsset = async (type?: string) => router.push({ name: 'new-asset', query: { type } });
const goToAsset = (event: any) => router.push({ name: 'asset', params: { id: event.data.id } });
const calculateTypeTotal = (type: string) => formatCurrencyBRL(assets.value?.filter(a => a.type === type).reduce((acc, cur) => acc + cur.latestPrice, 0) ?? 0);

watch(() => assets.value, async () => {
  await nextTick();
  document.querySelectorAll('.p-rowgroup-header > td').forEach((head) => {
    head.setAttribute('colspan', '6');
  });
});
</script>

<style scoped>
</style>
