package vo

import (
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type TopicJoinMap struct {
	topicmp map[string]*pubsub.Topic

	sync.RWMutex
}

func (t *TopicJoinMap) Load(key string) (*pubsub.Topic, bool) {
	t.RLock()
	defer t.RUnlock()

	tp, bl := t.topicmp[key]

	return tp, bl
}

func (t *TopicJoinMap) Store(key string, value *pubsub.Topic) {
	t.Lock()
	defer t.Unlock()

	t.topicmp[key] = value
}

func (t *TopicJoinMap) Delete(key string) {
	t.Lock()
	defer t.Unlock()

	delete(t.topicmp, key)
}
