VERSION=0.1.5-dev
GOVARS = -X main.Version=$(VERSION)
SYSTEM = ${GOOS}_${GOARCH}

build:
	rm dist/laravel-make
	go build -trimpath -ldflags "-s -w -X main.Version=0.1.5-dev" -o dist ./cmd/laravel-make

build-dist:
	go build -trimpath -ldflags "-s -w $(GOVARS)" -o build/bin/laravel-make-$(VERSION)-$(SYSTEM) ./cmd/laravel-make

laravel-make:
	go build -trimpath -ldflags "-s -w $(GOVARS)" -o dist ./cmd/laravel-make

build-dist-all:
	go run tools/build-all.go

package-setup:
	if [ ! -d "build/archives" ]; then\
		mkdir -p build/archives;\
	fi

package: build-dist package-setup

	mkdir -p build/laravel-make-$(VERSION)-$(SYSTEM);\
	cp README.md build/laravel-make-$(VERSION)-$(SYSTEM)
	if [ "${GOOS}" = "windows" ]; then\
		cp build/bin/laravel-make-$(VERSION)-$(SYSTEM) build/laravel-make-$(VERSION)-$(SYSTEM)/laravel-make.exe;\
		cd build;\
		zip -r -q -T archives/laravel-make-$(VERSION)-$(SYSTEM).zip laravel-make-$(VERSION)-$(SYSTEM);\
	else\
		cp build/bin/laravel-make-$(VERSION)-$(SYSTEM) build/laravel-make-$(VERSION)-$(SYSTEM)/laravel-make;\
		cd build;\
		tar -czf archives/laravel-make-$(VERSION)-$(SYSTEM).tar.gz laravel-make-$(VERSION)-$(SYSTEM);\
	fi

clean:
	rm -rf build

lint:
	golangci-lint run cmd/laravel-make
