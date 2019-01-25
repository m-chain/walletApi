(function() {
	var pluginName="imagedef";
    CKEDITOR.plugins.add(pluginName, {
        requires: ["dialog"],
        init: function(a) {
        	var icoPath=this.path + "dialogs/imagedef.png";
        	var pluginNameJs = this.path + "dialogs/"+pluginName+".js";
            a.addCommand(pluginName, new CKEDITOR.dialogCommand(pluginName));
            a.ui.addButton(pluginName, {
                label: "上传图片",//调用dialog时显示的名称
                command: pluginName,
                icon: icoPath//在toolbar中的图标
            });
            CKEDITOR.dialog.add(pluginName,pluginNameJs);
        }

    })

})();