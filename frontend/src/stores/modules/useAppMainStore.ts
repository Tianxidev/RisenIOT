import { ref } from "vue";
import { defineStore } from "pinia";

export type ThemeType = "light" | "dark" | "transparent";

export const useAppMainStore = defineStore("AppMain", () => {

    const scheme = window.matchMedia("(prefers-color-scheme: dark)");
    scheme.addEventListener("change", () => {
        theme.value = scheme.matches ? "dark" : "light";
    });
    const root = document.querySelector(":root");
    root?.setAttribute("theme", scheme.matches ? "dark" : "light");
    const theme = ref(scheme.matches ? "dark" : "light");

    // 切换主题颜色
    const onTheme = (typeVal?: ThemeType) => {
        if (typeVal !== undefined && typeVal === "transparent") {
            theme.value = typeVal;
            root?.setAttribute("theme", theme.value);
            return;
        }
        theme.value = theme.value === "light" ? "dark" : "light";
        root?.setAttribute("theme", theme.value);
    };

    const upVisible = ref(false);

    const onHideUp = () => {
        upVisible.value = false;
    };
    let versionCode = "";
    const onShowUp = (code: string) => {
        versionCode = code;
        upVisible.value = true;
    };
    const onUp = () => {
        upVisible.value = false;
    };

    // 组件颜色
    const themeAssemblyColor = (typeVal?: ThemeType) => {
        if (typeVal !== undefined && typeVal === "transparent") {
            return "#000000000";
        }
        return theme.value === "light" ? "#ffffff" : "#212121FF";
    };

    return { theme, themeAssemblyColor, onTheme, upVisible, onHideUp, onShowUp, onUp };
});
