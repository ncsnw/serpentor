package serpentor

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Cmd struct {
	cmd *cobra.Command
}

func NewCmd(use string) *Cmd {
	return &Cmd{&cobra.Command{Use: use}}
}

func FromCommand(cmd *cobra.Command) *Cmd {
	return &Cmd{cmd}
}

func (c *Cmd) Short(short string) *Cmd {
	c.cmd.Short = short
	return c
}

func (c *Cmd) Long(long string) *Cmd {
	c.cmd.Long = long
	return c
}

func (c *Cmd) Alias(aliases ...string) *Cmd {
	c.cmd.Aliases = aliases
	return c
}

func (c *Cmd) SuggestFor(suggestFor ...string) *Cmd {
	c.cmd.SuggestFor = suggestFor
	return c
}

func (c *Cmd) Example(example string) *Cmd {
	c.cmd.Example = example
	return c
}

func (c *Cmd) ValidArgs(validArgs ...string) *Cmd {
	c.cmd.ValidArgs = validArgs
	return c
}

func (c *Cmd) ValidArgsFunction(validArgsFunction func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)) *Cmd {
	c.cmd.ValidArgsFunction = validArgsFunction
	return c
}

func (c *Cmd) Args(args cobra.PositionalArgs) *Cmd {
	c.cmd.Args = args
	return c
}

func (c *Cmd) ArgAliases(argAliases ...string) *Cmd {
	c.cmd.ArgAliases = argAliases
	return c
}

func (c *Cmd) BashCompletionFunction(bashCompletionFunction string) *Cmd {
	c.cmd.BashCompletionFunction = bashCompletionFunction
	return c
}

func (c *Cmd) Deprecated(deprecated string) *Cmd {
	c.cmd.Deprecated = deprecated
	return c
}

func (c *Cmd) Annotations(annotations map[string]string) *Cmd {
	c.cmd.Annotations = annotations
	return c
}

func (c *Cmd) Version(version string) *Cmd {
	c.cmd.Version = version
	return c
}

func (c *Cmd) PersistentPreRun(persistentPreRun func(cmd *cobra.Command, args []string)) *Cmd {
	c.cmd.PersistentPreRun = persistentPreRun
	return c
}

func (c *Cmd) PersistentPreRunE(persistentPreRunE func(cmd *cobra.Command, args []string) error) *Cmd {
	c.cmd.PersistentPreRunE = persistentPreRunE
	return c
}

func (c *Cmd) PreRun(preRun func(cmd *cobra.Command, args []string)) *Cmd {
	c.cmd.PreRun = preRun
	return c
}

func (c *Cmd) PreRunE(preRunE func(cmd *cobra.Command, args []string) error) *Cmd {
	c.cmd.PreRunE = preRunE
	return c
}

func (c *Cmd) Run(run func(cmd *cobra.Command, args []string)) *Cmd {
	c.cmd.Run = run
	return c
}

func (c *Cmd) RunE(runE func(cmd *cobra.Command, args []string) error) *Cmd {
	c.cmd.RunE = runE
	return c
}

func (c *Cmd) PostRunE(postRunE func(cmd *cobra.Command, args []string) error) *Cmd {
	c.cmd.PostRunE = postRunE
	return c
}

func (c *Cmd) PersistentPostRun(persistentPostRun func(cmd *cobra.Command, args []string)) *Cmd {
	c.cmd.PersistentPostRun = persistentPostRun
	return c
}

func (c *Cmd) PersistentPostRunE(persistentPostRunE func(cmd *cobra.Command, args []string) error) *Cmd {
	c.cmd.PersistentPostRunE = persistentPostRunE
	return c
}

func (c *Cmd) FParseErrWhitelist(fParseErrWhitelist cobra.FParseErrWhitelist) *Cmd {
	c.cmd.FParseErrWhitelist = fParseErrWhitelist
	return c
}

func (c *Cmd) TraverseChildren(traverseChildren bool) *Cmd {
	c.cmd.TraverseChildren = traverseChildren
	return c
}

func (c *Cmd) SilenceErrors(silenceErrors bool) *Cmd {
	c.cmd.SilenceErrors = silenceErrors
	return c
}

func (c *Cmd) SilenceUsage(silenceUsage bool) *Cmd {
	c.cmd.SilenceUsage = silenceUsage
	return c
}

func (c *Cmd) DisableFlagParsing(disableFlagParsing bool) *Cmd {
	c.cmd.DisableFlagParsing = disableFlagParsing
	return c
}

func (c *Cmd) DisableAutoGenTag(disableAutoGenTag bool) *Cmd {
	c.cmd.DisableAutoGenTag = disableAutoGenTag
	return c
}

func (c *Cmd) DisableFlagsInUseLine(disableFlagsInUseLine bool) *Cmd {
	c.cmd.DisableFlagsInUseLine = disableFlagsInUseLine
	return c
}

func (c *Cmd) DisableSuggestions(disableSuggestions bool) *Cmd {
	c.cmd.DisableSuggestions = disableSuggestions
	return c
}

func (c *Cmd) SuggestionsMinimumDistance(suggestionsMinimumDistance int) *Cmd {
	c.cmd.SuggestionsMinimumDistance = suggestionsMinimumDistance
	return c
}

func (c *Cmd) Flags() *pflag.FlagSet {
	return c.cmd.Flags()
}

func (c *Cmd) PersistentFlags() *pflag.FlagSet {
	return c.cmd.PersistentFlags()
}

func (c *Cmd) AddCommand(cobraCmds ...*Cmd) *Cmd {
	cmds := make([]*cobra.Command, len(cobraCmds))
	for i := range cobraCmds {
		cmds[i] = cobraCmds[i].cmd
	}
	c.cmd.AddCommand(cmds...)
	return c
}

func (c *Cmd) Execute() {
	cobra.CheckErr(c.cmd.Execute())
}

func (c *Cmd) ExecuteE() error {
	return c.cmd.Execute()
}

func (c *Cmd) ExecuteContext(ctx context.Context) {
	cobra.CheckErr(c.cmd.ExecuteContext(ctx))
}

func (c *Cmd) ExecuteContextE(ctx context.Context) error {
	return c.cmd.ExecuteContext(ctx)
}

func (c *Cmd) GetCobra() *cobra.Command {
	return c.cmd
}
