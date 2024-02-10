<template>
  <div class="login-container">
    <div class="login-content-out">
      <div class="login-content">
        <div class="login-content-main">
          <div class="login-icon-group">
            <div class="login-icon-group-title">
              <div class="login-icon-group-title-text">{{ getThemeConfig.globalViceTitle }}</div>
            </div>
          </div>
          <el-tabs v-model="tabsActiveName" style="height: 350px;">
            <el-tab-pane label="账号密码登录" name="account">
              <Account />
            </el-tab-pane>
            <el-tab-pane label="手机号登录" name="mobile">
              <Mobile />
            </el-tab-pane>
            <el-tab-pane label="二维码登录" name="qrcode">
              <Scan />
            </el-tab-pane>
            <div class="font12 mt30 login-animation4 login-msg">* 温馨提示：建议使用谷歌、Microsoft Edge，版本 79.0.1072.62
              及以上浏览器，360浏览器请使用极速模式</div>
          </el-tabs>
        </div>
      </div>
    </div>
    <div class="login-footer">
      <div class="login-footer-content mt15">
        <div class="login-footer-content-warp">
          <div>Copyright &copy; 2024 riseniot.tianxidev.wiki All Rights Reserved.</div>
          <div class="mt5">RisenIOT · 物联网云平台</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { toRefs, reactive, computed, defineComponent, onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import loginIconTwo from '/@/assets/login-icon-two.svg';
import { NextLoading } from '/@/utils/loading';
import Account from '/@/views/login/component/account.vue';
import Mobile from '/@/views/login/component/mobile.vue';
import Scan from '/@/views/login/component/scan.vue';

// 定义接口来定义对象的类型
interface LoginState {
  tabsActiveName: string;
}

export default defineComponent({
  name: 'loginIndex',
  components: { Account, Mobile, Scan },
  setup() {
    const storesThemeConfig = useThemeConfig();
    const { themeConfig } = storeToRefs(storesThemeConfig);
    const state = reactive<LoginState>({
      tabsActiveName: 'account',
    });
    // 获取布局配置信息
    const getThemeConfig = computed(() => {
      return themeConfig.value;
    });
    // 页面加载时
    onMounted(() => {
      NextLoading.done();
    });
    return {
      loginIconTwo,
      getThemeConfig,
      ...toRefs(state),
    };
  },
});
</script>


<style scoped lang="scss">
.login-container {
  width: 100%;
  height: 100%;
  position: relative;
  color: #1c2518;
  background-color: #f7f7f7;

  .login-icon-group {
    width: 100%;
    height: 100%;
    position: relative;

    .login-icon-group-title {
      display: flex;
      align-items: center;

      img {
        width: 80px;
        height: 70px;
      }
    }

    &-icon {
      width: 60%;
      height: 70%;
      position: absolute;
      left: 0;
      bottom: 0;
    }
  }

  .login-content-out {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .login-content {
    width: 600px;
    padding: 10px;
    margin: auto;
    background-color: #ffffff;
    border: 2px solid #003422;
    border-radius: 5px;
    overflow: hidden;
    z-index: 1;
    position: relative;

    .login-content-main {
      margin: 0 auto;
      width: calc(100% - 30px);

      .login-content-title {
        color: var(--el-text-color-primary);
        font-weight: 500;
        font-size: 22px;
        text-align: center;
        letter-spacing: 4px;
        margin: 15px 0 30px;
        white-space: nowrap;
        z-index: 5;
        position: relative;
        transition: all 0.3s ease;
      }
    }

  }

  .login-icon-group-title-text {
    font-size: 50px;
    font-weight: 500;
    width: 100%;
    display: flex;
    justify-content: start;
    align-items: center;
    -webkit-text-stroke: 1.3px #000000;
    color: transparent;

  }

  .login-footer {
    position: absolute;
    bottom: 5px;
    width: 100%;

    &-content {
      width: 100%;
      display: flex;

      &-warp {
        margin: auto;
        color: #000000;
        text-align: center;
        animation: error-num 1s ease-in-out;
      }
    }
  }
}</style>
