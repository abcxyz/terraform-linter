#!/usr/bin/env bash

# Copyright 2025 The Authors (see AUTHORS file)
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eEuo pipefail

# These are expected to be set externally.
declare -r RUNNER_TEMP GITHUB_SHA RUNNER_OS RUNNER_ARCH GITHUB_OUTPUT

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
  echo "::error::Unsupported operating system ${RUNNER_OS}"
  exit 1
fi

ARCH=""
if [[ "${RUNNER_ARCH}" == "X64" ]]; then
  ARCH="amd64"
elif [[ "${RUNNER_ARCH}" == "ARM64" ]]; then
  ARCH="arm64"
else
  echo "::error::Unsupported system architecture ${RUNNER_ARCH}"
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
