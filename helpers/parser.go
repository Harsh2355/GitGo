package helpers

type Command interface {
	run(args []string)
}

// ------------------ Init a repo --------------------

type Init struct {
	// can add flags and stuff later or
}

func (init *Init) run(args []string) {
	// for now will only process gitgo init without any
	// args (including flags)

}

// ---------------------------------------------------

func Parser(command string, args []string) {

	switch command {
	case "init":
		initObj := &Init{}
		initObj.run(args)
	}
}
