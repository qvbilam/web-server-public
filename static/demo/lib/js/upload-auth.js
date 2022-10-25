const fileServer = "http://127.0.0.1:9501";
const CertificateRouter = "/v1/video/upload/certificate";
const FileStatusSuccess = "success"
let FileInfo = {}

//兼容IE11
if (!FileReader.prototype.readAsBinaryString) {
    FileReader.prototype.readAsBinaryString = function (fileData) {
        var binary = "";
        var pt = this;
        var reader = new FileReader();
        reader.onload = function (e) {
            var bytes = new Uint8Array(reader.result);
            var length = bytes.byteLength;
            for (var i = 0; i < length; i++) {
                binary += String.fromCharCode(bytes[i]);
            }
            //pt.result  - readonly so assign binary
            pt.content = binary;
            pt.onload()
        }
        reader.readAsArrayBuffer(fileData);
    }
}
$(document).ready(function () {
    /**
     * 请求凭证和上传地址
     */
    const getAddress=(params,type,uploadInfo)=>{
        var createUrl = fileServer + CertificateRouter;
        $.ajax({
            type: type,
            url: createUrl,
            data:JSON.stringify(params),
            contentType: "application/json;charset=UTF-8",
            dataType: "json",
            success: function (res) {
                if(res.data?.certificate){
                    const data=res.data?.certificate;
                    var uploadAuth = data.UploadAuth
                    var uploadAddress = data.UploadAddress
                    var videoId = data.VideoId
                    uploader.setUploadAuthAndAddress(uploadInfo, uploadAuth, uploadAddress,videoId)
                    FileInfo = {
                        id: res.data.id,
                    }
                }
            },
            error: function (err) {
                if(err.status===302){
                    const res=JSON.parse(err.responseText).errors;
                    if (res.status === FileStatusSuccess){ // 文件存在并且已上传成功
                        $('#status').text(`文件已存在!ID:${res.id}, BID:${res.business_id}`)
                    }else{
                        // 刷新凭证并且上传
                        getAddress({business_id:res.business_id},'put',uploadInfo);
                    }
                }
            }
        });
    }
    /**
     * 创建一个上传对象
     * 使用 UploadAuth 上传方式
     */
    function createUploader () {
        var u = $('#userId').val()
        var uploader = new AliyunUpload.Vod({
            timeout: $('#timeout').val() || 60000,
            partSize: $('#partSize').val() || 1048576,
            parallel: $('#parallel').val() || 5,
            retryCount: $('#retryCount').val() || 3,
            retryDuration: $('#retryDuration').val() || 2,
            region: $('#region').val(),
            userId: $('#userId').val(),
            // 添加文件成功
            addFileSuccess: function (uploadInfo) {
                console.log('addFileSuccess')
                $('#authUpload').attr('disabled', false)
                $('#resumeUpload').attr('disabled', false)
                $('#status').text('添加文件成功, 等待上传...')
                console.log("addFileSuccess: " + uploadInfo.file.name)
            },
            // 开始上传
            onUploadstarted: function (uploadInfo) {
                if (!uploadInfo.videoId) {
                    const params={
                        file_name:uploadInfo.file.name,
                        // file_sha1: uploadInfo.file.sha1.toString()+(new Date()).getTime() ,
                        file_sha1: uploadInfo.file.sha1.toString(),
                        file_type: uploadInfo.file.type,
                    }
                    getAddress(params,'post',uploadInfo);
                    $('#status').text('文件开始上传...')
                    console.log("onUploadStarted:" + uploadInfo.file.name + ", endpoint:" + uploadInfo.endpoint + ", bucket:" + uploadInfo.bucket + ", object:" + uploadInfo.object)
                } else {
                    console.log("refresh!")
                    getAddress({business_id:uploadInfo.videoId},'put',uploadInfo);
                }
            },
            // 文件上传成功
            onUploadSucceed: function (uploadInfo) {
                console.log("onUploadSucceed: " + uploadInfo.file.name + ", endpoint:" + uploadInfo.endpoint + ", bucket:" + uploadInfo.bucket + ", object:" + uploadInfo.object)
                $('#status').text(`文件上传成功!ID:${FileInfo.id}`)
            },
            // 文件上传失败
            onUploadFailed: function (uploadInfo, code, message) {
                console.log("onUploadFailed: file:" + uploadInfo.file.name + ",code:" + code + ", message:" + message)
                $('#status').text('文件上传失败!')
            },
            // 取消文件上传
            onUploadCanceled: function (uploadInfo, code, message) {
                console.log("Canceled file: " + uploadInfo.file.name + ", code: " + code + ", message:" + message)
                $('#status').text('文件上传已暂停!')
            },
            // 文件上传进度，单位：字节, 可以在这个函数中拿到上传进度并显示在页面上
            onUploadProgress: function (uploadInfo, totalSize, progress) {
                console.log("onUploadProgress:file:" + uploadInfo.file.name + ", fileSize:" + totalSize + ", percent:" + Math.ceil(progress * 100) + "%")
                var progressPercent = Math.ceil(progress * 100)
                $('#auth-progress').text(progressPercent)
                $('#status').text('文件上传中...')
            },
            // 上传凭证超时
            onUploadTokenExpired: function (uploadInfo) {
                // 上传大文件超时, 如果是上传方式一即根据 UploadAuth 上传时
                // 需要根据 uploadInfo.videoId 调用刷新视频上传凭证接口(https://help.aliyun.com/document_detail/55408.html)重新获取 UploadAuth
                // 然后调用 resumeUploadWithAuth 方法, 这里是测试接口, 所以我直接获取了 UploadAuth
                $('#status').text('文件上传超时!')

                getAddress({business_id:uploadInfo.videoId},'put',uploadInfo);

            },
            // 全部文件上传结束
            onUploadEnd: function (uploadInfo) {
                // $('#status').text('文件上传完毕!')
                $('#status').text(`文件上传完毕!ID:${FileInfo.id}`)
                console.log("onUploadEnd: uploaded all the files")
            }
        })
        return uploader
    }

    var uploader = null

    $('#fileUpload').on('change', function (e) {
        var file = e.target.files[0]
        if (!file) {
            alert("请先选择需要上传的文件!")
            return
        }
        var Title = file.name

        var userData = '{"Vod":{}}'

        $.hash.getSha1FromFile(file,(e, sha1)=>file.sha1 = sha1);

        if (uploader) {
            uploader.stopUpload()
            $('#auth-progress').text('0')
            $('#status').text("")
        }
        uploader = createUploader();
        // 首先调用 uploader.addFile(event.target.files[i], null, null, null, userData)
        // uploader.addFile(file, null, null, null, userData)
        uploader.addFile(file, null, null, null, userData)
        $('#authUpload').attr('disabled', false)
        $('#pauseUpload').attr('disabled', true)
        $('#resumeUpload').attr('disabled', true)
        // uploader.onUploadstarted(res.data.certificate)
    })


    // 第一种方式 UploadAuth 上传
    $('#authUpload').on('click', function () {
        // 然后调用 startUpload 方法, 开始上传
        if (uploader !== null) {
            uploader.startUpload()
            $('#authUpload').attr('disabled', true)
            $('#pauseUpload').attr('disabled', false)
        }
    })

    // 暂停上传
    $('#pauseUpload').on('click', function () {
        if (uploader !== null) {
            uploader.stopUpload()
            $('#resumeUpload').attr('disabled', false)
            $('#pauseUpload').attr('disabled', true)
        }
    })


    $('#resumeUpload').on('click', function () {
        if (uploader !== null) {
            uploader.startUpload()
            $('#resumeUpload').attr('disabled', true)
            $('#pauseUpload').attr('disabled', false)
        }
    })

})