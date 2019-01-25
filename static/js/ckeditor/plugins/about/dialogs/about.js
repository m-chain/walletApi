/*
 Copyright (c) 2003-2014, CKSource - Frederico Knabben. All rights reserved.
 For licensing, see LICENSE.md or http://ckeditor.com/license
*/
CKEDITOR.dialog.add("about",
function(a) {
    var a = a.lang.about,
    b = CKEDITOR.plugins.get("about").path + "dialogs/" + (CKEDITOR.env.hidpi ? "hidpi/": "") + "logo_ckeditor.png";
    return {
        title: CKEDITOR.env.ie ? a.dlgTitle: a.title,
        minWidth: 390,
        minHeight: 230,
        contents: [{
            id: "tab1",
            label: "",
            title: "",
            expand: !0,
            padding: 0,
            elements: [{
                type: "html",
                html: ""
            }]
        }],
        buttons: [CKEDITOR.dialog.cancelButton]
    }
});