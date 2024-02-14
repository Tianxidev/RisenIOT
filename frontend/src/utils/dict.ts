
// 回显数据字典
export function selectDictLabel(datas: any, value?: string | number): string {
    if (value === undefined) {
        return "";
    }
    const actions: string[] = [];
    Object.keys(datas).some((key) => {
        if (datas[key].value === value) {
            actions.push(datas[key].label);
            return true;
        }
    });
    if (actions.length === 0) {
        actions.push(String(value));
    }
    return actions.join('');
}
