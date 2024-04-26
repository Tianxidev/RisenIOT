<script setup lang="ts">

import TableView from "/@/components/table/TableView.vue";
import { DeviceStrategySearchReq } from "/@/api/system/device";
import { ElMessage } from "element-plus";
import { ref } from "vue";

// 添加
const drawerAdd = ref(false);

// 新增策略表单定义
class IAddStrategyForm {
  name = null; // 策略名称
  type = null; // 策略类型
  cron = null; // 定时时间
  deviceId = null; // 设备ID
  active = null; // 操作动作
}

const AddStrategyForm = ref(new IAddStrategyForm());

// 表格配置
let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: '策略ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '策略名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '策略类型', prop: 'type', minWidth: 100, align: 'center'},
      {label: '定时时间', prop: 'cron', minWidth: 100, align: 'center'},
      {label: '设备ID', prop: 'deviceId', minWidth: 100, align: 'center'},
      {label: '操作', prop: 'operate', slot: true, minWidth: 100, align: 'center'},
    ],
  },
};

// 表格数据加载函数
const tableDataFn = (params: any) => {
  return new Promise((resolve) => {
    DeviceStrategySearchReq(params).then((res) => {
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

const onEdit = (scope?: any) => {
  console.log('onEdit', scope);
};

const onDel = (scope?: any) => {
  console.log('onEdit', scope);
};


</script>

<template>
  <div class="device-model-container">
    <el-card shadow="hover">
      <div class="device-model-search mb15">
        <el-button type="info" @click="drawerAdd = true">添加策略</el-button>
      </div>
      <div class="device-model-search mb15">
        <TableView style="height: 100%" :dataFn="tableDataFn" :config="config">
          <template #operate="row">
            <el-button link type="primary" @click="onEdit(row.scope)">编辑</el-button>
            <el-button link type="primary" @click="onDel(row.scope)">删除</el-button>
          </template>
        </TableView>
      </div>
    </el-card>
    <el-drawer
      v-model="drawerAdd"
      title="添加策略"
      direction="rtl"
    >
      <!-- 添加策略 -->
      <div class="drawer-row">
        <el-form :model="AddStrategyForm" label-width="auto" style="max-width: 600px">
          <el-form-item label="策略名称">
            <el-input v-model="AddStrategyForm.name"/>
          </el-form-item>
          <el-form-item label="策略类型">
            <el-select v-model="AddStrategyForm.type" placeholder="请选择策略类型">
              <el-option label="基于时间的策略" value="1"/>
              <el-option label="基于数据的策略" value="2"/>
            </el-select>
          </el-form-item>
          <el-form-item label="策略时间 Cron 表达式">
            <el-col :span="11">
              <el-input v-model="AddStrategyForm.cron" :disabled="AddStrategyForm.type == 2"/>
            </el-col>
          </el-form-item>
          <el-form-item label="策略表达式">
            <el-input v-model="AddStrategyForm.active" type="textarea"/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary">创建</el-button>
            <el-button type="warning">重置</el-button>
            <el-button>取消</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>
  </div>
</template>

<style scoped lang="scss">
.drawer-row {
  padding: 20px;
}
</style>