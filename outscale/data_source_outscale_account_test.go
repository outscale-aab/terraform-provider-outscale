package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAccount_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAccountConfig(),
			},
		},
	})
}

func testAccDataSourceAccountConfig() string {
	return fmt.Sprintf(`
              data "outscale_account" "account" { }
	`)
}
