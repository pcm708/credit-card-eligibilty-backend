package constants

// constants for the application

// configs
var MIN_AGE int = 18
var MIN_INCOME int = 100000
var MAX_NUMBER_OF_CC int = 3
var ALLOWED_AREA_CODE = []int{0, 2, 5, 8}
var DESIRED_CREDIT_RISK_SCORE string = "LOW"

// rate limiter
var MAX_REQUESTS int = 10
var RATE_LIMITER_DURATION int = 1

// server properties
var CLOUD_DB_URL string = "192.168.1.107"
var DB_PORT string = "3306"
var REDIS_BASE_URL string = "192.168.1.107"
var REDIS_PORT string = "6379"

// constants for decision
var INVALID_AREA_CODE string = "area code is not valid"
var INVALID_CC_NUMBER string = "number of credit cards are not valid"
var INVALID_AGE string = "age is not valid"
var INVALID_INCOME string = "income is not valid"
var INVALID_CREDIT_RISK_SCORE string = "credit risk score is not low"
var POLITICALLY_EXPOSED string = "applicant is politically exposed"
var PREAPPROVED_NUMBER string = "number pre-approved"
var APPROVED string = "approved"
var DECLINED string = "declined"
var NUMBER_LOGGED string = "number stored in DB"

// constants for logging
var LOG_LEVEL_INFO string = "[INFO]    :: "
var LOG_LEVEL_ERROR string = "[ERROR]   :: "
var LOG_LEVEL_WARN string = "[WARNING] :: "

// file paths
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
