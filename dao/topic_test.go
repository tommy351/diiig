package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tommy351/diiig/model"
)

func TestMemoryTopicDAO_Create(t *testing.T) {
	set := &model.Set{
		Collection: new(model.LinkedList),
	}
	dao := &MemoryTopicDAO{
		set: set,
	}

	require.Nil(t, dao.Create("test"))
	assert.Equal(t, &model.Element{Key: "test", Value: 0}, set.Get("test"))
}

func TestMemoryTopicDAO_Vote(t *testing.T) {
	set := &model.Set{
		Collection: new(model.LinkedList),
	}
	dao := &MemoryTopicDAO{
		set: set,
	}

	require.Nil(t, dao.Create("test"))
	require.Nil(t, dao.Vote("test", 1))
	assert.Equal(t, &model.Element{Key: "test", Value: 1}, set.Get("test"))
}

func TestMemoryTopicDAO_Range(t *testing.T) {
	dao := NewMemoryTopicDAO()

	require.Nil(t, dao.Create("foo"))
	require.Nil(t, dao.Create("bar"))
	require.Nil(t, dao.Vote("bar", 1))

	topics, err := dao.Range(0, -1)
	require.Nil(t, err)
	assert.Equal(t, []model.Topic{
		{Name: "bar", Score: 1},
		{Name: "foo", Score: 0},
	}, topics)
}
