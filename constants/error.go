package constants

type MessageResult struct {
	Message string `bson:"message"`
	Status  bool   `bson:"status"`
}
