package inputer

import (
	"reflect"
	"testing"
)

func Test_inputer(t *testing.T) {
	type args struct {
		l         int
		inputType string
	}
	tests := []struct {
		name       string
		args       args
		wantChoice any
		wantErr    bool
	}{
		{
			name: "test_pos_01",
			args: args{
				l:         0,
				inputType: "int",
			},
			wantErr:    false,
			wantChoice: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChoice, err := inputer(tt.args.l, tt.args.inputType)
			if (err != nil) != tt.wantErr {
				t.Errorf("inputer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotChoice, tt.wantChoice) {
				t.Errorf("inputer() gotChoice = %v, want %v", gotChoice, tt.wantChoice)
			}
		})
	}
}
