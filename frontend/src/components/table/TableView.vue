<script setup>
import { computed, getCurrentInstance, onMounted, reactive } from 'vue';

const { proxy } = getCurrentInstance();

const props = defineProps({
  dataFn: {
    type: Function,
    default: () => {},
  }, // 获取数据的方法
  config: { type: Object },
  action: { type: String }, // 请求地址
  tableHeight: { type: Number },
  size: { type: String },
  pageSize: { type: Number, default: 10 },
  options: {
    // 表格配置项
    type: Object,
    default: () => {
      return {
        // height: 400,
        stripe: false, // 斑马纹
        highlightCurrentRow: true, // 是否要高亮当前行
        border: false, // 是否有纵向边框
        fit: true, // 列的宽度是否自撑开
        size: 'default', // Table 的尺寸 "default", "small", "large"
        maxHeight: 560, // Table 的最大高度。
        loading: false, //是否需要等待效果
        lazy: true, // 是否需要懒加载
      };
    },
  },
});

const state = reactive({
  isRowChange: false, // 表格列是否改变
  checkedValue: [], // 复选框已选数据
  tableDataValue: [], // 表格数据
  total: 0, // 表格数据总数
  mPageSize: props.pageSize, // 每页展现条数
  pageNum: 1, // 当前页
  isChange: false, // 是否选中
});

// 获取数据
const getData = async (params) => {
  const res = await props.dataFn({
    pageNum: state.pageNum,
    pageSize: state.mPageSize,
    ...params,
  });
  state.tableDataValue = res.tableDataValue;
  state.total = res.total;
};

// 过滤后的表格列
const filterColumns = computed(() => {
  if (!state.isChange) return props.config.table.columns;
  return props.config.table.columns.filter((i) => (!i.type && state.checkedValue.includes(i.label)) || i.type);
});

// 每页展现条数选择项
const pageSizeOpts = computed(() => {
  return [...new Set([10, 20, 30, 40, proxy.pageSize])].sort((a, b) => a - b);
});

// 页面展现条数事件
const handleSizeChange = (pageSize) => {
  state.mPageSize = pageSize;
  getData();
};

// 页面跳转事件
const handleCurrentChange = (val) => {
  state.pageNum = val;
  getData();
};

// 勾选复选框事件
const select = (row, column, event) => {
  let data = { row, column, event };
  proxy.mittBus.emit('select', data);
};

// 勾选全选事件
const selectAll = (row, column, event) => {
  let data = { row, column, event };
  proxy.mittBus.emit('selectAll', data);
};

// 事件总线触发表格渲染
proxy.mittBus.on('renderTable', (params) => {
  getData(params);
});

onMounted(() => {
  getData();
});
</script>

<template>
  <div class="table">
    <el-table
        v-loading="false"
        :size="props.options.size"
        :data="state.tableDataValue"
        :stripe="props.options.stripe"
        :max-height="props.options.maxHeight"
        :border="props.options.border"
        :fit="props.options.fit"
        :lazy="props.options.lazy"
        :highlight-current-row="props.options.highlightCurrentRow"
        @select="select"
        @select-all="selectAll"
    >
      <template v-for="(item, index) in filterColumns" :key="index">
        <!-- 复选框 (START)-->
        <el-table-column
            v-if="item.type === 'selection'"
            type="selection"
            :width="item.minWidth ? item.minWidth : 60"
        ></el-table-column>
        <!-- 复选框 (END) -->
        <!-- 序号 (START) -->
        <el-table-column
            v-else-if="item.type === 'index'"
            type="index"
            :label="item.label ? item.label : '序号'"
            :width="item.minWidth ? item.minWidth : 80"
            :align="item.align ? item.align : 'center'"
        ></el-table-column>
        <!-- 序号 (END) -->
        <!-- 默认渲染列 (START) -->
        <el-table-column
            v-else
            :prop="item.prop"
            :label="item.label"
            :align="item.align"
            :min-width="item.minWidth"
            :show-overflow-tooltip="true"
        >
          <template #default="scope">
            <template v-if="item.render">
              <slot :name="item.prop" :tit="index" :scope="scope.row"></slot>
            </template>
            <template v-else-if="item.slot">
              <slot :name="item.prop" :tit="index" :scope="scope.row"></slot>
            </template>
            <template v-else>
              <span>{{ scope.row[item.prop] }}</span>
            </template>
          </template>
        </el-table-column>
      </template>
    </el-table>
    <!-- 分页 (START) -->
    <div class="pager" style="text-align: right">
      <el-pagination
          v-model:currentPage="state.pageNum"
          :page-size="state.mPageSize"
          :page-sizes="pageSizeOpts"
          layout="total, sizes, prev, pager, next, jumper"
          :total="state.total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
      >
      </el-pagination>
    </div>
    <!-- 分页 (END) -->
  </div>
</template>

<style scoped lang="scss">

</style>