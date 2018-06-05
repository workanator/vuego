package server

type WsCloseCode = int

// WebSocket close codes
const (
	WsNormalClosure           = 1000
	WsGoingAway               = 1001
	WsProtocolError           = 1002
	WsUnsupportedData         = 1003
	WsInvalidFramePayloadData = 1007
	WsPolicyViolation         = 1008
	WsMessageTooBig           = 1009
	WsMissingExtension        = 1010
	WsInternalError           = 1011
	WsServiceRestart          = 1012
	WsTryAgainLater           = 1013
	WsBadGateway              = 1014
)
