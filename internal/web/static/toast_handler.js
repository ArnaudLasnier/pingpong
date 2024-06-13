// Technique taken from: https://blog.benoitblanchon.fr/django-htmx-toasts/

const toastEl = document.getElementById("toast");
const toastHeaderEl = document.getElementById("toast-header");
const toastHeaderTitleEl = document.getElementById("toast-header-title");
const toastBodyEl = document.getElementById("toast-body");

const toastHandler = new bootstrap.Toast(toastEl, { delay: 5000 });

document.body.addEventListener("showError", function(event) {
    toastHeaderEl.classList.add("text-danger", "bg-danger", "bg-opacity-25");
    toastHeaderTitleEl.innerText = event.detail.title;
    toastBodyEl.innerText = event.detail.detail;
    toastHandler.show()
});
