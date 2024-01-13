import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [];
const modules = import.meta.glob("./modules/*.router.ts", { eager: true, import: "default" });
Object.keys(modules).forEach((key) => {
  const mod = modules[key] || {};
  const modList = Array.isArray(mod) ? mod : [mod];
  routes.push(...modList);
});


// 路由排序
routes.sort((a, b) => {
  const sortA = a.meta?.sort as number || 0;
  const sortB = b.meta?.sort as number || 0;
  if (sortA === sortB) {
    return Number.MAX_SAFE_INTEGER;
  }
  return sortA - sortB;
});

const router = createRouter({
  history: createWebHistory(), routes
});


export default router;