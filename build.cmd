go env
set CGO_CFLAGS=-IF:\Software_backup\libvlc-sdk-3.0.8-win32\vlc-3.0.8\sdk\include
set CGO_LDFLAGS=-LF:\Software_backup\libvlc-sdk-3.0.8-win32\vlc-3.0.8
go env
go build && go install
