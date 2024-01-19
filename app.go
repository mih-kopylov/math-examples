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

	stat := NewStat()
	answersDistribution := NewDistribution[int]()
	generator, err := NewOperandGenerator(profile)
	if err != nil {
		return err
	}

	answerReader := NewConsoleAnswerReader()

	stat.PrintStartMessage(printer)
	for i := 0; i < profile.ExamplesCount; i++ {
		example, err := generator.GenerateExample(profile, answersDistribution)
		if err != nil {
			return err
		}

		answer, err := answerReader.Read(printer, i+1, example)
		if err != nil {
			return err
		}

		stat.AddAnswer(answer)

		if profile.ShowCorrectAnswerAfter == AfterEach {
			printer.Println(answer.GetAnalysis())
		}
	}

	printer.Println("===============")
	if profile.ShowCorrectAnswerAfter == AfterAll {
		stat.PrintAllAnswers(printer)
	}
	stat.PrintStatistics(printer)
	printer.Println("===============\n")

	return nil
}
