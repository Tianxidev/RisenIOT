export default [
  {
    path: "/",
    name: "bigdata",
    meta: {
      title: "可视化大屏",
      icon: "mdi-alpha-s",
      keepAlive: false,
      visible: true,
      sort: -1
    },
    component: () => import("@/views/home/HomeView.vue")
  }
];
