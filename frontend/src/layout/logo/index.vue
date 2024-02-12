<template>
	<div class="layout-logo" v-if="setShowLogo" @click="onThemeConfigChange">
		<span>{{ themeConfig.globalTitle }} - 物联网设备云平台 </span>
	</div>
	<div class="layout-logo-size" v-else @click="onThemeConfigChange">
	</div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';

export default defineComponent({
	name: 'layoutLogo',
	setup() {
		const storesThemeConfig = useThemeConfig();
		const { themeConfig } = storeToRefs(storesThemeConfig);
		// 设置 logo 的显示。classic 经典布局默认显示 logo
		const setShowLogo = computed(() => {
			let { isCollapse, layout } = themeConfig.value;
			return !isCollapse || layout === 'classic' || document.body.clientWidth < 1000;
		});

		// logo 点击实现菜单展开/收起
		const onThemeConfigChange = () => {
			if (themeConfig.value.layout === 'transverse') return false;
			themeConfig.value.isCollapse = !themeConfig.value.isCollapse;
		};

		return {
			setShowLogo,
			themeConfig,
			onThemeConfigChange,
		};
	},
});
</script>

<style scoped lang="scss">
.layout-logo {
	width: 320px;
  //background: rgba(28, 37, 24, 0.15);
	height: 50px;
	display: flex;
	align-items: center;
	justify-content: center;
	box-shadow: rgb(0 21 41 / 2%) 0px 1px 4px;
	//color: var(--el-color-primary);
	font-size: 20px;
  font-weight: bold;
	cursor: pointer;
	span {
		white-space: nowrap;
		display: inline-block;
	}
	&:hover {
		span {
			color: var(--color-primary-light-2);
		}
	}
	&-medium-img {
		width: 40px;
	}
}
.layout-logo-size {
	width: 100%;
	height: 50px;
	display: flex;
	cursor: pointer;
	&-img {
		width: 20px;
		margin: auto;
	}
}
</style>
