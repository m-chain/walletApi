(function() {
	var pluginName="imagedef";
    CKEDITOR.dialog.add(pluginName, 
    function(a) {  
        return {
            title: "多图上传",
            minWidth: 620,
            minHeight:100,
            buttons : [ CKEDITOR.dialog.okButton, CKEDITOR.dialog.cancelButton ],
            contents: [{
                id: "tab1",
                label: "",
                title: "上传图片",
                expand: true,
                padding: 0,
                elements: [{
                    type: "html",
                    style: "width:500px;height:500px",
                    html: uploadImageHtml(a)
                }]
            }],
            onOk: function() {
                //点击确定按钮后的操作               
            	var html="";
            	var bl = false;
            	var curObj = $("#localUpload",$(".cke_editor_"+a.name+"_dialog"));
            	
            	if($(".success",$("li",curObj)).size()==0)
            		bl=true;
            	$("li:not(:last)",curObj).each(function(i){
            		if($(".success",$(this)).length==0) {
            			bl = true;            			
            		}
            	});
            	if(!bl){
	            	$("img",curObj).each(function(){            		
	            		html+='<img src="'+$(this).attr("_src")+'" /><br/>';
	            	});	
	            	a.insertHtml(html);
            	}else{
            		alert("您有图片未上传，请先上传图片！");
            		return false;
            	}
            },
            onShow:function(){
            	$(".cke_dialog_ui_vbox_child").html(uploadImageHtml(a));
            }
        }
    })    
})();

/*$(document).ready(function(){
	$("progress").live("");
});*/

//存放图片的路径 

