<script lang="ts" setup>

import { onMounted, reactive } from "vue";
import { QueryPushDeerConfig, SavePushDeerConfig, TextMsgToPushDeer } from "/@/api/system/push";
import { ElMessage } from "element-plus";

const form = reactive({
  pushDeerKey: '',
});

// 提交配置
const submitForm = () => {
  SavePushDeerConfig(form.pushDeerKey).then(res => {
    if (res.code === 0) {
      ElMessage.success('保存成功');
    } else {
      ElMessage.error('保存失败');
    }
  })
};

// 测试推送
const testSubmit = () => {
  TextMsgToPushDeer(form.pushDeerKey, '测试推送').then(res => {
    if (res.code === 0) {
      ElMessage.success('推送成功');
    } else {
      ElMessage.error('推送失败');
    }
  });
};

onMounted(() => {
  QueryPushDeerConfig().then(res => {
    if (res.code === 0) {
      form.pushDeerKey = res.data.key;
    }
  })
})

</script>
<template>
  <div class="config-container">
    <el-card shadow="hover" header="PushDeer 推送">
      <img src="https://www.pushdeer.com/images/2022-02-03-17-55-30.png" alt="PushDeer">
      <p style="margin: 20px 0;">
        苹果手机（iOS 14+）用系统摄像头扫描上边的码即可拉起轻应用。亦可在苹果商店搜索「PushDeer」安装。
      </p>
      <el-form :model="form" label-width="120px">
        <el-form-item label="PushDeer Key">
          <el-input v-model="form.pushDeerKey"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="testSubmit">测试推送</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">保存配置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>
<style scoped lang="scss">
.email-content {
  .el-input__inner {
    height: 200px;
  }
}
</style>