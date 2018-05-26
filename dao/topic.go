package dao

import "github.com/tommy351/diiig/model"

type TopicDAO interface {
	Create(name string) error
	Vote(name string, vote int) error
	Range(start, end int) ([]model.Topic, error)
}

type MemoryTopicDAO struct {
	set *model.Set
}

func NewMemoryTopicDAO() TopicDAO {
	return &MemoryTopicDAO{
		set: &model.Set{
			Collection: new(model.LinkedList),
		},
	}
}

func (m *MemoryTopicDAO) Create(name string) error {
	m.set.Add(name)
	return nil
}

func (m *MemoryTopicDAO) Vote(name string, vote int) error {
	m.set.Increment(name, vote)
	return nil
}

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
