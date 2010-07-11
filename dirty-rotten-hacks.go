package now

import (
    "fmt"
    "path"
    "runtime"
)

func ResourceDir() string {
    _, file, _, _ := runtime.Caller(0)
    dir, _ := path.Split(file)
    fmt.Println(dir)
    return path.Join(dir, "human-friendly")
}
