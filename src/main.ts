import { createApp } from "vue";
import PrimeVue from "primevue/config";

import "./style.css";
import App from "./App.vue";
import { router } from "./router";
import { initPocketBase } from "./services/pocketbase";

initPocketBase();
createApp(App).use(router).use(PrimeVue).mount("#app");
