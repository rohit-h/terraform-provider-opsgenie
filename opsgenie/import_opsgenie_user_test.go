package opsgenie

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccOpsGenieUser_importBasic(t *testing.T) {
	resourceName := "opsgenie_user.test"

	ri := acctest.RandInt()
	config := testAccOpsGenieUser_basic(ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckOpsGenieUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccOpsGenieUser_importComplete(t *testing.T) {
	resourceName := "opsgenie_user.test"

	ri := acctest.RandInt()
	config := testAccOpsGenieUser_complete(ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckOpsGenieUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
