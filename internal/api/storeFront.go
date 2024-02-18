package api

import (
	"encoding/json"
	"fmt"

	"github.com/itsbradn/govalo/pkg/http"
)

type StoreFrontResponseBody struct {
	FeaturedBundle struct {
		Bundle struct {
			ID          string `json:"ID"`
			DataAssetID string `json:"DataAssetID"`
			CurrencyID  string `json:"CurrencyID"`
			Items       []struct {
				Item struct {
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
					Amount     uint8  `json:"Amount"`
				} `json:"Item"`
				CurrencyID      string  `json:"CurrencyID"`
				DiscountPercent float32 `json:"DiscountPercent"`
				DiscountPrice   uint32  `json:"DiscountPrice"`
				IsPromoItem     bool    `json:"IsPromoItem"`
			} `json:"Items"`
			ItemOffers []struct {
				BundleItemOfferID string `json:"BundleItemOfferID"`
				Offer             struct {
					OfferID          string           `json:"OfferID"`
					IsDirectPurchase bool             `json:"IsDirectPurchase"`
					StartDate        string           `json:"StartDate"` // ISO 8601 Date
					Cost             map[string]int32 `json:"Cost"`
					Rewards          []struct {
						ItemTypeID string `json:"ItemTypeID"`
						ItemID     string `json:"ItemID"`
						Quantity   uint16 `json:"Quantity"`
					} `json:"Rewards"`
				} `json:"Offer"`
				DiscountPercent float32           `json:"DiscountPercent"`
				DiscountedCost  map[string]uint32 `json:"DiscountedCost"`
			} `json:"ItemOffers"`
			TotalBaseCost              map[string]uint32 `json:"TotalBaseCost"`
			TotalDiscountedCost        map[string]uint32 `json:"TotalDiscountedCost"`
			TotalDiscountPercent       uint32            `json:"TotalDiscountPercent"`
			DurationRemainingInSeconds uint64            `json:"DurationRemainingInSeconds"`
			WholesaleOnly              bool              `json:"WholesaleOnly"`
		} `json:"Bundle"`
		Bundles []struct {
			ID          string `json:"ID"`
			DataAssetID string `json:"DataAssetID"`
			CurrencyID  string `json:"CurrencyID"`
			Items       []struct {
				Item struct {
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
					Amount     uint8  `json:"Amount"`
				} `json:"Item"`
				BasePrice       uint32  `json:"BasePrice"`
				CurrencyID      string  `json:"CurrencyID"`
				DiscountPercent float32 `json:"DiscountPercent"`
				DiscountedPrice uint32  `json:"DiscountedPrice"`
				IsPromoItem     bool    `json:"IsPromoItem"`
			} `json:"Items"`
			ItemOffers []struct {
				BundleItemOfferID string `json:"BundleItemOfferID"`
				Offer             struct {
					OfferID          string           `json:"OfferID"`
					IsDirectPurchase bool             `json:"IsDirectPurchase"`
					StartDate        string           `json:"StartDate"` // ISO 8601 Date
					Cost             map[string]int32 `json:"Cost"`
					Rewards          []struct {
						ItemTypeID string `json:"ItemTypeID"`
						ItemID     string `json:"ItemID"`
						Quantity   uint16 `json:"Quantity"`
					} `json:"Rewards"`
				} `json:"Offer"`
				DiscountPercent float32           `json:"DiscountPercent"`
				DiscountedCost  map[string]uint32 `json:"DiscountedCost"`
			} `json:"ItemOffers"`
			TotalBaseCost              map[string]uint32 `json:"TotalBaseCost"`
			TotalDiscountedCost        map[string]uint32 `json:"TotalDiscountedCost"`
			TotalDiscountPercent       float32           `json:"TotalDiscountPercent"`
			DurationRemainingInSeconds uint64            `json:"DurationRemainingInSeconds"`
			WholesaleOnly              bool              `json:"WholesaleOnly"`
		} `json:"Bundles"`
		BundleRemainingDurationInSeconds uint32
	} `json:"FeaturedBundle"`
	SkinsPanelLayout struct {
		SingleItemOffers      []string `json:"SingleItemOffers"`
		SingleItemStoreOffers []struct {
			OfferID          string            `json:"OfferID"`
			IsDirectPurchase bool              `json:"IsDirectPurchase"`
			StartDate        string            `json:"StartDate"` // ISO 8601 Date
			Cost             map[string]uint32 `json:"Cost"`
			Rewards          []struct {
				ItemTypeID string `json:"ItemTypeID"`
				ItemID     string `json:"ItemID"`
				Quantity   uint8  `json:"Quantity"`
			} `json:"Rewards"`
		} `json:"SingleItemStoreOffers"`
		SingleItemOffersRemainingDurationInSeconds uint32 `json:"SingleItemOffersRemainingDurationInSeconds"`
	} `json:"SkinsPanelLayout"`
	UpgradeCurrencyStore struct {
		UpgradeCurrencyOffers []struct {
			OfferID          string `json:"OfferID"`
			StorefrontItemID string `json:"StorefrontItemID"`
			Offer            struct {
				OfferID          string            `json:"OfferID"`
				IsDirectPurchase bool              `json:"IsDirectPurchase"`
				StartDate        string            `json:"StartDate"` // ISO 8601 Date
				Cost             map[string]uint32 `json:"Cost"`
				Rewards          []struct {
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
					Quantity   uint8  `json:"Quantity"`
				} `json:"Rewards"`
			} `json:"Offer"`
			DiscountedPercent float32 `json:"DiscountedPercent"`
		} `json:"UpgradeCurrencyOffers"`
	} `json:"UpgradeCurrencyStore"`
	AccessoryStore struct {
		AccessoryStoreOffers []struct {
			Offer struct {
				OfferID          string            `json:"OfferID"`
				IsDirectPurchase bool              `json:"IsDirectPurchase"`
				StartDate        string            `json:"StartDate"` // ISO 1806 Date
				Cost             map[string]uint32 `json:"Cost"`
				Rewards          []struct {
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
					Quantity   uint8  `json:"Quantity"`
				} `json:"Rewards"`
			} `json:"Offer"`
			ContractID string `json:"ContractID"`
		} `json:"AccessoryStoreOffers"`
		AccessoryStoreRemainingDurationInSeconds uint32 `json:"AccessoryStoreRemainingDurationInSeconds"`
		StorefrontID                             string `json:"StorefrontID"`
	} `json:"AccessoryStore"`
	NightMarket struct {
		Offers []struct {
			BonusOfferID string `json:"BonusOfferID"`
			Offer        struct {
				OfferID          string            `json:"OfferID"`
				IsDirectPurchase bool              `json:"IsDirectPurchase"`
				StartDate        string            `json:"StartDate"` // ISO 1806 Date
				Cost             map[string]uint32 `json:"Cost"`
				Rewards          []struct {
					ItemTypeID string `json:"ItemTypeID"`
					ItemID     string `json:"ItemID"`
					Quantity   uint8  `json:"Quantity"`
				} `json:"Rewards"`
			} `json:"Offer"`
			DiscountPercent float32           `json:"DiscountPercent"`
			DiscountCosts   map[string]uint32 `json:"DiscountCosts"`
			IsSeen          bool              `json:"IsSeen"`
		} `json:"BonusStoreOffers"`
		NightMarketRemainingDurationInSeconds uint32 `json:"BonusStoreRemainingDurationInSeconds"`
	} `json:"BonusStore"`
}

func GetStorefront(shard, puuid string) (*StoreFrontResponseBody, error) {
	res, err := http.SendRequest("GET", fmt.Sprintf("https://pd.%s.a.pvp.net/store/v2/storefront/%s", shard, puuid), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body *StoreFrontResponseBody
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
