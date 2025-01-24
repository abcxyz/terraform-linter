# Terraform Linter and GitHub Action

This repository contains the opinionated Terraform linter for abcxyz and the corresponding GitHub Action to run those linting checks.

## Usage

```yaml
name: 'terraform-lint'
on:
  push:
    branches:
      - 'main'
  pull_request:
    branches:
      - 'main'
  workflow_dispatch:

jobs:
  lint:
    uses: 'abcxyz/terraform-linter@main'
    with:
      terraform_version: '1.2'
      directory: './terraform'
```
