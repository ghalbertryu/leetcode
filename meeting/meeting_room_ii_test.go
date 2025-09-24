package meeting

import (
	"log"
	"testing"
)

func TestMinMeetingRoom(t *testing.T) {
	meetings0 := []MeetingInterval{
		{1, 5},
		{7, 9},
	}
	log.Println(minMeetingRoomOptimize(meetings0)) // 1

	meetings1 := []MeetingInterval{
		{1, 5},
		{7, 9},
		{2, 3},
	}
	log.Println(minMeetingRoomOptimize(meetings1)) // 2

	meetings3 := []MeetingInterval{
		{1, 10},
		{2, 8},
		{3, 5},
	}
	log.Println(minMeetingRoomOptimize(meetings3)) // 3

	meetings4 := []MeetingInterval{
		{1, 2},
		{3, 4},
		{1, 3},
	}
	log.Println(minMeetingRoomOptimize(meetings4)) // 2
}
