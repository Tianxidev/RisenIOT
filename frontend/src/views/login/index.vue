<template>
    <div class="login-wrapper">
        <div class="login-container">
            <div class="title-container">
                <h1 class="title margin-no">登录</h1>
                <h1 class="title">RisenIOT 物联网云平台</h1>
                <div class="sub-title">
                    <p class="tip">{{ type == 'register' ? '已有账号?' : '没有账号吗?' }}</p>
                    <p class="tip" style="color: #451fd3;" @click="switchType(type == 'register' ? 'login' : 'register')">
                        {{ type == 'register' ? '登录' : '注册新账号' }}
                    </p>
                </div>
            </div>
            <Login v-if="type === 'login'" />
            <div v-if="type === 'register'">
                很抱歉，暂未开放注册功能，如有需要请联系管理员
            </div>
            <tdesign-setting />
        </div>
        <footer class="copyright">
            <p>&copy;&nbsp;{{ new Date().getFullYear() }}
                <a href="https://riseniot.tianxidev.wiki">RisenIOT</a>&nbsp;
                <span class="ml-1">&nbsp;暂无备案号</span>
            </p>
        </footer>
    </div>
</template>
<script setup>
import Login from './components/Login.vue';
import { ref } from 'vue';
const type = ref('login');
const switchType = (val) => {
    type.value = val;
};
</script>
<style lang="less" scoped>
@import url('./index.less');
</style>
