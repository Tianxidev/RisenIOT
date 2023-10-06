
const frameIn = [
    {
        path: '/',
        redirect: {name: 'index'},
        component: () => import('@/layout/index.vue'),
        children: [
            {
                path: '/index',
                name: 'index',
                meta: {
                    cache: true,
                    title: '首页',
                    icon: 'shouye',
                    requiresAuth: true,
                },
                component: () => import('@/views/home/index.vue'),
            },
            {
                path: '/*',
                name: '404',
                hidden: true,
                meta: {
                    title: '404',
                },
                component: () => import('@/views/error/401.vue'),
            },
        ]
    },
]

/**
 * 在主框架之外显示
 */
const frameOut = [
    // 登录
    {
        path: '/login',
        name: 'login',
        meta: {
            title: '登录',
        },
        component: () => import('@/views/error/401.vue'),
    },
    {
        path: '/dataCenter',
        name: 'dataCenter',
        meta: {
            title: '大屏展示',
        },
        component: () => import('@/views/error/401.vue'),
    }
]

/**
 * 错误页面
 */
const errorPage = [
    {
        path: '/401',
        name: '401',
        component: () => import('@/views/error/401.vue'),
        meta: {
            title: '401',
        },
    },
    {
        path: '/:pathMatch(.*)*',
        name: '404',
        component: () => import('@/views/error/404.vue'),
        meta: {
            title: '404',
        },
    }
]

// 导出需要显示菜单的
export const frameInRoutes = frameIn

// 重新组织后导出
export default [
    ...frameIn,
    ...frameOut,
    ...errorPage
]
