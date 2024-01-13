<script setup lang="ts">
  import { useAppMainStore } from "@/stores/index";
  import { computed, reactive } from "vue";
  import { useRouter } from "vue-router";
  import { Config } from "@/config";

  const router = useRouter();
  const navState = reactive({
    menuVisible: true,
    rail: false,
    isMini: false,
    routes: router.options.routes
  });

  // navigationRail 用于切换导航栏的显示状态
  const navigationRail = (e: boolean) => {
    if (!navState.rail) return;
    navState.isMini = e;
  };

  const iconSmartSLClass = computed(() => (useAppMainStore().theme === "light" ? "icon-smartsl1" : "icon-smartsl"));
</script>

<template>
  <v-layout :class="{ isMini: !navState.menuVisible }">
    <v-navigation-drawer
      class="my-4 layout_navigation radius"
      :rail="navState.rail"
      expand-on-hover
      rail-width="77"
      @update:rail="navigationRail"
      v-model="navState.menuVisible"
      style="position: fixed"
    >
      <v-list class="py-4 mx-2 logo" nav>
        <v-list-item class="mx-1">
          <v-list-item-title class="title">智慧路灯云控平台</v-list-item-title>
          <v-list-item-subtitle>{{ Config().RecordNumber }}</v-list-item-subtitle>
        </v-list-item>
      </v-list>
      <v-divider></v-divider>

      <v-list nav class="mx-2">
        <v-list-subheader>Dashboard</v-list-subheader>

        <!--路由导航-->
        <template v-for="(item, key) in navState.routes" :key="key">
          <v-list-item
            v-if="item.meta?.visible && !item.children"
            :prepend-icon="item.meta?.icon as string"
            :title="item.meta?.title as string"
            :to="{ name: item.name }"
            class="mx-1"
            active-class="nav_active"
          ></v-list-item>

          <v-list-group v-if="item.meta?.visible && item.children && item.children.length > 0" class="mx-1">
            <template v-slot:activator="{ props }">
              <v-list-item v-bind="props" :prepend-icon="item.meta.icon as string" :title="item.meta.title as string"/>
            </template>
            <template v-for="(row, i) in item.children">
              <v-list-item
                v-if="row.meta?.visible as boolean"
                :title="row.meta?.title as string"
                :prepend-icon="navState.isMini ? (row.meta?.icon as string) : ''"
                :key="i"
                :to="{ name: row.name }"
              />
            </template>
          </v-list-group>
          <v-list-subheader v-if="item.name === 'bigdata'">MENU</v-list-subheader>
        </template>
      </v-list>
    </v-navigation-drawer>
    <main class="app_main">
      <header class="header radius">
        <div class="mx-4">
          <v-btn variant="text" icon="mdi-menu" @click="navState.menuVisible = !navState.menuVisible"></v-btn>
          <v-btn variant="text" icon="mdi-home-circle-outline" @click="router.replace('/home')"></v-btn>
        </div>
        <v-spacer></v-spacer>
        <div class="tool_btns">
          <v-btn
            @click="useAppMainStore().onTheme"
            variant="text"
            :icon="useAppMainStore().theme === 'light' ? 'mdi-weather-sunny' : 'mdi-weather-night'"
          />
          <v-btn variant="text" icon="mdi-bell-outline">
            <v-badge content="2" color="error">
              <v-icon size="small"></v-icon>
            </v-badge>
          </v-btn>
          <v-btn variant="text" append-icon="mdi-chevron-down" class="mr-2">
            <v-avatar size="x-small" class="avatar mr-2">
              <v-img :src="logo"></v-img>
            </v-avatar>
            <span v-if="true">管理员</span>
            <v-menu activator="parent">
              <v-list nav class="h_a_menu">
                <v-list-item title="退出登录" prepend-icon="mdi-login" to="/login"/>
              </v-list>
            </v-menu>
          </v-btn>
        </div>
        <div style="position: fixed; right: 20px; bottom: 100px; z-index: 99999">
          <v-btn icon="mdi-cog"/>
        </div>
      </header>
      <div class="router">
        <RouterView/>
      </div>
    </main>
  </v-layout>
</template>

