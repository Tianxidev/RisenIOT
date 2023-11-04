<template>
    <t-form ref="form" :class="['item-container', `login-${type}`]" :data="formData" :rules="FORM_RULES" label-width="0"
        @submit="onSubmit">
        <template v-if="type == 'password'">
            <t-form-item name="account">
                <t-input v-model="formData.account" size="large" placeholder="请输入账号">
                    <template #prefix-icon>
                        <t-icon name="user" />
                    </template>
                </t-input>
            </t-form-item>

            <t-form-item name="password">
                <t-input v-model="formData.password" size="large" :type="showPsw ? 'text' : 'password'" clearable
                    placeholder="请输入登录密码">
                    <template #prefix-icon>
                        <t-icon name="lock-on" />
                    </template>
                    <template #suffix-icon>
                        <t-icon :name="showPsw ? 'browse' : 'browse-off'" @click="showPsw = !showPsw" />
                    </template>
                </t-input>
            </t-form-item>

            <div class="check-container remember-pwd">
                <t-checkbox>记住账号</t-checkbox>
                <span class="tip">忘记账号？</span>
            </div>
        </template>
        <!-- 手机号登陆 -->
        <template v-else>
            <t-form-item class="verification-code" name="verifyCode">
                <t-input v-model="formData.phone" disabled size="large" placeholder="请输入手机号(登录入口暂时关闭)" />
            </t-form-item>
            <t-form-item class="verification-code" name="verifyCode">
                <t-input v-model="formData.verifyCode" size="large" placeholder="请输入验证码" />
            </t-form-item>
        </template>
        <t-form-item v-if="type !== 'qrcode'" class="btn-container">
            <t-button :disabled="loadingLogin" :loading="loadingLogin" block size="large" type="submit"> 登录 </t-button>
        </t-form-item>
        <div class="switch-container">
            <span v-if="type !== 'password'" class="tip" @click="switchType('password')">使用账号密码登录</span>
            <span v-if="type !== 'phone'" class="tip" @click="switchType('phone')">使用手机号登录</span>
        </div>
    </t-form>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';


const INITIAL_DATA = {
    phone: '',
    account: '',
    password: '',
    verifyCode: '',
    checked: false,
};

const FORM_RULES = {
    phone: [{ required: true, message: '手机号必填', type: 'error' }],
    account: [{ required: true, message: '账号必填', type: 'error' }],
    password: [{ required: true, message: '密码必填', type: 'error' }],
    verifyCode: [{ required: true, message: '验证码必填', type: 'error' }],
};

const type = ref('password');

const formData = ref({ ...INITIAL_DATA });
const showPsw = ref(false);

const switchType = (val: string) => {
    type.value = val;
};

const router = useRouter();
const loadingLogin = ref(false);
const onSubmit = async ({ validateResult }) => {
    if (loadingLogin.value) {
        return;
    }
    loadingLogin.value = true;
    if (validateResult === true) {
        console.log(formData.value);

        loadingLogin.value = false;
    }
};
</script>
  
<style lang="less" scoped>
@import url('../index.less');
</style>
  