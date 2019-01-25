/**
 * 基于Jquery的 CKeditor多图片上传插件
 * 
 * @Author xiechuanhai by 2017-01-18
 * 
 * 需要传递参数对象如下：
 * {
 * 	url:'',//上传图片url
 *  ckeditorObj:obj, //要插入图片的ckeditor
 * }
 */
;(function($,root){
	'use strict';
	function InsertCustomImageInCkEditor(config){
		this.params = config;
		var _self = this;
		$("body").append("<input id=\"insert_custom_image\" type=\"file\" multiple=\"true\" />");
		var fileObj = $("#insert_custom_image");
		fileObj.click();
		fileObj.on("change",function(event){
			var fileList = fileObj.get(0).files;
			var bl = true;
			for(var i=0;i<fileList.length;i++){
				if(fileList[i].name.toUpperCase().indexOf("JPG")==-1
					&&fileList[i].name.toUpperCase().indexOf("PNG")==-1
					&&fileList[i].name.toUpperCase().indexOf("BMP")==-1){
					bl = false;
					break;
				}
			}
			if(!bl){
				alert("图片上传格式错误</br>请上传 JPG,PNG,BMP 格式的图片");
				return;
			}
			_self.uploadImage();
		});
	}
	/**
	 * 更改CKEditor对象
	 * */
	InsertCustomImageInCkEditor.prototype.setCKEditorObj = function(ckObj){
		this.params.ckeditorObj = ckObj;
	}
	
	InsertCustomImageInCkEditor.prototype.progressHtml = function(randStrID){
		var html = [];
		html.push('<img src="" id="img_'+randStrID+'" style="display:none"/>');
		html.push('<div style="height:10px; border:2px solid #09F;margin:5px" id="parent_'+randStrID+'">');
		html.push('<div style="width:0; height:100%; background-color:#09F; text-align:center; line-height:10px; font-size:20px; font-weight:bold;" id="progess_'+randStrID+'"></div>');
		html.push('</div>');
		return html.join("");
	}
	/**
	 * 图片上传成功后
	 */
	InsertCustomImageInCkEditor.prototype.readystatechange = function(evt,xhr,randStrID){
		if(xhr.readyState == 4 && xhr.status == 200 ){
			var jsonObj = JSON.parse(xhr.responseText);
			var curEditor = this.params.ckeditorObj;
			curEditor.document.getById("img_"+randStrID)
								.setAttributes({"src":jsonObj.path,"data-cke-saved-src":jsonObj.path})
								.removeAttributes(["style","id"]);
			curEditor.document.getById("parent_"+randStrID).remove();
		}
	}
	
	/**
	 * 侦查附件上传情况 ,这个方法大概0.05-0.1秒执行一次
	 */
	InsertCustomImageInCkEditor.prototype.toOnprogress = function(evt,randStrID){
		var loaded = evt.loaded;     //已经上传大小情况 
	 	var tot = evt.total;      //附件总大小 
		var per = Math.floor(100*loaded/tot);//已经上传的百分比 
		var curEditor = this.params.ckeditorObj;
		$("#progess_"+randStrID,$("body",$(curEditor.document.$))).css("width",per+"%").text(per+"%");
	}
	
	InsertCustomImageInCkEditor.prototype.uploadImage = function(){
		// 获取上传文件，放到 formData对象里面
		var files = $("#insert_custom_image").get(0).files;
		var htmlAry = [];
		var _self = this;
		for(var i=0;i<files.length;i++){
			(function(file){
				var randStrID = Math.random().toString().substring(2);
				var formData = new FormData();
				formData.append("file" , file);
				htmlAry.push(_self.progressHtml(randStrID));
				$.ajax({
					type: "POST",
					url: _self.params.url,
					data: formData ,//这里上传的数据使用了formData 对象
					processData : false,
					//必须false才会自动加上正确的Content-Type 
					contentType : false ,
					//这里我们先拿到jQuery产生的 XMLHttpRequest对象，为其增加 progress 事件绑定，然后再返回交给ajax使用
					xhr: function(){
						var xhr = $.ajaxSettings.xhr();
						//if(onprogress && xhr.upload) {
						if(xhr.upload) {
							xhr.upload.addEventListener("progress" ,function(){_self.toOnprogress(event,randStrID)}, false);
							xhr.addEventListener("readystatechange" ,function(){_self.readystatechange(event,xhr,randStrID)}, false);
							return xhr;
						}
					} 	
				});
				if(i==files.length-1){
					_self.params.ckeditorObj.insertHtml(htmlAry.join(""));
					$("#insert_custom_image").remove();
				}
			})(files[i]);
		}
	}
	
	return root.InsertCustomImageInCkEditor = function(config){
		 new InsertCustomImageInCkEditor(config);
	}

})(jQuery,window);