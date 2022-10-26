package poll

type Poll struct {
	uuid     string
	question string
	answers  []Answer
}

type Answer struct {
	answer         string
	numberOfAnswer int
}
