// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package shield_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSShieldProactiveEngagementsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Shield::ProactiveEngagement", "awscc_shield_proactive_engagements", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}