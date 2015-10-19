package libclc

import (
  "text/template"
  "bufio"
  "os"
)

func WriteTemplatedFile(data interface{}, t *template.Template, path string) (error) {
  f, err := os.Create(path)
  if err != nil {
    return err
  }
  defer f.Close()
  w := bufio.NewWriter(f)
  err = t.Execute(w, data)
  if err != nil {
    return err
  }
  w.Flush()
  return nil
}
