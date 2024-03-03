package writer

type IWriter interface {
	StorePreApprovedNumber(phoneNumber string) error
	LogToJSON(uuid string, message string, status string, loglevel string) error
}
