import JSEncrypt from "encryptlong";

/**
 * 加解密相关工具
 */
const useSecretTool = () => {
  /**
   * Base64编码解码
   * @author Tianxidev<2200475850@qq.com>
   * @param {string} encodedString Base64编码的字符串
   * @returns
   */
  const decodeBase64 = (encodedString: string) => {
    const decodedString = atob(encodedString);
    const uint8Array = new Uint8Array(decodedString.length);

    for (let i = 0; i < decodedString.length; i++) {
      uint8Array[i] = decodedString.charCodeAt(i);
    }

    return uint8Array;
  };

  /**
   * RSA加密
   * @author Tianxidev<2200475850@qq.com>
   * @param {string} publicKey 公钥Base64编码的字符串
   * @param {string} data 要加密的数据
   * @returns {string} 加密后的数据
   */
  const rsaEncrypt = (publicKey: string, data: string): string => {
    const decrypt = new JSEncrypt();
    decrypt.setPublicKey(publicKey);
    return decrypt.encrypt(data);
  };

  return {
    decodeBase64,
    rsaEncrypt
  };
};

export default useSecretTool;
