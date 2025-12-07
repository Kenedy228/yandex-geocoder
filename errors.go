package geocoder

import "errors"

var ErrInternalServer = errors.New("internal server error")
var ErrInternalApi = errors.New("internal api error")
var ErrInvalidResponse = errors.New("invalid response from yandex geocoder")
var ErrBadParams = errors.New("bad credentials provided")
var ErrBadApiKey = errors.New("bad api key")
var ErrTooManyRequests = errors.New("too many requests")
var ErrUnsopportedData = errors.New("response from api contains unsupported data")
