package cmd

const (
	PREPROCESS_SHORT_DESCRIPTION = "preprocess is a fast data analysis preprocessing tool."
	PREPROCESS_LONG_DESCRIPTION  = `preprocess is a fast data analysis preprocessing tool built with Go.
	Complete documentation is available at https://github.com/agailloty/preprocess`
	INIT_SHORT_DESCRIPTION = "Generate a Prepfile"
	INIT_LONG_DESCRIPTION  = "This command is used to generate a Prepfile in which you can specify the preprocessing computations either on specifics columns of the dataset or on whole numeric or text columns."
	RUN_SHORT_DESCRIPTION  = "Run operations using Prepfile"
	RUN_LONG_DESCRIPTION   = `This command is used to run the preprocessing operations.
	You can optionally specify operations via the command line, but preferred way is to use Prepfile.`
	SUMMARY_SHORT_DESCRIPTION = "Generate dataset summary statistics."
	SUMMARY_LONG_DESCRIPTION  = `Generate dataset summary statistics.
	By default it generates a TOML file that contains all summary statistics. 
	You can choose to generate a HTML file with the --html flag.`
)
