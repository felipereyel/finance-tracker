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
      meta: {
        title: "Home",
      }
    },
    {
      name: "asset",
      path: "/assets/:id",
      component: () => import("./views/Asset.vue"),
      meta: {
        title: "Asset",
      }
    },
    {
      name: "new-asset-price",
      path: "/assets/:id/new-price",
      component: () => import("./views/NewAssetPrice.vue"),
      meta: {
        title: "New Asset Price",
      }
    },
    {
      name: "price",
      path: "/price/:id",
      component: () => import("./views/Price.vue"),
      meta: {
        title: "Price",
      }
    },
    {
      name: "new-asset",
      path: "/new-asset",
      component: () => import("./views/NewAsset.vue"),
      meta: {
        title: "New Asset",
      }
    },
  ],
});

createApp(App).use(router).use(PrimeVue).mount("#app");
