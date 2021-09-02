package sns

type Message struct {
	NewStateValue    string `json:"NewStateValue"`
	NewStateReason   string `json:"NewStateReason"`
	AlarmName        string `json:"AlarmName"`
	AlarmDescription string `json:"AlarmDescription"`
}