function uploadImageHtml(currentCkeditor){	
	//有多个时去掉最早的那个上传图片窗口
	if($(".cke_editor_ExpertTravelsContent_dialog").length>1){
		$(".cke_editor_ExpertTravelsContent_dialog")[0].remove();
	}
	var html = '';
	html+='<style type="text/css">';
	html+='ul, li {';
	html+='	margin: 0;';
	html+='	padding: 0;';
	html+='}';
	html+='.imageUpload_wrapper_area{';
	html+='	z-index: 10;';
	html+='	min-height: 100%;';
	html+='}';
	html+='.imageUpload_wrapper_area #tab_menu {';
	html+='	border-bottom: 0px solid #ccc;';
	html+='}';
	html+='.imageUpload_wrapper_area #tab_menu li {';
	html+='	list-style-type: none;';
	html+='	padding: 5px;';
	html+='	color: #646464;';
	html+='	display: inline-block;';
	html+='	height: 25px;';
	html+='	line-height: 25px;';
	html+='	font-size: 14px;';
	html+='	border: 1px solid #CCC;';
	html+='	border-bottom: 0;';
	html+='	z-index: 100;';
	html+='	cursor: pointer;';
	html+='	position: relative;';
	html+='	top: 1px;';
	html+='}';
	html+='.imageUpload_wrapper_area #tab_menu li.selected {';
	html+='	border-bottom: 1px solid #fff;';
	html+='	background-color: #fdfdfd;';
	html+='}';
	html+='.imageUpload_wrapper_area .imageUpload_context{';
	html+='	border: 1px solid #ccc;';
	html+='	padding-bottom:10px;';
	html+='}';
	html+='';
	html+='.imageUpload_context #localUpload{';
	html+='	min-height: 400px;';
	html+='	height:auto !important;';
	html+='	height:400px;';
	html+='	min-height:400px;';
	html+='}';
	html+='';
	html+='';
	html+='.webuploader-element-invisible {';
	html+='	position: absolute !important;';
	html+='	clip: rect(1px,1px,1px,1px);';
	html+='}';
	html+='.imageUpload_context #localUpload li.filePickerBlock {';
	html+='	width: 113px;';
	html+='	height: 113px;';
	html+='	background: url(../../images/image.png) no-repeat center 12px;';
	html+='	border: 1px solid #eeeeee;';
	html+='	border-radius: 0;';
	html+='}';
	html+='';
	html+='.webuploader-continue-pick{';
	html+='	font-size:12px;';
	html+='	background: #fff;';
	html+='	border-radius: 3px;';
	html+='	line-height: 30px;';
	html+='	padding: 0 20px;';
	html+='	display: inline-block;';
	html+='	margin: 0 auto;';
	html+='	cursor: pointer;';
	html+='	border:1px solid #eee;';
	html+='	box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);';
	html+='	color:#646464';
	html+='}';
	html+='.webuploader-continue-pick-hover{';
	html+='	background:#cff;';
	html+='	border-color:#cfc;';
	html+='}';
	html+='.webuploader-pick {';
	html+='	font-size: 18px;';
	html+='	background: #00b7ee;';
	html+='	border-radius: 3px;';
	html+='	line-height: 44px;';
	html+='	padding: 0 30px;';
	html+='	color: #fff;';
	html+='	display: inline-block;';
	html+='	margin: 0 auto;';
	html+='	cursor: pointer;';
	html+='	box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);';
	html+='}';
	html+='.webuploader-pick-hover {';
	html+='	background: #00a2d4;';
	html+='}';
	html+='.webuploader-container {';
	html+='	position: relative;';
	html+='	padding-top: 150px;';
	html+='}';
	html+='.placeholder {';
	html+='	margin: 10px;';
	html+='	border: 2px dashed #e6e6e6;';
	html+='	min-height: 400px;';
	html+='	text-align: center;';
	html+='	background: url(../../images/image.png) center 70px no-repeat;';
	html+='}';
	html+='.listPic {';
	html+='	padding: 5px;';
	html+='}';
	html+='.listPic li {';
	html+='	float: left;';
	html+='	margin: 5px 0 0 5px;';
	html+='	list-style-type: none;';
	html+='}';
	html+='.listPic li span.success{';
	html+='	display: block;';
	html+='	position: absolute;';
	html+='	left: 0;';
	html+='	bottom: 0;';
	html+='	height: 40px;';
	html+='	width: 100%;';
	html+='	z-index: 200;';
	html+='	background: url(../../images/success.png) no-repeat right bottom;';
	html+='}';
	html+='.listPic li p.error{';
	html+='	background: #f43838;';
	html+='	color: #fff;';
	html+='	position: absolute;';
	html+='	bottom: 0;';
	html+='	left: 0;';
	html+='	height: 28px;';
	html+='	line-height: 28px;';
	html+='	z-index: 100;';
	html+='}';
	html+='.hander {';
	html+='	cursor: pointer;';
	html+='}';
	html+='.upload-file-panel {';
	html+='	position: absolute;';
	html+='	height: 0;';
	html+=' filter: progid:DXImageTransform.Microsoft.gradient(GradientType=0, startColorstr=\'#80000000\', endColorstr=\'#80000000\') \\0;';
	html+='	background: rgba(0, 0, 0, 0.5);';
	html+='	width: 113px;';
	html+='	left: 0;';
	html+='	overflow: hidden;';
	html+='	z-index: 300;';
	html+='	height:30px;';
	html+='	display:none;';
	html+='}';
	html+='.upload-file-panel span {';
	html+='	width: 24px;';
	html+='	height: 24px;';
	html+='	display: inline;';
	html+='	float: right;';
	html+='	text-indent: -9999px;';
	html+='	overflow: hidden;';
	html+='	background: url(../../images/icons.png) no-repeat;';
	html+='	background: url(../../images/icons.gif) no-repeat \\9;';
	html+='	margin: 5px 1px 1px;';
	html+='	cursor: pointer;';
	html+='}';
	html+='.upload-file-panel span.cancel {';
	html+='	background-position: -49px -24px;';
	html+='}';
	html+='.upload-file-panel span.cancelover {';
	html+='	background-position: -49px 0;';
	html+='}';
	html+='.progress{';
	html+='	display: block;';
	html+='	position: absolute;';
	html+='	left: 0;';
	html+='	bottom: 0;';
	html+='	height:15px;';
	html+='	z-index: 200;';
	html+=' width:113px';
	html+='	display:none;';
	html+='}';
	html+='.uploadBtn{';
	html+='	font-size: 12px;';
	html+='	background: #00b7ee;';
	html+='	border-radius: 3px;';
	html+='	line-height: 30px;';
	html+='	padding: 0 30px;';
	html+='	color: #fff;';
	html+='	display: inline-block;';
	html+='	margin: 0 auto;';
	html+='	cursor: pointer;';
	html+='	box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);';
	html+='}';
	html+='.uploadBtn.disabled{';
	html+='	pointer-events: none;';
	html+='	filter: alpha(opacity=60);';
	html+='	-moz-opacity: 0.6;';
	html+='	-khtml-opacity: 0.6;';
	html+='	opacity: 0.6;';
	html+='}';
	html+='.uploadBtn-hover{';
	html+='	background:#00a2d4;';
	html+='}';
	html+='.zyh_statusBar{';
	html+='	height:40px;';
	html+='	border-bottom:1px solid #ccc;';
	html+='	margin:7px;';
	html+='	color:#646464;';
	html+='}';
	html+='.zyh_statusBar .info{';
	html+='	display:inline-block;';
	html+=' float:left;';
	html+='	height:40px;';
	html+='	line-height:40px;';
	html+=' color:#646464;';
	html+='}';
	html+='.zyh_statusBar .btns-area{';
	html+='	display:inline-block;';
	html+='	float:right;';
	html+='}';
	html+='.none {';
	html+='	display: none;';
	html+='}';
	html+='</style>';

	html+='<div class="imageUpload_wrapper_area">';
	html+='  <ul id="tab_menu">';
	html+='    <li onclick="tab(this)" class="selected">本地上传</li>';
	//html+='    <li onclick="tab(this)">在线管理</li>';
	html+='  </ul>';
	html+='  <div class="imageUpload_context" id="imageUpload_context">';
	html+='<input type="hidden" id="uploadImgType" value="travels"/>';
	html+='    <div id="localUpload">';
	html+='      <div class="placeholder">';
	html+='        <div id="filePickerReady" class="webuploader-container" style="text-align:center">';
	html+='          <div class="webuploader-pick">点击选择图片<span style="display:none">'+currentCkeditor.name+'</span></div>';
	html+='          <div>';
	html+='              <input type="file" name="upload_file" id="upload_file" class="webuploader-element-invisible" multiple="multiple" accept="image/*" />';
	html+='          </div>';
	html+='        </div>';
	html+='      </div>';
	html+='    </div>';
	html+='    <div style="clear:both;"></div>';
	html+='    <div id="onlinePic" class="none">这块功能暂时不能使用！</div>';
	html+='    <iframe width="0" height="0" name="uploadImg" frameborder="0" style="width:0;height:0"></iframe>';
	html+='  </div>';
	html+='</div>';

	return html;

}


