#!/usr/bin/env bash
set -eEuo pipefail

if [[ -z "${ACTIONS_STEP_DEBUG:-}" ]]; then
  set -x
fi

# Create a unique output file outside of the checkout.
BINARY_PATH="${RUNNER_TEMP}/terraform-linter-${GITHUB_SHA:0:7}"

OS=""
if [[ "${RUNNER_OS}" == "Linux" ]]; then
  OS="linux"
elif [[ "${RUNNER_OS}" == "macOS" ]]; then
  OS="darwin"
elif [[ "${RUNNER_OS}" == "Windows" ]]; then
  OS="windows"
  BINARY_PATH="${BINARY_PATH}.exe"
else
  echo "::error ::Unsupported operating system ${RUNNER_OS}"
  exit 1
fi

ARCH=""
if [[ "${RUNNER_ARCH}" == "X64" ]]; then
  ARCH="amd64"
elif [[ "${RUNNER_ARCH}" == "ARM64" ]]; then
  ARCH="arm64"
else
  echo "::error ::Unsupported system architecture ${RUNNER_ARCH}"
  exit 1
fi

# Download the file
gh --repo abcxyz/terraform-linter release download \
  --pattern "terraform-linter_${OS}_${ARCH}*" \
  --output "${BINARY_PATH}"

# Mark as executable
chmod +x "${BINARY_PATH}"

# Save the result to an output.
echo "::notice::Downloaded binary to ${BINARY_PATH}"
echo "binary-path=${BINARY_PATH}" >> "${GITHUB_OUTPUT}"
