function showSnackbar(mesg) {
    var x = document.getElementById("snackbar");
    x.innerHTML = mesg;
    x.className = "show";
    setTimeout(function () {
        x.className = x.className.replace("show", "");
    }, 3000);
}

function randomColor() {
    var classes = ["category-green", "category-blue", "category-purple", "category-red"];

    var cates = document.getElementsByClassName('category');
    Array.prototype.filter.call(cates, function (cate) {
        cate.classList.add(classes[~~(Math.random() * classes.length)])
    });
}