// 定义全局配置文件
export const Config = () => {
  let SystemName: string = "";
  let baseUrl: string = "";
  let WsUrl: string = "";

  // 读取环境变量
  const env: ImportMetaEnv = import.meta.env;

  // 是否开启调试模式
  let logDebug = false;

  const protocol: string = window.location.protocol;
  const hostname: string = window.location.hostname;
  const port: string = window.location.port;
  const Copyright: string = env.VITE_COPYRIGHT;
  const SystemDescription: string = env.VITE_SYSDESCRIPTION;
  const RecordNumber: string = env.VITE_RECORDNUMBER;

  switch (env.VITE_ENVIRONMENT) {
    case "development": // 开发环境
      SystemName = env.VITE_SYSNAME;
      baseUrl = "/api";
      WsUrl = env.VITE_WS_URL + "/ws/v1/lamp";
      logDebug = true;
      break;
    case "production": // 生产环境
      SystemName = env.VITE_SYSNAME;
      baseUrl = "/api";
      WsUrl = `${protocol === "https:" ? "wss:" : "ws:"}//${hostname}:${port}/ws/v1/lamp`;
      logDebug = false;
      break;
    default:
      break;
  }

  return {
    SystemName,
    SystemDescription,
    baseUrl,
    WsUrl,
    logDebug,
    Copyright,
    RecordNumber
  };
};
