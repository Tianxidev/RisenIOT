import {ref} from "vue"

export const useAdd = () => {
  const addNum = ref(0);

  const addFn = (num1, num2) => {
    addNum.value = num1 + num2
  }

  return {
    addNum,
    addFn
  }

}
