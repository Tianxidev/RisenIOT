<script setup lang="ts">

import {getCurrentInstance, reactive, ref} from "vue";
import TableView from "/@/components/table/TableView.vue";
import {DeviceCategoryList} from "/@/api/system/device";
import {ElMessage} from "element-plus";
import ManageEditCp from "/@/views/device/kind/component/ManageEditCp.vue";

const ManageEditCpRef = ref<InstanceType<typeof ManageEditCp>>()


const {proxy} = <any>getCurrentInstance();


interface IKind {
  id: number,
  name: string,
  mark: string,
  timeOut: number,
}

const state = reactive({
  isShowDialog: false,
  direction: 'btt',
  row: {} as IKind,
})

let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: 'ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '类型名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '数据类型', prop: 'dataType', minWidth: 100, align: 'center'},
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

const onEdit = (row?: any) => {
  ManageEditCpRef.value?.openDialog(row);
}

const onDel = (row?: any) => {

}

const handleClose = () => {
  Object.assign(state.row, {})
  proxy.mittBus.emit('RefreshPage', true);
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
        @close="handleClose"
    >
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
    </el-drawer>
    <ManageEditCp ref="ManageEditCpRef"/>
  </div>
</template>

<style scoped lang="scss">

</style>