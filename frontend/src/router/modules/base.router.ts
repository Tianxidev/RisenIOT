import BaseLayout from "@/layout/BaseLayout.vue";

export default [
  {
    path: "/",
    redirect: '/home',
    component: BaseLayout,
    children: [
      {
        path: '/home',
        name: "home",
        meta: {
          title: '概览',
          icon: 'mdi-alpha-s',
          keepAlive: false,
          visible: true,
        },
        component: () => import('@/views/home/HomeView.vue'),
        children: [],
      },
    ],
  }
];
