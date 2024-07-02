document.addEventListener("DOMContentLoaded", function() {
    console.log("Hello, World!");

    const messagesDiv = document.getElementById("messages");
    const messageInput = document.getElementById("messageInput");

    document.addEventListener("submit", function(event) {
        event.preventDefault();
        const message = messageInput.value;
        if (message) {
            const messageElement = document.createElement("div");
            messageElement.classList.add("p-2", "my-2", "border-b", "border-gray-200");
            messageElement.textContent = "User: " + message;
            messagesDiv.appendChild(messageElement);
            messagesDiv.scrollTop = messagesDiv.scrollHeight;
            messageInput.value = '';
        }
    });
});
