import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import PrimeVue from "primevue/config";

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
      name: "asset",
      path: "/assets/:id",
      component: () => import("./views/Asset.vue"),
    },
    {
      name: "new-asset-price",
      path: "/assets/:id/new-price",
      component: () => import("./views/NewAssetPrice.vue"),
    },
    {
      name: "new-asset",
      path: "/new-asset",
      component: () => import("./views/NewAsset.vue"),
    },
  ],
});

createApp(App).use(router).use(PrimeVue).mount("#app");
