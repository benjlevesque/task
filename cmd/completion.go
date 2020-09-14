package cmd

import (
	"os"

	"github.com/benjlevesque/task/pkg/util"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use: "completion [bash|zsh|fish|powershell]",
	Long: `To load completions:

Bash:

$ source <(task completion bash)

# To load completions for each session, execute once:
Linux:
  $ task completion bash > /etc/bash_completion.d/task
MacOS:
  $ task completion bash > /usr/local/etc/bash_completion.d/task

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ task completion zsh > "${fpath[1]}/_task"

# You will need to start a new shell for this setup to take effect.

Fish:

$ task completion fish | source

# To load completions for each session, execute once:
$ task completion fish > ~/.config/fish/completions/task.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	ValidArgsFunction:     util.NoFileCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
