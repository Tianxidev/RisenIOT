import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import mitt from "mitt";
import store from "@/stores";
import { vuetify } from "@/plugins/vuetify";
import router from "@/router";

const app = createApp(App);
const bus = mitt();

// 事件总线
app.config.globalProperties.$bus = bus;

app.use(store);
app.use(router);
app.use(vuetify);

app.mount("#app");
