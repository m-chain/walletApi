/*global Qiniu */
/*global plupload */

/*
 * 上传图片
 */
function UploadOSS(options, callback) {

  var container = null;
  if (isDom(options.containerId)) {
    container = options.containerId;
  } else {
    container = document.getElementById(options.containerId);
  }


  var uploader = Qiniu.uploader({
    disable_statistics_report: false,
    runtimes: 'html5,flash,silverlight,html4',
    browse_button: options.btnSelfile,
    multi_selection: options.multiSel,
    container: container,
    drop_element: container,
    max_file_size: '10mb',
    chunk_size: '4mb',
    flash_swf_url: '/static/js/plupload/Moxie.swf',
    silverlight_xap_url: '/static/js/plupload/Moxie.xap',
    dragdrop: true,
    uptoken_url: $('#uptoken_url').val(),
    // uptoken_func: function(){
    //     var ajax = new XMLHttpRequest();
    //     ajax.open('GET', $('#uptoken_url').val(), false);
    //     ajax.setRequestHeader("If-Modified-Since", "0");
    //     ajax.send();
    //     if (ajax.status === 200) {
    //         var res = JSON.parse(ajax.responseText);
    //         console.log('custom uptoken_func:' + res.uptoken);
    //         return res.uptoken;
    //     } else {
    //         console.log('custom uptoken_func err');
    //         return '';
    //     }
    // },
    domain: $('#domain').val(),
    get_new_uptoken: false,
    //downtoken_url: '/downtoken',
    unique_names: false,
    save_key: false,
    // x_vars: {
    //     'id': '1234',
    //     'time': function(up, file) {
    //         var time = (new Date()).getTime();
    //         // do something with 'time'
    //         return time;
    //     },
    // },
    auto_start: true,
    log_level: 5,
    resize: options.resize,
    filters: {
      mime_types: [ //只允许上传图片
        { title: "Image files", extensions: "jpg,gif,png,bmp,jpeg" }
      ],
      max_file_size: '10mb', //最大只能上传10mb的文件
      prevent_duplicates: true //不允许选取重复文件
    },
    init: {
      'BeforeChunkUpload': function (up, file) {
        console.log("before chunk upload:", file.name);
      },
      'PostInit': function () {
        if (isDom(options.ossfileId))
          options.ossfileId.innerHTML = '';
        else
          document.getElementById(options.ossfileId).innerHTML = '';
      },
      'FilesAdded': function (up, files) {
        var ossFileDom = null;
        if (isDom(options.ossfileId)) {
          ossFileDom = options.ossfileId;
        } else {
          ossFileDom = $('#' + options.ossfileId)[0];
        }
        $(ossFileDom).empty();
        plupload.each(files, function (file) {
          ossFileDom.innerHTML += '<div id="' + file.id + '" style="text-align:center;">' + file.name + ' (' + plupload.formatSize(file.size) + ')<b></b>'
            + '<div class="progress"><div class="progress-bar" style="width: 0%"></div></div>'
            + '</div>';
        });
        $(ossFileDom).show();
        up.start();
      },
      'BeforeUpload': function (up, file) {
        /*
        var suffix = get_suffix(file.name);
        filename = file.name + suffix;
        new_multipart_params = {
          'key': filename,
          'policy': policyBase64,
          'OSSAccessKeyId': accessid,
          'success_action_status': '200', //让服务端返回200,不然，默认会返回204
          'callback': callbackbody,
          'signature': signature,
        };
        up.setOption({
          'url': host,
          'multipart_params': new_multipart_params
        });*/
      },
      'UploadProgress': function (up, file) {
        var d = document.getElementById(file.id);
        d.getElementsByTagName('b')[0].innerHTML = '<span>' + file.percent + "%</span>";
        var prog = d.getElementsByTagName('div')[0];
        var progBar = prog.getElementsByTagName('div')[0]
        progBar.style.width = 2 * file.percent + 'px';
        progBar.setAttribute('aria-valuenow', file.percent);
      },
      'FileUploaded': function (up, file, info) {
        if (info.status == 200) {
          if (isDom(options.ossfileId)) {
            $(options.ossfileId).hide();
          } else {
            $('#' + options.ossfileId).hide();
          }

          if ($.isFunction(callback)) {
              var res = $.parseJSON(info.response);
              callback(res.key);
          }
        } else if (info.status == 203) {
          top.layer.alert("上传到OSS成功，但是oss访问用户设置的上传回调服务器失败!");
        } else {
          top.layer.alert(info.response);
        }
      },
      'Error': function (up, err, errTip) {
        if (err.code == -600) {
          top.layer.alert("选择的图片文件太大了！");
        }
        else if (err.code == -601) {
          top.layer.alert("图片上传失败！");
        }
        else if (err.code == -602) {
          top.layer.alert("图片已经上传过一遍了！");
        }
        else {
          top.layer.alert("图片上传失败！");
        }
      },
      'Key': function (up, file) {
        // 若想在前端对每个文件的key进行个性化处理，可以配置该函数
        // 该配置必须要在 unique_names: false , save_key: false 时才生效

        var suffix = get_suffix(file.name);
        var key = random_string2() + suffix;
        // do something with key
        return key
      }
    }
  });


  //uploader.init();
  uploader.bind('BeforeUpload', function () {
    console.log("hello man, i am going to upload a file");
  });

  uploader.bind('FileUploaded', function () {
    console.log('hello man,a file is uploaded');
  });

  return uploader;
}

function add0(m) {
  return m < 10 ? '0' + m : m;
}

function get_suffix(filename) {
  pos = filename.lastIndexOf('.')
  suffix = ''
  if (pos != -1) {
    suffix = filename.substring(pos)
  }
  return suffix;
}

function random_string2() {
  var time = new Date();
  var y = time.getFullYear();
  var m = time.getMonth() + 1;
  var d = time.getDate();
  var h = time.getHours();
  var mm = time.getMinutes();
  var s = time.getSeconds();
  var ms = time.getMilliseconds();

  return y + add0(m) + add0(d) + add0(h) + add0(mm) + add0(s) + ms;
}

function isDom(obj) {
  var isDOM = (typeof HTMLElement === 'object') ?
    function (obj) {
      return obj instanceof HTMLElement;
    } : function (obj) {
      return obj && typeof obj === 'object' && obj.nodeType === 1 && typeof obj.nodeName === 'string';
    }
  return isDOM(obj);
}