package process

type Master interface {
	NodeFinished(msg message)
}
