package domain

type MetadataResponse struct {
	TransactionID string `json:"transaction_id"`
	Timestamp     string `json:"timestamp"`
	TimeElapsed   string `json:"time_elapsed"`
}

type Response struct {
	Data map[string]interface{} `json:"data"`
	Meta MetadataResponse       `json:"meta"`
}

type SuccessResponse struct {
	StatusCode int `json:"-"`
	Response   Response
}

type FailureResponse struct {
	StatusCode int `json:"-"`
	Response   Response
}

func GenerateResponse(data map[string]interface{}, typeResponse, transactionID, timestamp, timeElapsed string, statusCode int) interface{} {
	meta := MetadataResponse{
		transactionID,
		timestamp,
		timeElapsed,
	}
	response := Response{
		data,
		meta,
	}

	if typeResponse == "failure" {
		return FailureResponse{
			statusCode,
			response,
		}
	}

	return SuccessResponse{
		statusCode,
		response,
	}
}
