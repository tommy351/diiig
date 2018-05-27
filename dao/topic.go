package dao

import "github.com/tommy351/diiig/model"

// TopicDAO is an interface to access topics.
type TopicDAO interface {
	Create(name string) error
	Vote(name string, vote int) error
	Range(start, end int) ([]model.Topic, error)
}

// MemoryTopicDAO implements TopicDAO interface and stores data in a set.
type MemoryTopicDAO struct {
	set *model.Set
}

// NewMemoryTopicDAO returns a new instance of MemoryTopicDAO.
func NewMemoryTopicDAO() TopicDAO {
	return &MemoryTopicDAO{
		set: &model.Set{
			Collection: new(model.LinkedList),
		},
	}
}

// Create adds the topic to the set.
func (m *MemoryTopicDAO) Create(name string) error {
	m.set.Add(name)
	return nil
}

// Vote increments the score of a topic in the set.
func (m *MemoryTopicDAO) Vote(name string, vote int) error {
	m.set.Increment(name, vote)
	return nil
}

// Range returns a slice of topics within the specified range.
func (m *MemoryTopicDAO) Range(start, end int) (topics []model.Topic, err error) {
	elements := m.set.Range(start, end)

	for _, e := range elements {
		topics = append(topics, model.Topic{
			Name:  e.Key,
			Score: e.Value,
		})
	}

	return
}
