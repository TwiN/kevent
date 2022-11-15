package kevent

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	typedv1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
)

var (
	scheme = runtime.NewScheme()

	broadcaster record.EventBroadcaster = nil
	recorder    record.EventRecorder    = nil
)

// CreateEvent creates a Kubernetes Event with the given parameters using the given kubernetes Client.
func CreateEvent(kubernetesClient kubernetes.Interface, resourceNamespace, resourceKind, resourceName, reason, message string, isWarning bool) {
	if broadcaster == nil {
		broadcaster = record.NewBroadcaster()
		broadcaster.StartStructuredLogging(4)
		broadcaster.StartRecordingToSink(&typedv1core.EventSinkImpl{Interface: kubernetesClient.CoreV1().Events("")})
		recorder = broadcaster.NewRecorder(scheme, corev1.EventSource{})
	}
	var eventType string
	if isWarning {
		eventType = corev1.EventTypeWarning
	} else {
		eventType = corev1.EventTypeNormal
	}
	us := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": resourceKind,
			"metadata": map[string]interface{}{
				"name":      resourceName,
				"namespace": resourceNamespace,
			},
		},
	}
	recorder.Event(us, eventType, reason, message)
}
