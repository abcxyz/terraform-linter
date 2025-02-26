// Copyright 2025 The Authors (see AUTHORS file)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"context"
	"fmt"

	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"

	"github.com/abcxyz/pkg/cli"
	"github.com/abcxyz/terraform-linter/internal/terraformlinter"
)

var _ cli.Command = (*LintCommand)(nil)

// LintCommand lints terraform configurations.
type LintCommand struct {
	cli.BaseCommand

	excludePaths []string
	ignoreRules  []string
}

func (c *LintCommand) Desc() string {
	return `Lint Terraform configurations`
}

func (c *LintCommand) Help() string {
	return `
Usage: {{ COMMAND }} [options] FILE_OR_DIRECTORY [...FILE_OR_DIRECTORY]

Lint Terraform file:

      {{ COMMAND }} main.tf

Lint all Terraform configurations in a directory:

      {{ COMMAND }} terraform/

The command will only lint .tf and .tf.json files.
`
}

func (c *LintCommand) PredictArgs() complete.Predictor {
	return predict.Or(
		predict.Dirs(""),
		predict.Files("*.tf"),
		predict.Files("*.tf.json"),
	)
}

func (c *LintCommand) Flags() *cli.FlagSet {
	set := cli.NewFlagSet()

	// Command options
	f := set.NewSection("COMMAND OPTIONS")

	f.StringSliceVar(&cli.StringSliceVar{
		Name: "exclude-path",
		Usage: "List of Terraform files or directories to exclude from linting. " +
			"This option can be specified multiple times to exclude multiple paths " +
			"or directories.",
		Example: "<pattern>",
		EnvVar:  "TFLINT_EXCLUDE_PATHS",
		Predict: predict.Or(
			predict.Dirs(""),
			predict.Files("*.tf"),
			predict.Files("*.tf.json"),
		),
		Target: &c.excludePaths,
	})

	f.StringSliceVar(&cli.StringSliceVar{
		Name: "ignore-rule",
		Usage: "List of linter rules to ignore. This option can be specified " +
			"multiple times to ignore multiple rules.",
		Example: "<rule>",
		EnvVar:  "TFLINT_IGNORE_RULES",
		Target:  &c.ignoreRules,
	})

	return set
}

func (c *LintCommand) Run(ctx context.Context, args []string) error {
	f := c.Flags()
	if err := f.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	args = f.Args()
	if got := len(args); got < 1 {
		return fmt.Errorf("expected at least one argument, got %d", got)
	}

	linter, err := terraformlinter.New(&terraformlinter.Config{
		ExcludePaths: c.excludePaths,
		IgnoreRules:  c.ignoreRules,
	})
	if err != nil {
		return fmt.Errorf("failed to configure linter: %w", err)
	}

	if err := linter.Run(ctx, args); err != nil {
		return fmt.Errorf("failed to run linter: %w", err)
	}

	findings := linter.Findings()

	// Print out each violation.
	// TODO(sethvargo): support other formats
	for _, finding := range findings {
		fmt.Fprintln(c.Stdout(), finding.String())
	}

	switch l := len(findings); l {
	case 0:
		return nil
	case 1:
		return fmt.Errorf("found 1 violation")
	default:
		return fmt.Errorf("found %d violations", l)
	}
}
