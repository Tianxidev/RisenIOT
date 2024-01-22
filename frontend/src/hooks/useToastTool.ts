export const useToastTool = () => {
  function info(...msg: any[]) {
    console.log(msg);
  }

  function success(...msg: any[]) {
    console.log(msg);
  }

  function error(...msg: any[]) {
    console.log(msg);
  }

  function warning(...msg: any[]) {
    console.log(msg);
  }

  return {
    info,
    success,
    error,
    warning
  };
};