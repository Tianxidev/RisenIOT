<template>
	<div class="home-container">
		<el-row :gutter="15" class="home-card-one mb15">
			<el-col
				:xs="24"
				:sm="12"
				:md="12"
				:lg="6"
				:xl="6"
				v-for="(v, k) in homeOne"
				:key="k"
				:class="{ 'home-media home-media-lg': k > 1, 'home-media-sm': k === 1 }"
			>
				<div class="home-card-item flex">
					<div class="flex-margin flex w100" :class="` home-one-animation${k}`">
						<div class="flex-auto">
							<span class="font30">{{ v.num1 }}</span>
							<span class="ml5 font16" :style="{ color: v.color1 }">{{ v.num2 }}</span>
							<div class="mt10">{{ v.num3 }}</div>
						</div>
						<div class="home-card-item-icon flex" :style="{ background: `var(${v.color2})` }">
							<i class="flex-margin font32" :class="v.num4" :style="{ color: `var(${v.color3})` }"></i>
						</div>
					</div>
				</div>
			</el-col>
		</el-row>
		<el-row :gutter="15" class="home-card-two mb15">
      <el-col :xs="24">
        <div class="home-card-item">
        </div>
      </el-col>
		</el-row>
	</div>
</template>

<script lang="ts">
import { defineComponent, nextTick, onActivated, onMounted, reactive, ref, toRefs, watch } from 'vue';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import { useTagsViewRoutes } from '/@/stores/tagsViewRoutes';

export default defineComponent({
	name: 'home',
	setup() {
		const homeLineRef = ref();
		const homePieRef = ref();
		const homeBarRef = ref();
		const storesTagsViewRoutes = useTagsViewRoutes();
		const storesThemeConfig = useThemeConfig();
		const { themeConfig } = storeToRefs(storesThemeConfig);
		const { isTagsViewCurrenFull } = storeToRefs(storesTagsViewRoutes);
		const state = reactive({
			homeOne: [
				{
					num1: '125,12',
					num2: '台',
					num3: '接入设备数量',
					num4: 'iconfont iconfont-shebei',
					color1: '#6690F9',
					color2: '--next-color-primary-lighter',
					color3: '--el-color-primary',
				},
				{
					num1: '653,33',
					num2: '台',
					num3: '在线设备数量',
					num4: 'iconfont iconfont-zaixianshebei',
					color1: '#6690F9',
					color2: '--next-color-success-lighter',
					color3: '--el-color-success',
				},
				{
					num1: '125,65',
					num2: '台',
					num3: '离线设备数量',
					num4: 'iconfont iconfont-lixianshebei',
					color1: '#6690F9',
					color2: '--next-color-warning-lighter',
					color3: '--el-color-warning',
				},
				{
					num1: '520,43',
					num2: '个',
					num3: '分布式节点数量',
					num4: 'iconfont iconfont-fenbushi',
					color1: '#FF6462',
					color2: '--next-color-danger-lighter',
					color3: '--el-color-danger',
				},
			],
			homeThree: [
				{
					icon: 'iconfont icon-yangan',
					label: '浅粉红',
					value: '2.1%OBS/M',
					iconColor: '#F72B3F',
				},
				{
					icon: 'iconfont icon-wendu',
					label: '深红(猩红)',
					value: '30℃',
					iconColor: '#91BFF8',
				},
				{
					icon: 'iconfont icon-shidu',
					label: '淡紫红',
					value: '57%RH',
					iconColor: '#88D565',
				},
				{
					icon: 'iconfont icon-shidu',
					label: '弱紫罗兰红',
					value: '107w',
					iconColor: '#88D565',
				},
				{
					icon: 'iconfont icon-zaosheng',
					label: '中紫罗兰红',
					value: '57DB',
					iconColor: '#FBD4A0',
				},
				{
					icon: 'iconfont icon-zaosheng',
					label: '紫罗兰',
					value: '57PV',
					iconColor: '#FBD4A0',
				},
				{
					icon: 'iconfont icon-zaosheng',
					label: '暗紫罗兰',
					value: '517Cpd',
					iconColor: '#FBD4A0',
				},
				{
					icon: 'iconfont icon-zaosheng',
					label: '幽灵白',
					value: '12kg',
					iconColor: '#FBD4A0',
				},
				{
					icon: 'iconfont icon-zaosheng',
					label: '海军蓝',
					value: '64fm',
					iconColor: '#FBD4A0',
				},
			],
			myCharts: [],
			charts: {
				theme: '',
				bgColor: '',
				color: '#303133',
			},
		});

		// 批量设置 echarts resize
		const initEchartsResizeFun = () => {
			nextTick(() => {
				for (let i = 0; i < state.myCharts.length; i++) {
					setTimeout(() => {
						(<any>state.myCharts[i]).resize();
					}, i * 1000);
				}
			});
		};
		// 批量设置 echarts resize
		const initEchartsResize = () => {
			window.addEventListener('resize', initEchartsResizeFun);
		};
		// 页面加载时
		onMounted(() => {
			initEchartsResize();
		});
		// 由于页面缓存原因，keep-alive
		onActivated(() => {
			initEchartsResizeFun();
		});
		// 监听 vuex 中的 tagsview 开启全屏变化，重新 resize 图表，防止不出现/大小不变等
		watch(
			() => isTagsViewCurrenFull.value,
			() => {
				initEchartsResizeFun();
			}
		);
		// 监听 vuex 中是否开启深色主题
		watch(
			() => themeConfig.value.isIsDark,
			(isIsDark) => {
				nextTick(() => {
					state.charts.theme = isIsDark ? 'dark' : '';
					state.charts.bgColor = isIsDark ? 'transparent' : '';
					state.charts.color = isIsDark ? '#dadada' : '#303133';
				});
			},
			{
				deep: true,
				immediate: true,
			}
		);
		return {
			homeLineRef,
			homePieRef,
			homeBarRef,
			...toRefs(state),
		};
	},
});
</script>

<style scoped lang="scss">
$homeNavLength: 8;
.home-container {
	overflow: hidden;
	.home-card-one,
	.home-card-two,
	.home-card-three {
		.home-card-item {
			width: 100%;
			height: 130px;
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
	}
	.home-card-one {
		@for $i from 0 through 3 {
			.home-one-animation#{$i} {
				opacity: 0;
				animation-name: error-num;
				animation-duration: 0.5s;
				animation-fill-mode: forwards;
				animation-delay: calc($i/10) + s;
			}
		}
	}
	.home-card-two,
	.home-card-three {
		.home-card-item {
			height: 400px;
			width: 100%;
			overflow: hidden;
			.home-monitor {
				height: 100%;
				.flex-warp-item {
					width: 25%;
					height: 111px;
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
					}
					@for $i from 0 through $homeNavLength {
						.home-animation#{$i} {
							opacity: 0;
							animation-name: error-num;
							animation-duration: 0.5s;
							animation-fill-mode: forwards;
							animation-delay: calc($i/10) + s;
						}
					}
				}
			}
		}
	}
}
</style>
