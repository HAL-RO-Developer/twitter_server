package model

var tweetsRep = Tweets{}

func (t *Tweet) Save() {
	t.ID = len(tweetsRep) + 1
	tweetsRep = append(tweetsRep, *t)
}

func GetTweets() Tweets {
	return tweetsRep
}
