import {createRouter, createWebHashHistory} from 'vue-router'
import routes from './routes'

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition;
        } else {
            return {x: 0, y: 0};
        }
    },
    routes
})

function hasPermission(perms, route) {
    if (perms.includes('*')) return true
    if (route.meta && route.meta.perms) {
        return perms.some(perm => route.meta.perms.includes(perm))
    } else {
        return true
    }
}

router.beforeEach(async (to, from, next) => {

    // 检查是否需要登录
    if (to.meta.requiresAuth) {

        const isLoggedIn = true

        // 检查用户是否已经登录
        if (isLoggedIn) {

            // 已经登录 访问登录页则直接跳到主页
            if (to.name === 'login') {
                return next('/');
            }

            // 已经登录 访问非登录页则需要验证权限
            const permissionList = true;

            // 如果还未找到缓存的权限数组则返回登录页
            if (!permissionList) {
                // 重定向到登录页
                return next({name: 'login'});
            }

            // 如果找到了缓存的权限数组则验证权限
            const hasRoles = true;
            if (hasRoles) {
                return next();
            }

            // 无权限则重定向到401
            return next({name: '401'});
        }

        // 未登录
        // 重定向到登录页
        return next({name: 'login'});
    }

    // 无需登录的页面 直接访问即可
    next();
});

router.afterEach((to, from) => {
    window.document.title = to.meta.title + " - RisenIOT 物联网云平台";
})

router.onError((error) => {
    // 在这里处理路由错误
    console.error('路由错误:', error);
});

export default router
