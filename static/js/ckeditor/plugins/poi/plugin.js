(function() {
	 var todoFun= { 
        exec:function(editor){ 
        	showAllRoute();
        } 
	};
	var pluginName="poi";
    CKEDITOR.plugins.add(pluginName, {
        requires: ["button"],
        init: function(a) {
        	var icoPath=this.path + "dialogs/extract_foreground_objects.png";
        	var pluginNameJs = this.path + "dialogs/"+pluginName+".js";
            //a.addCommand(pluginName, new CKEDITOR.dialogCommand(pluginName));
        	a.addCommand(pluginName, todoFun);
            a.ui.addButton(pluginName, {
                label: "插入指定线路的所有POI图文详情",//调用dialog时显示的名称
                command: pluginName,
                icon: icoPath//在toolbar中的图标
            });
        }

    })

})();