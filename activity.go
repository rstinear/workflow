package sample

import (
	"context"
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/rstinear/workflow/client"
	conk "github.com/rstinear/workflow/proto"
	md "google.golang.org/grpc/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	fmt.Println("settings: ", settings)

	act := &Activity{settings: settings} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings *Settings
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)

	fmt.Printf("my input value: %v\n", ctx.GetInput("projectID"))
	fmt.Printf("gidday, my input object is: %v and my settings object is: %v\n", input, a.settings)

	cc := client.New(a.settings.Token, a.settings.URI, true)

	gctx := md.AppendToOutgoingContext(context.Background(), "r-cs", input.Callsign, "r-p", input.ProjectID, "grpc-timeout", "0")

	msg := &conk.ControlMessage{
		Destination: "/http",
		Payload: []byte(
			`{
				"address": "https://c49c66b2a1762a4acad8d883ee445089.m.pipedream.net",
				"method": "POST",
				"timeout": 5000
			}`)}

	fmt.Printf("about to call control robot with msg yo %v\n", msg)
	stream, err := cc.ControlRobotCallWithAckStream(gctx, msg)
	if err != nil {
		fmt.Printf("failed to send control: %v\n", err)
		return false, err
	}

	var result string
	for {
		ack, err := stream.Recv()
		if err != nil {
			fmt.Printf("received err on ack stream: %v\n", err)
			return false, err
		}

		fmt.Printf("ack: %v\n", ack)
		result = ack.Stage

		if ack.Stage == "final" {
			break
		}
	}

	if err != nil {
		return true, err
	}

	output := &Output{AnOutput: result}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
