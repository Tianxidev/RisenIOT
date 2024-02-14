<template>
  <div class="layout-navbars-breadcrumb-user pr15" :style="{ flex: layoutUserFlexNum }">
    <div class="layout-navbars-breadcrumb-user-icon" @click="onSearchClick">
      <el-icon :title="$t('message.user.title2')">
        <ele-Search/>
      </el-icon>
    </div>
    <div class="layout-navbars-breadcrumb-user-icon" @click="removeCacheClick">
      <i class="fa-trash fa" title="清除缓存"></i>
    </div>
    <div class="layout-navbars-breadcrumb-user-icon">
      <el-popover placement="bottom" trigger="click" transition="el-zoom-in-top" :width="300" :persistent="false">
        <template #reference>
          <el-badge :is-dot="true">
            <el-icon :title="$t('message.user.title4')">
              <ele-Bell/>
            </el-icon>
          </el-badge>
        </template>
        <template #default>
          <UserNews/>
        </template>
      </el-popover>
    </div>
    <div class="layout-navbars-breadcrumb-user-icon mr10" @click="onScreenfullClick">
      <i
          class="iconfont"
          :title="isScreenfull ? '退出全屏' : '全屏'"
          :class="!isScreenfull ? 'icon-fullscreen' : 'icon-tuichuquanping'"
      ></i>
    </div>
    <el-dropdown trigger="click" :show-timeout="70" :hide-timeout="50" @command="onHandleCommandClick">
			<span class="layout-navbars-breadcrumb-user-link">
				「&nbsp;{{ userInfos.userName === '' ? 'common' : userInfos.userName }}&nbsp;」
			</span>d
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item command="/personal">个人中心</el-dropdown-item>
          <el-dropdown-item divided command="logOut">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
    <Search ref="searchRef"/>
  </div>
</template>

<script lang="ts">
import { ref, getCurrentInstance, computed, reactive, toRefs, onMounted, defineComponent } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import screenfull from 'screenfull';
import { useI18n } from 'vue-i18n';
import { storeToRefs } from 'pinia';
import { useUserInfo } from '/@/stores/userInfo';
import { useThemeConfig } from '/@/stores/themeConfig';
import { Session, Local } from '/@/utils/storage';
import UserNews from '/@/layout/navBars/breadcrumb/userNews.vue';
import Search from '/@/layout/navBars/breadcrumb/search.vue';
import { logout } from "/@/api/login";
import { removeCache } from "/@/api/system/cache";

