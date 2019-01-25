
CKEDITOR.editorConfig = function( config ) {
	config.height = 400;
	config.extraPlugins="font,imagedef,colorbutton,panelbutton,poi,insertbgcolor,insertCustomImage";
	config.minimumChangeMilliseconds = 100;
	config.uiColor = '#eef5fd';

	config.forcePasteAsPlainText =false;	
	config.filebrowserUploadUrl="actions/Photo_uploadPhotos.do";
	
	config.disableObjectResizing = false;
	config.tabSpaces=4;

	var pathName = window.document.location.pathname;
	var projectName = pathName.substring(0, pathName.substr(1).indexOf('/') + 1);
	config.filebrowserImageUploadUrl = projectName+'/Photo_upload.do';  
	
	config.removeDialogTabs = 'image:advanced;image:Upload;';	
};
