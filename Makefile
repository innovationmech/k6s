.DEFAULT_GOAL := all

GO:=go
BIN_NAME=k6s
OUTPUT_DIR=_output
ROOT_PACKAGE := $(shell pwd)

# The OS can be linux/windows/darwin when building binaries
PLATFORMS ?= darwin_amd64 darwin_arm64 windows_amd64 windows_arm64 linux_amd64 linux_arm64

.PHONY: all
all: tidy build

define USAGE_OPTIONS

Options:
  tidy             Format the code.
  build            Build the binary(multi arch), output binary is in _output directory.
  clean            Delete the output binary.
 endef
 export USAGE_OPTIONS

.PHONY: tidy
tidy:
	@echo "go mod tidy"
	@$(GO) mod tidy

.PHONY: build
build:
	@echo "build go program:"
	$(MAKE) build.multiarch

.PHONY: build.multiarch
build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(p)))


.PHONY: go.build.%
go.build.%:
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	$(eval GO_OUT_EXT := $(if $(findstring windows,$(OS)),.exe,))
	@echo "===========> Building binary for $(OS) $(ARCH)"
	@mkdir -p $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) -o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(BIN_NAME)$(GO_OUT_EXT) $(ROOT_PACKAGE)/cmd/.

.PHONY: clean
clean:
	@echo "clean go program"
	@$(RM) -rf $(OUTPUT_DIR)

.PHONY: help
help:
	@echo "$$USAGE_OPTIONS"

.PHONY: test
test:
	@echo $(shell pwd)
