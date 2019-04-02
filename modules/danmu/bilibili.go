package danmu

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/themoonbear/api-server/utils"
)

type bilibili struct {
}

var (
	roomReg = regexp.MustCompile(`"room_id":(\d+),`)
	//uidReg = regexp.MustCompile(`"uid":(\d+),`)
	addressReg        = regexp.MustCompile(`https://live\.bilibili\.com/(\d+)`)
	serverReg         = regexp.MustCompile("<dm_server>(.*?)</dm_server>")
	stateReg          = regexp.MustCompile("<state>(.*?)</state>")
	bilibiliServerAPI = "http://live.bilibili.com/api/player?id=cid:"
	bilibiliRoomAPI   = "https://api.live.bilibili.com/room/v1/Room/room_init?id="
)

func (bili bilibili) MatchAddress(address string) bool {
	return addressReg.MatchString(strings.TrimSpace(address))
}

func (bili bilibili) ParseAddress(address string) (interface{}, error) {
	roomID, err := bili.getRoomIDAndUID(address)
	if err != nil {
		return nil, err
	}
	server, state, err := bili.getServerAndLiveState(roomID)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"platform": "bilibili",
		"server":   "wss://" + server + "/sub",
		"state":    state == "LIVE",
		"roomID":   utils.String2Int(roomID),
		"uid":      rand.Intn(4e7-1e5) + 1e5,
	}, nil
}

func (bili bilibili) getRoomIDAndUID(address string) (string, error) {
	matchs := addressReg.FindStringSubmatch(address)
	if len(matchs) < 2 {
		return "", fmt.Errorf("ADDRESS submatch %q", matchs)
	}
	if len(matchs[1]) > 3 {
		return matchs[1], nil
	}
	//room data
	body, err := utils.HTTPGet(bilibiliRoomAPI+matchs[1], nil)
	if err != nil {
		return "", err
	}
	//roomID
	matchs = roomReg.FindStringSubmatch(body)
	if len(matchs) < 2 {
		return "", fmt.Errorf("ROOMID submatch %q", matchs)
	}
	return matchs[1], nil
}

func (bili bilibili) getServerAndLiveState(room string) (string, string, error) {
	uri := bilibiliServerAPI + room
	body, err := utils.HTTPGet(uri, nil)
	if err != nil {
		return "", "", err
	}
	matchs := serverReg.FindStringSubmatch(body)
	if len(matchs) < 2 {
		return "", "", fmt.Errorf("server submatch %q", matchs)
	}
	server := matchs[1]

	matchs = stateReg.FindStringSubmatch(body)
	if len(matchs) < 2 {
		return "", "", fmt.Errorf("state submatch %q", matchs)
	}
	state := matchs[1]
	return server, state, nil
}

func init() {
	registerParser(&bilibili{})
}
