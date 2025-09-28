
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
)

const (
    logFile   = "/var/log/syslog"
    outputLog = "/home/stephen/code/sre/health_check.log"

    cpuThresh  = 80.0
    ramThresh  = 80.0
    diskThresh = 90.0
)

func logMessage(message string) {
    f, _ := os.OpenFile(outputLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    line := fmt.Sprintf("[%s] %s\n", timestamp, message)
    f.WriteString(line)
    fmt.Print(line)
}

func checkSystem() {
    cpuPercent, _ := cpu.Percent(0, false)
    if cpuPercent[0] > cpuThresh {
        logMessage(fmt.Sprintf("ALERT: High CPU usage: %.2f%%", cpuPercent[0]))
    }

    vm, _ := mem.VirtualMemory()
    if vm.UsedPercent > ramThresh {
        logMessage(fmt.Sprintf("ALERT: High RAM usage: %.2f%%", vm.UsedPercent))
    }

    du, _ := disk.Usage("/")
    if du.UsedPercent > diskThresh {
        logMessage(fmt.Sprintf("ALERT: High Disk usage: %.2f%%", du.UsedPercent))
    }
}

func checkLogs() {
    f, err := os.Open(logFile)
    if err != nil {
        return
    }
    defer f.Close()

    lines := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if len(lines) > 20 {
        lines = lines[len(lines)-20:]
    }

    for _, line := range lines {
        lower := strings.ToLower(line)
        if strings.Contains(lower, "error") || strings.Contains(lower, "fail") {
            logMessage("Log Alert: " + line)
        }
    }
}

func main() {
    logMessage("Health check script started (Go lean version).")
    for {
        checkSystem()
        checkLogs()
        time.Sleep(60 * time.Second)
    }
}
