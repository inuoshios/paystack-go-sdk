package paystack

type InitializeTransaction struct {
	Amount            string   `json:"amount,omitempty"`
	Email             string   `json:"email,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	CallbackURL       string   `json:"callback_url,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      uint     `json:"invoice_limit,omitempty"`
	Metadata          string   `json:"metadata,omitempty"`
	Channels          []string `json:"channels,omitempty"`
	SplitCode         string   `json:"split_code,omitempty"`
	Subaccount        string   `json:"subaccount,omitempty"`
	TransactionCharge uint     `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
}
