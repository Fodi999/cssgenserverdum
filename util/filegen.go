package util

import (
    "fmt"
    "os"
    "path/filepath"
)

// CreateHTMLFile создает HTML файл с заданным именем и содержимым.
func CreateHTMLFile(fileName, content string) error {
    dir := filepath.Dir(fileName)
    err := os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating directory %v: %v", dir, err)
    }

    _, err = os.Stat(fileName)
    if os.IsNotExist(err) {
        file, err := os.Create(fileName)
        if err != nil {
            return fmt.Errorf("error creating HTML file: %v", err)
        }
        defer file.Close()

        _, err = file.WriteString(content)
        if err != nil {
            return fmt.Errorf("error writing to HTML file: %v", err)
        }
    } else if err == nil {
        // Файл уже существует, не возвращаем ошибку, просто выходим из функции
        return nil
    } else {
        return fmt.Errorf("error checking if file exists: %v", err)
    }
    return nil
}

// CreateJSFile создает JS файл с заданным именем и содержимым.
func CreateJSFile(fileName, content string) error {
    dir := filepath.Dir(fileName)
    err := os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating directory %v: %v", dir, err)
    }

    _, err = os.Stat(fileName)
    if os.IsNotExist(err) {
        file, err := os.Create(fileName)
        if err != nil {
            return fmt.Errorf("error creating JS file: %v", err)
        }
        defer file.Close()

        _, err = file.WriteString(content)
        if err != nil {
            return fmt.Errorf("error writing to JS file: %v", err)
        }
    } else if err == nil {
        // Файл уже существует, не возвращаем ошибку, просто выходим из функции
        return nil
    } else {
        return fmt.Errorf("error checking if file exists: %v", err)
    }
    return nil
}

// CreateDefaultHTMLFile создает стандартный HTML файл hello.html, если он не существует.
func CreateDefaultHTMLFile() error {
    content := `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="description" content="Default description">
        <link href="/static/css/style.css" rel="stylesheet">
        <title>Hello</title>
        <meta name="theme-color" content="#007bff">
    </head>
    <body class="bg-gray-100">
        <nav class="bg-blue-600 p-4 text-white">
            <a href="/" class="mr-4">Home</a>
            <a href="/about" class="mr-4">About</a>
            <a href="/contact" class="mr-4">Contact</a>
            <a href="/user" class="mr-4">User</a>
        </nav>
        <div class="mx-auto mt-4 p-4 bg-white shadow rounded w-96">
            <h1 class="text-2xl font-bold mb-4">WebSocket Chat</h1>
            <div id="messages" class="border border-gray-300 rounded p-4 h-64 overflow-y-auto bg-gray-50 mb-4"></div>
            <form class="flex mb-4">
                <input type="text" id="messageInput" class="border border-gray-300 rounded p-2 flex-grow" placeholder="Type your message...">
                <button type="submit" class="bg-blue-600 text-white p-2 rounded ml-2">Send</button>
            </form>
            <h2 class="text-xl font-bold mb-2">Upload Image</h2>
            <form id="uploadForm" action="/upload" method="post" enctype="multipart/form-data" class="flex flex-col">
                <input type="file" name="file" class="mb-2">
                <button type="submit" class="bg-blue-600 text-white p-2 rounded">Upload</button>
            </form>
        </div>
        <script src="/static/js/script.js"></script>
        <script src="/static/js/checkStyles.js"></script>
    </body>
</html>`
    return CreateHTMLFile("templates/hello.html", content)
}

// CreateDefaultJSFile создает стандартный JS файл script.js, если он не существует.
func CreateDefaultJSFile() error {
    content := `document.addEventListener("DOMContentLoaded", function() {
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
`
    return CreateJSFile("static/js/script.js", content)
}




