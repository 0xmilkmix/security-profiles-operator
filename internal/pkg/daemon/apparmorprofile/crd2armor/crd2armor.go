package crd2armor

import (
	"bytes"
	"errors"
	"text/template"

	apparmorprofileapi "sigs.k8s.io/security-profiles-operator/api/apparmorprofile/v1alpha1"
)

var appArmorTemplate = `
#include <tunables/global>
profile {{.Name}} flags=(attach_disconnected,mediate_deleted) {
  #include <abstractions/base>

  # Executable rules
{{ if ne .Abstract.Executable nil }}{{ if ne .Abstract.Executable.AllowedExecutables nil }}
{{range $allowed := .Abstract.Executable.AllowedExecutables}}  {{$allowed}} ix,
{{end}}{{end}}
{{ if ne .Abstract.Executable.AllowedLibraries nil }}
{{range $allowedlib := .Abstract.Executable.AllowedLibraries}}  {{$allowedlib}} mr,
{{end}}{{end}}{{end}}

  # Filesystem rules
{{ if ne .Abstract.Filesystem nil }}{{ if ne .Abstract.Filesystem.ReadOnlyPaths nil }}
{{range $readonly := .Abstract.Filesystem.ReadOnlyPaths}}  {{$readonly}} r,
{{end}}
{{range $readonly := .Abstract.Filesystem.ReadOnlyPaths}}  deny {{$readonly}} wl,
{{end}}{{end}}
{{ if ne .Abstract.Filesystem.WriteOnlyPaths nil }}
{{range $writeonly := .Abstract.Filesystem.WriteOnlyPaths}}  {{$writeonly}} wl,
{{end}}
{{range $writeonly := .Abstract.Filesystem.WriteOnlyPaths}}  deny {{$writeonly}} r,
{{end}}{{end}}
{{ if ne .Abstract.Filesystem.ReadWritePaths nil }}
{{range $readwrite := .Abstract.Filesystem.ReadWritePaths}}  {{$readwrite}} rwl,
{{end}}{{end}}{{end}}

  # Network rules
{{ if ne .Abstract.Network nil }}{{ if ne .Abstract.Network.AllowRaw nil }}
{{ if .Abstract.Network.AllowRaw}}{{else}}  deny network raw,
{{end}}{{end}}
{{ if ne .Abstract.Network.Protocols nil }}
{{if ne .Abstract.Network.Protocols.AllowTCP nil }}
{{if .Abstract.Network.Protocols.AllowTCP}}  network inet tcp,
{{end}}{{end}}{{if ne .Abstract.Network.Protocols.AllowUDP nil }}
{{if .Abstract.Network.Protocols.AllowUDP}}  network inet udp,
{{end}}{{end}}{{end}}{{end}}

  # Capabilities rules
{{ if ne .Abstract.Capability nil}}{{range $cap := .Abstract.Capability.AllowedCapabilities}}  capability {{$cap}},
{{end}}{{end}}

  # Raw rules placeholder

  # Add default deny for known information leak/priv esc paths
  deny @{PROC}/* w,   # deny write for all files directly in /proc (not in a subdir)
  deny @{PROC}/{[^1-9],[^1-9][^0-9],[^1-9s][^0-9y][^0-9s],[^1-9][^0-9][^0-9][^0-9]*}/** w,
  deny @{PROC}/sys/[^k]** w,  # deny /proc/sys except /proc/sys/k* (effectively /proc/sys/kernel)
  deny @{PROC}/sys/kernel/{?,??,[^s][^h][^m]**} w,  # deny everything except shm* in /proc/sys/kernel/
  deny @{PROC}/sysrq-trigger rwklx,
  deny @{PROC}/mem rwklx,
  deny @{PROC}/kmem rwklx,
  deny @{PROC}/kcore rwklx,
  deny mount,
  deny /sys/[^f]*/** wklx,
  deny /sys/f[^s]*/** wklx,
  deny /sys/fs/[^c]*/** wklx,
  deny /sys/fs/c[^g]*/** wklx,
  deny /sys/fs/cg[^r]*/** wklx,
  deny /sys/firmware/efi/efivars/** rwklx,
  deny /sys/kernel/security/** rwklx,
}
`

type apparmorTemplateArgs struct {
	Name     string
	Abstract *apparmorprofileapi.AppArmorAbstract
}

// GenerateProfile uses the CRD representation of an abstracted profile to generate a
// full AppArmor profile.
func GenerateProfile(name string, abstract *apparmorprofileapi.AppArmorAbstract) (string, error) {
	var generated bytes.Buffer
	templateArgs := apparmorTemplateArgs{
		Name:     name,
		Abstract: abstract,
	}

	if abstract == nil {
		return "", errors.New("abstract cannot be nil")
	}

	tpl, err := template.New("apparmor").Parse(appArmorTemplate)
	if err != nil {
		return "", err
	}
	if err := tpl.Execute(&generated, templateArgs); err != nil {
		return "", err
	}
	return generated.String(), nil
}
