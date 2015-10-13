package libclc

import (
  "text/template"
  "bufio"
)

type UnitConfig struct {
  Units []*Unit
}

type Unit struct {
  Name string
  Image string
  Command string
  Type string
  Filename string
  Restart string
  Envs []string
  Ports []string
  Xfleet []string
}

func WriteUnit(unit Unit, t *template.Template) (error) {
  return nil
}

func MakeUnit(t *template.Template, service *Unit, w *bufio.Writer) (error) {
  return t.Execute(w, service)
}
