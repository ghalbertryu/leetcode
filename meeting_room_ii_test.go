package leetcode

import (
	"log"
	"testing"
)

func TestMinMeetingRoom(t *testing.T) {
	meetings0 := []MeetingInterval{
		{1, 5},
		{7, 9},
	}
	log.Println(minMeetingRoomOptimize(meetings0))

	meetings1 := []MeetingInterval{
		{1, 5},
		{7, 9},
		{2, 3},
	}
	log.Println(minMeetingRoomOptimize(meetings1))

	meetings3 := []MeetingInterval{
		{1, 10},
		{2, 8},
		{3, 5},
	}
	log.Println(minMeetingRoomOptimize(meetings3))
}
