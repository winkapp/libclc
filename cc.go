package libclc

import (
  "text/template"
  "bufio"
)

type File struct {
  HostPath string `yaml:"host_path"`
  Path string
  Content string
  Owner string
  Permissions string
}

type CloudConfig struct {
  DiscoveryUrl string
  Files []*File
}

func WriteCloudConfig(config *CloudConfig, t *template.Template, path string) (error) {
  return WriteTemplate("cc", config, t, path)
}

func BufferCloudConfig(config *CloudConfig, t *template.Template, w *bufio.Writer) (error) {
  return BufferTemplate("cc", config, t, w)
}
