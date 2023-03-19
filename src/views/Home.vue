<template>
  <h1>Notes</h1>
  <div v-if="loading">
    <h2>Loading...</h2>
  </div>
  <div v-else>
    <ul>
      <li v-for="note in notes" :key="note.id">
        <router-link :to="{ name: 'noteEditor', params: { id: note.id } }">
          {{ note.title }}
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { NoteModel } from '../models/notes';
import { asyncComputed } from '../utils/vue';

const { result: notes, loading } = asyncComputed(() => NoteModel.getAllNotes());
</script>

<style scoped></style>
