package modules

type SalesType struct {
	DomesticSalesST      bool `json:"domestic_sales"`
	InternationalSalesST bool `json:"international_sales"`
}

type PurchaseType struct {
	DomesticSales      bool `json:"domestic_sales"`
	InternationalSales bool `json:"international_sales"`
}

type Count struct {
	Warp int `json:"warp"`
	Weft int `json:"weft"`
}
type Density struct {
	WPI int `json:"wpi"`
	CPI int `json:"cpi"`
}
type Weight struct {
	BeforeWash int `json:"before_wash"`
	AfterWash  int `json:"after_wash"`
}
type Width struct {
	Value             int    `json:"value"`
	UnitofMeasurement string `json:"unit_of_measurement"`
}
type SellarInformation struct {
	ID          int
	Address     string
	PhoneNumber string
	Email       string
	SalesType
	PurchaseType
	SalesPrice    int
	PurchasePrice int
	MaterialType  string
	UnitMetric    int
}

type QualityDetails struct {
	ID                     int
	SellarID               int
	Name                   string
	Logo                   string
	ArtNumber              int
	CompositionFibres      string
	Composition            string
	CountWarp              int
	CountWeft              int
	DensityWPI             int
	DensityCPI             int
	WeightBeforeWash       int
	WeightAfterWash        int
	WidthValue             int
	WidthUnitofMeasurement string
	Finish                 string
	Remarks                string
	QRCodeGeneration       string
}

type TechnicalRequirement struct {
	ID                  int    `json:"id"`
	SellarID            int    `json:"sellar_id"`
	Season              string `json:"season"`
	ArticleNumber       string `json:"article_number"`
	MaterialSubType     string `json:"material_sub_type"`
	DyeStuff            string `json:"dye_stuff"`
	MechanicalFinish    string `json:"mechanical_finish"`
	ChemicalFinish      string `json:"chemicalfinish"`
	PerformanceClaims   string `json:"performance_claims"`
	WarpSpinningMethod  string `json:"warp_spinning_method"`
	WeftSpenningMethod  string `json:"weft_spenning_method"`
	WeaveType           string `json:"weave_type"`
	SupplierDeclaration string `json:"supplier_declaration"`
	SpecialInstruction  string `json:"special_instruction"`
}

type SellarInformationReq struct {
	ID            int    `json:"id"`
	SellarName    string `json:"sellar_name"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	SalesType     string `json:"sales_type"`
	PurchaseType  string `json:"purchase_type"`
	SalesPrice    int    `json:"sales_price"`
	PurchasePrice int    `json:"purchase_price"`
	MaterialType  string `json:"material_type"`
	UnitMetric    int    `json:"unit_metric"`
	AdminID       int    `json:"admin_id"`
	// QualityDetails QualityDetailsReq `json:"quality_details"`
}

type QualityDetailsReq struct {
	ID                int     `json:"id"`
	SellarID          int     `json:"sellar_id"`
	Name              string  `json:"name"`
	Logo              string  `json:"logo"`
	ArtNumber         int     `json:"art_number"`
	CompositionFibres string  `json:"composition_fibres"`
	Composition       string  `json:"composition"`
	Count             Count   `json:"count"`
	Density           Density `json:"density"`
	Weight            Weight  `json:"weight"`
	Width             Width   `json:"width"`
	Finish            string  `json:"finish"`
	Remarks           string  `json:"remarks"`
	QRCodeGeneration  string  `json:"qr_code"`
}
