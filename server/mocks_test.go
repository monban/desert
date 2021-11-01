package server

import "github.com/monban/desert"

type MockPublisher struct {
	calls struct {
		Publish struct {
			receives struct {
				subj string
				data []byte
			}
			returns error
		}
	}
}

func (mp *MockPublisher) Publish(subj string, data []byte) error {
	mp.calls.Publish.receives.subj = subj
	mp.calls.Publish.receives.data = data
	return mp.calls.Publish.returns
}

type MockDaWatcher struct {
	calls struct {
		Watch struct {
			receives struct {
				fn func(desert.DeckEvent)
			}
		}
	}
}

func (w *MockDaWatcher) Watch(fn func(desert.DeckEvent)) {
	w.calls.Watch.receives.fn = fn
}
