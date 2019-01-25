/*global Qiniu */
/*global plupload */
/*global FileProgress */
/*global hljs */
function initImgUpload(op) {
  //如果是查看，不需要获取上传的token
  if (op != 'view') {
    var uploader = Qiniu.uploader({
      disable_statistics_report: false,
      runtimes: 'html5,flash,html4',
      browse_button: 'pickfiles',
      container: 'container',
      drop_element: 'container',
      max_file_size: '4mb',
      flash_swf_url: '/static/js/plupload/Moxie.swf',
      dragdrop: true,
      chunk_size: '4mb',
      multi_selection: !(moxie.core.utils.Env.OS.toLowerCase() === "ios"),
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
      filters: {
        mime_types: [ //只允许上传图片和zip文件
          { title: "Image files", extensions: "jpg,jpeg,gif,png,bmp" }
        ],
        max_file_size: '4mb', //最大只能上传400kb的文件
        prevent_duplicates: true //不允许选取重复文件
      },
      init: {
        'BeforeChunkUpload': function (up, file) {
          console.log("before chunk upload:", file.name);
        },
        'FilesAdded': function (up, files) {
          $('table').show();
          //$('#success').hide();
          //文件限制
          plupload.each(files, function (file) {

            if (file.type == 'image/jpeg' || file.type == 'image/jpg' || file.type == 'image/png'
              || file.type == 'image/gif' || file.type == 'image/bmp') {
              isUpload = true;
              var progress = new FileProgress(file,
                'fsUploadProgress');
              progress.setStatus("等待...");
              progress.bindUploadCancel(up);
            } else {
              alert(file.type)
              isUpload = false;
              up.removeFile(file);
              top.layer.alert('上传类型只能是.jpg,.jpeg,.png,.gif,.bmp');
              return false;
            }
          });
        },
        'BeforeUpload': function (up, file) {
          var progress = new FileProgress(file, 'fsUploadProgress');
          var chunk_size = plupload.parseSize(this.getOption(
            'chunk_size'));
          if (up.runtime === 'html5' && chunk_size) {
            progress.setChunkProgess(chunk_size);
          }
        },
        'UploadProgress': function (up, file) {
          var progress = new FileProgress(file, 'fsUploadProgress');
          var chunk_size = plupload.parseSize(this.getOption(
            'chunk_size'));
          progress.setProgress(file.percent + "%", file.speed,
            chunk_size);
        },
        'UploadComplete': function () {
          //$('#success').show();
        },
        'FileUploaded': function (up, file, info) {
          var progress = new FileProgress(file, 'fsUploadProgress');
          progress.setComplete(up, info.response);
        },
        'Error': function (up, err, errTip) {
          $('table').show();
          var progress = new FileProgress(err.file, 'fsUploadProgress');
          if (err.code == -600) {
            progress.bindUploadCancel(up);
          }
          progress.setError();
          progress.setStatus(errTip);
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


    $('#container').on(
      'dragenter',
      function (e) {
        e.preventDefault();
        $('#container').addClass('draging');
        e.stopPropagation();
      }
    ).on('drop', function (e) {
      e.preventDefault();
      $('#container').removeClass('draging');
      e.stopPropagation();
    }).on('dragleave', function (e) {
      e.preventDefault();
      $('#container').removeClass('draging');
      e.stopPropagation();
    }).on('dragover', function (e) {
      e.preventDefault();
      $('#container').addClass('draging');
      e.stopPropagation();
    });
  }
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

function progressHtml(randStrID) {
  var html = [];
  html.push('<img src="" id="img_' + randStrID + '" style="display:none"/>');
  html.push('<div style="height:10px; border:2px solid #09F;margin:5px" id="parent_' + randStrID + '">');
  html.push('<div style="width:0; height:100%; background-color:#09F; text-align:center; line-height:10px; font-size:20px; font-weight:bold;" id="progess_' + randStrID + '"></div>');
  html.push('</div>');
  return html.join("");
}

function initCKEditorUpload(ckeditorObj) {
  var Q2 = new QiniuJsSDK();
  var uploader2 = Q2.uploader({
    disable_statistics_report: false,
    runtimes: 'html5,flash,html4',
    browse_button: 'cke_32',
    container: 'container',
    drop_element: 'container',
    max_file_size: '4mb',
    flash_swf_url: '/static/js/plupload/Moxie.swf',
    dragdrop: true,
    chunk_size: '4mb',
    multi_selection: !(moxie.core.utils.Env.OS.toLowerCase() === "ios"),
    uptoken_url: $('#uptoken_url').val(),
    domain: $('#domain').val(),
    get_new_uptoken: false,
    unique_names: false,
    save_key: false,
    auto_start: true,
    log_level: 5,
    filters: {
      mime_types: [ //只允许上传图片和zip文件
        { title: "Image files", extensions: "jpg,jpeg,gif,png,bmp" }
      ],
      max_file_size: '4mb', //最大只能上传400kb的文件
      prevent_duplicates: true //不允许选取重复文件
    },
    init: {
      'BeforeChunkUpload': function (up, file) {
        console.log("before chunk upload:", file.name);
      },
      'FilesAdded': function (up, files) {
        var htmlAry = [];
        for (var i = 0; i < files.length; i++) {
          var randStrID = files[i].id;
          htmlAry.push(progressHtml(randStrID));

          if (i == files.length - 1) {
            ckeditorObj.insertHtml(htmlAry.join(""));
          }
        }
      },
      'BeforeUpload': function (up, file) {

      },
      'UploadProgress': function (up, file) {
        var randStrID = file.id;
        $("#progess_" + randStrID, $("body", $(ckeditorObj.document.$))).css("width", file.percent + "%").text(file.percent + "%");
      },
      'UploadComplete': function () {
        //所有文件上传完成
      },
      'FileUploaded': function (up, file, info) {
        if (info.status == 200) {
          var randStrID = file.id;
          var domain = up.getOption('domain');
          var res = $.parseJSON(info.response);
          url = domain + encodeURI(res.key);

          ckeditorObj.document.getById("img_" + randStrID)
            .setAttributes({ "src": url, "data-cke-saved-src": url })
            .removeAttributes(["style", "id"]);
          ckeditorObj.document.getById("parent_" + randStrID).remove();
        } else {
          top.layer.alert("status:=" + info.status + " info:=" + info.response);
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

  uploader2.bind('BeforeUpload', function () {
    console.log("hello man, i am going to upload a file");
  });

  uploader2.bind('FileUploaded', function () {
    console.log('hello man,a file is uploaded');
  });
}