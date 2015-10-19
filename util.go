package libclc

import (
  "text/template"
  "bufio"
  "os"
  "io/ioutil"
  "runtime"
  "path"
)

var template_dir string = ""



func WriteTemplate(ttype string, data interface{}, t *template.Template, path string) (error) {
  t, err := checkTemplate(ttype, t)
  if err != nil {
    return err
  }
  return writeTemplatedFile(data, t, path)
}

func BufferTemplate(ttype string, data interface{}, t *template.Template, w *bufio.Writer) (error) {
  t, err := checkTemplate(ttype, t)
  if err != nil {
    return err
  }
  return t.Execute(w, data)
}

func checkTemplate(ttype string, t *template.Template) (*template.Template, error) {
  if t != nil {
    return t, nil
  }
  switch ttype {
  case "cc":
    return getTemplate("Cloud Config", "cloud-config.template")
  case "ud":
    return getTemplate("User Data", "user-data.template")
  case "unit":
    return getTemplate("Service", "service.template")
  case "vf":
    return getTemplate("Vagrantfile", "Vagrantfile.template")
  case "crb":
    return getTemplate("Config.rb", "config.rb.template")
  }
  return nil, nil
}

func getTemplate(name string, filename string) (t *template.Template, err error) {
  templ, err := getFile(path.Join(templateDir(), filename))
  if err != nil {
    return nil, err
  }
  t = template.New(name)
  return t.Parse(templ)
}

func writeTemplatedFile(data interface{}, t *template.Template, path string) (error) {
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

func getFile(path string) (string, error) {
  dat, err := ioutil.ReadFile(path)
  return string(dat), err
}

func templateDir() (string) {
  if template_dir == "" {
    _, filename, _, _ := runtime.Caller(1)
    template_dir = path.Join(path.Dir(filename), "templates")
  }
  return template_dir
}
