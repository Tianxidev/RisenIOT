<script setup lang="ts">

import TableView from "/@/components/table/TableView.vue";
import { ElMessage } from "element-plus";
import { DeviceList } from "/@/api/system/device";

let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: '设备ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '设备名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '设备分组', prop: 'group', minWidth: 100, align: 'center'},
      {label: '设备SN', prop: 'sn', minWidth: 100, align: 'center'},
      {label: '设备密码', prop: 'pwd', minWidth: 100, align: 'center'},
      {label: '在线状态', prop: 'monitor', minWidth: 100, align: 'center'},
      {label: '状态ID', prop: 'statusId', minWidth: 100, align: 'center'},
      {label: '设备LOGO', prop: 'logo', minWidth: 100, align: 'center'},
      {label: '设备SN', prop: 'timeOut', minWidth: 100, align: 'center'},
      {label: '创建时间', prop: 'createdAt', minWidth: 100, align: 'center'},
      {label: '更新时间', prop: 'updatedAt', minWidth: 100, align: 'center'},
      {label: '最后一次数据上报时间', prop: 'lastDataUpdateTime', minWidth: 100, align: 'center'},
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


const onEdit = (scope: any) => {
  console.log(scope);
};

</script>

<template>
  <div class="device-model-container">
    <el-card shadow="hover">
      <div class="device-model-search mb15">
        <el-button type="info">添加设备</el-button>
      </div>
      <div class="device-model-search mb15">
        <TableView style="height: 100%" :dataFn="tableDataFn" :config="config">
          <template #operate="row">
            <el-button type="text" size="mini" @click="onEdit(row.scope)">编辑</el-button>
            <el-button type="text" size="mini">删除</el-button>
          </template>
        </TableView>
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">

</style>