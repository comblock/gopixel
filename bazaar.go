package gopixel

import (
	"encoding/json"

	"errors"

	structs "github.com/comblock/gopixel/structs"
)

func (client *Client) BazaarData() (structs.Bazaar, error) {
	var bazaar structs.Bazaar

	data, err := get("api.hypixel.net/skyblock/bazaar?key=" + client.Key)
	if err != nil {
		return bazaar, err
	}

	err = json.Unmarshal(data, &bazaar)

	if !bazaar.Success {
		err = errors.New(bazaar.Cause)
	}

	return bazaar, err
}
