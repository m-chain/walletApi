(function() {
	 var todoFun= { 
        exec:function(editor){ 
        	//InsertCustomImageInCkEditor({url:"../../Upload_photo?type=routes",ckeditorObj:editor});
        } 
	};
	var pluginName="insertCustomImage";
    CKEDITOR.plugins.add(pluginName, {
        requires: ["button"],
        init: function(a) {
        	var icoPath=this.path + "insertImage.png";
        	a.addCommand(pluginName, todoFun);
            a.ui.addButton(pluginName, {
                label: "上传图片(按ctrl键可多选)",//调用dialog时显示的名称
                command: pluginName,
                icon: icoPath//在toolbar中的图标
            });
        }

    })

})();