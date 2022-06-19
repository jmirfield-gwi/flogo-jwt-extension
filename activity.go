package jwt

import (
	"github.com/project-flogo/core/activity"
	"fmt"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	return act, nil
}

type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	fmt.Println(input)

	sharedEncryptionKey := []byte(input.Secret)

	fmt.Println(sharedEncryptionKey)

	if err != nil {
		return true, err
	}

	switch input.Mode {
	
	case "Sign":
		{
			fmt.Println(input.Mode)
		}
	}

	return true, nil
}