package fabricHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bhaskarhc/admin-nestiin/db"
	"github.com/bhaskarhc/admin-nestiin/modules"
	apiresponse "github.com/bhaskarhc/admin-nestiin/utils/apiResponse"
)

func Fabric(w http.ResponseWriter, r *http.Request) {
	resp := apiresponse.Message(200, true, "fabric Endpoint working")
	apiresponse.Respond(w, resp)
}

func SellarDetailsHandler(w http.ResponseWriter, r *http.Request) {
	rb := modules.SellarInformationReq{}
	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		fmt.Println("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sellarStruct := modules.SellarInformationReq{
		SellarName:    rb.SellarName,
		Address:       rb.Address,
		PhoneNumber:   rb.PhoneNumber,
		Email:         rb.Email,
		SalesType:     rb.SalesType,
		PurchaseType:  rb.PurchaseType,
		SalesPrice:    rb.SalesPrice,
		PurchasePrice: rb.PurchasePrice,
		MaterialType:  rb.MaterialType,
		UnitMetric:    rb.UnitMetric,
		AdminID:       r.Context().Value("user").(int),
	}
	err = db.GetDB().Create(&sellarStruct).Error
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusConflict)
		return
	}
	fmt.Print("\n --inserted-- \n")
}

func QualityRequirements(w http.ResponseWriter, r *http.Request) {

	Qr := modules.QualityDetailsReq{}
	err := json.NewDecoder(r.Body).Decode(&Qr)
	if err != nil {
		fmt.Println("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	QualityReq := modules.QualityDetails{
		SellarID:               Qr.SellarID,
		Name:                   Qr.Name,
		Logo:                   Qr.Logo,
		ArtNumber:              Qr.ArtNumber,
		CompositionFibres:      Qr.CompositionFibres,
		Composition:            Qr.Composition,
		CountWarp:              Qr.Count.Warp,
		CountWeft:              Qr.Count.Weft,
		DensityWPI:             Qr.Density.WPI,
		DensityCPI:             Qr.Density.CPI,
		WeightBeforeWash:       Qr.Weight.BeforeWash,
		WeightAfterWash:        Qr.Weight.AfterWash,
		WidthValue:             Qr.Width.Value,
		WidthUnitofMeasurement: Qr.Width.UnitofMeasurement,
		Finish:                 Qr.Finish,
		Remarks:                Qr.Remarks,
		QRCodeGeneration:       "QR code generation URl",
	}

	err = db.GetDB().Create(&QualityReq).Error
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusConflict)
		return
	}
	fmt.Print("\n --inserted-- \n")

}

func TechnicalRequirement(w http.ResponseWriter, r *http.Request) {

	Tr := modules.TechnicalRequirement{}
	err := json.NewDecoder(r.Body).Decode(&Tr)
	if err != nil {
		fmt.Println("invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.GetDB().Create(&Tr).Error
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusConflict)
		return
	}
	fmt.Print("\n --inserted-- \n")

}
