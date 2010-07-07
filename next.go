package next

import (
    "bufio"
    "os"
    "strings"
)

type Thing struct {
    Name string
}

const filename = ".next.things"

func openList() (*os.File, bool) {
    home := os.Getenv("HOME")
    if home == "" {
        return nil, false
    }

    filePath := home + "/" + filename
    file, err := os.Open(filePath, os.O_APPEND|os.O_CREAT|os.O_RDWR, 0600)

    if err != nil {
        return file, false
    }

    return file, true
}

func AddThing(addition string) bool {
    listFile, ok := openList()
    defer listFile.Close()

    if !ok { return false }

    _, err := listFile.WriteString(addition + "\n")
    if err != nil { return false }

    return true
}

func GetNext() string {
    listFile, ok := openList()
    defer listFile.Close()

    if !ok { return "" }

    reader := bufio.NewReader(listFile)

    nextThing, err := reader.ReadString('\n')
    if err != nil { return "" }

    return strings.Split(nextThing, "\n", 0)[0]
}