<style lang="scss" scoped>
  .v-application {
    // color: #344767!important;
    background: var(--theme-background) !important;

    .layout_navigation {
      border-radius: 12px;
      border: var(--header-bg) 1px solid !important;
      margin: 16px 0 16px 16px !important;
      overflow: hidden;
      height: calc(100vh - 2rem) !important;
      background: var(--header-bg);

      .v-list-item__prepend > .v-icon,
      .v-list-item__append > .v-icon {
        opacity: 1 !important;
      }

      .v-list-item__prepend {
        width: 40px;
      }

      .v-list-item-title {
        font-size: 14px;
        font-weight: 500;
        // color: #ffffff;
        opacity: 1;

        .link {
          color: inherit;
        }
      }

      .v-list-group__items {
        .v-list-item {
          padding-inline-start: 47px !important;
        }

        .v-list-item--active {
          //background-color: rgb(var(--v-theme-primary));
        }
      }

      .v-icon {
        font-size: 24px !important;
      }

      .logo {
        .v-avatar {
          border-radius: 0;
        }

        .v-list-item__content {
          .title {
            margin-top: 4px;
            font-size: 16px;
            //font-weight: 700;
            font-weight: bold;
          }

          .v-list-item-subtitle {
            margin-top: 4px;
          }
        }

        .v-list-item__prepend {
          width: auto;
          padding-left: 0;
        }
      }
    }

    .app_main {
      flex: 1;
      margin-left: 288px;
      transition: all 0.2s;
      margin-top: 92px;
      margin-right: 10px;
      height: var(--content-height);

      .header {
        // background: ;
        background: var(--header-bg);
        // filter: blur(25px);
        backdrop-filter: saturate(200%) blur(30px);
        // box-shadow: inset 0 0 1px 1px hsla(0,0%,99.6%,.9),0 20px 27px 0 rgba(0,0,0,.05)!important;
        height: 76px;
        transition: all 0.2s;
        display: flex;
        align-items: center;
        border-radius: 10px;
        position: fixed;
        top: 16px;
        left: 288px;
        width: calc(100vw - 304px);
        z-index: 2;

        .header_title {
          .v-breadcrumbs {
            padding: 0 !important;

            .imglogo {
              height: 22px;
            }

            .v-breadcrumbs-item--disabled {
              opacity: 0.6;
            }

            .v-breadcrumbs-item {
              font-weight: 500;
            }

            .v-breadcrumbs-divider {
              // color: rgb(204, 204, 204);
              opacity: 0.6;
              padding: 0;
            }

            .v-breadcrumbs__prepend {
              opacity: 0.6;

              .v-icon {
                margin-left: 5px;
                // width: 5px;
              }

              // width: 25px;
            }

            .link {
              color: inherit;
            }
          }

          .page_title {
            margin-top: 8px;
            font-size: 16px;
            font-weight: 700;
          }
        }

        .gamepad {
          cursor: pointer;
          align-self: flex-start;
          position: relative;
          z-index: 2;
        }

        .search_ip {
          position: relative;
          z-index: 2;

          .v-field--variant-solo {
            box-shadow: none;
          }
        }

        .tool_btns {
          position: relative;
          z-index: 2;
        }
      }

      .router {
        padding-top: 16px;
      }
    }
  }

  .isMini {
    .layout_navigation {
      .v-list-item__prepend {
        padding-left: 0;
      }

      .v-list-group__items .v-list-item {
        padding-inline-start: 8px !important;
      }
    }

    .app_main {
      margin: 92px 5px 0 10px;
      height: var(--content-height);

      .header {
        left: 10px;
        width: calc(100vw - 20px);
      }
    }
  }

  .h_a_menu {
    .v-list-item--density-default.v-list-item--one-line {
      min-height: 36px !important;
    }

    .v-list-item__prepend > .v-icon {
      margin-inline-end: 16px;
    }
  }

  /* 图标动画 */
  @keyframes fill {
    from {
      fill: transparent;
    }
    to {
      fill: #ccc;
      stroke-dashoffset: 0;
    }
  }

  @keyframes fill1 {
    from {
      fill: transparent;
    }
    to {
      fill: #14b369;
      stroke-dashoffset: 0;
    }
  }

  @keyframes fillHover {
    from {
      fill: transparent;
    }
    to {
      fill: red;
      stroke-dashoffset: 0;
    }
  }

  @mixin iconStyle {
    width: 100%;
    stroke-dasharray: 1000;
    stroke-dashoffset: 1000;
    &:hover {
      animation: fillHover 1s ease both;
    }
  }

  .icon-smartsl {
    @include iconStyle;
    animation: fill 2s ease-in-out both;
  }

  .icon-smartsl1 {
    @include iconStyle;
    animation: fill1 2s ease-in-out both;
  }
</style>
