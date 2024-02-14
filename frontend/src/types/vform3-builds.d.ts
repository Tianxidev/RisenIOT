declare module 'vform3-builds' {

    export interface VFormRender {
        getFormData: ()=> Promise,
        resetForm: ()=>{},
        loadOptions: (options) => {},
    }

    const VForm3: any;
    export default VForm3;
}