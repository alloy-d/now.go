package now

import (
    "bufio"
    "container/list"
    "fmt"
    "os"
    "strings"
)

type Thing struct {
    Name string
}

var (
    theList *list.List = list.New()
)

const filename = ".next.things"

func init() { readList() }

func openList(fresh bool) (*os.File, bool) {
    home := os.Getenv("HOME")
    if home == "" {
        return nil, false
    }

    filePath := home + "/" + filename
    flags := os.O_CREAT|os.O_RDWR
    if fresh {
        flags |= os.O_TRUNC
    } else {
        flags |= os.O_APPEND
    }
    file, err := os.Open(filePath, flags, 0600)

    if err != nil {
        return file, false
    }

    return file, true
}

func readList() {
    listFile, ok := openList(false)
    defer listFile.Close()

    if !ok { return }

    reader := bufio.NewReader(listFile)

    theList.Init()
    nextThing, err := reader.ReadString('\n')
    for ; err == nil; nextThing, err = reader.ReadString('\n') {
        theList.PushBack(strings.Split(nextThing, "\n", 0)[0])
    }
}

func writeList() {
    listFile, ok := openList(true)
    defer listFile.Close()

    if !ok { return }

    things := theList.Iter()
    for thing := range things {
        listFile.WriteString(fmt.Sprintf("%s\n", thing))
    }
}

func AddThing(addition string) {
    readList()
    theList.PushBack(addition)
    writeList()
}

func GetNext() string {
    readList()
    if front := theList.Front(); front != nil {
        return theList.Front().Value.(string)
    }
    return ""
}

func Pop() {
    readList()
    if front := theList.Front(); front != nil {
        theList.Remove(front)
    }
    writeList()
}

func PushNext(addition string) {
    readList()
    theList.PushFront(addition)
    writeList()
}
