package libclc

import (
  "text/template"
  "bufio"
  "os"
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

func WriteUnit(unit *Unit, t *template.Template, path string) (error) {
  fileName := unitFileName(unit)
  f, err := os.Create(path + "/" + fileName)
  if err != nil {
    return err
  }
  defer f.Close()
  w := bufio.NewWriter(f)
  err = MakeUnit(unit, t, w)
  if err != nil {
    return err
  }
  w.Flush()
  return nil
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
