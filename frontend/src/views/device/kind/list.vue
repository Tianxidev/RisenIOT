<script setup lang="ts">

import {ElMessage, ElMessageBox} from "element-plus";
import {DeviceKindDel, DeviceKindList} from "/@/api/system/device";
import {getCurrentInstance, onMounted, reactive, ref} from "vue";
import AddCp from "./component/AddCp.vue";
import EditCp from "./component/EditCp.vue";
import ManageCp from "./component/ManageCp.vue";

const {proxy} = <any>getCurrentInstance();
const AddCpRef = ref<InstanceType<typeof AddCp>>()
const EditCpRef = ref<InstanceType<typeof EditCp>>()
const ManageCpRef = ref<InstanceType<typeof ManageCp>>()

interface IKindList {
  id: number, // 产品ID
  name: string, // 产品名称
  mark: string, // 产品标记
  timeOut: number, // 心跳超时
  createdAt: string, // 创建时间
  updatedAt: string, // 更新时间
  createBy: number, // 创建人
  public: number, // 是否公开
}

const state = reactive({
  list: [] as IKindList[], // 产品列表
})

// 添加产品
const onAdd = () => {
  AddCpRef.value?.openDialog();
};

// 编辑
const onEdit = (row?: any) => {
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
      ElMessage.success(res.msg);
      proxy.mittBus.emit('RefreshPage', true);
    });
  });
}

const onManage = (row?: any) => {
  console.log("管理数据:", row)
  ManageCpRef.value?.openDialog(row);
}

const initData = () => {
  DeviceKindList().then(res => {
    if (res.code !== 0) {
      ElMessage.error(res.msg);
      return;
    }
    state.list = res.data.list;
  });
}

onMounted(() => {
  initData();
})

</script>

<template>
  <div class="device-model-container">
    <div class="device-model-search mb15">
      <el-button type="info" @click="onAdd">添加产品</el-button>
    </div>
    <div>
      <el-row :gutter="20">
        <el-col :span="6" v-for="(item, index) in state.list" :key="index">
          <el-card class="box-card">
            <template #header>
              <div class="card-header">
                <h1>{{ item.name }}</h1>
              </div>
            </template>
            <div class="text item">
              <p><b>产品ID:</b> {{ item.id }}</p>
              <p><b>标识:</b> {{ item.mark }}</p>
              <p><b>创建时间:</b> {{ item.createdAt }}</p>
              <p><b>产品数据心跳超时:</b> {{ item.timeOut }} 秒</p>
              <p><b>是否公开产品:</b> {{ item.public ? "公开" : "私有" }}</p>
            </div>
            <template #footer>
              <el-button class="button" @click="onManage(item)" type="primary" size="small" plain>管理</el-button>
              <el-button class="button" @click="onEdit(item)" type="info" size="small" plain>编辑</el-button>
              <el-button class="button" @click="onDel(item)" type="danger" size="small" plain>删除</el-button>
            </template>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <ManageCp ref="ManageCpRef"/>
    <AddCp ref="AddCpRef"/>
    <EditCp ref="EditCpRef"/>
  </div>
</template>

<style scoped lang="scss">

</style>