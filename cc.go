package libclc

import (
  "text/template"
  "bufio"
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
  err := WriteTemplatedFile(config, t, path)
  return err
}

func MakeUserData(service *CloudConfig, t *template.Template, w *bufio.Writer) (error) {
  return t.Execute(w, service)
}
