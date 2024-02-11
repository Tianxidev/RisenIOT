<script setup lang="ts">
import { getCurrentInstance, reactive, ref } from "vue";

const ruleFormRef = ref<HTMLElement | null>(null);
const {proxy} = getCurrentInstance() as any;

const props = {
  acType: {
    type: String,
    default: () => 'add'
  },
  row: {
    type: Object,
    default: () => {
      return {};
    },
  },
};

const state = reactive({
  loading: false,
  isShowDialog: false,
  roles: [],
  ruleForm: {
    id: undefined,
    pid: 0, // 上级菜单
    menuType: '0', // 菜单类型
    menuName: '', // 菜单名称
    name: '', // 接口规则
    component: '', // 组件路径
    isLink: 0, // 是否外链
    menuSort: 0, // 菜单排序
    path: '', // 路由路径
    redirect: '', // 路由重定向，有子集 children 时
    icon: '', // 菜单图标
    roles: [], // 权限标识，取角色管理
    isHide: '0', // 是否隐藏
    isKeepAlive: 1, // 是否缓存
    isAffix: 0, // 是否固定
    linkUrl: '', // 外链/内嵌时链接地址（http:xxx.com），开启外链条件，`1、isLink:true 2、链接地址不为空`
    isIframe: 0, // 是否内嵌，开启条件，`1、isIframe:true 2、链接地址不为空`
  },
  // 表单校验
  rules: {
    parentId: [
      {required: true, message: "父菜单不能为空", trigger: "blur"},
    ],
    name: [
      {required: true, message: "接口规则不能为空", trigger: "blur"},
    ],
    path: [
      {required: true, message: "路由地址不能为空", trigger: "blur"},
    ],
    menuName: [
      {required: true, message: "菜单名称不能为空", trigger: "blur"},
    ],
    menuType: [
      {required: true, message: "菜单类型不能为空", trigger: "blur"},
    ],
  },
  menuData: [], // 上级菜单数据
});

// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};

// 打开菜单
const openDialog = (row: any) => {

}

// 取消
const onCancel = () => {
  closeDialog();
};

// 新增
const onSubmit = () => {

}

</script>

<template>
  <div class="device-add-dialog-container">
    <el-dialog :title="(props.acType==='add'?'新增':'修改')+'菜单'" v-model="state.isShowDialog">

    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>