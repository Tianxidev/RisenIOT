<script setup lang="ts">
import {getCurrentInstance, reactive, ref} from "vue";
import {DeviceInfoEdit} from "/@/api/system/device";
import {ElMessage} from "element-plus";
import {VFormRender} from "vform3-builds";
import formEdit from "./formEdit.json"

const {proxy} = <any>getCurrentInstance();
const formJson = reactive(formEdit)
const formData = reactive({
  id: null,
  name: null,
  mark: null,
  timeOut: null,
})
const vFormRef = ref<VFormRender | null>(null);
const optionData = reactive({})

const state = reactive({
  isShowDialog: false,
})

const closeDialog = () => {
  state.isShowDialog = false;
};

const openDialog = (row?: any) => {
  if (row) {
    formData.id = row.id;
    formData.name = row.name;
    formData.mark = row.mark;
    formData.timeOut = row.timeOut;
  }
  state.isShowDialog = true;
}

const submitForm = () => {
  vFormRef.value?.getFormData().then((data: any) => {
    DeviceInfoEdit(data).then(res => {
      if (res.code != 0) {
        ElMessage.error(res.msg)
      }
      proxy.mittBus.emit('renderTable', {});
      vFormRef.value?.resetForm();
      ElMessage.success(res.msg);
      closeDialog();
    });
  }).catch((e: any) => {
    ElMessage.error(e)
  })
}

const handleClose = () => {
  Object.assign(formData, {})
  proxy.mittBus.emit('RefreshPage', true);
}

defineExpose({
  name: "EditCp",
  openDialog,
  closeDialog,
})
</script>

<template>
  <div class="device-edit-dialog-container">
    <el-dialog title="编辑产品" v-model="state.isShowDialog" @close="handleClose">
      <v-form-render :form-json="formJson" :form-data="formData" :option-data="optionData" ref="vFormRef"/>
      <el-button type="primary" class="submit" @click="submitForm">提交</el-button>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>