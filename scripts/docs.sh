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
