package flexibleengine

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDNSV2Zone_importBasic(t *testing.T) {
	var zoneName = fmt.Sprintf("ACPTTEST%s.com.", acctest.RandString(5))
	resourceName := "flexibleengine_dns_zone_v2.zone_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSV2ZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config:             testAccDNSV2Zone_basic(zoneName),
				ExpectNonEmptyPlan: true,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
