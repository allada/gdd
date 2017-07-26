package config

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

type Config struct {
  Port string
  DlvPath string
  DebugSession struct {
    File string
    Args []string
  }
}

func fileFromArg(c *Config, fileToDebug string) bool {
  if !fileExists(fileToDebug) {
    fmt.Println("Invalid file path for first argument")
    return false
  }
  c.DebugSession.File = fileToDebug
  return true
}

func portFromArg(c *Config, port string) bool {
  if _, err := strconv.Atoi(port); err != nil {
    fmt.Println("Value for 'port' must be a valid number.")
    return false
  }
  c.Port = port
  return true
}

func fileExists(path string) bool {
  if finfo, err := os.Stat(path); err != nil || finfo.IsDir() {
    return false
  }
  return true
}

func dlvFromArg(c *Config, dlvPath string) bool {
  if !fileExists(dlvPath) {
    fmt.Println("Invalid file for 'dlv' argument.")
    return false
  }
  c.DlvPath = dlvPath
  return true
}

var boundArgsInfo = map[string]func(*Config, string)bool{
  "0": fileFromArg,
  "--port": portFromArg,
  "--dlv": dlvFromArg,
  "--help": func (_ *Config, _ string) bool {
    printHelp()
    return false
  },
}

func NewFromArgs(args []string) *Config {
  if len(args) <= 0 {
    printHelp();
    return nil
  }

  defaultDlvPath := ""
  if val, ok := os.LookupEnv("PATH"); ok {
    for _, folder := range strings.Split(val, ":") {
      file := folder + "/dlv"
      if fileExists(file) {
        defaultDlvPath = file
        break
      }
    }
  }


  config := &Config{
    Port: "9922",
    DlvPath: defaultDlvPath,
  }

  curNumbPos := 0
  i := 0
  for ; i < len(args); i++ {
    arg := args[i]
    if !strings.HasPrefix(arg, "--") {
      fn, ok := boundArgsInfo[strconv.Itoa(curNumbPos)]
      curNumbPos++
      if ok {
        if fn(config, arg) {
          continue
        }
      } else {
        fmt.Println("Unknown arguemnt '" + arg + "'")
      }
      return nil
    }
    keyValue := strings.SplitN(arg, "=", 2)
    fn, ok := boundArgsInfo[keyValue[0]]
    if !ok {
      fmt.Println("Unknown argument '" + args[i] + "'")
      return nil
    }

    value := keyValue[0]
    if len(keyValue) >= 2 {
      value = keyValue[1]
    }
    ok = fn(config, value)
    if !ok {
      return nil
    }
  }
  return config
}

func printHelp() {
  data := []string{
    "GDD is a simple debuger that binds Chrome Devtools debugger as a front-end and",
    "uses DLV as a backend.",
    "",
    "Usage:",
    "",
    "  gdd go-file-to-debug [options] [-- args-to-pass-to-file ...]",
    "",
    "Options:",
    "",
    "  --port=PORT        Port Number to use to communicate to Chrome Devtools",
    "                     front-end. Default 9922.",
    "  --dlv=FILE         Location of DLV installed on machine. Default (tries to",
    "                     find 'dlv' in system path.",
    "  --help             Prints this help dialog.",
    "",
  }
  for _, line := range data {
    fmt.Println(line)
  }
}
