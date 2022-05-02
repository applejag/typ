# SPDX-FileCopyrightText: 2022 Kalle Fagerberg
#
# SPDX-License-Identifier: CC0-1.0

.PHONY: test clean tidy deps \
	lint lint-md lint-go lint-license \
	lint-fix lint-md-fix

test:
	go test ./...

tidy:
	go mod tidy

deps:
	go install github.com/mgechev/revive@latest
	go install golang.org/x/tools/cmd/goimports@latest
	python3 -m pip install --upgrade --user reuse
	npm install

lint: lint-md lint-go lint-license
lint-fix: lint-md-fix

lint-md:
	npx remark . .github

lint-md-fix:
	npx remark . .github -o

lint-go:
	revive -formatter stylish -config revive.toml ./...

lint-license:
	reuse lint
