package meeting

import (
	"slices"
)

type MeetingInterval struct {
	startTime int
	endTime   int
}

func compareTo(a MeetingInterval, b MeetingInterval) int {
	return a.startTime - b.startTime
}

// 每個meeting有一個開始時間(start)與結束時間(end)
// 問這些meeting時間有沒有衝突，衝突的意思為同一個時間點需要同時出現在兩場meeting
// 若某一場meeting的結束時間與下一場meeting的開始時間相同不視為衝突
func meetingNotConflict(meetings []MeetingInterval) bool {
	slices.SortFunc(meetings, compareTo) // O(nlogn)

	endTime := -1
	for _, meeting := range meetings {
		if meeting.startTime < endTime {
			return false
		}
		endTime = meeting.endTime
	}
	return true
}
