/*******************************************************************************
 * jQuery hash算法插件
 *
 * @author:Allen
 * @date:2016/10/22
 ******************************************************************************/
(function() {
    // 检查是否引入了jquery库
    if (typeof jQuery == 'undefined')
        return false;
    // 检查是否支持File等对象
    if (!window.File || !window.FileReader || !window.FileList || !window.Blob || !File.prototype.slice)
        return false;
    // 处理不同浏览器中的File、Blob对象分割方法slice
    if ((typeof File == 'undefined'))
        return false;
    if (!File.prototype.slice) {
        if (File.prototype.webkitSlice)
            File.prototype.slice = File.prototype.webkitSlice;
        else if (File.prototype.mozSlice)
            File.prototype.slice = File.prototype.mozSlice;
    }
    return true;
})() && (function($) {
    /**
     * 工具方法集
     */
    var util = g = {
        rotl: function(a, b) {
            return a << b | a >>> 32 - b;
        },
        rotr: function(a, b) {
            return a << 32 - b | a >>> b;
        },
        endian: function(a) {
            if (a.constructor == Number)
                return g.rotl(a, 8) & 16711935 | g.rotl(a, 24) & 4278255360;
            for (var b = 0; b < a.length; b++)
                a[b] = g.endian(a[b]);
            return a;
        },
        randomBytes: function(a) {
            for (var b = []; a > 0; a--)
                b.push(Math.floor(Math.random() * 256));
            return b;
        },
        bytesToWords: function(a) {
            for (var b = [], c = 0, d = 0; c < a.length; c++, d += 8)
                b[d >>> 5] |= a[c] << 24 - d % 32;
            return b;
        },
        wordsToBytes: function(a) {
            for (var b = [], c = 0; c < a.length * 32; c += 8)
                b.push(a[c >>> 5] >>> 24 - c % 32 & 255);
            return b;
        },
        bytesToHex: function(a) {
            for (var b = [], c = 0; c < a.length; c++)
                b.push((a[c] >>> 4).toString(16)),
                b.push((a[c] & 15).toString(16));
            return b.join("");
        },
        hexToBytes: function(a) {
            for (var b = [], c = 0; c < a.length; c += 2)
                b.push(parseInt(a.substr(c, 2), 16));
            return b;
        },
        bytesToBase64: function(a) {
            if (typeof btoa == "function")
                return btoa(f.bytesToString(a));
            for (var b = [], c = 0; c < a.length; c += 3)
                for (var d = a[c] << 16 | a[c + 1] << 8 | a[c + 2], e = 0; e < 4; e++)
                    c * 8 + e * 6 <= a.length * 8 ? b.push("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".charAt(d >>> 6 * (3 - e) & 63)) : b.push("=");
            return b.join("");
        },
        base64ToBytes: function(a) {
            if (typeof atob == "function")
                return f.stringToBytes(atob(a));
            for (var a = a.replace(/[^A-Z0-9+/]/ig, ""),
                    b = [],
                    c = 0,
                    d = 0; c < a.length; d = ++c % 4)
                d != 0 && b.push(("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".indexOf(a.charAt(c - 1)) & Math.pow(2, -2 * d + 8) - 1) << d * 2 | "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".indexOf(a.charAt(c)) >>> 6 - d * 2);
            return b;
        }
    };
    /**
     * sha1算法处理
     *
     * @param (Array){block}分块数组
     * @param (Array){hash}当前计算出的hash数组
     */
    var sha1 = function(block, hash) {
            var words = [];
            var count_parts = 16;
            var h0 = hash[0],
                h1 = hash[1],
                h2 = hash[2],
                h3 = hash[3],
                h4 = hash[4];
            for (var i = 0; i < block.length; i += count_parts) {
                var th0 = h0,
                    th1 = h1,
                    th2 = h2,
                    th3 = h3,
                    th4 = h4;
                for (var j = 0; j < 80; j++) {
                    if (j < count_parts)
                        words[j] = block[i + j] | 0;
                    else {
                        var n = words[j - 3] ^ words[j - 8] ^ words[j - 14] ^ words[j - count_parts];
                        words[j] = (n << 1) | (n >>> 31);
                    }
                    var f, k;
                    if (j < 20) {
                        f = (h1 & h2 | ~h1 & h3);
                        k = 1518500249;
                    } else if (j < 40) {
                        f = (h1 ^ h2 ^ h3);
                        k = 1859775393;
                    } else if (j < 60) {
                        f = (h1 & h2 | h1 & h3 | h2 & h3);
                        k = -1894007588;
                    } else {
                        f = (h1 ^ h2 ^ h3);
                        k = -899497514;
                    }
                    var t = ((h0 << 5) | (h0 >>> 27)) + h4 + (words[j] >>> 0) + f + k;
                    h4 = h3;
                    h3 = h2;
                    h2 = (h1 << 30) | (h1 >>> 2);
                    h1 = h0;
                    h0 = t;
                }
                h0 = (h0 + th0) | 0;
                h1 = (h1 + th1) | 0;
                h2 = (h2 + th2) | 0;
                h3 = (h3 + th3) | 0;
                h4 = (h4 + th4) | 0;
            }
            return [h0, h1, h2, h3, h4];
        }
        /**
         * 通过字符数组对应的Unicode数组得到相应的hash值
         *
         * @param (Uint8Array){charCodeArr}字符数组对应的Unicode数组，例如：[97,98,99]
         * @param (Function){method}计算hash所用的加密方法，如sha1等
         * @param (Number)buffer分块处理的块大小
         * @return (Array){hash}hash数组
         */
    var getHash = function(charCodeArr, method, buffer) {
            (buffer && buffer > 0) || (buffer = 1024 * 16 * 64); // 处理默认的buffer
            var hash = [1732584193, -271733879, -1732584194, 271733878, -1009589776];
            var len = charCodeArr.length;
            var start = 0;
            var end = buffer - 1;
            var block;
            while (start + 1 <= len) {
                end = Math.min(end, len - 1);
                block = util.bytesToWords(charCodeArr.subarray(start, end + 1));
                if (end === len - 1) {
                    var bTotal, bLeft, bTotalH, bTotalL;
                    bTotal = len * 8;
                    bLeft = (end - start + 1) * 8;
                    bTotalH = Math.floor(bTotal / 0x100000000);
                    bTotalL = bTotal & 0xFFFFFFFF;
                    // Padding
                    block[bLeft >>> 5] |= 0x80 << (24 - bLeft % 32);
                    block[((bLeft + 64 >>> 9) << 4) + 14] = bTotalH;
                    block[((bLeft + 64 >>> 9) << 4) + 15] = bTotalL;
                }
                hash = method(block, hash);
                start += buffer;
                end += buffer;
            }
            return hash;
        }
        /**
         * 通过字符数组对应的Unicode数组得到sha1值
         *
         * @param (Uint8Array){charCodeArr}字符数组对应的Unicode数组，例如：[97,98,99]
         * @param (Number){buffer}分块处理的块大小
         * @return (String){sha1}sha1值
         */
    var getSha1 = function(charCodeArr, buffer) {
        var hash = getHash(charCodeArr, sha1, buffer);
        return util.bytesToHex(hash);
    }

    /**
     * 获取字符串的sha1
     *
     * @param (String){str}字符串
     * @return (String){sha1}字符串对应的sha1值
     */
    var getSha1FromStr = function(str) {
            var len = str.length;
            var b = [len];
            for (var i = 0; i < len; i++) {
                b[i] = str.charCodeAt(i);
            }
            return getSha1(new Uint8Array(b), 10240);
        }
        /**
         * 获取文件的sha1
         *
         * @param (File){file}文件对象
         * @return (String){sha1}sha1值
         */
    var getSha1FromFile = function(file, callback) {
        var reader = new FileReader();
        reader.readAsArrayBuffer(file);
        reader.onload = function(e) {
            var u8Arr = new Uint8Array(e.target.result);
            var sha1 = jQuery.hash.getSha1(u8Arr, 10240);
            file.sha1 = sha1;
            callback(e, sha1);
        }
    }


    /**
     * 切割大文件
     *
     * @param (Uint8Array){file}file
     * @param (Number){cutSize}分块处理的块大小
     * @return (Uint8Array){fileArr}fileArr
     */
    var cutFile = function(file, cutSize) {
        var count = file.size / cutSize | 0,
            fileArr = [];
        for (var i = 0; i < count; i++) {
            fileArr.push({
                name: file.name + ".part" + (i + 1),
                file: file.slice(cutSize * i, cutSize * (i + 1))
            });
        };
        fileArr.push({
            name: file.name + ".part" + (count + 1),
            file: file.slice(cutSize * count, file.size)
        });
        return fileArr;
    }


    /*
     *   ajax大文件切割上传(支持单个文件)  -- level2的新特性，请保证你的项目支持新的特性再使用
     *       url                 文件上传地址
     *       fileSelector        input=file 选择器
     *       cutSize             切割文件大小
     *       fileType            文件限制类型 mime类型
     *       successEvent        上传成功处理
     *       progressEvent       上传进度事件
     *       errorEvent          上传失败处理
     *       timeoutEvent        超时处理事件
     *
     *   return: status:  0      请选择文件
     *                    1      非允许文件格式
     * */
    upload_big = function(url, fileSelector, cutSize, fileType, successEvent, progressEvent, errorEvent, timeoutEvent) {
        var file = document.querySelector(fileSelector).files,
            result = {};
        //以下为上传文件限制检查
        if (file.length === 1) {
            if (fileType != "*") {
                if (fileType.indexOf(file.type) === -1) {
                    result["status"] = 1;
                    result["errMsg"] = "非允许文件格式";
                }
            }
        } else {
            result["status"] = 0;
            result["errMsg"] = "请选择文件/只能上传一个文件";
        };

        if (result.status !== undefined) return result; //如果有错误信息直接抛出去,结束运行

        //判断上传文件是否超过需要切割的大小
        if (file[0].size > cutSize) {
            var fileArr = $.hash.cutFile(file[0], cutSize); //切割文件
            cutFile_upload(fileArr);
        } else {
            return tempObj.upload(url, fileSelector, file[0].size, fileType, successEvent, errorEvent, timeoutEvent);
        };

        /*
         *   切割文件上传，配合后台接口进行对接
         *       传输参数：
         *           count   -- 当前传输part的次数
         *           name    -- 做过处理的文件名称
         *           file    -- 上传的.part的切割文件
         *           isLast  -- 是否为最后一次切割文件上传（默认值："true"  字符串，只有最后一次才附加）
         * */
        function cutFile_upload(fileArr, count) {
            var formData = new FormData();
            if (count == undefined) {
                count = 0;
                formData.append("count", count);
                formData.append("name", fileArr[0].name);
                formData.append("file".name, fileArr[0].file);
            } else {
                if (count === fileArr.length - 1) {
                    formData.append("isLast", "true")
                };
                formData.append("count", count);
                formData.append("name", fileArr[count].name);
                formData.append("file".name, fileArr[count].file);
            };
            var ajaxParam = {
                type: "post",
                url: url,
                data: formData,
                isFormData: true,
                success: function(data) {
                    /*
                     *   data 参数设置  需要后台接口配合
                     *       建议：如果后台成功保存.part文件，建议返回下次所需要的部分，比如当前发送count为0，则data返回下次为1。
                     *             如果保存不成功，则可false，或者返回错误信息，可在successEvent中处理
                     *
                     * */
                    progressEvent(count + 1, fileArr.length); //上传进度事件，第一个参数：当前上传次数；第二个参数：总共文件数

                    var currCount = Number(data);
                    if (currCount) {
                        if (currCount != fileArr.length) {
                            cutFile_upload(fileArr, currCount);
                        };
                    };
                    successEvent(data); //成功处理事件
                },
                error: errorEvent,
                timeout: timeoutEvent
            };
            $.param(ajaxParam)
                // ajax.common(ajaxParam);
        }
    }

    $.hash = {
        util: util,
        getSha1: getSha1,
        getSha1FromStr: getSha1FromStr,
        getSha1FromFile: getSha1FromFile,
        cutFile: cutFile,
        upload_big: upload_big
    };
})(jQuery);