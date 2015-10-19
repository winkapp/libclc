package libclc

import (
  "text/template"
  "bufio"
)

func WriteUserData(config *CloudConfig, t *template.Template, path string) (error) {
  return WriteTemplate("ud", config, t, path)
}

func BufferUserData(config *CloudConfig, t *template.Template, w *bufio.Writer) (error) {
  return BufferTemplate("ud", config, t, w)
}
