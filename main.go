package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, n.importance()
	case groupMessage:
		return v.groupName, n.importance()
	case systemAlert:
		return v.alertCode, n.importance()
	}
	return "", 0
}
