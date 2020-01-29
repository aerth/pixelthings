version     := 0.0.1

ldflags := '-w -s -X main.version=$(version) -extldflags "-fno-PIC"'
buildflags := -v -tags 'osusergo netgo static_build'
winflags := GOARCH=386 GOOS=windows CC=i686-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1

bin/pixelgif: assets/assets_vfsdata.go buttonid_string.go *.go */*.go
	go build $(buildflags) -o $@ --ldflags $(ldflags)
	#upx --best bin/pixelgif
bin/pixelgif.exe: assets/assets_vfsdata.go buttonid_string.go *.go */*.go
	env $(winflags) go build $(buildflags) -o $@ --ldflags $(ldflags)
	#upx --best bin/pixelgif.exe

release: bin/pixelgif bin/pixelgif.exe
	upx --best bin/pixelgif
	upx --best bin/pixelgif.exe
	cd bin && zip pixelgif_linux.zip pixelgif
	cd bin && zip pixelgif_windows.zip pixelgif.exe
	cd bin && sha256sum *.zip >> checksums.txt

buttonid_string.go: buttons.go
	go generate -v .

assets/assets_vfsdata.go: assets/
	go generate -v ./assets/

clean:
	rm -f ./assets/assets_vfsdata.go
	rm -f buttonid_string.go
	rm -f bin/*
	rmdir bin