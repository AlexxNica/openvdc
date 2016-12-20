package cmd

import (
	"fmt"
	"strconv"
	"strings"

	mlog "github.com/ContainX/go-mesoslog/mesoslog"
	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/cmd/openvdc/internal/util"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log [Instance ID]",
	Short: "Print logs of an instance",
	Long:  "Print logs of an instance",
	Example: `
	% openvdc log i-xxxxxxx
	`,
	DisableFlagParsing: true,
	PreRunE:            util.PreRunHelpFlagCheckAndQuit,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			log.Fatalf("Please provide an Instance ID.")
		}

		instanceID := "VDC_" + args[0]

		split := strings.Split(util.MesosMasterAddr, ":")
		mesosMasterAddr := split[0]
		mesosMasterPort, err := strconv.Atoi(split[1])

		if err != nil {
			log.Errorln("Couldn't convert string to int") //  <--- lol
		}

		cl, err := mlog.NewMesosClientWithOptions(mesosMasterAddr, mesosMasterPort, &mlog.MesosClientOptions{SearchCompletedTasks: false, ShowLatestOnly: true})
		if err != nil {
			log.Infoln(err)
		}

		result, err := cl.GetLog(instanceID, mlog.STDERR, "")
		if err != nil {
			log.Errorln("Error getting log")
		}

		for _, log := range result {
			fmt.Printf(log.Log)
		}

		return err
	},
}
