#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

function exists() {
  command -v "${@}" > /dev/null 2>&1
}

function run() {
  if exists gum; then
    local prefix="$(gum style --background=14 --padding="0 1" RUN)"
    local message="$(gum style --foreground=14 "${*}")"
    gum join --horizontal "${prefix}" " " "${message}"
  fi
  "${@}"
}

workspace="$(git rev-parse --show-toplevel || pwd)"
cd "${workspace}"

GOARCH="$(go env GOARCH)"
GOEXE="$(go env GOEXE)"
GOOS="$(go env GOOS)"

for f in bin/*; do
  if [[ ! ${f} =~ .*-[[:alnum:]]+-[[:alnum:]]+(.exe)? ]]; then
    run mv "${f}" "${f%${GOEXE}}-${GOOS}-${GOARCH}${GOEXE}"
  fi
done
