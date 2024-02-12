<template>
  <el-form ref="loginForm" size="large" class="login-content-form" :model="ruleForm"  :rules="formRules">
    <el-form-item class="login-animation1" prop="username">
      <el-input
          type="text"
          placeholder="请输入用户名"
          v-model="ruleForm.username"
          clearable autocomplete="off">
        <template #prefix>
          <el-icon class="el-input__icon"><ele-User /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation2" prop="password">
      <el-input
          :type="isShowPassword ? 'text' : 'password'"
          placeholder="请输入密码"
          v-model="ruleForm.password"
          autocomplete="off"
          @keyup.enter="onSignIn"
      >
        <template #prefix>
          <el-icon class="el-input__icon"><ele-Unlock /></el-icon>
        </template>
        <template #suffix>
          <i
              class="iconfont el-input__icon login-content-password"
              :class="isShowPassword ? 'icon-yincangmima' : 'icon-xianshimima'"
              @click="isShowPassword = !isShowPassword"
          >
          </i>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation4">
      <el-button type="primary" class="login-content-submit" round @click="onSignIn" :loading="loading.signIn">
        <span>登录</span>
      </el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts">
import { computed, defineComponent, getCurrentInstance, onMounted, reactive, ref, toRefs, unref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { useI18n } from 'vue-i18n';
import Cookies from 'js-cookie';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import { initBackEndControlRoutes } from '/@/router/backEnd';
import { Session } from '/@/utils/storage';
import { formatAxis } from '/@/utils/formatTime';
import { NextLoading } from '/@/utils/loading';
import { LoginApi } from "/@/api/login";

export default defineComponent({
	name: 'loginAccount',
	setup() {
		const { t } = useI18n();
    const {proxy} = <any>getCurrentInstance();
		const storesThemeConfig = useThemeConfig();
		const { themeConfig } = storeToRefs(storesThemeConfig);
		const route = useRoute();
		const router = useRouter();
    const loginForm = ref(null)
		const state = reactive({
			isShowPassword: false,
			ruleForm: {
        username: '',
				password: '',
			},
      formRules:{
        username: [
          { required: true, trigger: "blur", message: "用户名不能为空" }
        ],
        password: [
          { required: true, trigger: "blur", message: "密码不能为空" }
        ],
      },
			loading: {
				signIn: false,
			},
      captchaSrc:'',
		});
    onMounted(() => {
    });

		// 时间获取
		const currentTime = computed(() => {
			return formatAxis(new Date());
		});
		// 登录
		const onSignIn = async () => {
      if(state.loading.signIn){
        return
      }
      const formWrap = unref(loginForm) as any;
      if (!formWrap) return;
      formWrap.validate((valid: boolean) => {
        if(valid){
          state.loading.signIn = true;

          // 请求登录
          LoginApi(state.ruleForm).then(async (res:any)=>{

            // 检查响应状态码是否正确
            if(res.code != 0){
              state.loading.signIn = false;
              return
            }

            // 获取用户信息
            const userInfo = res.data.userInfo
            // 存储 token 到浏览器缓存
            Session.set('token', res.data.token);
            // 存储用户信息到浏览器缓存
            Session.set('userInfo', userInfo);
            // 设置用户菜单
            Session.set('userMenu',res.data.menuList)
            // 设置按钮权限
            Session.set('permissions',res.data.permissions)
            // 模拟数据，对接接口时，记得删除多余代码及对应依赖的引入。用于 `/src/stores/userInfo.ts` 中不同用户登录判断（模拟数据）
            Cookies.set('username', state.ruleForm.username);
            // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
            await initBackEndControlRoutes();
            // 执行完 initBackEndControlRoutes，再执行 signInSuccess
            signInSuccess();
          }).catch(()=>{
            state.loading.signIn = false;
          })
        }
      })
		};
		// 登录成功后的跳转
		const signInSuccess = () => {
			// 初始化登录成功时间问候语
			let currentTimeInfo = currentTime.value;
			// 登录成功，跳到转首页
			// 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
			if (route.query?.redirect) {
				router.push({
					path: <string>route.query?.redirect,
					query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
				});
			} else {
				router.push('/');
			}
			// 登录成功提示
			// 关闭 loading
			state.loading.signIn = true;
			const signInText = t('message.signInText');
			ElMessage.success(`${currentTimeInfo}，${signInText}`);
			// 添加 loading，防止第一次进入界面时出现短暂空白
			NextLoading.start();
		};
		return {
			onSignIn,
      loginForm,
			...toRefs(state),
		};
	},
});
</script>


<style scoped lang="scss">
.login-content-form {
  margin-top: 20px;
  @for $i from 1 through 4 {
    .login-animation#{$i} {
      opacity: 0;
      animation-name: error-num;
      animation-duration: 0.5s;
      animation-fill-mode: forwards;
      animation-delay: calc($i/10) + s;
    }
  }
  .login-content-password {
    display: inline-block;
    width: 20px;
    cursor: pointer;
    &:hover {
      color: #909399;
    }
  }
  .login-content-code {
    display: flex;
    align-items: center;
    justify-content: space-around;
    .login-content-code-img {
      width: 100%;
      height: 40px;
      line-height: 40px;
      background-color: #ffffff;
      border: 1px solid rgb(220, 223, 230);
      cursor: pointer;
      transition: all ease 0.2s;
      border-radius: 4px;
      user-select: none;
      &:hover {
        border-color: #c0c4cc;
        transition: all ease 0.2s;
      }
    }
  }
  .login-content-submit {
    width: 100%;
    letter-spacing: 2px;
    font-weight: 300;
    margin-top: 15px;
  }
}
</style>
