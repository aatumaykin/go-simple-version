package version

import (
	"os"
	"runtime"
	"runtime/debug"
)

const (
	unknownVersion  = "(devel)"
	unknownProperty = "N/A"
)

var (
	Name      = unknownProperty
	Version   = unknownVersion
	Commit    = unknownProperty
	BuildTime = unknownProperty
)

type Info struct {
	BuildTime string
	Version   string
	Os        string
	Arch      string
	GoVersion string
	Commit    string
}

func Get() Info {
	info, _ := debug.ReadBuildInfo()

	if Version == unknownVersion && info.Main.Version != "" {
		Version = info.Main.Version
	}

	if Name == unknownProperty {
		Name = os.Args[0]
	}

	return Info{
		BuildTime: BuildTime,
		Version:   Version,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		Commit:    Commit,
	}
}
