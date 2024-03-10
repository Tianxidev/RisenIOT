<script setup lang="ts">
import { getCurrentInstance, reactive, ref } from "vue";
import formAdd from './formAdd.json'
import { ElMessage } from "element-plus";
import { type VFormRender } from 'vform3-builds';
import { DeviceKindAdd } from "/@/api/system/device";

const {proxy} = <any> getCurrentInstance();

const formJson = reactive(formAdd)
const formData = reactive({})
const optionData = reactive({})
const vFormRef = ref<VFormRender | null>(null);

const state = reactive({
  loading: false,
  isShowDialog: false,
  row: null,
});

const closeDialog = () => {
  state.isShowDialog = false;
};

const openDialog = (row?: any) => {
  if (row) {
    state.row = row;
  }
  state.isShowDialog = true;
}

// 提交
const submitForm = () => {
  vFormRef.value?.getFormData().then((data: any) => {
    DeviceKindAdd(data).then(res => {
      if (res.code != 0) {
        ElMessage.error(res.msg)
      }
      proxy.mittBus.emit('initView', true);
      vFormRef.value?.resetForm();
      ElMessage.success(res.msg);
      closeDialog();
    })

  }).catch((e: any) => {
    ElMessage.error(e)
  })
}

const handleClose = () => {
  Object.assign(formData, {})
  proxy.mittBus.emit('initView', true);
}

defineExpose({
  name: "AddCp",
  openDialog,
  closeDialog,
})

</script>

<template>
  <div class="device-add-dialog-container">
    <el-dialog title="添加产品" v-model="state.isShowDialog" @close="handleClose">
      <v-form-render :form-json="formJson" :form-data="formData" :option-data="optionData" ref="vFormRef"/>
      <el-button type="primary" class="submit" @click="submitForm">添加</el-button>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.submit {
  margin-top: 20px;
}
</style>