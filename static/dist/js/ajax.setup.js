$.ajaxSetup({
    async: true,
    beforeSend: function (xhr) {
        // let token = window.localStorage.site_token;
        xhr.setRequestHeader("Accept", "application/json");
        // xhr.setRequestHeader("Content-Type", "application/json");
        // xhr.setRequestHeader('Authorization', 'Bearer ' + token);
    },
    dataFilter: function (data, type) {
        let obj = JSON.parse(data);

        if(obj.code === 401 && obj.msg === "缺失认证token，您需要先登录"){
            showToast(obj.msg, "/member/sign_in.html");
            return false;
        }

        if (obj.code !== 200) {
            if (typeof(errorJump) !== "undefined") {
                showToast(obj.msg, errorJump);
                return false;
            }

            showToast(obj.msg);
            return false;
        }
        return data;
    },
    complete: function (xhr, status) {
        if (typeof xhr.responseJSON !== "undefined" && xhr.responseJSON.hasOwnProperty('code') && xhr.responseJSON.code === 401) {
            showToast(xhr.responseJSON.msg, window.location.origin + "/member/sign_in.html"); //, window.location.origin + "/member/sign_in.html"
        }
    }
});