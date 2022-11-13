package kevent

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func init() {
	_ = corev1.AddToScheme(scheme)
}

// CreateEvent creates a Kubernetes Event with the given parameters using the given kubernetes Client.
func CreateEvent(kubernetesClient kubernetes.Interface, namespace, podName, reason, message string, isWarning bool) {
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
	recorder.Event(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
	}, eventType, reason, message)
}
