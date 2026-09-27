package dto

type IsOnlineInput struct {
	UserIDs []uint
}

type IsOnlineOutput struct {
	Presences []Presence `json:"presences"`
}

type Presence struct {
	UserID    uint   `json:"user_id"`
	Timestamp uint64 `json:"timestamp"`
}
