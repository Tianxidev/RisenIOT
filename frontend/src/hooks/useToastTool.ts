import { toast } from "vuetify-sonner";

export const useToastTool = () => {
  function info(...msg: any[]) {
    toast.info(msg.join(" "), {
      duration: 1000
    });
  }

  function success(...msg: any[]) {
    toast.success(msg.join(" "), {
      duration: 1000
    });
  }

  function error(...msg: any[]) {
    toast.error(msg.join(" "), {
      duration: 1000
    });
  }

  function warning(...msg: any[]) {
    toast.warning(msg.join(" "), {
      duration: 1000
    });
  }

  return {
    info,
    success,
    error,
    warning
  };
};