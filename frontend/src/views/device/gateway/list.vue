<script setup lang="ts">


const homeThree = [
  {
    icon: 'iconfont iconfont-Http-get',
    label: 'HTTP Gateway',
    value: 'HTTP 数据解析网关',
    remark: '点击配置',
    iconColor: '#F72B3F',
  },
  {
    icon: 'iconfont iconfont-MQTT',
    label: 'MQTT Gateway',
    value: 'MQTT 数据解析网关',
    remark: '点击配置',
    iconColor: '#F72B3F',
  },
  // {
  //   icon: 'iconfont iconfont-haikangshipinjiankongxitong',
  //   label: 'HIK Gateway',
  //   value: '海康威视 SDAP 被动设备发现网关',
  //   remark: '点击配置',
  //   iconColor: '#F72B3F',
  // },
]

</script>

<template>
  <div>
    <div class="gateway-card-item">
      <div class="gateway-card-item-title">数据网关配置</div>
      <div class="gateway-monitor">
        <div class="flex-warp">
          <div class="flex-warp-item" v-for="(v, k) in homeThree" :key="k">
            <div class="flex-warp-item-box" :class="`gateway-animation${k}`">
              <div class="flex-margin">
                <div class="title">
                  <i :class="v.icon" :style="{ color: v.iconColor }"></i>
                  <span class="label">{{ v.label }}</span>
                </div>
                <div class="mt10">{{ v.value }}</div>
                <div class="remark">{{ v.remark }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
$gatewayNavLength: 8;

.gateway-card-item {
  width: 100%;
  border-radius: 4px;
  transition: all ease 0.3s;
  padding: 20px;
  overflow: hidden;
  background: var(--el-color-white);
  color: var(--el-text-color-primary);
  border: 1px solid var(--next-border-color-light);

  &:hover {
    box-shadow: 0 2px 12px var(--next-color-dark-hover);
    transition: all ease 0.3s;
  }

  &-icon {
    width: 70px;
    height: 70px;
    border-radius: 100%;
    flex-shrink: 1;

    i {
      color: var(--el-text-color-placeholder);
    }
  }

  &-title {
    font-size: 15px;
    font-weight: bold;
    height: 30px;
  }
}

.gateway-monitor {
  height: 100%;
  width: 100%;

  .flex-warp-item {
    width: 25%;
    height: 120px;
    display: flex;

    .flex-warp-item-box {
      margin: auto;
      text-align: center;
      color: var(--el-text-color-primary);
      display: flex;
      border-radius: 5px;
      background: var(--next-bg-color);
      cursor: pointer;
      transition: all 0.3s ease;

      &:hover {
        background: var(--el-color-primary-light-9);
        transition: all 0.3s ease;
      }

      .flex-margin {
        display: flex;
        flex-direction: column;
        margin: auto;
        padding-top: 10px;

        .title {
          display: flex;
          align-items: center;
          justify-content: start;

          i {
            font-size: 40px;
            margin-left: 5px;
            margin-right: 5px;
          }

          .label {
            font-size: 30px;
          }

        }

      }

    }

    @for $i from 0 through $gatewayNavLength {
      .gateway-animation#{$i} {
        opacity: 0;
        animation-name: error-num;
        animation-duration: 0.5s;
        animation-fill-mode: forwards;
        animation-delay: calc($i/10) + s;
      }
    }
  }
}

.remark {
  color: #070e94;
  font-size: 11px;
  margin: 10px 0;
}
</style>