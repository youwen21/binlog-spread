$.ajaxSetup({
    async: true,
    // beforeSend: function (xhr) {
    //     let token = window.localStorage.site_admin_token;
    //     xhr.setRequestHeader("Accept", "application/json");
    //     // xhr.setRequestHeader("Content-Type", "application/json");
    //     xhr.setRequestHeader('Authorization', 'Bearer ' + token);
    // },
    complete: function (xhr, status) {
        // if (typeof (xhr.responseJSON.code) !== "undefined" && xhr.responseJSON.code === 600) {
        //     showToast(xhr.responseJSON.msg);
        // }
    }
});

function getId() {
    return window.location.hash.substring(1);
}

function getSearchId(name="id") {
    let searchParams = new URLSearchParams(window.location.search);
    let id = searchParams.get(name);
    if (id == null) {
        return 0;
    }

    return id;
}

function genButtonInfo(url = '', text = "text") {
    return '<a class="btn btn-sm btn-info" href="' + url + '">' + text + '</a>';
}

function getButton(url, text, type = 'default', attr = '') {
    return '<a class="btn btn-sm ' + type + '" href="' + url + '" ' + attr + ' >' + text + '</a>';
}

function genButtonView(url = '') {
    return '<a class="btn btn-sm btn-info" href="' + url + '">查看</a>';
}

function genButtonEdit(url = '') {
    return '<a class="btn btn-sm btn-success" href="' + url + '">编辑</a>';
}

function genButtonDel(url = '', restType = 'DELETE') {
    return '<button onclick="ajaxRest(this);" data-rest-method="' + restType + '" class="ajax btn btn-sm btn-danger" href="' + url + '">删除</button>';
}

function genButtonPermission(url = '') {
    return '<a class="btn btn-sm btn-warning" href="' + url + '">有此权限的角色</a>';
}

function genButtonRolePermissions(url = '') {
    return '<a class="btn btn-sm btn-warning" href="' + url + '">授权管理</a>';
}

function genButtonResetPwd(url = '') {
    return '<a class="btn btn-sm btn-warning" href="' + url + '">重置密码</a>';
}

function getButtonRole(url = '') {
    return '<a class="btn btn-sm btn-primary" href="' + url + '">配置角色</a>';
}

function genButtonRoleUsers(url) {
    return '<a class="btn btn-sm btn-primary" href="' + url + '">用户列表</a>';
}

function ajaxRest(element) {
    if (!confirm("确认删除吗？")) {
        return false;
    }

    $.ajax({
        url: $(element).attr('href'),
        method: $(element).attr('data-rest-method'),
        success: function (res) {
            showToast("删除成功", window.location.href);
        }
    })
}

/**
 * @see https://www.cnblogs.com/RobotTech/p/5737370.html
 * @param element
 */
function logout(element) {
    event.preventDefault();
    // event.stopImmediatePropagation();
    $.ajax({
        url: $(element).attr('href'),
        method: $(element).attr('data-rest-method'),
        success: function (res) {
            if (res.code == 200) {
                window.localStorage.site_admin_token = '';
                showToast("退出成功", window.location.origin + "/admin/entrance/login.html");
            }
        }
    })
}

/*
@see https://kamranahmed.info/toast#toasts-events
 */
function showToast(msg, url) {
    if ($.toast) {
        $.toast({
            text: msg,
            showHideTransition: 'slide',  // It can be plain, fade or slide
            bgColor: 'blue',              // Background color for toast
            textColor: '#eee',            // text color
            // allowToastClose : false,       // Show the close button or not
            hideAfter: 3000,              // `false` to make it sticky or time in miliseconds to hide after
            stack: 5,                     // `fakse` to show one stack at a time count showing the number of toasts that can be shown at once
            textAlign: 'left',            // Alignment of text i.e. left, right, center
            position: 'top-center',       // bottom-left or bottom-right or bottom-center or top-left or top-right or top-center or mid-center or an object representing the left, right, top, bottom values to position the toast on page
            afterHidden: function () {
                if (url === 'go-back') {
                    window.history.back(-1);
                } else if (url !== undefined) {
                    window.location.href = url;
                }
            }
        })
    } else {
        alert(msg);
        if(url === "undefined"){
            url = "/admin/entrance/login.html";
        }
        setTimeout(function () {
            window.location.href = url;
        }, 500);
    }

}

function initSideBar() {
    $.ajax({
        url: window.location.origin + "/admin-api/v1/sidebar",
        async: true,
        method: 'GET',
        data: $(".ajax-form").serialize(),
        success: function (res) {
            if (res.code === 401) {
                showToast(res.msg, window.location.origin + "/admin/entrance/login.html");
                return;
            }
            // $.get('/dist/js/sideBar.js', function (template) {
            //     $.tmpl(template, res.data).insertAfter('#sidebar-menu');
            // });
            $('#sidebar-level-1').tmpl(res.data).insertAfter('#sidebar-menu');

            if (typeof window.sideBarPath === "undefined") {
                window.sideBarPath = window.location.pathname;
            }
            $("a[href='" + window.sideBarPath + "']").find("i").addClass("fa-dot-circle");
            $("a[href='" + window.sideBarPath + "']").parents(".nav-treeview").show();
            $("a[href='" + window.sideBarPath + "']").parents(".has-treeview").addClass("menu-open");
        }
    });
}

$(function () {
    if ($('#sidebar-menu').length > 0) {
        initSideBar();
    }
// ajax form
//https://www.cnblogs.com/mg007/p/10145582.html
// https://blog.csdn.net/m0_37505854/article/details/79639046
//方式1
    if ($.fn.ajaxForm) {
        $("#ajaxFormAdd,.ajaxFormAdd").ajaxForm({
            async: true,
            dataType: 'json',
            success: function (res) {
                if (res.code === 200) {
                    showToast(res.msg, 'go-back');
                } else {
                    showToast(res.msg);
                }
            }
        });

        $("#ajaxFormEdit,.ajaxFormEdit").ajaxForm({
            async: true,
            dataType: "JSON",
            beforeSubmit: function (arr, $form, options) {
                options.url += '/' + getId();
            },
            success: function (res) {
                if (res.code === 200) {
                    showToast(res.msg, 'go-back');
                } else {
                    showToast(res.msg);
                }
            }
        });
    }

});


function render(info) {
    Object.keys(info).forEach(function (key) {
        if ($("#" + key).length > 0) {
            console.log($("#" + key));
            let ele = $("#" + key)[0].tagName;
            if (ele === "INPUT") {
                $("#" + key).val(info[key]);
                if ($("#" + key).attr("data-image-preview")) {
                    $("#" + key + '-preview').attr('src', info[key]);
                }

                // if($("#" + key)[0].type==="file"){
                //     $("#" + key+'-text').text(info[key]);
                // }else{
                // }
            } else if (ele === "SELECT") {
                $("#" + key).val(info[key]);
                $("#" + key).change();
            } else if (ele === "TEXTAREA") {
                let editorName = $("#" + key).attr('data-editor');
                if (undefined !== editorName) {
                    $('.' + editorName).summernote('code', info[key])
                } else {
                    $("#" + key).val(info[key]);
                }
            } else if (ele === "IMG") {
                $("#" + key).attr('src', info[key]);
            } else {
                $("#" + key).text(info[key]);
            }
        }
    });
}

function ajaxGetInfo(url) {
    $.ajax({
        url: url,
        method: 'GET',
        async: true, // 这个地方设置成false 会导致编辑器上传图片不好用
        success: function (res) {
            render(res.data)
        }
    });
}

function loadContent(url) {
    ajaxGetInfo(url);
}