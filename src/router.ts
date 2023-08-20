import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: "home",
      path: "/",
      component: () => import("./views/Home/Page.vue"),
      meta: {
        title: "Home",
      }
    },
    {
      name: "asset",
      path: "/assets/:id",
      component: () => import("./views/Asset/Page.vue"),
      meta: {
        title: "Asset",
      }
    },
    {
      name: "new-asset-price",
      path: "/assets/:id/new-price",
      component: () => import("./views/NewAssetPrice/Page.vue"),
      meta: {
        title: "New Asset Price",
      }
    },
    {
      name: "price",
      path: "/price/:id",
      component: () => import("./views/Price/Page.vue"),
      meta: {
        title: "Price",
      }
    },
    {
      name: "new-asset",
      path: "/new-asset",
      component: () => import("./views/NewAsset/Page.vue"),
      meta: {
        title: "New Asset",
      }
    },
  ],
});

export const setTitle = (title?: string) => {
  let docTitle = `${router.currentRoute.value.meta.title} | Finance Tracker`;
  if (title) docTitle = `${title} | ${docTitle}`;
  window.document.title = docTitle;
}