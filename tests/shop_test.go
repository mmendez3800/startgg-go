// Package tests performs a review of the API wrapper methods implemented and confirms that their
// functionality can be utilized as needed.
package tests

import (
	"log"
	"os"
	"testing"

	"github.com/mmendez3800/startgg-go"
)

// TestShop checks the functionality of the methods and confirms that they provide a valid output.
func TestShop(t *testing.T) {
	// Sets the authentication token based on parameter passed for tests.
	startgg.SetAuthToken(os.Getenv("STARTGG_TOKEN"))

	// Implements a level of logging to indicate type.
	ShopLogger := log.New(os.Stdout, "SHOP TESTS - ", log.LstdFlags)

	// Performs test for Shop object by ID.
	idResult, err := startgg.ShopByID(3244)
	if err != nil {
		ShopLogger.Fatalf("- ID Error - %s", err.Error())
	}
	ShopLogger.Printf("- ID Results - %+v", idResult)

	// Performs test for Shop object by slug.
	slugResult, err := startgg.ShopBySlug("shop/mainshop-1")
	if err != nil {
		ShopLogger.Fatalf("- Slug Error - %s", err.Error())
	}
	ShopLogger.Printf("- Slug Results - %+v\n\n", slugResult)
}
