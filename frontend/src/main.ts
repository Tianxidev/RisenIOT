import {createApp} from 'vue';
import pinia from '/@/stores/index';
import App from './App.vue';
import router from './router';
import {directive} from '/@/utils/directive';
import {i18n} from '/@/i18n';
import other from '/@/utils/other';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import '/@/theme/index.scss';
import mitt from 'mitt';
import VueGridLayout from 'vue-grid-layout';
import 'default-passive-events';
import {getUpFileUrl, handleTree, parseTime, selectDictLabel} from '/@/utils/gfast';
import {useDict} from '/@/api/system/dict/data';
import {getItems, setItems, getOptionValue, isEmpty} from '/@/api/items'

import VForm3 from 'vform3-builds'  // 引入VForm3库
import 'vform3-builds/dist/designer.style.css'  //引入VForm3样式

// 分页组件
import pagination from '/@/components/pagination/index.vue'

// 大文件上传组件
// @ts-ignore
import uploader from 'vue-simple-uploader'
import 'vue-simple-uploader/dist/style.css'


const app = createApp(App);

directive(app);
other.elSvg(app);

app.component('pagination', pagination)
app.use(pinia)
app.use(uploader)
app.use(router)
app.use(ElementPlus, {i18n: i18n.global.t})
app.use(VForm3)  // 全局注册VForm3(同时注册了v-form-designe、v-form-render等组件)
app.use(i18n)
app.use(VueGridLayout)
app.mount('#app');

app.config.globalProperties.getUpFileUrl = getUpFileUrl
app.config.globalProperties.handleTree = handleTree
app.config.globalProperties.useDict = useDict
app.config.globalProperties.selectDictLabel = selectDictLabel

app.config.globalProperties.getItems = getItems
app.config.globalProperties.setItems = setItems
app.config.globalProperties.getOptionValue = getOptionValue
app.config.globalProperties.isEmpty = isEmpty
app.config.globalProperties.parseTime = parseTime

const globalProperties = {
    mittBus: mitt(),
    i18n
}


//必须合并vue默认的变量，否则有问题
app.config.globalProperties = Object.assign(
    app.config.globalProperties,
    globalProperties
);
