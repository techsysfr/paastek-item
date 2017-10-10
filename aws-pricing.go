package item

// AWSBillingItem is used to represent any line in the billing file
type AWSBillingItem struct {
	identityLineItemID            string         `csv:"identity/LineItemId"`
	identityTimeInterval          string         `csv:"identity/TimeInterval"`
	billInvoiceID                 string         `csv:"bill/InvoiceId"`
	billBillingEntity             string         `csv:"bill/BillingEntity"`
	billBillType                  string         `csv:"bill/BillType"`
	billPayerAccountID            string         `csv:"bill/PayerAccountId"`
	billBillingPeriodStartDate    string         `csv:"bill/BillingPeriodStartDate"`
	billBillingPeriodEndDate      string         `csv:"bill/BillingPeriodEndDate"`
	lineItemUsageAccountID        string         `csv:"lineItem/UsageAccountId"`
	lineItemLineItemType          string         `csv:"lineItem/LineItemType"`
	lineItemUsageStartDate        string         `csv:"lineItem/UsageStartDate"`
	lineItemUsageEndDate          string         `csv:"lineItem/UsageEndDate"`
	lineItemProductCode           string         `csv:"lineItem/ProductCode"`
	lineItemUsageType             string         `csv:"lineItem/UsageType"`
	lineItemOperation             string         `csv:"lineItem/Operation"`
	lineItemAvailabilityZone      string         `csv:"lineItem/AvailabilityZone"`
	lineItemResourceID            string         `csv:"lineItem/ResourceId"`
	lineItemUsageAmount           string         `csv:"lineItem/UsageAmount"`
	lineItemNormalizationFactor   string         `csv:"lineItem/NormalizationFactor"`
	lineItemNormalizedUsageAmount string         `csv:"lineItem/NormalizedUsageAmount"`
	lineItemCurrencyCode          string         `csv:"lineItem/CurrencyCode"`
	lineItemUnblendedRate         string         `csv:"lineItem/UnblendedRate"`
	lineItemUnblendedCost         string         `csv:"lineItem/UnblendedCost"`
	lineItemBlendedRate           string         `csv:"lineItem/BlendedRate"`
	lineItemBlendedCost           string         `csv:"lineItem/BlendedCost"`
	lineItemLineItemDescription   string         `csv:"lineItem/LineItemDescription"`
	lineItemTaxType               string         `csv:"lineItem/TaxType"`
	Product                       AWSProduct     `csv:",inline"`
	Pricing                       AWSPricing     `csv:",inline"`
	Reservation                   AWSReservation `csv:",inline"`
}

// AWSReservation ...
type AWSReservation struct {
	NormalizedUnitsPerReservation string `csv:"reservation/NormalizedUnitsPerReservation"`
	TotalReservedNormalizedUnits  string `csv:"reservation/TotalReservedNormalizedUnits"`
	TotalReservedUnits            string `csv:"reservation/TotalReservedUnits"`
	UnitsPerReservation           string `csv:"reservation/UnitsPerReservation"`
}

// AWSPricing ...
type AWSPricing struct {
	PublicOnDemandCost string `csv:"pricing/publicOnDemandCost"`
	PublicOnDemandRate string `csv:"pricing/publicOnDemandRate"`
	Term               string `csv:"pricing/term"`
	Unit               string `csv:"pricing/unit"`
}

// AWSProduct ...
type AWSProduct struct {
	ProductName             string `csv:"product/ProductName"`
	Availability            string `csv:"product/availability"`
	ClockSpeed              string `csv:"product/clockSpeed"`
	CurrentGeneration       string `csv:"product/currentGeneration"`
	Durability              string `csv:"product/durability"`
	Ecu                     string `csv:"product/ecu"`
	FromLocation            string `csv:"product/fromLocation"`
	FromLocationType        string `csv:"product/fromLocationType"`
	Group                   string `csv:"product/group"`
	GroupDescription        string `csv:"product/groupDescription"`
	InstanceFamily          string `csv:"product/instanceFamily"`
	InstanceType            string `csv:"product/instanceType"`
	InstanceTypeFamily      string `csv:"product/instanceTypeFamily"`
	LicenseModel            string `csv:"product/licenseModel"`
	Location                string `csv:"product/location"`
	LocationType            string `csv:"product/locationType"`
	MaxIopsBurstPerformance string `csv:"product/maxIopsBurstPerformance"`
	MaxIopsvolume           string `csv:"product/maxIopsvolume"`
	MaxThroughputvolume     string `csv:"product/maxThroughputvolume"`
	MaxVolumeSize           string `csv:"product/maxVolumeSize"`
	Memory                  string `csv:"product/memory"`
	NetworkPerformance      string `csv:"product/networkPerformance"`
	NormalizationSizeFactor string `csv:"product/normalizationSizeFactor"`
	OperatingSystem         string `csv:"product/operatingSystem"`
	Operation               string `csv:"product/operation"`
	PhysicalProcessor       string `csv:"product/physicalProcessor"`
	PreInstalledSw          string `csv:"product/preInstalledSw"`
	ProcessorArchitecture   string `csv:"product/processorArchitecture"`
	ProcessorFeatures       string `csv:"product/processorFeatures"`
	ProductFamily           string `csv:"product/productFamily"`
	Provisioned             string `csv:"product/provisioned"`
	Region                  string `csv:"product/region"`
	Servicecode             string `csv:"product/servicecode"`
	Servicename             string `csv:"product/servicename"`
	Sku                     string `csv:"product/sku"`
	Storage                 string `csv:"product/storage"`
	StorageClass            string `csv:"product/storageClass"`
	StorageMedia            string `csv:"product/storageMedia"`
	Tenancy                 string `csv:"product/tenancy"`
	ToLocation              string `csv:"product/toLocation"`
	ToLocationType          string `csv:"product/toLocationType"`
	TransferType            string `csv:"product/transferType"`
	Usagetype               string `csv:"product/usagetype"`
	Vcpu                    string `csv:"product/vcpu"`
	VolumeType              string `csv:"product/volumeType"`
}
