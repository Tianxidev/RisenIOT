<script setup lang="ts">
import {getCurrentInstance, reactive, ref} from "vue";
import {ItemLabel} from "./interface";
import {DeviceGroupList, DeviceInfoEdit, DeviceKindList} from "/@/api/system/device";
import {ElMessage} from "element-plus";
import {VFormRender} from "vform3-builds";
import formEdit from "./formEdit.json"

const {proxy} = <any>getCurrentInstance();
const formJson = reactive(formEdit)
const formData = reactive({
  id: null,
  name: null,
  group: null,
  sn: null,
  pwd: null,
  kind: null,
})
const vFormRef = ref<VFormRender | null>(null);
const optionData = reactive({
  kind: [] as ItemLabel[],
  group: [] as ItemLabel[],
})

const state = reactive({
  isShowDialog: false,
  row: null,
})

// 加载产品类型字典
DeviceKindList().then(res => {
  const data = res.data;
  data?.list.forEach((item: any) => {
    optionData.kind.push({
      label: item.name,
      value: item.id
    })
  })
})

// 加载设备分组字典
DeviceGroupList().then(res=>{
  const data = res.data;
  data?.list.forEach((item: any) => {
    optionData.group.push({
      label: item.name,
      value: item.id
    })
  })
})

const closeDialog = () => {
  state.isShowDialog = false;
};

const openDialog = (row?: any) => {
  if (row) {
    state.row = row;
    formData.id = row.id;
    formData.name = row.name;
    formData.group = row.group;
    formData.sn = row.sn;
    formData.pwd = row.pwd;
    formData.kind = row.kind;
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


defineExpose({
  name: "EditCp",
  openDialog,
  closeDialog,
})
</script>

<template>
  <div class="device-edit-dialog-container">
    <el-dialog title="编辑设备" v-model="state.isShowDialog">
      <v-form-render :form-json="formJson" :form-data="formData" :option-data="optionData" ref="vFormRef"/>
      <el-button type="primary" class="submit" @click="submitForm">提交</el-button>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>