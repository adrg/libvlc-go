package vlc

// #include <vlc/libvlc_version.h>
import "C"
import "fmt"

// VersionInfo contains details regarding the version of the libVLC module.
type VersionInfo struct {
	Major int
	Minor int
	Patch int
}

// String returns a string representation of the version.
func (v VersionInfo) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

var moduleVersion = VersionInfo{
	Major: C.LIBVLC_VERSION_MAJOR,
	Minor: C.LIBVLC_VERSION_MINOR,
	Patch: C.LIBVLC_VERSION_REVISION,
}

func assertVersion(major, minor, patch int) error {
	if v := &moduleVersion; major > v.Major || minor > v.Minor || patch > v.Patch {
		return fmt.Errorf("operation requires libVLC v%d.%d.%d or later: v%s installed",
			major, minor, patch, v.String())
	}

	return nil
}
