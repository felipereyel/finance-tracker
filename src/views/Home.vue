<template>
  <h1>Assets</h1>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else>
    <ul>
      <li v-for="note in assets" :key="note.id">
        <router-link :to="{ name: 'asset', params: { id: note.id } }">
          {{ note.name }}
        </router-link>
      </li>
    </ul>
  </div>
  <button @click="createAsset">Create Asset</button>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { AssetModel } from '../models';
import { asyncComputed } from '../utils/vue';
const router = useRouter();

const { result: assets, loading } = asyncComputed(() => AssetModel.getLiveAssets());

const createAsset = async () => router.push({ name: 'new-asset' });
</script>

<style scoped></style>
