package libclc

import (
  "text/template"
  "bufio"
  "os"
)

type File struct {
  Path string
  Content string
  Owner string
  Permissions string
}

type CloudConfig struct {
  DiscoveryUrl string
  Files []*File
}

func WriteUserData(config *CloudConfig, t *template.Template, path string) (error) {
  f, err := os.Create(path)
  if err != nil {
    return err
  }
  defer f.Close()
  w := bufio.NewWriter(f)
  err = MakeUserData(config, t, w)
  if err != nil {
    return err
  }
  w.Flush()
  return nil
}

func MakeUserData(service *CloudConfig, t *template.Template, w *bufio.Writer) (error) {
  return t.Execute(w, service)
}
