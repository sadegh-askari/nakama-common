package api_test

import (
	"github.com/heroiclabs/nakama-common/api"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

func TestRegionalStoreProviders(t *testing.T) {
	for _, store := range []struct {
		value api.StoreProvider
		name  string
	}{{api.StoreProvider_CAFE_BAZAAR, "CAFE_BAZAAR"}, {api.StoreProvider_MYKET, "MYKET"}} {
		if store.value.String() != store.name {
			t.Fatalf("unexpected store name: %s", store.value)
		}
		data, err := protojson.Marshal(&api.ValidatedPurchase{Store: store.value})
		if err != nil {
			t.Fatal(err)
		}
		var purchase api.ValidatedPurchase
		if err := protojson.Unmarshal(data, &purchase); err != nil || purchase.Store != store.value {
			t.Fatalf("store round trip failed: %s %v", data, err)
		}
	}
	if api.StoreProvider_CAFE_BAZAAR != 10 || api.StoreProvider_MYKET != 11 || api.StoreProvider_SAMSUNG_GALAXY_STORE != 4 {
		t.Fatal("store IDs changed")
	}
}
