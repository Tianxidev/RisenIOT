/**
 * 参数字符串转换成对象形式，如：a=1&b=2 转换成 {a:1, b:2}
 * @param {String} str 需要转换的字符串
 * @param {String} [sep=&] 连接符，可选，默认 &
 * @param {String} [eq==] 键值间隔符，可选，默认 =
 * @returns {Object}
 */
export function parse(str: string, sep?: string, eq?: string): any {
    // eslint-disable-next-line
    let obj: any = {};
    str = (str || location.search).replace(/^[^]*\?/, '');
    sep = sep || '&';
    eq = eq || '=';
    let arr,
        // eslint-disable-next-line prefer-const
        reg = new RegExp('(?:^|\\' + sep + ')([^\\' + eq + '\\' + sep + ']+)(?:\\' + eq + '([^\\' + sep + ']*))?', 'g');
    while ((arr = reg.exec(str)) !== null) {
        if (arr[1] !== str) {
            obj[decodeURIComponent(arr[1])] = decodeURIComponent(arr[2] || '');
        }
    }
    return obj;
}

/**
 * 将对象转换成参数字符串形式，如：{a: 1, b: 2} 转换成 "a=1&b=2"
 * @param {Object} obj 需要转换的对象
 * @param {String} [sep=&] 连接符，可选，默认 &
 * @param {String} [eq==] 键值间隔符，可选，默认 =
 * @returns {String}
 */
export function stringify(obj: Record<string, any>, sep?: string, eq?: string): string {
    sep = sep || '&';
    eq = eq || '=';
    const pairs: any[] = [];

    for (const key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) {
            const value = obj[key];
            if (value !== null && value !== undefined) {
                pairs.push(encodeURIComponent(key) + eq + encodeURIComponent(value));
            }
        }
    }

    return pairs.join(sep);
}

/**
 * 在url追加参数
 * @param {string} url 原本的url
 * @param {object} query 需要追加的参数，Object|String
 * @returns {string} 追加参数后的url
 */
export function appendQuery(url: string, query: object): string {
    const path: string = url.split('?')[0];
    const originalQuery = parse(url);
    const joinQuery = Object.assign({}, originalQuery, query);
    const queryStr: string = stringify(joinQuery);
    return queryStr ? [path, queryStr].join('?') : url;
}