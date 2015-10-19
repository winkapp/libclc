package libclc

import (
  "text/template"
  "bufio"
  "os"
  "io/ioutil"
)

var template_dir string = os.ExpandEnv("$GOPATH/src/github.com/winkapp/libclc/templates")

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

func GetTemplate(name string, filename string) (t *template.Template, err error) {
  templ, err := getFile(template_dir + "/" + filename)
  if err != nil {
    return nil, err
  }
  t = template.New(name)
  return t.Parse(templ)
}

func getFile(path string) (string, error) {
  dat, err := ioutil.ReadFile(path)
  return string(dat), err
}
