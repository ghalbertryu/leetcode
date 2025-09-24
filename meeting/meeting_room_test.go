package meeting

import (
	"log"
	"testing"
)

func TestMeetingNotConflict(t *testing.T) {
	meetings0 := []MeetingInterval{
		{1, 5},
		{7, 9},
	}
	log.Println(meetingNotConflict(meetings0))

	meetings1 := []MeetingInterval{
		{1, 5},
		{7, 9},
		{2, 3},
	}
	log.Println(meetingNotConflict(meetings1))

	meetings2 := []MeetingInterval{
		{1, 5},
		{5, 8},
		{7, 9},
	}
	log.Println(meetingNotConflict(meetings2))
}
