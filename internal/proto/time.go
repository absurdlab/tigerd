package proto

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func protoDurationOfSeconds(sec uint64) *durationpb.Duration {
	if sec <= 0 {
		return nil
	}
	return durationpb.New(time.Duration(sec) * time.Second)
}

func protoTimestampOfEpochSeconds(ts uint64) *timestamppb.Timestamp {
	return timestamppb.New(time.Unix(int64(ts), 0))
}
