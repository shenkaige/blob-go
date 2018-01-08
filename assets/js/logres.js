function getNotify() {
    var notify = getParameterByName("notify");
    var elenotify = document.getElementById("logres-notify");
    if (notify == null || notify === '') {
        elenotify.hidden = true;
    } else {
        elenotify.hidden = false;
        elenotify.innerText = notify;
    }
}

function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}