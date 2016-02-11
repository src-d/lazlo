package modules

import (
	"strings"

	lazlo "github.com/src-d/lazlo/lib"
)

var PrettyHelp = &lazlo.Module{
	Name:  `Help`,
	Usage: `"%BOTNAME% help": prints the usage information of every registered plugin`,
	Run: func(b *lazlo.Broker) {
		cb := b.MessageCallback(`(?i)help`, true)
		for {
			pm := <-cb.Chan
			go getPrettyHelp(b, &pm)
		}
	},
}

func getPrettyHelp(b *lazlo.Broker, pm *lazlo.PatternMatch) {
	a := []lazlo.Attachment{
		lazlo.Attachment{
			Color:  "#ff0000",
			Title:  "Modules In use",
			Fields: []lazlo.AttachmentField{},
		},
	}
	for _, m := range b.Modules {
		if strings.Contains(m.Usage, `%HIDDEN%`) {
			continue
		}
		usage := strings.Replace(m.Usage, `%BOTNAME%`, b.Config.Name, -1)

		a[0].Fields = append(a[0].Fields, lazlo.AttachmentField{
			Title: m.Name,
			Value: usage,
		})
	}
	pm.Event.RespondAttachments(a)
}
