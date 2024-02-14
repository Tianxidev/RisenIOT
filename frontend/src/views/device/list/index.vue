<script setup lang="ts">

import TableView from "/@/components/table/TableView.vue";
import { ElMessage } from "element-plus";
import {DeviceKindList, DeviceList} from "/@/api/system/device";
import {reactive, ref} from "vue";
import AddCp from "./component/AddCp.vue";
import { ItemLabel } from "./component/interface"
import {selectDictLabel} from "/@/utils/dict";
const AddCpRef = ref<InstanceType<typeof AddCp>>()

const state = reactive({
  kind: [] as ItemLabel[],
})

let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: '设备ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '设备名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '产品类型', prop: 'kind', slot: true, minWidth: 100, align: 'center'},
      {label: '设备分组', prop: 'group', minWidth: 100, align: 'center'},
      {label: '设备SN', prop: 'sn', minWidth: 100, align: 'center'},
      {label: '在线状态', prop: 'monitor', slot: true, minWidth: 100, align: 'center'},
      {label: '创建时间', prop: 'createdAt', minWidth: 100, align: 'center'},
      {label: '更新时间', prop: 'updatedAt', minWidth: 100, align: 'center'},
      {label: '最后一次数据上报', prop: 'lastDataUpdateTime', minWidth: 100, align: 'center'},
      {label: '操作', prop: 'operate', slot: true, minWidth: 100, align: 'center'},
    ],
  },
};

const tableDataFn = (params: any) => {
  return new Promise((resolve) => {
    DeviceList(params).then((res) => {
      if (res.code !== 0) {
        ElMessage.error(res.msg);
        resolve({
          tableDataValue: [],
          total: 0,
        });
        return;
      }
      resolve({
        tableDataValue: res.data.list,
        total: res.data.total,
      });
    }).catch((err) => {
      ElMessage.error(err.message);
      resolve({
        tableDataValue: [],
        total: 0,
      });
    });
  });
};

DeviceKindList().then(res => {
  const data = res.data;
  data?.list.forEach((item: any) => {
    state.kind.push({
      label: item.name,
      value: item.id
    })
  })
})



const onEdit = (scope: any) => {
  // console.log(scope);
};

// 添加设备
const onAdd = (row?: any) => {
  AddCpRef.value?.openDialog(row);
};

</script>

<template>
  <div class="device-model-container">
    <el-card shadow="hover">
      <div class="device-model-search mb15">
        <el-button type="info" @click="onAdd">添加设备</el-button>
        <AddCp ref="AddCpRef"></AddCp>
      </div>
      <div class="device-model-search mb15">
        <TableView style="height: 100%" :dataFn="tableDataFn" :config="config">
          <template #kind="row">
            <el-tag type="info">{{ selectDictLabel(state.kind, row.scope.kind) }}</el-tag>
          </template>
          <template #monitor="row">
            <el-tag type="success" v-if="row.scope['monitor'] === 1">在线</el-tag>
            <el-tag type="info" v-else>离线</el-tag>
          </template>
          <template #operate="row">
            <el-button link type="primary" @click="onEdit(row.scope)">查看</el-button>
            <el-button link type="primary" @click="onEdit(row.scope)">编辑</el-button>
            <el-button link type="primary">删除</el-button>
          </template>
        </TableView>
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">

</style>