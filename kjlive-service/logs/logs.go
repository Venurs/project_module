package logs

import (
    "os"
    "runtime"
    "fmt"

    "github.com/kpango/glg"

    "kjlive-service/conf"
)

func LogInit() (infoLog *os.File, errorLog *os.File, actionLog *os.File)  {
    infoLog = glg.FileWriter(conf.Settings.InfoLogFile, 0666)
    errorLog = glg.FileWriter(conf.Settings.ErrorLogFile, 0666)
    actionLog = glg.FileWriter(conf.Settings.ActionLogFile, 0666)
    glg.Get().SetMode(glg.WRITER).
        AddLevelWriter(glg.INFO, infoLog).
        AddLevelWriter(glg.ERR, errorLog).
        AddLevelWriter(glg.WARN, actionLog)
    return
}

func Error(err error)  {
    pc,file,line,ok := runtime.Caller(1)
    if ok {
        funcName := runtime.FuncForPC(pc)
        glg.Errorf("pc:%v file:%v line:%v func_name:%v error:%v",  pc, file, line, funcName.Name(), err)
    }
}

func Errorf(format string, val ...interface{})  {
    pc,file,line,ok := runtime.Caller(1)
    if ok {
        funcName := runtime.FuncForPC(pc)
        glg.Errorf("pc:%v file:%v line:%v func_name:%v error:%v",  pc, file, line, funcName.Name(), fmt.Errorf(format, val...))
    }
}

func Info(err string)  {
    pc,file,line,ok := runtime.Caller(1)
    if ok {
        funcName := runtime.FuncForPC(pc)
        glg.Infof("pc:%v file:%v line:%v func_name:%v error:%v",  pc, file, line, funcName.Name(), err)
    }
}

func Infof(format string, val ...interface{})  {
    pc,file,line,ok := runtime.Caller(1)
    if ok {
        funcName := runtime.FuncForPC(pc)
        glg.Infof("pc:%v file:%v line:%v func_name:%v error:%v",  pc, file, line, funcName.Name(), fmt.Errorf(format, val...))
    }
}

