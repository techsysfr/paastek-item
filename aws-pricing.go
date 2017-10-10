package item

// AWSBillingItem is used to represent any line in the billing file
type AWSBillingItem struct {
	IdentityLineItemID            string `csv:"identity/LineItemId"`
	IdentityTimeInterval          string `csv:"identity/TimeInterval"`
	BillInvoiceID                 string `csv:"bill/InvoiceId"`
	BillBillingEntity             string `csv:"bill/BillingEntity"`
	BillBillType                  string `csv:"bill/BillType"`
	BillPayerAccountID            string `csv:"bill/PayerAccountId"`
	BillBillingPeriodStartDate    string `csv:"bill/BillingPeriodStartDate"`
	BillBillingPeriodEndDate      string `csv:"bill/BillingPeriodEndDate"`
	LineItemUsageAccountID        string `csv:"lineItem/UsageAccountId"`
	LineItemLineItemType          string `csv:"lineItem/LineItemType"`
	LineItemUsageStartDate        string `csv:"lineItem/UsageStartDate"`
	LineItemUsageEndDate          string `csv:"lineItem/UsageEndDate"`
	LineItemProductCode           string `csv:"lineItem/ProductCode"`
	LineItemUsageType             string `csv:"lineItem/UsageType"`
	LineItemOperation             string `csv:"lineItem/Operation"`
	LineItemAvailabilityZone      string `csv:"lineItem/AvailabilityZone"`
	LineItemResourceID            string `csv:"lineItem/ResourceId"`
	LineItemUsageAmount           string `csv:"lineItem/UsageAmount"`
	LineItemNormalizationFactor   string `csv:"lineItem/NormalizationFactor"`
	LineItemNormalizedUsageAmount string `csv:"lineItem/NormalizedUsageAmount"`
	LineItemCurrencyCode          string `csv:"lineItem/CurrencyCode"`
	LineItemUnblendedRate         string `csv:"lineItem/UnblendedRate"`
	LineItemUnblendedCost         string `csv:"lineItem/UnblendedCost"`
	LineItemBlendedRate           string `csv:"lineItem/BlendedRate"`
	LineItemBlendedCost           string `csv:"lineItem/BlendedCost"`
	LineItemLineItemDescription   string `csv:"lineItem/LineItemDescription"`
	LineItemTaxType               string `csv:"lineItem/TaxType"`
	NormalizedUnitsPerReservation string `csv:"reservation/NormalizedUnitsPerReservation"`
	TotalReservedNormalizedUnits  string `csv:"reservation/TotalReservedNormalizedUnits"`
	TotalReservedUnits            string `csv:"reservation/TotalReservedUnits"`
	UnitsPerReservation           string `csv:"reservation/UnitsPerReservation"`
	PublicOnDemandCost            string `csv:"pricing/publicOnDemandCost"`
	PublicOnDemandRate            string `csv:"pricing/publicOnDemandRate"`
	Term                          string `csv:"pricing/term"`
	Unit                          string `csv:"pricing/unit"`
	ProductName                   string `csv:"product/ProductName"`
	Availability                  string `csv:"product/availability"`
	ClockSpeed                    string `csv:"product/clockSpeed"`
	CurrentGeneration             string `csv:"product/currentGeneration"`
	Durability                    string `csv:"product/durability"`
	Ecu                           string `csv:"product/ecu"`
	FromLocation                  string `csv:"product/fromLocation"`
	FromLocationType              string `csv:"product/fromLocationType"`
	Group                         string `csv:"product/group"`
	GroupDescription              string `csv:"product/groupDescription"`
	InstanceFamily                string `csv:"product/instanceFamily"`
	InstanceType                  string `csv:"product/instanceType"`
	InstanceTypeFamily            string `csv:"product/instanceTypeFamily"`
	LicenseModel                  string `csv:"product/licenseModel"`
	Location                      string `csv:"product/location"`
	LocationType                  string `csv:"product/locationType"`
	MaxIopsBurstPerformance       string `csv:"product/maxIopsBurstPerformance"`
	MaxIopsvolume                 string `csv:"product/maxIopsvolume"`
	MaxThroughputvolume           string `csv:"product/maxThroughputvolume"`
	MaxVolumeSize                 string `csv:"product/maxVolumeSize"`
	Memory                        string `csv:"product/memory"`
	NetworkPerformance            string `csv:"product/networkPerformance"`
	NormalizationSizeFactor       string `csv:"product/normalizationSizeFactor"`
	OperatingSystem               string `csv:"product/operatingSystem"`
	Operation                     string `csv:"product/operation"`
	PhysicalProcessor             string `csv:"product/physicalProcessor"`
	PreInstalledSw                string `csv:"product/preInstalledSw"`
	ProcessorArchitecture         string `csv:"product/processorArchitecture"`
	ProcessorFeatures             string `csv:"product/processorFeatures"`
	ProductFamily                 string `csv:"product/productFamily"`
	Provisioned                   string `csv:"product/provisioned"`
	Region                        string `csv:"product/region"`
	Servicecode                   string `csv:"product/servicecode"`
	Servicename                   string `csv:"product/servicename"`
	Sku                           string `csv:"product/sku"`
	Storage                       string `csv:"product/storage"`
	StorageClass                  string `csv:"product/storageClass"`
	StorageMedia                  string `csv:"product/storageMedia"`
	Tenancy                       string `csv:"product/tenancy"`
	ToLocation                    string `csv:"product/toLocation"`
	ToLocationType                string `csv:"product/toLocationType"`
	TransferType                  string `csv:"product/transferType"`
	Usagetype                     string `csv:"product/usagetype"`
	Vcpu                          string `csv:"product/vcpu"`
	VolumeType                    string `csv:"product/volumeType"`
}
