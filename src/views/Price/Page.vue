<template>
  <h2>Edit Price</h2>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="result" class="col ai-s">
    <div class="col w-100">
      <div class="col ai-s mb-1">
        <span>Value</span>
        <input :value="result.price.value" type="number" disabled>
      </div>
      <div class="col ai-s mb-1">
        <span>Logget At</span>
        <input :value="result.price.loggedAt" type="date" disabled>
      </div>
      <div class="col ai-s mb-1">
        <span>Comment</span>
        <textarea cols="30" rows="3" v-model="result.price.comment"></textarea>
      </div>
      <div class="col ai-s mb-1">
        <span>Gain</span>
        <input type="number" v-model="result.price.gain">
      </div>
    </div>
    <button @click="back">Back</button>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';

import { query } from '.';

const route = useRoute();
const router = useRouter();
const { result, loading } = query(route.params.id as string);

const back = () =>{
  if (!result.value) return;
  router.push({ name: 'asset', params: { id: result.value.asset.id } });
}
</script>
