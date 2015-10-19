package libclc

import (
  "text/template"
  "bufio"
)

type Configrb struct {
}

func WriteConfigrb(config *Configrb, t *template.Template, path string) (error) {
  return WriteTemplate("crb", config, t, path)
}

func BufferConfigrb(config *Configrb, t *template.Template, w *bufio.Writer) (error) {
  return BufferTemplate("crb", config, t, w)
}
