package event_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niallyoung/goNDK/event"
)

func TestNewEvent(t *testing.T) {
	t.Run("NewEvent() returns an Event", func(t *testing.T) {
		e := ValidEvent()
		assert.NotNil(t, e)
	})
}

func TestEvent_MarshalJSON(t *testing.T) {
	t.Run("unmarshal ValidEventJSON to Event{}", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEventJSON), &e)
		assert.NoError(t, err)
	})

	t.Run("unmarshal ValidEvent2JSON to Event{}", func(t *testing.T) {
		var e event.Event
		err := json.Unmarshal([]byte(ValidEvent2JSON), &e)
		assert.NoError(t, err)
	})
}

func TestEvent_Validate(t *testing.T) {
	t.Run("valid Event", func(t *testing.T) {
		e := ValidEvent()
		err := e.Validate()
		assert.NoError(t, err)
	})

	t.Run("invalid Event", func(t *testing.T) {
		e := InValidEvent()
		err := e.Validate()
		assert.Error(t, err)
	})
}
