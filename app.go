package main

func runApplication() error {
	printer, err := NewLogPrinter()
	if err != nil {
		return err
	}

	profile, err := ReadProfile(printer)
	if err != nil {
		return err
	}

	st := NewStat(profile.ExamplesCount)
	generator := NewGenerator(profile)
	answerReader := NewConsoleAnswerReader()

	st.PrintStartMessage(printer)
	for i := 0; i < profile.ExamplesCount; i++ {
		example, err := generator.GenerateExample(profile, st, i+1)
		if err != nil {
			return err
		}

		answer, err := answerReader.Read(printer, i+1, example)
		if err != nil {
			return err
		}

		st.AddAnswer(example, answer)

		if profile.ShowCorrectAnswerAfter == AfterEach {
			printer.Println(answer.GetAnalysis())
		}
	}

	printer.Println("===============")
	if profile.ShowCorrectAnswerAfter == AfterAll {
		st.PrintAllAnswers(printer)
	}
	st.PrintStatistics(printer)
	printer.Println("===============\n")

	return nil
}
