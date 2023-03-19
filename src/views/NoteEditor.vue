<template>
  <h1>Notes</h1>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else-if="note">
    <h2>{{ note.title }}</h2>
    <textarea cols="30" rows="10" v-model="note.content"></textarea>
  </div>
  <div v-else>
    <h2>Not found</h2>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import { NoteModel } from '../models/notes';
import { asyncComputed } from '../utils/vue';

const route = useRoute();

const { result: note, loading } = asyncComputed(() => NoteModel.getNoteById(route.params.id as string));
</script>

<style scoped></style>
