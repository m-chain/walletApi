(function() {
	 var todoFun= {
        exec:function(editor){ 
        	btn(editor);
        } 
	};
	var pluginName="insertbgcolor";
    CKEDITOR.plugins.add(pluginName, {
        requires: ["button"],
        init: function(a) {
        	var icoPath=this.path + "dialogs/dopplr.png";
        	var pluginNameJs = this.path + "dialogs/"+pluginName+".js";
            //a.addCommand(pluginName, new CKEDITOR.dialogCommand(pluginName));
        	a.addCommand(pluginName, todoFun);
            a.ui.addButton(pluginName, {
                label: "插入背景",//调用dialog时显示的名称
                command: pluginName,
                icon: icoPath//在toolbar中的图标
            });
        }
    })
})();