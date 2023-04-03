<template>
  <div class="col ai-s" v-if="route.name != 'home'">
    <router-link to="/">←</router-link>
  </div>
  <router-view></router-view>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
const route = useRoute();
</script>