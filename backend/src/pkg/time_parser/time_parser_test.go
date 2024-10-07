package time_parser

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestStringToDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name       string
		args       args
		wantParsed time.Time
		wantErr    bool
	}{
		{
			name: "test_pos_01",
			args: args{
				date: "2020-09-01 12:45:55",
			},
			wantErr:    false,
			wantParsed: time.Date(2020, 9, 1, 12, 45, 55, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParsed, err := StringToDate(tt.args.date)

			fmt.Println(gotParsed)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParsed, tt.wantParsed) {
				t.Errorf("StringToDate() gotParsed = %v, want %v", gotParsed, tt.wantParsed)
			}
		})
	}
}
