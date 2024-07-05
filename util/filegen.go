package util

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "log"
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

// CreateImageDirectories создает директории static/img и static/img_compressed.
func CreateImageDirectories() error {
    dirs := []string{"static/img", "static/img_compressed"}

    for _, dir := range dirs {
        err := os.MkdirAll(dir, os.ModePerm)
        if err != nil {
            return fmt.Errorf("error creating directory %v: %v", dir, err)
        }
    }

    return nil
}

// Проверка и конвертация файла PNG, если он существует
func ConvertOrCreateImages() {
    // Создание директории static/img, если она не существует
    if err := os.MkdirAll("static/img", os.ModePerm); err != nil {
        log.Fatalf("Failed to create directory: %v", err)
    }

    outputFilePath := filepath.Join("static/img", "output.png")

    if _, err := os.Stat(outputFilePath); err == nil {
        log.Println("output.png already exists.")
    } else {
        // Если файла PNG нет, создаем и сохраняем тестовые изображения
        createAndSaveImages(outputFilePath)
    }
}

// Создание и сохранение тестовых изображений
func createAndSaveImages(outputFilePath string) {
    img := createTestImage(100, 100)
    err := savePNG(outputFilePath, img)
    if err != nil {
        log.Fatalf("Failed to save PNG: %v", err)
    }
    log.Println("Images created and saved successfully:", outputFilePath)
}

// Создание тестового изображения с градиентом
func createTestImage(width, height int) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            clr := color.RGBA{uint8(x % 256), uint8(y % 256), uint8((x + y) % 256), 255}
            img.Set(x, y, clr)
        }
    }
    return img
}

// Сохранение изображения в формате PNG
func savePNG(filename string, img image.Image) error {
    outFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer outFile.Close()

    return png.Encode(outFile, img)
}

// Чтение изображения из файла PNG
func loadPNG(filename string) (image.Image, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    img, err := png.Decode(file)
    if err != nil {
        return nil, err
    }
    return img, nil
}

// Уменьшение изображения до указанного максимального размера (по большей стороне)
func resizeImageToMaxSize(img image.Image, maxSize int) image.Image {
    bounds := img.Bounds()
    width := bounds.Dx()
    height := bounds.Dy()

    if width <= maxSize && height <= maxSize {
        return img
    }

    var newWidth, newHeight int
    if width > height {
        newWidth = maxSize
        newHeight = (height * maxSize) / width
    } else {
        newHeight = maxSize
        newWidth = (width * maxSize) / height
    }

    dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
    scaleImage(dst, img)
    return dst
}

// Масштабирование изображения без использования внешних библиотек
func scaleImage(dst *image.RGBA, src image.Image) {
    srcBounds := src.Bounds()
    dstBounds := dst.Bounds()
    xRatio := float64(srcBounds.Dx()) / float64(dstBounds.Dx())
    yRatio := float64(srcBounds.Dy()) / float64(dstBounds.Dy())

    for y := 0; y < dstBounds.Dy(); y++ {
        for x := 0; x < dstBounds.Dx(); x++ {
            srcX := int(float64(x) * xRatio)
            srcY := int(float64(y) * yRatio)
            dst.Set(x, y, src.At(srcX, srcY))
        }
    }
}

// Уменьшение всех изображений в директории и сохранение в новую директорию
func ResizeImagesInDirectory(inputDir, outputDir string, maxSize int) {
    // Создание директории outputDir, если она не существует
    if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
        log.Fatalf("Failed to create directory: %v", err)
    }

    // Перебор всех файлов в inputDir
    err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Пропуск директорий
        if info.IsDir() {
            return nil
        }

        // Только файлы с расширением .png
        if filepath.Ext(path) == ".png" {
            relativePath, err := filepath.Rel(inputDir, path)
            if err != nil {
                return err
            }
            outputFilePath := filepath.Join(outputDir, relativePath)

            img, err := loadPNG(path)
            if err != nil {
                log.Printf("Failed to load PNG: %v", err)
                return nil // Продолжаем, несмотря на ошибку
            }

            resizedImg := resizeImageToMaxSize(img, maxSize)

            err = savePNG(outputFilePath, resizedImg)
            if err != nil {
                log.Printf("Failed to save resized PNG: %v", err)
                return nil // Продолжаем, несмотря на ошибку
            }

            log.Println("Image resized and saved successfully:", outputFilePath)

            // Удаление исходного файла
            err = os.Remove(path)
            if err != nil {
                log.Printf("Failed to remove original file: %v", err)
                return nil // Продолжаем, несмотря на ошибку
            }
        }
        return nil
    })

    if err != nil {
        log.Fatalf("Error walking the path %q: %v", inputDir, err)
    }
}

// CompressImagesInDirectory вызывает ResizeImagesInDirectory с фиксированным размером
func CompressImagesInDirectory(inputDir, outputDir string) {
    ResizeImagesInDirectory(inputDir, outputDir, 2000)
}



























