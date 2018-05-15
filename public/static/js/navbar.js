document.addEventListener('DOMContentLoaded', function () {

    var rootEl = document.documentElement;
    var navbarTopEl = document.getElementById('navbar');
    var navbarBottomEl = document.getElementById('navbarBottom');
    var fixBottomEl = document.getElementById('navbarFixBottom');
    if (!!fixBottomEl) {
        var fixBottomElIcon = !!fixBottomEl ? fixBottomEl.querySelector('.fa') : null;
        var fixedBottom = false;

        fixBottomEl.addEventListener('click', function (event) {
            fixedBottom = !fixedBottom;

            if (fixedBottom) {
                fixBottomEl.className = 'button is-success';
                fixBottomElIcon.className = 'fa fa-check-square-o';
                rootEl.classList.add('has-navbar-fixed-bottom');
                navbarBottomEl.classList.remove('is-hidden');
            } else {
                fixBottomEl.className = 'button is-link';
                fixBottomElIcon.className = 'fa fa-square-o';
                rootEl.classList.remove('has-navbar-fixed-bottom');
                navbarBottomEl.classList.add('is-hidden');
            }
        });
    }
});