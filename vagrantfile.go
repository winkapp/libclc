package libclc

import (
  "text/template"
  "bufio"
)

type Vagrantfile struct {
}

func WriteVagrantfile(config *Vagrantfile, t *template.Template, path string) (error) {
  return WriteTemplate("vf", config, t, path)
}

func BufferVagrantfile(config *Vagrantfile, t *template.Template, w *bufio.Writer) (error) {
  return BufferTemplate("vf", config, t, w)
}
