<script setup lang="ts">

import TableView from "/@/components/table/TableView.vue";
import { ElMessage } from "element-plus";
import { DeviceKindList } from "/@/api/system/device";

let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: '类型ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '类型名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '类型标记', prop: 'mark', minWidth: 100, align: 'center'},
      {label: '类型操作超时时间', prop: 'timeOut', minWidth: 100, align: 'center'},
      {label: '创建时间', prop: 'createdAt', minWidth: 100, align: 'center'},
      {label: '更新时间', prop: 'updatedAt', minWidth: 100, align: 'center'},
      {label: '操作', prop: 'operate', slot: true, minWidth: 100, align: 'center'},
    ],
  },
};

const tableDataFn = (params: any) => {
  return new Promise((resolve) => {
    DeviceKindList(params).then((res) => {
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

</script>

<template>
  <div class="device-model-container">
    <el-card shadow="hover">
      <div class="device-model-search mb15">
        设备类型
      </div>
      <div class="device-model-search mb15">
        <TableView style="height: 100%" :dataFn="tableDataFn" :config="config">

        </TableView>
      </div>
    </el-card>
  </div>
</template>

<style scoped lang="scss">

</style>