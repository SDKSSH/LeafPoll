package poll

type Poll struct {
	uuid        string
	question    string
	answers     []Answer
	totalAnswer int
}

type Answer struct {
	answer         string
	numberOfAnswer int
}
