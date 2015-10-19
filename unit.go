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
  Environment []string
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
  fileName := UnitFileName(unit)
  unitPath := path.Join(p, fileName)

  return WriteTemplate("unit", unit, t, unitPath)
}

func BufferUnit(unit *Unit, t *template.Template, w *bufio.Writer) (error) {
  return BufferTemplate("unit", unit, t, w)
}

func UnitFileName(service *Unit) string {
  switch service.Type {
  case "single":
    service.Filename = (service.Name + ".service")
  case "multi":
    service.Filename = (service.Name + "@.service")
  }

  return service.Filename
}
