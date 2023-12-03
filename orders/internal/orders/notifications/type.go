package notifications

type NotifPayload struct {
	Title string
	Body  string
	CTA   map[string]any
}
