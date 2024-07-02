package main

import (
    "bufio"
    "fmt"
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

    startServer()

    go monitorCommands()

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    fmt.Println("Shutting down server...")
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
        default:
            fmt.Println("\033[31mUnknown command: \033[0m", command)
        }
    }
}

func printPrompt() {
    fmt.Print("> ")
}

func handleCreateFiles() {
    // Create HTML file
    err := util.CreateDefaultHTMLFile()
    if err != nil {
        fmt.Println("\033[31mError creating HTML file: \033[0m", err)
        return
    }

    // Create JS file
    err = util.CreateDefaultJSFile()
    if err != nil {
        fmt.Println("\033[31mError creating JS file: \033[0m", err)
        return
    }

    // Generate CSS file
    util.GenerateCSS()

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

        port := util.GetSitePort(siteName)
        cmd := exec.Command("go", "run", "sites/"+siteName+"/main.go")
        cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err := cmd.Start()
        if err != nil {
            fmt.Printf("\033[31mError starting site %s: %v\033[0m\n", siteName, err)
            continue
        }

        fmt.Printf("\033[32mSite %s started on http://localhost:%d\033[0m\n", siteName, port)
    }
}

func loadEnvVariables() {
    // Dummy function to emulate loading environment variables if necessary
}

func startServer() {
    // Dummy function to emulate starting the main server
}

func watchFiles() {
    // Dummy function to emulate watching files for changes and reloading the server
}






