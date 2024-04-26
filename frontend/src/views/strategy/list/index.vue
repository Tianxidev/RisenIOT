<script setup lang="ts">

import TableView from "/@/components/table/TableView.vue";
import {
  DeviceInfoList,
  DeviceStrategyAddReq, DeviceStrategyDelReq,
  DeviceStrategyEditReq,
  DeviceStrategySearchReq
} from "/@/api/system/device";
import { ElMessage, ElMessageBox } from "element-plus";
import { onMounted, ref, watch } from "vue";

const TableRef = ref<typeof TableView>();

// 设备信息结构体
interface IDeviceInfo {
  id: number; // 设备ID
  name: string; // 设备名称
  kind: number; // 产品类型
  group: number; // 设备分组
  sn: string; // 设备SN
  monitor: number; // 在线状态
  upTime: string; // 最后一次数据上报
  updatedAt: string; // 更新时间
  pwd: string; // 设备密码
}


const deviceInfoList = ref([] as IDeviceInfo[]); // 设备信息列表
const drawerAdd = ref(false); // 添加抽屉
const drawerEdit = ref(false); // 编辑抽屉

// 策略表单定义
class IStrategyForm {
  name = null; // 策略名称
  type = null; // 策略类型
  cron = null; // 定时时间
  deviceId = null; // 设备ID
  action = null; // 操作动作
}

// 表单
const AddStrategyForm = ref(new IStrategyForm()); // 添加策略表单
const EditStrategyForm = ref(new IStrategyForm()); // 编辑策略表单


// 表格配置
let config = {
  table: {
    columns: [
      {type: 'selection', minWidth: 40},
      {label: '策略ID', prop: 'id', minWidth: 100, align: 'center'},
      {label: '策略名称', prop: 'name', minWidth: 100, align: 'center'},
      {label: '策略类型', prop: 'type', minWidth: 100, align: 'center'},
      {label: '定时时间', prop: 'cron', minWidth: 100, align: 'center'},
      {label: '目标设备', prop: 'deviceId', slot: true, minWidth: 100, align: 'center'},
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

// 编辑策略
const onHandleEdit = (scope?: IStrategyForm, subject?: boolean) => {
  if (!scope) return;
  if (subject) {
    DeviceStrategyEditReq(EditStrategyForm.value).then(res => {
      if (res.code !== 0) {
        ElMessage.error(res.msg);
        return;
      }
      ElMessage.success('编辑成功');
      drawerEdit.value = false;
      Object.assign(EditStrategyForm.value, new IStrategyForm());
      TableRef.value?.getData();
    });
    return;
  }
  console.log('onEdit', scope);
  EditStrategyForm.value = scope;
  drawerEdit.value = true;
};

// 删除策略
const onHandleDel = (scope?: any) => {
  console.log('onEdit', scope);
  ElMessageBox.confirm("确定要删除吗?", '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    DeviceStrategyDelReq([scope.id]).then((res) => {
      if (res.code != 0) {
        ElMessage.error(res.msg)
      }
      ElMessage.success(res.msg);
      TableRef.value?.getData();
    });
  });
};

// 添加策略
const onHandleAdd = () => {
  console.log('onHandleAdd', AddStrategyForm.value);
  DeviceStrategyAddReq(AddStrategyForm.value).then(res => {
    console.log(res)
    if (res.code !== 0) {
      ElMessage.error(res.msg);
      return;
    }
    ElMessage.success('添加成功');
    drawerAdd.value = false;
    Object.assign(AddStrategyForm.value, new IStrategyForm());
    TableRef.value?.getData();
  })
};

// 加载设备列表
const loadDeviceList = () => {
  DeviceInfoList().then(res => {
    // console.log('加载设备列表', res);
    deviceInfoList.value = res.data.list;
  })
}

// 监听编辑策略类型变化
watch(EditStrategyForm, () => {
  if (EditStrategyForm.value.type == 2) {
    EditStrategyForm.value.cron = null;
  }
});

// 监听添加策略类型变化
watch(AddStrategyForm, () => {
  if (AddStrategyForm.value.type == 2) {
    AddStrategyForm.value.cron = null;
  }
});

onMounted(() => {
  loadDeviceList();
});

</script>

<template>
  <div class="device-model-container">
    <el-card shadow="hover">
      <div class="device-model-search mb15">
        <el-button type="info" @click="drawerAdd = true">添加策略</el-button>
      </div>
      <div class="device-model-search mb15">
        <TableView ref="TableRef" style="height: 100%" :dataFn="tableDataFn" :config="config">
          <template #deviceId="row">
            <template v-for="item in deviceInfoList" :key="item.id">
              <el-tag v-if="item.id == row.scope.deviceId" type="info">{{ item.name }}</el-tag>
            </template>
          </template>
          <template #operate="row">
            <el-button link type="primary" @click="onHandleEdit(row.scope)">编辑</el-button>
            <el-button link type="primary" @click="onHandleDel(row.scope)">删除</el-button>
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
          <el-form-item label="目标设备">
            <el-select
              v-model="AddStrategyForm.deviceId"
              placeholder="选择目标设备"
            >
              <el-option
                v-for="item in deviceInfoList"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="Cron 表达式">
            <el-input v-model="AddStrategyForm.cron" :disabled="AddStrategyForm.type == 2"/>
          </el-form-item>
          <el-form-item label="策略表达式">
            <el-input v-model="AddStrategyForm.action" type="textarea"/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onHandleAdd">创建</el-button>
            <el-button @click="drawerAdd = false">取消</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>

    <el-drawer
      v-model="drawerEdit"
      title="编辑策略"
      direction="rtl"
    >
      <!-- 编辑策略 -->
      <div class="drawer-row">
        <el-form :model="EditStrategyForm" label-width="auto" style="max-width: 600px">
          <el-form-item label="策略名称">
            <el-input v-model="EditStrategyForm.name"/>
          </el-form-item>
          <el-form-item label="策略类型">
            <el-select v-model="EditStrategyForm.type" placeholder="请选择策略类型">
              <el-option label="基于时间的策略" value="1"/>
              <el-option label="基于数据的策略" value="2"/>
            </el-select>
          </el-form-item>
          <el-form-item label="Cron 表达式">
            <el-input v-model="EditStrategyForm.cron" :disabled="EditStrategyForm.type == 2"/>
          </el-form-item>
          <el-form-item label="策略表达式">
            <el-input v-model="EditStrategyForm.action" type="textarea"/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onHandleEdit(EditStrategyForm, true)">更新</el-button>
            <el-button @click="drawerEdit=false">取消</el-button>
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