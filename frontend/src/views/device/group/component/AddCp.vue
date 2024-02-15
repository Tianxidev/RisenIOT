<script setup lang="ts">
import {getCurrentInstance, reactive, ref} from "vue";
import formAdd from './formAdd.json'
import {ElMessage} from "element-plus";
import {type VFormRender} from 'vform3-builds';
import {DeviceGroupAdd} from "/@/api/system/device";

const { proxy } = <any>getCurrentInstance();

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

const submitForm = () => {
  vFormRef.value?.getFormData().then((data: any) => {
    DeviceGroupAdd(data).then(res => {
      if (res.code != 0) {
        ElMessage.error(res.msg)
      }
      proxy.mittBus.emit('renderTable', {});
      vFormRef.value?.resetForm();
      ElMessage.success(res.msg);
      closeDialog();
    })

  }).catch((e: any) => {
    ElMessage.error(e)
  })
}


defineExpose({
  name: "AddCp",
  openDialog,
  closeDialog,
})

</script>

<template>
  <div class="device-add-dialog-container">
    <el-dialog title="添加设备" v-model="state.isShowDialog">
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