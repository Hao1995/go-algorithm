package designtwitter

import (
	"container/heap"
	"time"
)

type post struct {
	Id        int
	CreatedAt time.Time
}

type recentPQ []*post

func (pq recentPQ) Len() int           { return len(pq) }
func (pq recentPQ) Less(i, j int) bool { return pq[i].CreatedAt.After(pq[j].CreatedAt) }
func (pq recentPQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *recentPQ) Push(item any)     { (*pq) = append((*pq), item.(*post)) }
func (pq *recentPQ) Pop() any {
	item := (*pq)[pq.Len()-1]
	(*pq) = (*pq)[:pq.Len()-1]
	return item
}

type Twitter struct {
	// {user_id: [tweet_id]}. from most recent to least recent
	posts map[int][]*post

	// {curr_user_id: [followed_user_id]}.
	following map[int]map[int]bool
}

// n: 500
// m: 10^4
func Constructor() Twitter {
	return Twitter{
		posts:     make(map[int][]*post),
		following: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	post := &post{tweetId, time.Now()}
	this.posts[userId] = append(this.posts[userId], post)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	posts, _ := this.posts[userId]
	for following, _ := range this.following[userId] {
		posts = append(posts, this.posts[following]...)
	}

	result := make([]int, 0, 10)
	pq := recentPQ(posts)
	heap.Init(&pq)
	for pq.Len() > 0 {
		post := heap.Pop(&pq).(*post)
		result = append(result, post.Id)
		if len(result) == 10 {
			return result
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.following[followerId]; !ok {
		this.following[followerId] = make(map[int]bool)
	}
	this.following[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.following[followerId], followeeId)
}
