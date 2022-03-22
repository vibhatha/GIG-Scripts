package GIG_Scripts

import "github.com/lsflk/gig-sdk/client"

/**
Set the GIG server API url here for crawlers
 */

var GigClient = client.GigClient{
	ApiUrl:                 "http://localhost:9000/api/",
	ApiKey:                 "$2a$12$dcKw7SVbheBUwWSupp1Pze7zOHJBlcgW2vuQGSEh0QVHC/KUeRgwW",
	NerServerUrl:           "http://localhost:8080/classify",
	NormalizationServerUrl: "http://localhost:9000/api/",
	OcrServerUrl:           "http://localhost:8081/extract?url=",
}
