package jwt

import (
	"github.com/project-flogo/core/activity"
	"github.com/golang-jwt/jwt/v4"
	"encoding/json"
	"fmt"
	"time"
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
	output := &Output{}
	err = ctx.GetInputObject(input)

	if err != nil {
		return false, err
	}

	sharedEncryptionKey := []byte(input.Secret)

	if err != nil {
		return true, err
	}

	switch input.Mode {
	
	case "Sign":
		{
			claims := jwt.MapClaims{"exp":time.Now().Unix() + 60}
			var header map[string]interface{}

			// take the payload (claims) string and unmarshall it into a byte slice
			if err := json.Unmarshal([]byte(input.Payload), &claims); err != nil {
				ctx.Logger().Info("Invalid Payload: ", err)
				return false, err
			}

			fmt.Println(claims)

			// Take the header string and unmarshall
			if err := json.Unmarshal([]byte(input.Header), &header); err != nil {
				ctx.Logger().Info("Invalid Header: ", err)
				return false, err
			}


			// get the alg value from the header
			alg := header["alg"].(string)

			// if the header and the passed algo method the same
			if input.Algorithm != alg {
				ctx.Logger().Info("Header algo doesn't match algorithm parm")
				return false, nil
			}

			// use the alg name to get the signing method
			signwith := jwt.GetSigningMethod(alg)

			token := jwt.NewWithClaims(signwith, claims)

			var key interface{}

			//  Depending on the algorithm type we need to convert  the format of the private string
			key, err = jwt.ParseRSAPrivateKeyFromPEM(sharedEncryptionKey)
				if err != nil {
					ctx.Logger().Info("Bad RSA key", err)
					return false, err
				}

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString(key)
			fmt.Println(tokenString)
			if err == nil {
				ctx.Logger().Debug("Token String created", tokenString)
				output.Token = tokenString
				ctx.SetOutputObject(output)
				return true, nil
			} else {
				ctx.Logger().Info("Signing error: ", err)
				return false, err
			}
		}
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}