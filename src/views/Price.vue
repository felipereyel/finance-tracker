<template>
  <h2>Edit Price</h2>
  <div v-if="priceLoadig">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="price" class="col ai-s">
    <div class="col w-100">
      <div class="col ai-s mb-1">
        <span>Value</span>
        <input :value="price.value" type="number" disabled>
      </div>
      <div class="col ai-s mb-1">
        <span>Logget At</span>
        <input :value="price.loggedAt" type="date" disabled>
      </div>
      <div class="col ai-s mb-1">
        <span>Comment</span>
        <textarea cols="30" rows="3" v-model="price.comment"></textarea>
      </div>
      <div class="col ai-s mb-1">
        <span>Gain</span>
        <input type="number" v-model="price.gain">
      </div>
    </div>
    <button @click="router.push({ name: 'asset', params: { id: price?.assetID } })">Back</button>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';

import { watch } from 'vue';
import { setTitle } from '../utils/title';
import { AssetPriceModel } from '../models';
import { asyncComputed } from '../utils/vue';

const route = useRoute();
const router = useRouter();
const { result: price, loading: priceLoadig } = asyncComputed(() => AssetPriceModel.getPriceById(route.params.id as string));

watch(() => price.value, () => {
  if (!price.value?.asset) return
  setTitle(route, price.value.asset.name)
})
</script>

<style scoped></style>
