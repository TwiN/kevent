package kevent

import (
	"testing"

	fakekubernetes "k8s.io/client-go/kubernetes/fake"
)

// TestNewEventManager makes sure that the EventManager is created correctly.
func TestNewEventManager(t *testing.T) {
	kubernetesClient := fakekubernetes.NewSimpleClientset()
	em := NewEventManager(kubernetesClient, "component")
	if em == nil {
		t.Error("EventManager is nil")
	}
}

// TestEventManager_Create makes sure that the Create method works correctly.
// Until I figure out how to test this, it doesn't really do much besides for making sure that nothing panics
func TestEventManager_Create(t *testing.T) {
	kubernetesClient := fakekubernetes.NewSimpleClientset()
	em := NewEventManager(kubernetesClient, "component")
	em.Create("default", "pod", "name", "Reason", "Message", true)
}