/**
*  基于jquery的多图片上传js对象

*/
var UploadMultImg=function(){}
/**
 * 初始化绑定属性
 * */
UploadMultImg.prototype.init = function(){
	var curCkeditorDialog = null;
	var topObj = this;
	//点击选择图片鼠标移过事件
	$(".webuploader-pick").hover(
		function(e){
			$(this).addClass("webuploader-pick-hover");
		},function(e){
			$(this).removeClass("webuploader-pick-hover");
		}
	);
	$(document).on("click",".webuploader-pick",function(){
		var curDialogClass = $("span",$(this)).text();
		curCkeditorDialog = $(".cke_editor_"+curDialogClass+"_dialog");
		$("input[name='upload_file']",curCkeditorDialog).click();
	});
	//上传图片
	$(document).on("change","input[name='upload_file']",function(r){	
		$("#uploadBtn").removeClass("disabled");
		if(window.FileReader){//HTML5实现预览，兼容chrome、火狐7+等 
			//var filesObj = document.getElementById("upload_file");
			var filesObjJquery = $("input[name='upload_file']",curCkeditorDialog);
			var filesObj = filesObjJquery[0];
			var curSize=0;			
			for(var i=0;i<filesObj.files.length;i++){
				var reader = new FileReader();
				var j=0,m=0;					
           	 	reader.onload = function(e){
           	 		var mm=m++;
               	 	var name=filesObj.files[mm].name;
					var size=filesObj.files[mm].size;
					var type=filesObj.files[mm].type;
					curSize+=size;
					topObj.preView({src:e.target.result,length:filesObj.files.length,currNum:++j,name:name,size:size,type:type,ccd:curCkeditorDialog});						
				}					
				reader.readAsDataURL(filesObj.files[i]);					
            }				
			
		}
	});
	
	//绑定增加图片按钮事件
	$(document).on("click","#addPicBtn",function(){
		$("#upload_file",curCkeditorDialog)[0].click();
		//document.getElementById("upload_file").click();
	});
	
	//继续添加鼠标移入移出事件
	$(document).on("mouseover mouseout click",".webuploader-continue-pick",function(e){
		if(e.type=="mouseover"){
			$(this).addClass("webuploader-continue-pick-hover");
		}else if(e.type=="mouseout"){
			$(this).removeClass("webuploader-continue-pick-hover");
		}else if(e.type=="click"){
			$("#addPicBtn",curCkeditorDialog).click();
		}
	});
	
	//上传图片按钮鼠标移入移出事件
	$(document).on("mouseover mouseout",".uploadBtn",function(e){
		if(e.type=="mouseover"){
			$(this).addClass("uploadBtn-hover");
		}else{
			$(this).removeClass("uploadBtn-hover");
		}
	});
	
	//鼠标移入移出事件
	$(document).on("mouseover mouseleave","#listPicDiv>ul>li",function(event){
		if(!$(this).hasClass("filePickerBlock")){ 
			if(event.type== 'mouseover'){					
				topObj.liHoverOp($(this),curCkeditorDialog);					
			}else if(event.type == 'mouseleave'){
				topObj.liOutOp($(this));
			}
		}
	}); 
	
	//上传图片点击事件
	//$("#uploadBtn",curCkeditorDialog).live("click",function(){
	  $(document).on("click","#uploadBtn",function(){	
		if(topObj.CheckPicSize(curCkeditorDialog)) return;
		
		//上传总的进度条		
		var progress = '<div style="background:url(../../images/loadingbar.gif) no-repeat center center;width:120px;height:8px;position:absolute"></div>';
		$("#picStatusInfo",curCkeditorDialog).html(progress);
		
		$("li:not(:last)",$("#localUpload",curCkeditorDialog)).each(function(i){			
			if($(".success",$(this)).length==0){
				//获取窗口相对位置
				var posW = $(".cke_dialog",curCkeditorDialog).position();
				$(".loadingLi",$(this)).css({"top":$(this).offset().top-posW.top+($(this).height()-15),"left":$(this).offset().left-posW.left,"width":$(this).width()}).show();
				//$("progress",$(this)).val(30);
				$("#uploadBtn",curCkeditorDialog).addClass("disabled");
				var m=0;
				(function(obj,imgStr){
					m++;
					$.ajax({
					   async:true,	 
					   type: "POST",
					   url: "../../MultPicUpload_uploadMult",
					   data: "type="+$("#uploadImgType",curCkeditorDialog).val()+"&imgStr="+imgStr,
					   dataType: "json",
					   success:function(msg){
						   if(msg.state=="0"){							  
							  
							  $("img",$(obj)).attr("_src",msg.file);
							  var pos = $(obj).offset();
							  var html='<span class="success"></span>';							  
							  var htmlObj = $(html).css({"top":pos.top-posW.top-3+$(obj).height()-40,"left":pos.left-posW.left,"width":$(obj).width()});
							  $(obj).append($(htmlObj));							  
							  $(".loadingLi",$(obj)).hide();							  
						   }else{
							  var html= '<p class="error" style="display: block;">IO错误</p>';
							  var htmlObj = $(html).css({"top":pos.top+$(obj).height()-28-posW.top-3,"left":pos.left-posW.left,"width":$(obj).width()});
							  $(obj).append($(htmlObj));
							  $(".loadingLi",$(obj)).hide();
						   }
					   }
					});
				 })($(this),$("img",$(this)).attr("src"));
				
			}				
		});
		topObj.uploadPicInfo(curCkeditorDialog);
	});
	
}
/**
 * 检测图片是否超过指定大小
 * */
