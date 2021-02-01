package ncloud

import (
	"fmt"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccResourceNcloudAutoScalingPolicy_classic_basic(t *testing.T) {
	var policy AutoScalingPolicy
	name := fmt.Sprintf("terraform-testacc-asp-%s", acctest.RandString(5))
	resourceCHANG := "ncloud_auto_scaling_policy.test-policy-CHANG"
	resourceEXACT := "ncloud_auto_scaling_policy.test-policy-EXACT"
	resourcePRCNT := "ncloud_auto_scaling_policy.test-policy-PRCNT"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccClassicProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckNcloudAutoScalingPolicyDestroy(state, testAccClassicProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNcloudAutoScalingPolicyClassicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNcloudAutoScalingPolicyExists(resourceCHANG, &policy, testAccClassicProvider),
					resource.TestCheckResourceAttr(resourceCHANG, "adjustment_type_code", "CHANG"),
					resource.TestCheckResourceAttr(resourceCHANG, "scaling_adjustment", "2"),
					resource.TestCheckResourceAttr(resourceCHANG, "name", name+"-chang"),

					testAccCheckNcloudAutoScalingPolicyExists(resourceEXACT, &policy, testAccClassicProvider),
					resource.TestCheckResourceAttr(resourceEXACT, "adjustment_type_code", "EXACT"),
					resource.TestCheckResourceAttr(resourceEXACT, "scaling_adjustment", "2"),
					resource.TestCheckResourceAttr(resourceEXACT, "name", name+"-exact"),

					testAccCheckNcloudAutoScalingPolicyExists(resourcePRCNT, &policy, testAccClassicProvider),
					resource.TestCheckResourceAttr(resourcePRCNT, "adjustment_type_code", "PRCNT"),
					resource.TestCheckResourceAttr(resourcePRCNT, "scaling_adjustment", "2"),
					resource.TestCheckResourceAttr(resourcePRCNT, "name", name+"-prcnt"),
				),
			},
		},
	})
}

func TestAccResourceNcloudAutoScalingPolicy_vpc_basic(t *testing.T) {
	var policy AutoScalingPolicy
	name := fmt.Sprintf("terraform-testacc-asp-%s", acctest.RandString(5))
	resourceCHANG := "ncloud_auto_scaling_policy.test-policy-CHANG"
	resourceEXACT := "ncloud_auto_scaling_policy.test-policy-EXACT"
	resourcePRCNT := "ncloud_auto_scaling_policy.test-policy-PRCNT"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckNcloudAutoScalingPolicyDestroy(state, testAccProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNcloudAutoScalingPolicyVpcConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNcloudAutoScalingPolicyExists(resourceCHANG, &policy, testAccProvider),
					resource.TestCheckResourceAttr(resourceCHANG, "adjustment_type_code", "CHANG"),
					resource.TestCheckResourceAttr(resourceCHANG, "scaling_adjustment", "2"),
					resource.TestCheckResourceAttr(resourceCHANG, "name", name+"-chang"),

					testAccCheckNcloudAutoScalingPolicyExists(resourceEXACT, &policy, testAccProvider),
					resource.TestCheckResourceAttr(resourceEXACT, "adjustment_type_code", "EXACT"),
					resource.TestCheckResourceAttr(resourceEXACT, "scaling_adjustment", "2"),
					resource.TestCheckResourceAttr(resourceEXACT, "name", name+"-exact"),

					testAccCheckNcloudAutoScalingPolicyExists(resourcePRCNT, &policy, testAccProvider),
					resource.TestCheckResourceAttr(resourcePRCNT, "adjustment_type_code", "PRCNT"),
					resource.TestCheckResourceAttr(resourcePRCNT, "scaling_adjustment", "2"),
					resource.TestCheckResourceAttr(resourcePRCNT, "name", name+"-prcnt"),
				),
			},
		},
	})
}

