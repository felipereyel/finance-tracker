<template>
  <div class="col ai-s">
    <h1>Assets</h1>
  </div>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="sections">
    <div class="row ai-s">
      <h3>Total: {{ formatCurrencyBRL(total) }}</h3>
    </div>
    <div v-for="section in sections" :key="section.type">
      <div class="row ai-c jc-sb">
        <h2>{{ formatAssetType(section.type) }}</h2>
        <button class="btn-sm" @click="createAsset(section.type)">Create {{ formatAssetType(section.type) }}</button>
      </div>
      <DataTable :value="section.assets" tableStyle="min-width: 50rem" @row-click="goToAsset">
        <Column field="name" header="Name"></Column>
        <Column field="initialPrice" header="Buy Price">
          <template #body="{ data }">
            {{ formatCurrencyBRL(data.initialPrice) }}
          </template>
        </Column>
        <Column field="buyDate" header="Buy Date">
          <template #body="{ data }">
            {{ formatDate(data.buyDate) }}
          </template>
        </Column>
        <Column field="latestPrice" header="Latest Price">
          <template #body="{ data }">
            {{ formatCurrencyBRL(data.latestPrice) }}
          </template>
        </Column>
        <Column field="latestDate" header="Latest Date">
          <template #body="{ data }">
            {{ formatDate(data.latestDate) }}
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import { AssetAggregatedModel, assetTypeOptions } from '../models';
import { asyncComputed } from '../utils/vue';
import { formatDate } from '../utils/date';
import { formatCurrencyBRL } from '../utils/currency';
import { computed } from 'vue';
import { formatAssetType } from '../utils/types';

const router = useRouter();
const { result: assets, loading } = asyncComputed(() => AssetAggregatedModel.getLiveAssets());
const createAsset = async (type?: string) => router.push({ name: 'new-asset', query: { type } });
const goToAsset = (event: any) => router.push({ name: 'asset', params: { id: event.data.id } });

const sections = computed(
  () => assets.value 
        ? assetTypeOptions.map((ato) => {
          return {assets: assets.value?.filter((a) => a.type === ato) ?? [], type: ato}
        }).filter((s) => s.assets.length > 0) 
        : null
);

const total = computed(() => assets.value?.reduce((acc, cur) => acc + cur.latestPrice, 0) ?? 0);
</script>

<style scoped>
</style>
