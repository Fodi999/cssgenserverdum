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

// CreateDefaultJSFile создает стандартный JS файл script.js, если он не существует.
func CreateDefaultJSFile() error {
    content := `document.addEventListener("DOMContentLoaded", function() {
    console.log("Hello, World!");
});
`
    return CreateJSFile("static/js/script.js", content)
}