func TestAccResourceNcloudAutoScalingPolicy_classic_disappears(t *testing.T) {
	var policy AutoScalingPolicy
	name := fmt.Sprintf("terraform-testacc-asp-%s", acctest.RandString(5))
	resourceCHANG := "ncloud_auto_scaling_policy.test-policy-CHANG"
	resourceEXACT := "ncloud_auto_scaling_policy.test-policy-EXACT"
	resourcePRCNT := "ncloud_auto_scaling_policy.test-policy-PRCNT"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccClassicProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckNcloudAutoScalingPolicyDestroy(state, testAccClassicProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNcloudAutoScalingPolicyClassicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNcloudAutoScalingPolicyExists(resourceCHANG, &policy, testAccClassicProvider),
					testAccCheckNcloudAutoScalingPolicyExists(resourceEXACT, &policy, testAccClassicProvider),
					testAccCheckNcloudAutoScalingPolicyExists(resourcePRCNT, &policy, testAccClassicProvider),

					testAccCheckResourceDisappears(testAccClassicProvider, resourceNcloudAutoScalingPolicy(), resourceCHANG),
					testAccCheckResourceDisappears(testAccClassicProvider, resourceNcloudAutoScalingPolicy(), resourceEXACT),
					testAccCheckResourceDisappears(testAccClassicProvider, resourceNcloudAutoScalingPolicy(), resourcePRCNT),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceNcloudAutoScalingPolicy_vpc_disappears(t *testing.T) {
	var policy AutoScalingPolicy
	name := fmt.Sprintf("terraform-testacc-asp-%s", acctest.RandString(5))
	resourceCHANG := "ncloud_auto_scaling_policy.test-policy-CHANG"
	resourceEXACT := "ncloud_auto_scaling_policy.test-policy-EXACT"
	resourcePRCNT := "ncloud_auto_scaling_policy.test-policy-PRCNT"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckNcloudAutoScalingPolicyDestroy(state, testAccProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNcloudAutoScalingPolicyVpcConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNcloudAutoScalingPolicyExists(resourceCHANG, &policy, testAccProvider),
					testAccCheckNcloudAutoScalingPolicyExists(resourceEXACT, &policy, testAccProvider),
					testAccCheckNcloudAutoScalingPolicyExists(resourcePRCNT, &policy, testAccProvider),

					testAccCheckResourceDisappears(testAccProvider, resourceNcloudAutoScalingPolicy(), resourceCHANG),
					testAccCheckResourceDisappears(testAccProvider, resourceNcloudAutoScalingPolicy(), resourceEXACT),
					testAccCheckResourceDisappears(testAccProvider, resourceNcloudAutoScalingPolicy(), resourcePRCNT),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckNcloudAutoScalingPolicyExists(n string, p *AutoScalingPolicy, provider *schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AutoScalingPolicy ID is set: %s", n)
		}

		config := provider.Meta().(*ProviderConfig)
		autoScalingPolicy, err := getAutoScalingPolicy(config, rs.Primary.ID, rs.Primary.Attributes["auto_scaling_group_no"])
		if err != nil {
			return err
		}
		if autoScalingPolicy == nil {
			return fmt.Errorf("Not found AutoScalingPolicy : %s", rs.Primary.ID)
		}
		*p = *autoScalingPolicy
		return nil
	}
}

func testAccCheckNcloudAutoScalingPolicyDestroy(s *terraform.State, provider *schema.Provider) error {
	config := provider.Meta().(*ProviderConfig)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ncloud_auto_scaling_policy" {
			continue
		}
		autoScalingPolicy, err := getAutoScalingPolicy(config, rs.Primary.ID, rs.Primary.Attributes["auto_scaling_group_no"])
		if err != nil {
			return err
		}

		if autoScalingPolicy != nil {
			return fmt.Errorf("AutoScalingPolicy(%s) still exists", ncloud.StringValue(autoScalingPolicy.AutoScalingPolicyName))
		}
	}
	return nil
}

