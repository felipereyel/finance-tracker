<template>
  <div class="col ai-s">
    <h1>Assets</h1>
  </div>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="result">
    <div class="row ai-s">
      <h3>Total: {{ formatCurrencyBRL(result.total) }}</h3>
    </div>
    <div>
      <DataTable    
        :value="result.assets" 
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
              <span class="ml-0-5">({{ result.totalByType[data.type as AssetType] }})</span>
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
import { nextTick } from 'vue';
import { useRouter } from 'vue-router';
import DataTable from 'primevue/datatable';

import { query } from "."
import { AssetType } from '../../models';
import { formatDate } from '../../utils/date';
import { formatAssetType } from '../../utils/types';
import { formatCurrencyBRL } from '../../utils/currency';

const router = useRouter();
const { result, loading, onResult } = query();

onResult(async () => {
  await nextTick();
  document.querySelectorAll('.p-rowgroup-header > td').forEach((head) => {
    head.setAttribute('colspan', '6');
  });
});

const assetUrl = (id: string) => router.resolve({ name: 'asset', params: { id } }).href;
const goToAsset = (event: any) => router.push({ name: 'asset', params: { id: event.data.id } });
const createAsset = async (type?: string) => router.push({ name: 'new-asset', query: { type } });
</script>
