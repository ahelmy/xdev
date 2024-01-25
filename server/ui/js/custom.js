
const showAlert = (message, type = 'danger') => {
    const alertPlaceholder = document.getElementById('liveAlertPlaceholder')
    const wrapper = document.createElement('div')
    wrapper.innerHTML =
        `<div class="alert-message">
            <div class='alert alert-${type} alert-dismissible fade show shadow' role="alert">
                ${message}
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
            </div>
        </div>`;
    alertPlaceholder.innerHTML = '';
    alertPlaceholder.append(wrapper);
}

const copyToClipboard = (input, copyButton) => {
    input.focus();
    input.select();

    navigator.clipboard
        .writeText(input.value)
        .then(() => {
            // Show tooltip for a brief period
            copyButton.text("Copied!");
            copyButton.tooltip("show");

            // Hide tooltip after 1.5 seconds (adjust as needed)
            setTimeout(function () {
                copyButton.tooltip("hide");
                copyButton.text("Copy!");
            }, 1500);
        })
        .catch(() => {
            alert("something went wrong");
        });
}