function showSnackbar(mesg) {
    var x = document.getElementById("snackbar");
    x.innerHTML = mesg;
    x.className = "show";
    setTimeout(function () {
        x.className = x.className.replace("show", "");
    }, 3000);
}