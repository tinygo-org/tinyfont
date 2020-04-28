# md5sum binary
MD5SUM = md5sum

# tinygo binary for tests
TINYGO ?= tinygo

.PHONY: all clean fmt fmt-check

all:

clean:
	@rm -f $(TARGET)

FMT_PATHS = ./*.go cmd freemono freesans freeserif gophers proggy examples
fmt:
	@gofmt -l -w $(FMT_PATHS)
fmt-check:
	@unformatted=$$(gofmt -l $(FMT_PATHS)); [ -z "$$unformatted" ] && exit 0; echo "Unformatted:"; for fn in $$unformatted; do echo "  $$fn"; done; exit 1

TARGET = examples_epd.hex \
		 examples_hub75.hex \
		 examples_pybadge.hex \
		 examples_pyportal.hex
.PHONY: smoketest $(TARGET)
smoketest: $(TARGET)
	$(TINYGO) version

examples_epd.hex:
	$(TINYGO) build -size short -o $@ -target=microbit            ./examples/epd
	@$(MD5SUM) $@

examples_hub75.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge             ./examples/hub75
	@$(MD5SUM) $@

examples_pybadge.hex:
	$(TINYGO) build -size short -o $@ -target=pybadge             ./examples/pybadge
	@$(MD5SUM) $@

examples_pyportal.hex:
	$(TINYGO) build -size short -o $@ -target=pyportal            ./examples/pyportal
	@$(MD5SUM) $@
