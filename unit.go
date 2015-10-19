package libclc

import (
  "text/template"
  "bufio"
  "path"
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

func WriteUnits(units *UnitConfig, t *template.Template, path string) (error) {
  for _, service := range units.Units {
    err := WriteUnit(service, t, path)
    if err != nil {
      return err
    }
  }
  return nil
}

func WriteUnit(unit *Unit, t *template.Template, p string) (error) {
  fileName := unitFileName(unit)
  var err error

  if t == nil {
    t, err = GetTemplate("Service Template", "service.template")
  }
  err = WriteTemplatedFile(unit, t, path.Join(p, fileName))
  return err
}

func MakeUnit(service *Unit, t *template.Template, w *bufio.Writer) (error) {
  return t.Execute(w, service)
}

func unitFileName(service *Unit) string {
  switch service.Type {
  case "single":
    service.Filename = (service.Name + ".service")
  case "multi":
    service.Filename = (service.Name + "@.service")
  }

  return service.Filename
}
