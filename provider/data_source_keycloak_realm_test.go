package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccKeycloakDataSourceRealm_basic(t *testing.T) {
	realm := "terraform-" + acctest.RandString(10)

	resourceName := "keycloak_realm.realm"
	dataSourceName := "data.keycloak_realm.realm"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckKeycloakRealmDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testDataSourceKeycloakRealm_basic(realm),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "realm", resourceName, "realm"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
				),
			},
		},
	})
}

func testDataSourceKeycloakRealm_basic(realm string) string {
	return fmt.Sprintf(`
resource "keycloak_realm" "realm" {
	realm        = "%s"
	display_name = "foo"
}

data "keycloak_realm" "realm" {
	realm = "${keycloak_realm.realm.realm}"
}`, realm)
}
