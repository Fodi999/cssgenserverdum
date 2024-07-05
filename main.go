package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "os/exec"
    "os/signal"
    "strings"
    "syscall"
    "github.com/Fodi999/cssgenserverdum/util"
)

var quit = make(chan os.Signal, 1)

func main() {
    loadEnvVariables()

    go monitorCommands()

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    fmt.Println("Shutting down server...")
}

func getFreePort() (int, error) {
    addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
    if err != nil {
        return 0, err
    }

    listener, err := net.ListenTCP("tcp", addr)
    if err != nil {
        return 0, err
    }
    defer listener.Close()

    return listener.Addr().(*net.TCPAddr).Port, nil
}

func monitorCommands() {
    reader := bufio.NewReader(os.Stdin)
    for {
        printPrompt()
        command, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("\033[31mError reading command: \033[0m", err)
            continue
        }
        command = strings.TrimSpace(command)
        args := strings.Split(command, " ")

        switch args[0] {
        case "quit":
            fmt.Println("\033[32mShutting down server...\033[0m")
            quit <- syscall.SIGTERM
            return
        case "reload":
            fmt.Println("\033[32mReloading server...\033[0m")
            go watchFiles()
        case "status":
            fmt.Println("\033[32mServer is running...\033[0m")
        case "create":
            handleCreateFiles()
        case "create_site":
            handleCreateSite(reader)
        case "start_site":
            if len(args) > 1 {
                handleStartSite(args[1:])
            } else {
                fmt.Println("\033[31mPlease specify site names to start.\033[0m")
            }
        case "compress_images":
            handleCompressImages()
        default:
            fmt.Println("\033[31mUnknown command: \033[0m", command)
        }
    }
}

func printPrompt() {
    fmt.Print("> ")
}

func handleCreateFiles() {
    // Create JS file
    err := util.CreateDefaultJSFile()
    if err != nil {
        fmt.Println("\033[31mError creating JS file: \033[0m", err)
        return
    }

    // Generate CSS file
    util.GenerateCSS()

    // Create directories
    err = util.CreateImageDirectories()
    if err != nil {
        fmt.Println("\033[31mError creating image directories: \033[0m", err)
        return
    }

    // Create image
    util.ConvertOrCreateImages()

    fmt.Println("\033[32mFiles created successfully.\033[0m")
}

func handleCreateSite(reader *bufio.Reader) {
    fmt.Print("Enter site name: ")
    siteName, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("\033[31mError reading site name: \033[0m", err)
        return
    }
    siteName = strings.TrimSpace(siteName)
    if siteName == "" {
        fmt.Println("\033[31mInvalid site name. Please try again.\033[0m")
        return
    }

    err = util.CreateSite(siteName)
    if err != nil {
        fmt.Println("\033[31mError creating site: \033[0m", err)
        return
    }

    fmt.Printf("\033[32mSite %s created successfully.\033[0m\n", siteName)
}

func handleStartSite(siteNames []string) {
    for _, siteName := range siteNames {
        siteName = strings.TrimSpace(siteName)
        if siteName == "" {
            fmt.Println("\033[31mInvalid site name. Please try again.\033[0m")
            continue
        }

        port, err := getFreePort()
        if err != nil {
            fmt.Printf("Error getting free port for site %s: %v\n", siteName, err)
            continue
        }

        cmd := exec.Command("go", "run", "sites/"+siteName+"/main.go")
        cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err = cmd.Start()
        if err != nil {
            fmt.Printf("Error starting site %s: %v\n", siteName, err)
            continue
        }

        fmt.Printf("Site %s started on http://localhost:%d\n", siteName, port)
    }
}

func handleCompressImages() {
    util.CompressImagesInDirectory("static/img", "static/img_compressed")
    fmt.Println("\033[32mImages compressed successfully.\033[0m")
}

func loadEnvVariables() {
    // Dummy function to emulate loading environment variables if necessary
}

func watchFiles() {
    // Dummy function to emulate watching files for changes and reloading the server
}






















































