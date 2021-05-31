package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID string `json:"id"`
	Title         string `json:"title"`
	ArtistName    string `json:"artist_name"`
	CreationDate  string `json:"creation_date"`
	ArtMedium     string `json:"art_medium"`
	Dimension     string `json:"dimension"`
	Description   string `json:"description"`
	PriceValue    string `json:"price_value"`
	LastSoldPrice string `json:"last_sold_price"`
	LastSoldDate  string `json:"last_sold_date"`
	Owner         string `json:"owner"`
	PreviousOwner string `json:"previous_owner"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "asset1" , Title: "Fleurs dans un pot (Roses et brouillard)", ArtistName: "Claude Monet", CreationDate: "1878", ArtMedium: "Oil on canvas", Dimension: "33 x 24 in", Description: "Depicts flowers in a pot.", PriceValue: "5,000,000 USD", LastSoldPrice: "2,500,000 USD", LastSoldDate: "1978", Owner: "James Bond", PreviousOwner: "Dr. No"},
		{ID: "asset2" ,Title: "Le Bassin aux nympheas", ArtistName: "Claude Monet", CreationDate: "1917", ArtMedium: "Oil on canvas", Dimension: "39 by 79 in", Description: "The work, which dates from circa 1917-20, is a powerful testament to Monet’s enduring vision and creativity in his mature years. Le Bassin aux nympheas triumphantly achieves monuments of color; the water reflects the skies’ shifting hues and the lilies themselves are elegant touches of paint applied with bravura. ", PriceValue: "50,000,000 USD", LastSoldPrice: "17,940,000 USD", LastSoldDate: "1983", Owner: "Dr. No", PreviousOwner: "Goldfinger"},
		{ID: "asset3" ,Title: "La Seine a Lavacourt, debacle", ArtistName: " Claude Monet", CreationDate: "1881", ArtMedium: "Oil on canvas", Dimension: "24 by 39 in", Description: "", PriceValue: "7,000,000 USD", LastSoldPrice: "N/A", LastSoldDate: "N/A", Owner: "N/A", PreviousOwner: "George Vanderbilt II"},
		{ID: "asset4" ,Title: "Koropokkur in the Forest", ArtistName: "Takashi Murakami", CreationDate: "2020", ArtMedium: "Offset Print", Dimension: "30 x 36.2 in", Description: "na", PriceValue: "1,387 - 1,859 USD", LastSoldPrice: "na", LastSoldDate: " 5/20/2021", Owner: "na", PreviousOwner: "na"},
		{ID: "asset5" ,Title: "Waterfall 5", ArtistName: "Barry Masteller", CreationDate: "1997", ArtMedium: "Oil on canvas", Dimension: "62 x 50 in", Description: "na", PriceValue: "23,000 USD", LastSoldPrice: "na", LastSoldDate: "na", Owner: "Barry Masteller", PreviousOwner: "na"},
		{ID: "asset6" ,Title: "Pimp My Ride", ArtistName: "Vanilla Ice", CreationDate: "2001", ArtMedium: "Colored pencil on notebook paper", Dimension: "50 by 50", Description: "Pimp My Ride is an American television series produced by MTV and hosted by rapper Xzibit, which ran on MTV for six seasons from 2004 to 2007. Each episode consists of taking one car in poor condition and restoring it, as well as customizing it.", PriceValue: "678,000 USD", LastSoldPrice: "2,000 USD", LastSoldDate: "2009", Owner: "DMX", PreviousOwner: "Sheeeeesh"},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface,
	id string,
	title string,
	artistName string,
	creationDate string,
	artMedium string,
	dimension string,
	description string,
	priceValue string,
	lastSoldPrice string,
	lastSoldDate string,
	owner string,
	previousOwner string) error {

	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID: id,
		Title:         title,
		ArtistName:    artistName,
		CreationDate:  creationDate,
		ArtMedium:     artMedium,
		Dimension:     dimension,
		Description:   description,
		PriceValue:    priceValue,
		LastSoldPrice: lastSoldPrice,
		LastSoldDate:  lastSoldDate,
		Owner:         owner,
		PreviousOwner: previousOwner,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface,
	id string,
	title string,
	artistName string,
	creationDate string,
	artMedium string,
	dimension string,
	description string,
	priceValue string,
	lastSoldPrice string,
	lastSoldDate string,
	owner string,
	previousOwner string) error {

	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ID: id,
		Title:         title,
		ArtistName:    artistName,
		CreationDate:  creationDate,
		ArtMedium:     artMedium,
		Dimension:     dimension,
		Description:   description,
		PriceValue:    priceValue,
		LastSoldPrice: lastSoldPrice,
		LastSoldDate:  lastSoldDate,
		Owner:         owner,
		PreviousOwner: previousOwner,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string, date string, price string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.PreviousOwner = asset.Owner
	asset.Owner = newOwner
	asset.LastSoldDate = date
	asset.LastSoldPrice = price

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
