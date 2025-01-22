package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var javaURLs = []string{
	"https://jdk.java.net/java-se-ri/7",
	"https://jdk.java.net/java-se-ri/8-MR6",
	"https://jdk.java.net/java-se-ri/9",
	"https://jdk.java.net/java-se-ri/10",
	"https://jdk.java.net/java-se-ri/11-MR3",
	"https://jdk.java.net/java-se-ri/12",
	"https://jdk.java.net/java-se-ri/13",
	"https://jdk.java.net/java-se-ri/14",
	"https://jdk.java.net/java-se-ri/15",
	"https://jdk.java.net/java-se-ri/16",
	"https://jdk.java.net/java-se-ri/17-MR1",
	"https://jdk.java.net/java-se-ri/18",
	"https://jdk.java.net/java-se-ri/19",
	"https://jdk.java.net/java-se-ri/20",
	"https://jdk.java.net/java-se-ri/21",
	"https://jdk.java.net/java-se-ri/22",
	"https://jdk.java.net/java-se-ri/23",
}

const (
	path = "/usr/lib/jvm"
)

func main() {
	// Ensure the script is run with sufficient permissions
	if os.Geteuid() != 0 {
		fmt.Println("This script requires root permissions to install Java.")
		fmt.Println("Please run the script using 'sudo' or as the root user.")
		os.Exit(1)
	}

	// Ensure the script is running on a Debian-compatible system
	if _, err := os.Stat("/usr/bin/update-alternatives"); os.IsNotExist(err) {
		fmt.Println("This script is designed for Debian-compatible systems.")
		fmt.Println("Please ensure you are using a compatible distribution.")
		os.Exit(1)
	}

	// Create the JVM directory
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("Failed to create JVM directory: %v\n", err)
		os.Exit(1)
	}

	// Show currently installed Java versions
	fmt.Println("Currently installed Java versions:")
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Failed to read JVM directory: %v\n", err)
		os.Exit(1)
	}
	if len(files) == 0 {
		fmt.Println("No Java versions installed.")
	} else {
		for _, file := range files {
			javaBinPath := filepath.Join(path, file.Name(), "bin")
			fmt.Printf("- %s (%s)\n", file.Name(), javaBinPath)
		}
	}

	// Display available versions
	fmt.Println("Available Java versions:")
	for i, url := range javaURLs {
		fmt.Printf("[%d] Java SE %s\n", i+1, strings.Split(url, "/")[4])
	}

	fmt.Print("Enter the number of the version to explore: ")
	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(javaURLs) {
		fmt.Println("Invalid choice.")
		os.Exit(1)
	}

	selectedURL := javaURLs[choice-1]

	// Fetch the selected version page
	resp, err := http.Get(selectedURL)
	if err != nil {
		fmt.Printf("Failed to fetch selected Java version: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	links := []string{}
	linkRegex := regexp.MustCompile(`href="(https?://[^"\s]+\.tar\.gz)"`)

	for scanner.Scan() {
		matches := linkRegex.FindStringSubmatch(scanner.Text())
		if len(matches) > 1 {
			links = append(links, matches[1])
		}
	}

	if len(links) == 0 {
		fmt.Println("No 64-bit Linux tar.gz links found for the selected version.")
		fmt.Println("Here are the available tar.gz links on the site:")
		for _, link := range links {
			fmt.Println("-", link)
		}
		os.Exit(1)
	}

	selectedTarURL := links[0]
	fmt.Printf("Selected URL: %s\n", selectedTarURL)

	// Download the selected tar.gz
	tempFile := "/tmp/java.tar.gz"
	out, err := os.Create(tempFile)
	if err != nil {
		fmt.Printf("Failed to create temporary file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	resp, err = http.Get(selectedTarURL)
	if err != nil {
		fmt.Printf("Failed to download Java: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		fmt.Printf("Failed to save Java archive: %v\n", err)
		os.Exit(1)
	}

	// Extract the tar.gz file
	extractDir := path
	if err := exec.Command("tar", "-xzf", tempFile, "-C", extractDir).Run(); err != nil {
		fmt.Printf("Failed to extract Java archive: %v\n", err)
		os.Exit(1)
	}

	// Configure alternatives
	files, err = os.ReadDir(path)
	if err != nil {
		fmt.Printf("Failed to read JVM directory: %v\n", err)
		os.Exit(1)
	}

	currentDefault := ""
	for _, file := range files {
		javaPath := filepath.Join(path, file.Name(), "bin")
		for _, app := range []string{"java", "javac", "javadoc", "javap"} {
			appPath := filepath.Join(javaPath, app)
			if _, err := os.Stat(appPath); err == nil {
				cmd := exec.Command("update-alternatives", "--install", "/usr/bin/"+app, app, appPath, "1")
				if err := cmd.Run(); err != nil {
					fmt.Printf("Failed to configure alternative for %s: %v\n", app, err)
				}

				// Set the newly installed Java as default
				cmd = exec.Command("update-alternatives", "--set", app, appPath)
				if err := cmd.Run(); err != nil {
					fmt.Printf("Failed to set %s as default: %v\n", app, err)
				} else {
					currentDefault = appPath
				}
			}
		}
	}

	fmt.Println("Java installation completed successfully.")

	// Show updated list of installed Java versions
	fmt.Println("Updated installed Java versions:")
	for _, file := range files {
		javaBinPath := filepath.Join(path, file.Name(), "bin")
		if javaBinPath == filepath.Dir(currentDefault) {
			fmt.Printf("- %s (%s) [default]\n", file.Name(), javaBinPath)
		} else {
			fmt.Printf("- %s (%s)\n", file.Name(), javaBinPath)
		}
	}

	fmt.Println("To change the default Java version, use the 'update-alternatives --config java' and 'update-alternatives --config javac' commands.")
}
