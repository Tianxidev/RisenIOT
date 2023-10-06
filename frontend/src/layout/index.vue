<template>
    <mi-layout>
        <template v-slot:headerExtra>
            <mi-search :data="searchData" search-key="title" :width="120" :height="48" placeholder="搜索组件" :list-width="320"
                :list-height="335" :gap="4" :page-size="3" :list-radius="8" border-color="transparent"
                background-color="transparent" />
        </template>
    </mi-layout>
</template>

<script setup>
import { getCurrentInstance } from 'vue'
import { DashboardOutlined, LayoutOutlined } from '@ant-design/icons-vue'

const { appContext: { config: { globalProperties: vm } } } = getCurrentInstance()

vm.$g.menus.items = [{
    name: 'dashboard',
    path: '/dashboard',
    meta: {
        title: '控制中心',
        subTitle: 'Dashboard',
        icon: DashboardOutlined,
        tag: { color: '#f50', content: 'Hot' }
    }
}, {
    // ... sider menu
}]

const searchData = [{
    title: '页面布局',
    content: '基于 Layout 组件的二次定制',
    link: '/layout',
    icon: LayoutOutlined
}, {
    // ... search data
}]
</script>