func testAccNcloudAutoScalingPolicyVpcConfigBase(name string) string {
	return fmt.Sprintf(`
resource "ncloud_vpc" "test" {
	ipv4_cidr_block    = "10.0.0.0/16"
}

resource "ncloud_subnet" "test" {
	vpc_no             = ncloud_vpc.test.vpc_no
	subnet             = "10.0.0.0/24"
	zone               = "KR-2"
	network_acl_no     = ncloud_vpc.test.default_network_acl_no
	subnet_type        = "PUBLIC"
	usage_type         = "GEN"
}

resource "ncloud_launch_configuration" "test" {
    name = "%[1]s"
    server_image_product_code = "SW.VSVR.OS.LNX64.CNTOS.0703.B050"
    server_product_code = "SVR.VSVR.HICPU.C002.M004.NET.SSD.B050.G002"
}

resource "ncloud_auto_scaling_group" "test" {
	name = "%[1]s"
	access_control_group_no_list = [ncloud_vpc.test.default_access_control_group_no]
	subnet_no = ncloud_subnet.test.subnet_no
	launch_configuration_no = ncloud_launch_configuration.test.launch_configuration_no
	min_size = 1
	max_size = 1
}
`, name)
}

func testAccNcloudAutoScalingPolicyVpcConfig(name string) string {
	return testAccNcloudAutoScalingPolicyVpcConfigBase(name) + fmt.Sprintf(`
resource "ncloud_auto_scaling_policy" "test-policy-CHANG" {
    name = "%[1]s-chang"
    adjustment_type_code = "CHANG"
    scaling_adjustment = 2
    auto_scaling_group_no = ncloud_auto_scaling_group.test.auto_scaling_group_no
}

resource "ncloud_auto_scaling_policy" "test-policy-EXACT" {
    name = "%[1]s-exact"
    adjustment_type_code = "EXACT"
    scaling_adjustment = 2
    auto_scaling_group_no = ncloud_auto_scaling_group.test.auto_scaling_group_no
}

resource "ncloud_auto_scaling_policy" "test-policy-PRCNT" {
    name = "%[1]s-prcnt"
    adjustment_type_code = "PRCNT"
    scaling_adjustment = 2
    auto_scaling_group_no = ncloud_auto_scaling_group.test.auto_scaling_group_no
}
`, name)
}

func testAccNcloudAutoScalingPolicyClassicConfigBase(name string) string {
	return fmt.Sprintf(`
resource "ncloud_launch_configuration" "test" {
    name = "%[1]s"
    server_image_product_code = "SPSW0LINUX000046"
    server_product_code = "SPSVRSSD00000003"
}

resource "ncloud_auto_scaling_group" "test" {
	name = "%[1]s"
	launch_configuration_no = ncloud_launch_configuration.test.launch_configuration_no
	min_size = 1
	max_size = 1
	zone_no_list = ["2"]
}
`, name)
}

func testAccNcloudAutoScalingPolicyClassicConfig(name string) string {
	return testAccNcloudAutoScalingPolicyClassicConfigBase(name) + fmt.Sprintf(`
resource "ncloud_auto_scaling_policy" "test-policy-CHANG" {
    name = "%[1]s-chang"
    adjustment_type_code = "CHANG"
    scaling_adjustment = 2
    auto_scaling_group_no = ncloud_auto_scaling_group.test.auto_scaling_group_no
}

resource "ncloud_auto_scaling_policy" "test-policy-EXACT" {
    name = "%[1]s-exact"
    adjustment_type_code = "EXACT"
    scaling_adjustment = 2
    auto_scaling_group_no = ncloud_auto_scaling_group.test.auto_scaling_group_no
}

resource "ncloud_auto_scaling_policy" "test-policy-PRCNT" {
    name = "%[1]s-prcnt"
    adjustment_type_code = "PRCNT"
    scaling_adjustment = 2
    auto_scaling_group_no = ncloud_auto_scaling_group.test.auto_scaling_group_no
}
`, name)
}
