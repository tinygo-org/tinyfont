# md5sum binary
MD5SUM = md5sum

# tinygo binary for tests
TINYGO ?= tinygo

.PHONY: all clean fmt fmt-check

all: fmt fmt-check smoke-test

clean:
	@mkdir -p build
	@rm -f $(TARGET)

FMT_PATHS = ./*.go cmd freemono freesans freeserif gophers proggy examples
fmt:
	@gofmt -l -w $(FMT_PATHS)
fmt-check:
	@unformatted=$$(gofmt -l $(FMT_PATHS)); [ -z "$$unformatted" ] && exit 0; echo "Unformatted:"; for fn in $$unformatted; do echo "  $$fn"; done; exit 1

TARGET = build/examples_epd.hex \
		 build/examples_hub75.hex \
		 build/examples_displays_pybadge.hex \
		 build/examples_displays_pyportal.hex \
		 build/examples_displays_wioterminal.hex \
		 build/examples_unicode_font_pyportal.hex \
		 build/examples_unicode_font_wioterminal.hex \
		 build/examples_unicode_font2_pyportal.hex \
		 build/examples_unicode_font2_wioterminal.hex
.PHONY: smoke-test $(TARGET)
smoke-test: clean $(TARGET)

build/examples_epd.hex:
	$(TINYGO) build -size short -o $@ -target=microbit ./examples/epd
	@$(MD5SUM) $@

build/examples_hub75.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge ./examples/hub75
	@$(MD5SUM) $@

build/examples_displays_pybadge.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge ./examples/displays
	@$(MD5SUM) $@

build/examples_displays_pyportal.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/displays
	@$(MD5SUM) $@

build/examples_displays_wioterminal.hex:
	$(TINYGO) build -size short -o $@ -target=wioterminal ./examples/displays
	@$(MD5SUM) $@

build/examples_unicode_font_pyportal.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/unicode_font
	@$(MD5SUM) $@

build/examples_unicode_font_wioterminal.hex:
	$(TINYGO) build -size short -o $@ -target=wioterminal ./examples/unicode_font
	@$(MD5SUM) $@

build/examples_unicode_font2_pyportal.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal ./examples/unicode_font2
	@$(MD5SUM) $@

build/examples_unicode_font2_wioterminal.hex:
	$(TINYGO) build -size short -o $@ -target=wioterminal ./examples/unicode_font2
	@$(MD5SUM) $@
