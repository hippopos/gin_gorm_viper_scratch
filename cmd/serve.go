package cmd

import (
	"fmt"
	"os"

	"scratch/src/server"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	port int

	postgresHost     string
	postgresPort     int
	postgresUser     string
	postgresPass     string
	postgresDatabase string
	postgresSSL      string
	//http
)

func newServeCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "serve",
		Short: "Provide API service",
		// Long:  `All software has versions. This is Scheduler's`,
		Run: serveRun,
		// TraverseChildren: true,
	}
	cobra.OnInitialize(initConfig)
	//http
	command.PersistentFlags().StringVar(&cfgFile, "config", "./config.yaml", "config file (default is ./config.yaml)")
	command.PersistentFlags().IntVarP(&port, "port", "p", 8099, "restful serve port")
	command.PersistentFlags().StringVar(&postgresHost, "postgres.host", "localhost", "postgres server ip")
	command.PersistentFlags().IntVar(&postgresPort, "postgres.port", 5432, "postgres server port")
	command.PersistentFlags().StringVar(&postgresUser, "postgres.user", "postgres", "postgres server username")
	command.PersistentFlags().StringVar(&postgresPass, "postgres.password", "postgres", "postgres server password")
	command.PersistentFlags().StringVar(&postgresDatabase, "postgres.database", "postgres", "postgres database")
	command.PersistentFlags().StringVar(&postgresSSL, "postgres.ssl", "disable", "postgres disable ssl")

	viper.BindPFlags(command.PersistentFlags())
	return command
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	jww.SetStdoutThreshold(jww.LevelDebug)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err, ". use default setting")
		// os.Exit(1)
	}
}

func serveRun(cmd *cobra.Command, args []string) {
	// cf := viper.GetString("config")

	s := server.NewServer()
	defer s.Close()

	s.Run()

}
