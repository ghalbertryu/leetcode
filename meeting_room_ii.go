package leetcode

import (
	"slices"
)

type MeetingEvent struct {
	eventTime int
	isStart   bool
}

// 給定一些meeting的時間，每個meeting有一個開始時間(start)與結束時間(end)
// 問要完成這些meeting最少需要幾間會議室
// 若同一個時間點有兩場meeting在進行，則這兩場meeting需要用不同間會議室
// 若某一場meeting的結束時間與下一場meeting的開始時間相同，則他們可以使用同一間會議室
func minMeetingRoomOptimize(meetings []MeetingInterval) int {
	meetingEvents := make([]MeetingEvent, 0)
	for _, meeting := range meetings {
		meetingEvents = append(meetingEvents, MeetingEvent{
			eventTime: meeting.startTime,
			isStart:   true,
		})
		meetingEvents = append(meetingEvents, MeetingEvent{
			eventTime: meeting.endTime,
			isStart:   false,
		})
	}

	slices.SortFunc(meetingEvents, func(a, b MeetingEvent) int {
		return a.eventTime - b.eventTime
	})

	ans := 0
	roomCount := 0
	for _, event := range meetingEvents {
		if event.isStart {
			roomCount++
		} else {
			roomCount--
		}
		if roomCount > ans {
			ans = roomCount
		}
	}
	return ans
}

type Room struct {
	start int
	end   int
}

func minMeetingRoomFirst(meetings []MeetingInterval) int {
	slices.SortFunc(meetings, compareTo)
	rooms := make([]Room, 0)

	for _, meeting := range meetings {
		if !findAvailableMeetingRoom(rooms, meeting) {
			rooms = append(rooms, Room{
				start: meeting.startTime,
				end:   meeting.endTime,
			})
		}
	}
	return len(rooms)
}

func findAvailableMeetingRoom(rooms []Room, meeting MeetingInterval) bool {
	for i, _ := range rooms {
		room := &rooms[i]
		if meeting.endTime <= room.start ||
			meeting.startTime >= room.end {
			// use this room
			room.start = meeting.startTime
			room.end = meeting.endTime
			return true
		}
	}
	return false
}
