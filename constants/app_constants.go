package constants

var INVALID_AREA_CODE string = "area code is not valid"
var INVALID_CC_NUMBER string = "number of credit cards are not valid"
var INVALID_AGE string = "age is not valid"
var INVALID_INCOME string = "income is not valid"
var INVALID_CREDIT_RISK_SCORE string = "credit risk score is not low"
var POLITICALLY_EXPOSED string = "applicant is politically exposed"
var PREAPPROVED_NUMBER string = "number is preapproved"
var APPROVED string = "approved"
var DECLINED string = "declined"
var NUMBER_LOGGED string = "number logged"

// constants for logging
var LOG_LEVEL_INFO string = "[INFO]   "
var LOG_LEVEL_ERROR string = "[ERROR]  "
var LOG_LEVEL_WARN string = "[WARNING]"

// file paths
var CONFIG_FILE string = "CONFIG_PATH"
var NUMBERS_FILE string = "APPROVED_NUMBERS_FILE_PATH"
var LOG_FILE_PATH string = "LOG_FILE_PATH"

// validator constants
var NO_FIELDS_BLANK string = "please check the input fields"
var INCOME_NEGATIVE string = "income cannot be negative"
var CC_NEGATIVE string = "number of credit cards cannot be negative"
var AGE_NEGATIVE string = "age cannot be negative"
var INVALID_PHONE string = "phone number is not valid"

// json record location
var JSON_RECORDS_5 string = ".././5-records.json"
var JSON_RECORDS_1000 string = ".././1000-records.json"
