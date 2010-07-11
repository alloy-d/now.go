package now

import (
    "path"
    "runtime"
)

func ResourceDir() string {
    _, file, _, _ := runtime.Caller(0)
    dir, _ := path.Split(file)
    return path.Join(dir, "human-friendly")
}
