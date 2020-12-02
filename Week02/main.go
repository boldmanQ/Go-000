package main

import (
    "github.com/pkg/errors"
    //buildInErr "errors"
    "fmt"
//    "os"
    "database/sql"
)


type CrashPodLog struct {
    PodName string
    NameSpace string
}

func DAO(podName string) (string, error) {
    //dosomethins, if get sql.ErrNoRows, return it
    // ····
    var err = sql.ErrNoRows
    //err := buildInErr.New("sql: no rows in result set")
    //instance := sql.Row{}
    //err := instance.Err()
    //_, err := os.Open(podName)
    return "", errors.Wrap(err, "Custome Wrap Err Messages")
}

func queryCrashLog(podName string) (string, error) {
    log, err := DAO(podName)
    if err != nil {
        return log, err
    }
    return log, err
}

func main() {
    setPodName := "testPod"
    log, err := queryCrashLog(setPodName)
    if err != nil {
        fmt.Printf("origin err: %T %v\n", errors.Cause(err), errors.Cause(err))
        fmt.Printf("stack trace:\n%+v\n", err)
    }
}