export default defineComponent({
  name: 'layoutBreadcrumbUser',
  components: {UserNews, Search},
  setup() {
    const {t} = useI18n();
    const {proxy} = <any> getCurrentInstance();
    const router = useRouter();
    const stores = useUserInfo();
    const storesThemeConfig = useThemeConfig();
    const {userInfos} = storeToRefs(stores);
    const {themeConfig} = storeToRefs(storesThemeConfig);
    const searchRef = ref();
    const state = reactive({
      isScreenfull: false,
      disabledI18n: 'zh-cn',
      disabledSize: 'large',
    });
    // 设置分割样式
    const layoutUserFlexNum = computed(() => {
      let num: string | number = '';
      const {layout, isClassicSplitMenu} = themeConfig.value;
      const layoutArr: string[] = ['defaults', 'columns'];
      if (layoutArr.includes(layout) || (layout === 'classic' && !isClassicSplitMenu)) num = '1';
      else num = '';
      return num;
    });
    // 全屏点击时
    const onScreenFullClick = () => {
      if (!screenfull.isEnabled) {
        ElMessage.warning('暂不不支持全屏');
        return false;
      }
      screenfull.toggle();
      screenfull.on('change', () => {
        if (screenfull.isFullscreen) state.isScreenfull = true;
        else state.isScreenfull = false;
      });
    };
    //清除缓存
    const removeCacheClick = () => {
      //清除浏览器缓存
      Session.remove('userMenu');
      //清除后端缓存
      removeCache().then(() => {
        ElMessage.success('缓存清除成功');
        window.location.reload();
      })
    };
    // 下拉菜单点击时
    const onHandleCommandClick = (path: string) => {
      if (path === 'logOut') {
        ElMessageBox({
          closeOnClickModal: false,
          closeOnPressEscape: false,
          title: t('message.user.logOutTitle'),
          message: t('message.user.logOutMessage'),
          showCancelButton: true,
          confirmButtonText: t('message.user.logOutConfirm'),
          cancelButtonText: t('message.user.logOutCancel'),
          buttonSize: 'default',
          beforeClose: (action, instance, done) => {
            if (action === 'confirm') {
              //后端退出
              logout().then(() => {
                instance.confirmButtonLoading = true;
                instance.confirmButtonText = t('message.user.logOutExit');
                setTimeout(() => {
                  done();
                  setTimeout(() => {
                    instance.confirmButtonLoading = false;
                  }, 300);
                }, 500);
              })
            } else {
              done();
            }
          },
        }).then(async () => {
          // 清除缓存/token等
          Session.clear();
          // 使用 reload 时，不需要调用 resetRoute() 重置路由
          window.location.reload();
        }).catch(() => {
        });
      } else {
        router.push(path);
      }
    };
    // 菜单搜索点击
    const onSearchClick = () => {
      searchRef.value.openSearch();
    };

    // 设置 element plus 组件的国际化
    const setI18nConfig = (locale: string) => {
      proxy.mittBus.emit('getI18nConfig', proxy.i18n.global.messages.value[locale]);
    };
    // 初始化言语国际化
    const initI18n = () => {
      switch (Local.get('themeConfig').globalI18n) {
        case 'zh-cn':
          state.disabledI18n = 'zh-cn';
          setI18nConfig('zh-cn');
          break;
        case 'en':
          state.disabledI18n = 'en';
          setI18nConfig('en');
          break;
        case 'zh-tw':
          state.disabledI18n = 'zh-tw';
          setI18nConfig('zh-tw');
          break;
      }
    };
    // 初始化全局组件大小
    const initComponentSize = () => {
      switch (Local.get('themeConfig').globalComponentSize) {
        case 'large':
          state.disabledSize = 'large';
          break;
        case 'default':
          state.disabledSize = 'default';
          break;
        case 'small':
          state.disabledSize = 'small';
          break;
      }
    };
    // 页面加载时
    onMounted(() => {
      if (Local.get('themeConfig')) {
        initI18n();
        initComponentSize();
      }
    });
    return {
      userInfos,
      onHandleCommandClick,
      onScreenfullClick: onScreenFullClick,
      onSearchClick,
      removeCacheClick,
      searchRef,
      layoutUserFlexNum,
      ...toRefs(state),
    };
  },
});
</script>

<style scoped lang="scss">
.layout-navbars-breadcrumb-user {
  display: flex;
  align-items: center;
  justify-content: flex-end;

  &-link {
    height: 100%;
    display: flex;
    align-items: center;
    white-space: nowrap;

    &-photo {
      width: 25px;
      height: 25px;
      border-radius: 100%;
    }
  }

  &-icon {
    padding: 0 10px;
    cursor: pointer;
    color: var(--next-bg-topBarColor);
    height: 50px;
    line-height: 50px;
    display: flex;
    align-items: center;

    &:hover {
      background: var(--next-color-user-hover);

      i {
        display: inline-block;
        animation: logoAnimation 0.3s ease-in-out;
      }
    }
  }

  :deep(.el-dropdown) {
    color: var(--next-bg-topBarColor);
  }

  :deep(.el-badge) {
    height: 40px;
    line-height: 40px;
    display: flex;
    align-items: center;
  }

  :deep(.el-badge__content.is-fixed) {
    top: 12px;
  }
}

.layout-navbars-breadcrumb-user-link {
  margin-right: 10px;
}
</style>
