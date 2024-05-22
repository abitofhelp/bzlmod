BZL_CMD=bazel
BAZEL_BUILD_OPTS:=
GO_CMD=$(BZL_CMD) run @rules_go//go
PROJECT_NAME=go_connect_bazelmod
CLIENT_DIR=github.com/abitofhelp/$(PROJECT_NAME)
PROJECT_DIR=$(GOPATH)/src/$(CLIENT_DIR)

clean:
	pushd $(PROJECT_DIR) && \
	rm -rf internal/gen/*.go && \
	popd

list_endpoints:
	pushd $(PROJECT_DIR) && \
	grpcurl -plaintext -proto proto/abitofhelp/helloworld/v1/helloworld.proto localhost:8080 list && \
	popd

buf_generate: clean
	pushd $(PROJECT_DIR) && \
	buf generate proto && \
	popd

go_mod_download:
	pushd $(PROJECT_DIR) && \
	$(GO_CMD) -- mod download && \
	popd;

go_mod_tidy:
	pushd $(PROJECT_DIR) && \
	$(GO_CMD) -- mod tidy && \
	popd;
go_mod_upgrade_dependencies:
	pushd $(PROJECT_DIR) && \
	$(GO_CMD) -- get -u ./... && \
	popd;

go_mod_vendor:
	pushd $(PROJECT_DIR) && \
	$(GO_CMD) -- mod vendor -v && \
	popd;

go_mod_verify:
	## Verify that the go.sum file matches what was downloaded to prevent someone “git push — force” over a tag being used.
	pushd $(PROJECT_DIR) && \
	$(GO_CMD) -- mod verify && \
	popd;

go_targets:
	$(BZL_CMD) query "@rules_go//go:*"

zap:
	pushd $(PROJECT_DIR) && \
	find . -type f -name "go.sum" -delete && \
	go clean -modcache -cache && \
	make go_mod_download && \
	make go_mod_tidy && \
	make go_mod_vendor && \
	make go_mod_verify && \
	popd;