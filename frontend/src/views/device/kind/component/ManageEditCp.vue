<script setup lang="ts">

import {reactive} from "vue";

const state = reactive({
  dialogFormVisible: false,
  formLabelWidth: '120px',
  form: {
    id: 0,
    name: '',
    dataType: null,
    mark: '',
    unit: '',
  },
})

const openDialog = (row?: any) => {
  state.dialogFormVisible = true;
  if (row) {
    state.form.id = row.id;
    state.form.name = row.name;
    state.form.mark = row.mark;
    state.form.unit = row.unit;
    state.form.dataType = row.dataType;
  }
}

const closeDialog = () => {
  state.dialogFormVisible = false;
}

defineExpose({
  name: "ManageEditCp",
  openDialog,
  closeDialog,
})
</script>

<template>
  <div>
    <el-dialog v-model="state.dialogFormVisible" title="编辑数据类型" width="500">
      <el-form :model="state.form">
        <el-form-item label="ID" :label-width="state.formLabelWidth">
          <el-input v-model="state.form.id" autocomplete="off" disabled/>
        </el-form-item>
        <el-form-item label="类型名称" :label-width="state.formLabelWidth">
          <el-input v-model="state.form.name" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="类型标记" :label-width="state.formLabelWidth">
          <el-input v-model="state.form.mark" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="单位" :label-width="state.formLabelWidth">
          <el-input v-model="state.form.unit" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="数据类型" :label-width="state.formLabelWidth">
          <el-select v-model="state.form.dataType" placeholder="请选择数据类型">
            <el-option label="整形" value="1"/>
            <el-option label="浮点型" value="2"/>
            <el-option label="字符串" value="3"/>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="state.dialogFormVisible = false">取消</el-button>
          <el-button type="primary" @click="state.dialogFormVisible = false">
            保存
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>