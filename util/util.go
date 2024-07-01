package util

import (
    "fmt"
    "os"
    "path/filepath"
)

// CreateSite создает директорию сайта с HTML, CSS и JS файлами
func CreateSite(name string) error {
    // Создание директории для сайта
    siteDir := filepath.Join("sites", name)
    err := os.MkdirAll(siteDir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating site directory: %v", err)
    }

    // Создание HTML файла
    htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="` + name + ` site">
    <link href="/static/css/style.css" rel="stylesheet">
    <title>` + name + `</title>
</head>
<body>
    <nav>
        <a href="/">Home</a>
    </nav>
    <h1>Welcome to ` + name + `</h1>
    <script src="/static/js/` + name + `.js"></script>
</body>
</html>`
    err = CreateFile(filepath.Join(siteDir, "index.html"), htmlContent)
    if err != nil {
        return err
    }

    // Создание CSS файла
    cssContent := `@dum base;
@dum components;
@dum utilities;`
    err = CreateFile(filepath.Join(siteDir, "static/css/style.css"), cssContent)
    if err != nil {
        return err
    }

    // Создание JS файла
    jsContent := `document.addEventListener("DOMContentLoaded", function() {
    console.log("` + name + ` site loaded");
});`
    err = CreateFile(filepath.Join(siteDir, "static/js/"+name+".js"), jsContent)
    if err != nil {
        return err
    }

    return nil
}

// CreateFile создает файл с заданным содержимым
func CreateFile(filePath, content string) error {
    err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating directory: %v", err)
    }

    file, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("error creating file: %v", err)
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return fmt.Errorf("error writing to file: %v", err)
    }

    return nil
}


