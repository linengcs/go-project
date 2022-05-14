package service

import (
	"errors"
	"github.com/linengcs/go-project/repository"
	"sync"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}
func NewQueryPageInfoFlow(topID int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topID,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

// 参数校验
func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

// 准备数据
func (f *QueryPageInfoFlow) prepareInfo() error {
	// get the information of topic
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

// 组装数据
func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}
