package dispatch

import (
  "fmt"
  "os"
  "syscall"
  "io/ioutil"
  "bufio"
  "strconv"
)


func pidFilename(name string) string {
  return fmt.Sprintf("%s.pid", name)
}

func killCommand(name string) {
  pid := getPID(name)
  if pid != 0 {
    syscall.Kill(pid, 9)
    clearPID(name)
  }
}

func getPID(name string) int {
  fn := pidFilename(name)
  file, err := os.Open(fn)
  if err != nil {
    return 0
  }
  reader := bufio.NewReader(file)
  pid, _ := reader.ReadString('\n')
  pid = pid[0: len(pid)]
  pidInt, _ := strconv.Atoi(pid)
  return int(pidInt)
}

func clearPID(name string) {
  fn := pidFilename(name)
  os.Remove(fn)
}

func writePID(name string) {
  pid := os.Getpid()
  fn := pidFilename(name)
  ioutil.WriteFile(fn, []byte(fmt.Sprintf("%d", pid)), 0777)
}

