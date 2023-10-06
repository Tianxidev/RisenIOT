import { createApp } from 'vue'
import './style.css'
import router from './router';
import App from './App.vue'
import MakeitAdminPro from 'makeit-admin-pro'
import 'makeit-admin-pro/dist/miitvip.min.css'

const app = createApp(App);


// 使用路由
app.use(router);

// 使用组件库
app.use(MakeitAdminPro);

// 挂载应用
app.mount('#app');