UploadMultImg.prototype.CheckPicSize = function(ccd){
	var bl = false;
	var picCount =$("li:not(:last)",$("#localUpload",ccd)).length;
	$("li:not(:last)",$("#localUpload",ccd)).each(function(i){
		var info = $("p:first",$(this)).html().split("||");
		if(parseInt(info[1])>=1024*1024*50){
			alert("第 "+(i+1)+" 张图片超过了单张图片50M限制大小！");
			bl = true;
		}		
	});
	return bl;
}
/**上传图片统计信息*/
UploadMultImg.prototype.uploadPicInfo = function(ccd){	
	var totalSize=0;
	var picCount =$("li:not(:last)",$("#localUpload",ccd)).length;
	$("li:not(:last)",$("#localUpload",ccd)).each(function(i){
		var info = $("p:first",$(this)).html().split("||");
		totalSize+=parseInt(info[1]);
		
	});
	var total = "共"+(totalSize/1024).toFixed(2)+"K";
	if((totalSize/1024).toFixed(2)>1024)total = "共"+((totalSize/1024).toFixed(2)/1024).toFixed(2)+"M"
	$("#picStatusInfo",ccd).html("共"+picCount+"张（"+total+"），"+picCount+"张上传成功");
	
}
/**
* 计算当前图片数量及总的大小
*/
UploadMultImg.prototype.countPicInfo = function(ccd){
	var totalSize=0;
	var picCount =$("li:not(:last)",$("#localUpload",ccd)).length;
	$("li:not(:last)",$("#localUpload",ccd)).each(function(i){
		var info = $("p:first",$(this)).html().split("||");
		totalSize+=parseInt(info[1]);
		
	});
	var total = "共"+(totalSize/1024).toFixed(2)+"K";
	if((totalSize/1024).toFixed(2)>1024)total = "共"+((totalSize/1024).toFixed(2)/1024).toFixed(2)+"M"
	$("#picStatusInfo",ccd).html("选中"+picCount+"张图片，"+total);
}
/**
*提交前预览图片
*/
UploadMultImg.prototype.preView = function(obj){
	var ccd = obj.ccd;
	var topObj = this;
	var picInfo = obj.name+"||"+obj.size+"||"+obj.type;
	if($("ul",$("#localUpload",ccd)).length>0){
		//先删除
		if($("#addPicBtn",ccd).length>0){
			$("#addPicBtn",ccd).parent().remove();
		}
		
		var html='<li><p style="display:none">'+picInfo;
		html+='</p><img src="'+obj.src+'" width="113" height="113" style="width:113px;height:113px;">';
		//html+='<progress class="progress" value="0" max="100"></progress>';
		html+= '<div class="loadingLi" style="background:url(../../images/loadingbar.gif) repeat-x center center;width:100%;height:15px; display:none;background-size:115px 8px;position:absolute"></div>';
		html+='</li>';
	    $("ul",$("#localUpload",ccd)).append(html);	
	}else{
		var html = '<div id="listPicDiv" style="max-height:500px;overflow-y:auto">';
			html+= '<ul class="listPic"><li>';
			html+= '<p style="display:none">'+picInfo+'</p>';
			html+='<img src="'+obj.src+'" width="113" height="113" style="width:113px;height:113px;">';
			html+= '<div class="loadingLi" style="background:url(../../images/loadingbar.gif) repeat-x center center;width:100%;height:15px; display:none;background-size:115px 8px;position:absolute"></div>';
			//html+='<progress class="progress" value="0" max="100"></progress>';
			html+='';
			html+='</li>';
			html+="</ul>";
			html+="</div>"
		$("#localUpload",ccd).html(html).prepend(topObj.picStatusAndBtnArea());
		
	}
	//图片区域中的继续添加图片区
	if(obj.length==obj.currNum && $("input[name='upload_file']",ccd).length==0){			
		$("ul",$("#localUpload",ccd)).append('<li class="filePickerBlock hander">'+topObj.addPicBtn()+'</li>');
	}		
	topObj.countPicInfo(ccd);
}
/**
*	鼠标移入li可以选择 删除操作
*/
UploadMultImg.prototype.liHoverOp = function(liObj,ccd){
	var topObj = this;
	if($(".upload-file-panel",$(liObj)).length==0&&$(".success",$(liObj)).length==0){
		var html='<div class="upload-file-panel">';
		html+='<span class="cancel">删除</span></div>';
		var htmlObj = $(html);
		//获取窗口相对位置
		var pos = $(".cke_dialog",ccd).position();
		$($(htmlObj)).css({"top":$(liObj).offset().top-pos.top,"left":$(liObj).offset().left-pos.left});			
		$(liObj).append($(htmlObj));
		$(".cancel",$(liObj)).hover(
			function(){$(this).addClass("cancelover");},
			function(){$(this).removeClass("cancelover");}).on("click",function(){				
			$(liObj).remove();				
			topObj.countPicInfo();
		});
		$(".upload-file-panel",$(liObj)).slideDown();
		
	}
}
new UploadMultImg().init();
/**
*	鼠标移出li删除div
*/
UploadMultImg.prototype.liOutOp = function(liObj){
	$(".upload-file-panel",$(liObj)).slideUp().remove();
}

