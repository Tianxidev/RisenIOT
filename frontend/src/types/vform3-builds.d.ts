declare module 'vform3-builds' {

    export interface VFormRender {
        getFormData: Function,
        loadOptions: (options) => {},
    }

    const VForm3: any;
    export default VForm3;
}