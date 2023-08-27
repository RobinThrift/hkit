staticcheck_version := "2023.1.5"
golangci_lint_version := "v1.54.2"
git_chglog_version := "v0.15.4"
cyclonedx_version := "v1.4.1"

_default:
    @just --list

test +flags="-failfast":
	go test {{flags}} ./...

alias tw := test-watch
test-watch +flags="-failfast":
	gotestsum --watch -- {{flags}} ./...

lint:
	go run honnef.co/go/tools/cmd/staticcheck@{{staticcheck_version}} ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@{{golangci_lint_version}} run ./...

fmt:
	@go fmt ./...

clean:
	go clean -cache

alias gen := generate
generate:
	@go run ./generate

sbom:
    go run github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@{{cyclonedx_version}} mod -json -licenses -std -test -output sbom.json

changelog tag:
	go run github.com/git-chglog/git-chglog/cmd/git-chglog@{{git_chglog_version}} --next-tag {{tag}} --sort semver -o CHANGELOG.md

release tag:
    just changelog {{tag}}
    git add CHANGELOG.md
    git commit -m "Generated changelog for {{tag}}"
    git tag {{tag}}
    git push
    git push origin {{tag}}