/**
* 图片状态信息及按钮区
*/
UploadMultImg.prototype.picStatusAndBtnArea = function(){
	var temp = '';
	temp+='<div class="zyh_statusBar">';
	temp+='<input type="hidden" name="savePicSizeTemp" id="savePicSizeTemp" />';
	temp+='  <div class="info" id="picStatusInfo">选中张图片，共K。</div>';
	temp+='  <div class="btns-area">';
	temp+='    <div class="webuploader-continue-pick">继续添加</div>';
	temp+='    <div class="uploadBtn" id="uploadBtn">开始上传</div>';
	temp+='  </div>';
	temp+='</div>';
	return temp;
}
/**加入图片增加按钮*/
UploadMultImg.prototype.addPicBtn = function(){
	var html='<div id="addPicBtn" style="top: 0px; left: 0px;';
	html+=' width: 113px; height: 113px; overflow: hidden; bottom: auto; right: auto;">';
	html+='<input type="file" name="upload_file" id="upload_file" class="webuploader-element-invisible" multiple="multiple" accept="image/*">';
	html+='<label style="opacity: 0; width: 100%; height: 100%; display: block;';
	html+=' cursor: pointer; background: rgb(255, 255, 255);"></label></div>';
	return html;
	
}
function tab(o){
	var lis=document.getElementsByTagName("li");
	for(var i=0;i<lis.length;i++){
		lis[i].className="";
		if(o==lis[i]){
			if(i==0){
				document.getElementById("localUpload").className="";
				document.getElementById("onlinePic").className="none";
			}else if(i==1){
				document.getElementById("localUpload").className="none";
				document.getElementById("onlinePic").className="";
			}
		}
	}
	o.className="selected";
}
