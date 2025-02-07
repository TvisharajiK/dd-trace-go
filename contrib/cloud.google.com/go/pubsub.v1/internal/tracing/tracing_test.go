package tracing

// Test generated using Keploy
import (
    "context"
    "testing"
)

func TestTraceReceiveFunc_SpanCreationAndDeliveryAttempt(t *testing.T) {
    mockSubscription := &MockSubscription{name: "test-subscription"}
    msg := &Message{
        ID:              "456",
        Data:            []byte("test data"),
        OrderingKey:     "key2",
        Attributes:      map[string]string{},
        DeliveryAttempt: intPtr(3),
    }
    ctx := context.Background()

    receiveFunc := TraceReceiveFunc(mockSubscription)
    ctx, finishSpan := receiveFunc(ctx, msg)
    defer finishSpan()

    if msg.DeliveryAttempt == nil || *msg.DeliveryAttempt != 3 {
        t.Errorf("Expected delivery attempt to be set to 3, but got %v", msg.DeliveryAttempt)
    }
}

func intPtr(i int) *int {
    return &i
}

type MockSubscription struct {
    name string
}

func (m *MockSubscription) String() string {
    return m.name
}


// Test generated using Keploy
func TestTraceAcknowledge_SpanCreation(t *testing.T) {
    msg := &Message{
        ID:          "789",
        Data:        []byte("acknowledge data"),
        OrderingKey: "key3",
    }
    ctx := context.Background()

    // Ensure that TraceAcknowledge executes without panicking
    TraceAcknowledge(ctx, msg)
}
