package vlc

// #cgo LDFLAGS: -lvlc
// #include <vlc/vlc.h>
// #include <vlc/libvlc_version.h>
import "C"
import "fmt"

// VersionInfo contains details regarding the version of the libVLC module.
type VersionInfo struct {
	Major uint
	Minor uint
	Patch uint
	Extra uint
}

// String returns a string representation of the version.
func (v VersionInfo) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Runtime returns the runtime version of libVLC, usually including
// the codename of the build.
// NOTE: Due to binary backward compatibility, the runtime version may be more
// recent than the build version.
func (v VersionInfo) Runtime() string {
	return C.GoString(C.libvlc_get_version())
}

// Changeset returns the changeset identifier for the current libVLC build.
func (v VersionInfo) Changeset() string {
	return C.GoString(C.libvlc_get_changeset())
}

// Compiler returns information regarding the compiler used to build libVLC.
func (v VersionInfo) Compiler() string {
	return C.GoString(C.libvlc_get_compiler())
}

var moduleVersion = VersionInfo{
	Major: C.LIBVLC_VERSION_MAJOR,
	Minor: C.LIBVLC_VERSION_MINOR,
	Patch: C.LIBVLC_VERSION_REVISION,
	Extra: C.LIBVLC_VERSION_EXTRA,
}
