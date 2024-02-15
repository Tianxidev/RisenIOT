<script setup lang="ts">

import TableView from "/@/components/table/TableView.vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {DeviceKindDel, DeviceKindList} from "/@/api/system/device";
import {getCurrentInstance, ref} from "vue";
import AddCp from "./component/AddCp.vue";
import EditCp from "./component/EditCp.vue";

const {proxy} = <any>getCurrentInstance();
const AddCpRef = ref<InstanceType<typeof AddCp>>()
const EditCpRef = ref<InstanceType<typeof EditCp>>()

let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: '产品ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '产品名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '产品标记', prop: 'mark', minWidth: 100, align: 'center'},
      {label: '可见', prop: 'public', slot: true, minWidth: 100, align: 'center'},
      {label: '心跳超时', prop: 'timeOut', minWidth: 100, align: 'center'},
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

// 添加产品
const onAdd = () => {
  AddCpRef.value?.openDialog();
};

// 编辑
const onEdit = (row?:any) => {
  EditCpRef.value?.openDialog(row);
}

// 删除
const onDel = (row?: any) => {
  ElMessageBox.confirm("确定要删除吗?", '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    DeviceKindDel([row.id]).then((res) => {
      if (res.code != 0) {
        ElMessage.error(res.msg)
      }
      proxy.mittBus.emit('renderTable', {});
      ElMessage.success(res.msg);
    });
  });
}


</script>

<template>
  <div class="device-model-container">
    <el-card shadow="hover">
      <div class="device-model-search mb15">
        <el-button type="info" @click="onAdd">添加产品</el-button>
      </div>
      <div class="device-model-search mb15">
        <TableView style="height: 100%" :dataFn="tableDataFn" :config="config">
          <template #public="row">
            <el-tag type="success" v-if="row.scope['public'] === 0">私有</el-tag>
            <el-tag type="warning" v-else>公开</el-tag>
          </template>
          <template #operate="row">
            <el-button link type="primary" @click="onEdit(row.scope)">编辑</el-button>
            <el-button link type="primary" @click="onDel(row.scope)">删除</el-button>
          </template>
        </TableView>
      </div>
    </el-card>
    <AddCp ref="AddCpRef"></AddCp>
    <EditCp ref="EditCpRef"></EditCp>
  </div>
</template>

<style scoped lang="scss">

</style>