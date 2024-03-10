<script setup lang="ts">

import { getCurrentInstance, reactive, unref } from "vue";
import TableView from "/@/components/table/TableView.vue";
import { DeviceCategoryAdd, DeviceCategoryDel, DeviceCategoryList } from "/@/api/system/device";
import { ElMessage } from "element-plus";
import { NButton, NForm, NFormItem, NInput, NSelect } from "naive-ui";

const {proxy} = <any> getCurrentInstance();

const {device_data_type} = proxy.useDict('device_data_type');

// 获取设备数据类型
const getDeviceDataType = (val: string) => {
  return proxy.selectDictLabel(unref(device_data_type), val);
}

interface IKind {
  id: number,
  name: string,
  mark: string,
  timeOut: number,
}

// 设备数据类型
interface IData {
  id: number, // 数据类型ID
  kindId: number, // 产品ID
  name: string, // 数据类型名称
  mark: string, // 数据类型标记
  dataType: number, // 数据类型
  unit: string, // 单位
}

const state = reactive({
  isShowDialog: false,
  direction: 'btt',
  row: {} as IKind,
  addData: {} as IData,
})

let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: 'ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '类型名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '数据类型', prop: 'dataType', slot: true, minWidth: 100, align: 'center'},
      {label: '类型标记', prop: 'mark', minWidth: 100, align: 'center'},
      {label: '单位', prop: 'unit', minWidth: 100, align: 'center'},
      {label: '创建时间', prop: 'createdAt', minWidth: 100, align: 'center'},
      {label: '更新时间', prop: 'updatedAt', minWidth: 100, align: 'center'},
      {label: '操作', prop: 'operate', slot: true, minWidth: 100, align: 'center'},
    ],
  },
};

const tableDataFn = (params: any) => {
  params.kindId = state.row.id;
  return new Promise((resolve) => {
    DeviceCategoryList(params).then(res => {
      if (res.code !== 0) {
        ElMessage.error(res.msg);
        resolve({tableDataValue: [], total: 0});
        return;
      }
      resolve({
        tableDataValue: res.data.list,
        total: res.data.total,
      });
    }).catch(() => {
      resolve({tableDataValue: [], total: 0});
    });
  });
};

const closeDialog = () => {
  state.isShowDialog = false;
};

const openDialog = (row?: any) => {
  state.row = row;
  state.isShowDialog = true;
  proxy.mittBus.emit('renderTable', {});
}

const onHandleDel = (row?: any) => {
  DeviceCategoryDel([row?.id]).then(res => {
    if (res.code !== 0) {
      ElMessage.error(res.msg);
      return;
    }
    ElMessage.success(res.msg);
    proxy.mittBus.emit('renderTable', {});
  });
}

const handleClose = () => {
  // Object.assign(state.row, {})
  state.row = {} as IKind;
  proxy.mittBus.emit('clearTable', true);
}

const onHandleAdd = () => {
  // console.log("添加数据", state.addData)

  // 设置产品类型ID
  state.addData.kindId = state.row.id;

  DeviceCategoryAdd(state.addData).then(res => {
    if (res.code !== 0) {
      ElMessage.error(res.msg);
      return;
    }
    ElMessage.success(res.msg);
    state.addData = {} as IData;
    proxy.mittBus.emit('renderTable', {});
  });
}


defineExpose({
  name: "ManageCp",
  openDialog,
  closeDialog,
})
</script>

<template>
  <div class="device-add-dialog-container">
    <el-drawer
        v-model="state.isShowDialog"
        title="管理产品数据结构"
        size="90%"
        :close-on-click-modal="false"
        :direction="state.direction"
        @close="handleClose">
      <div class="tool-area">
        <n-form
            ref="formRef"
            inline
            :label-width="80"
            :model="state.addData"
            size="large"
        >
          <n-form-item class="form-item" label="类型名称">
            <n-input v-model:value="state.addData.name" placeholder="键入数据类型名称"/>
          </n-form-item>
          <n-form-item class="form-item" label="类型标识">
            <n-input v-model:value="state.addData.mark" placeholder="键入数据类型标识"/>
          </n-form-item>
          <n-form-item class="form-item" label="数据类型">
            <n-select v-model:value="state.addData.dataType" placeholder="选择数据类型" :options="device_data_type"/>
          </n-form-item>
          <n-form-item class="form-item" label="类型单位">
            <n-input v-model:value="state.addData.unit" placeholder="键入数据类型单位"/>
          </n-form-item>
          <n-form-item class="form-item">
            <n-button attr-type="button" @click="onHandleAdd">
              添加
            </n-button>
          </n-form-item>
        </n-form>
      </div>
      <TableView style="height: 100%" :dataFn="tableDataFn" :config="config">
        <template #dataType="row">
          <el-tag type="warning">{{ getDeviceDataType(row.scope['dataType']) }}</el-tag>
        </template>
        <template #public="row">
          <el-tag type="success" v-if="row.scope['public'] === 0">私有</el-tag>
          <el-tag type="warning" v-else>公开</el-tag>
        </template>
        <template #operate="row">
          <el-button link type="primary" @click="onHandleDel(row.scope)">删除</el-button>
        </template>
      </TableView>
    </el-drawer>
  </div>
</template>

<style scoped lang="scss">
.tool-area {
  display: flex;
  padding: 20px;
  border-bottom: #EBEEF5 1px solid;
}

.form-item {
  min-width: 200px;
}
</style>