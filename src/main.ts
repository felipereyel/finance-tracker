import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";

import "./style.css";
import App from "./App.vue";
import { initPocketBase } from "./services/pocketbase";

initPocketBase();

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: "home",
      path: "/",
      component: () => import("./views/Home.vue"),
    },
    {
      name: "about",
      path: "/about",
      component: () => import("./views/About.vue"),
    },
    {
      name: "noteEditor",
      path: "/notes/:id",
      component: () => import("./views/NoteEditor.vue"),
    },
  ],
});

createApp(App).use(router).mount("#app");
