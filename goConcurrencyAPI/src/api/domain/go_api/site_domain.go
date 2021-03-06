package go_api

type Site struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryId          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyId  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	/*Settings           struct {
		IdentificationTypes      []string `json:"identification_types"`
		TaxpayerTypes            []string `json:"taxpayer_types"`
		IdentificationTypesRules []struct {
			IdentificationType string `json:"identification_type"`
			Rules              []struct {
				EnabledTaxpayerTypes []string `json:"enabled_taxpayer_types"`
				BeginsWith           string   `json:"begins_with"`
				Type                 string   `json:"type"`
				MinLength            int      `json:"min_length"`
				MaxLength            int      `json:"max_length"`
			} `json:"rules"`
		} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`*/
}

func (s *Site) SaberNombreSitio() string{
	return s.Name
}
