// Copyright 2025 The Authors (see AUTHORS file)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rules

var (
	// Naming.
	HyphenInName = &Rule{"TF001", `Resource name must not contain a "-". Prefer underscores ("_") instead.`}

	// Ordering.
	ProviderNewline  = &Rule{"TF050", `Provider-specific attributes must have an additional newline separating them from the next section.`}
	MetaBlockNewline = &Rule{"TF051", `Meta block must have an additional newline separating it from the next section.`}

	// Ordering.
	LeadingMetaBlockAttribute  = &Rule{"TF100", `Attribute must be in the meta block at the top of the definition.`}
	ProviderAttributes         = &Rule{"TF101", `Attribute must be below any meta attributes (e.g. "for_each", "count") but above all other attributes. Attributes must be ordered organization > folder > project.`}
	TrailingMetaBlockAttribute = &Rule{"TF199", `Attribute must be at the bottom of the resource definition and in the order "depends_on" then "lifecycle."`}
)

// Rule represents a single rule entry.
type Rule struct {
	// ID is the unique identifier for the rule, used in ignore clauses.
	ID string

	// Description is the human-friendly description of the rule.
	Description string
}
