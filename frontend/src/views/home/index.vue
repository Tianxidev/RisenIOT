<script setup lang="ts">
import { NCard, NGi, NGrid } from 'naive-ui';
import { onMounted, reactive } from "vue";
import { OverviewApi } from "/@/api/home";

const state = reactive({
  deviceTotal: 0,
  onlineDeviceTotal: 0,
  offlineDeviceTotal: 0,
  emqxTotal: 0
})

const initView = () => {
  OverviewApi().then(res=>{
    state.deviceTotal = res?.data['totalDevice']
  })
}

onMounted(() => {
  initView()
})

</script>

<template>
  <div class="dashboard">
    <n-grid x-gap="12" :cols="4">
      <n-gi>
        <n-card title="接入设备" class="top-grid" hoverable>
          <template #header-extra>台</template>
          <span class="num">{{ state.deviceTotal }}</span>
        </n-card>
      </n-gi>
      <n-gi>
        <n-card title="在线设备" class="top-grid" hoverable>
          <template #header-extra>台</template>
          <span class="num">{{ state.onlineDeviceTotal }}</span>
        </n-card>
      </n-gi>
      <n-gi>
        <n-card title="掉线设备" class="top-grid" hoverable>
          <template #header-extra>台</template>
          <span class="num">{{ state.offlineDeviceTotal }}</span>
        </n-card>
      </n-gi>
      <n-gi>
        <n-card title="EMQX-集群" class="top-grid" hoverablea>
          <template #header-extra>台</template>
          <span class="num">{{ state.emqxTotal }}</span>
        </n-card>
      </n-gi>
    </n-grid>
  </div>
</template>

<style lang="scss" scoped>
.top-grid {
  background-color: #ffffff;

  .num {
    font-size: 40px;
    font-weight: bold;
  }
}
</style>
