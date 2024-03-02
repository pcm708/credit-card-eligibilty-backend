package writer

type IWriter interface {
	StorePreApprovedNumber(phoneNumber string) error
	LogToJSON(phoneNumber string, message string, status string, loglevel string) error
}
