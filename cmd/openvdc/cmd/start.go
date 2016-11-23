package cmd

import (
	"context"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	// TODO: Remove --server option from sub-command.
	startCmd.PersistentFlags().StringVarP(&serverAddr, "server", "s", "localhost:5000", "gRPC API server address")
	startCmd.PersistentFlags().SetAnnotation("server", cobra.BashCompSubdirsInDir, []string{})
}

var startCmd = &cobra.Command{
	Use:   "start [Instance ID]",
	Short: "Start an instance",
	Long:  `Start an instance in REGISTERED or STOPPED state to RUNNING.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			log.Fatalf("Please provide an Instance ID.")
		}

		instanceID := args[0]

		req := &api.StartRequest{
			InstanceId: instanceID,
		}
		return remoteCall(func(conn *grpc.ClientConn) error {
			c := api.NewInstanceClient(conn)
			res, err := c.Start(context.Background(), req)
			if err != nil {
				log.WithError(err).Fatal("Disconnected abnormaly")
				return err
			}
			fmt.Println(res)
			return err
		})
	}}
